package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/lionsoul2014/ip2region/binding/golang/ip2region"
)

var (
	region *ip2region.Ip2Region
	err    error
)

func init() {
	region, err = ip2region.New("./ip2region.db")
	defer region.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func isExist(name string) bool {
	_, err := os.Stat(name)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func newExcelOutput(name, sheet string) (outfd *excelize.File, err error) {
	if isExist(name) {
		outfd, err = excelize.OpenFile(name)
	} else {
		outfd = excelize.NewFile()
		outfd.SetCellValue(sheet, "A1", "IP")
		outfd.SetCellValue(sheet, "B1", "国家")
		outfd.SetCellValue(sheet, "C1", "省份")
		outfd.SetCellValue(sheet, "D1", "城市")
		outfd.SetCellValue(sheet, "E1", "运营商")
		outfd.SetCellValue(sheet, "F1", "次数")
	}
	return
}

func readTxt(name, output string) {
	fd, err := os.Open(name)
	if err != nil {
		fmt.Println(err)
		return
	}

	row := 2
	sheet := "Sheet1"
	outfd, err := newExcelOutput(output, sheet)
	if err != nil {
		fmt.Println("open failed ", err)
		return
	}

	br := bufio.NewReader(fd)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		ip := strings.TrimSpace(string(a))
		loc, err := region.MemorySearch(ip)
		if err != nil {
			fmt.Println(err)
			return
		}

		outfd.SetCellValue(sheet, "A"+strconv.Itoa(row), ip)
		outfd.SetCellValue(sheet, "B"+strconv.Itoa(row), loc.Country)
		outfd.SetCellValue(sheet, "C"+strconv.Itoa(row), loc.Province)
		outfd.SetCellValue(sheet, "D"+strconv.Itoa(row), loc.City)
		outfd.SetCellValue(sheet, "E"+strconv.Itoa(row), loc.ISP)
		row++
	}

	if err := outfd.SaveAs(output); err != nil {
		println(err.Error())
	}
}

// IPList for ip list
type IPList struct {
	IP  string `json:"key"`
	Cnt int    `json:"doc_count"`
}

func readJSON(name string, ipmap map[string]int) {
	data, err := ioutil.ReadFile(name)
	if err != nil {
		fmt.Println("open failed  ", err)
		return
	}

	var ipl []IPList

	err = json.Unmarshal(data, &ipl)
	if err != nil {
		fmt.Println("json un failed  ", err)
		return
	}

	for _, v := range ipl {
		ip := strings.TrimSpace(v.IP)
		if _, ok := ipmap[ip]; ok {
			ipmap[ip] += v.Cnt
		} else {
			ipmap[ip] = v.Cnt
		}
	}

}

func outputExcel(ipmap map[string]int, output string) {
	row := 2
	sheet := "Sheet1"
	outfd, err := newExcelOutput(output, sheet)
	if err != nil {
		fmt.Println("open failed ", err)
		return
	}

	for k, v := range ipmap {
		loc, err := region.MemorySearch(k)
		if err != nil {
			fmt.Println(err)
			return
		}

		outfd.SetCellValue(sheet, "A"+strconv.Itoa(row), k)
		outfd.SetCellValue(sheet, "B"+strconv.Itoa(row), loc.Country)
		outfd.SetCellValue(sheet, "C"+strconv.Itoa(row), loc.Province)
		outfd.SetCellValue(sheet, "D"+strconv.Itoa(row), loc.City)
		outfd.SetCellValue(sheet, "E"+strconv.Itoa(row), loc.ISP)
		outfd.SetCellValue(sheet, "F"+strconv.Itoa(row), v)
		row++
	}

	if err := outfd.SaveAs(output); err != nil {
		println(err.Error())
	}
}

func readJSONDir(pathname, output string) {
	ipmap := make(map[string]int, 0)

	err := filepath.Walk(pathname, func(pathname string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		if strings.HasSuffix(pathname, ".json") {
			readJSON(pathname, ipmap)
		}
		return nil
	})

	outputExcel(ipmap, output)

	if err != nil {
		fmt.Println(err)
		return
	}
}

func readString(pathname string) {
	iplist := strings.Split(pathname, ";")
	for _, v := range iplist {
		ip := strings.TrimSpace(v)
		if len(ip) == 0 {
			continue
		}
		loc, err := region.MemorySearch(ip)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("%s,%s,%s,%s,%s\n", ip, loc.Country, loc.Province, loc.City, loc.ISP)
	}

}

func main() {
	pathname := "input"
	output := "output.xlsx"
	flag.StringVar(&pathname, "i", "input", "input file name(all.txt | all.json | input/ | ip1;ip2)")
	flag.StringVar(&output, "o", "output.xlsx", "output file name")
	flag.Parse()

	if strings.Contains(pathname, ";") {
		readString(pathname)
		return
	}

	fi, err := os.Stat(pathname)
	if err == nil {
		if fi.IsDir() {
			fmt.Println("parse json file in dir : ", pathname)
			readJSONDir(pathname, output)
		} else if strings.HasSuffix(pathname, ".json") {
			fmt.Println("parse json file : ", pathname)
			ipmap := make(map[string]int, 0)
			readJSON(pathname, ipmap)
			outputExcel(ipmap, "json-"+output)
		} else {
			fmt.Println("parse txt file : ", pathname)
			readTxt(pathname, "txt-"+output)
		}
	} else {
		fmt.Println(err)
	}
}
