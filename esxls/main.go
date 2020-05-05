package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

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

// Dstip for dst ip
type Dstip struct {
	IP  string `json:"ip_dst_addr"`
	Pro string `json:"enrichments:geo:ip_src_addr:province"`
}

type Hits1 struct {
	Hits []Hits2 `json:"hits"`
}
type Hits2 struct {
	Source Dstip   `json:"_source"`
	Time   []int64 `json:"sort"`
}

type Timestamp struct {
	Hits Hits1 `json:"hits"`
}

// Record for record
type Record struct {
	IP        string    `json:"key"`
	Cnt       int       `json:"doc_count"`
	Timestamp Timestamp `json:"timestamp1"`
}

type Info struct {
	Provice string
	Srcip   string
	Dstip   string
	Time    string
	Cnt     int
}

func readJSON(name string, ipmap map[string]Info) {
	var rd []Record

	data, err := ioutil.ReadFile(name)
	if err != nil {
		fmt.Println("open failed  ", err)
		return
	}
	err = json.Unmarshal(data, &rd)
	if err != nil {
		fmt.Println("json failed  ", err)
		return
	}
	// fmt.Printf("%+v", rd)

	for _, v := range rd {
		t := v.Timestamp.Hits.Hits[0].Time[0] / 1000
		ts := time.Unix(t, 0).Format("2006-01-02 15:04")
		// fmt.Printf("%s,%s,%s,%s,%d\n", v.Timestamp.Hits.Hits[0].Source.Pro, v.IP,
		// 	v.Timestamp.Hits.Hits[0].Source.IP,
		// 	ts, v.Cnt)
		sip := strings.TrimSpace(v.IP)
		dip := strings.TrimSpace(v.Timestamp.Hits.Hits[0].Source.IP)
		if _, ok := ipmap[sip+dip]; !ok {
			info := Info{}
			loc, err := region.MemorySearch(sip)
			if err != nil {
				info.Provice = v.Timestamp.Hits.Hits[0].Source.Pro
			} else {
				if loc.Country == "中国" {
					info.Provice = loc.Province
				} else {
					info.Provice = loc.Country
				}
			}
			info.Srcip = sip
			info.Dstip = dip

			info.Time = ts
			info.Cnt = v.Cnt
			ipmap[sip+dip] = info
		}
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
		outfd.SetCellValue(sheet, "A1", "省份")
		outfd.SetCellValue(sheet, "B1", "源IP")
		outfd.SetCellValue(sheet, "C1", "目的IP")
		outfd.SetCellValue(sheet, "D1", "访问时间")
		outfd.SetCellValue(sheet, "E1", "次数")
	}
	return
}

func outputExcel(ipmap map[string]Info, output string) {
	row := 2
	sheet := "Sheet1"
	outfd, err := newExcelOutput(output, sheet)
	if err != nil {
		fmt.Println("open failed ", err)
		return
	}

	for _, v := range ipmap {
		outfd.SetCellValue(sheet, "A"+strconv.Itoa(row), v.Provice)
		outfd.SetCellValue(sheet, "B"+strconv.Itoa(row), v.Srcip)
		outfd.SetCellValue(sheet, "C"+strconv.Itoa(row), v.Dstip)
		outfd.SetCellValue(sheet, "D"+strconv.Itoa(row), v.Time)
		outfd.SetCellValue(sheet, "E"+strconv.Itoa(row), v.Cnt)
		row++
	}

	if err := outfd.SaveAs(output); err != nil {
		println(err.Error())
	}
}

func readDir(pathname string, ipmap map[string]Info) {
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

	if err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	ipmap := make(map[string]Info, 0)
	readDir("input", ipmap)
	outputExcel(ipmap, "output.xlsx")
	// fmt.Printf("%+v\n", ipmap)
}
