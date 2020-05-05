package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// Rendata is data in res
type Rendata struct {
	Tname string `json:"tname"`
	Time  int64  `json:"time"`
	Link  string `json:"link"`
}

// DYRes is Douyu result
type DYRes struct {
	Type  string  `json:"type"`
	Data  Rendata `json:"Rendata"`
	Info  string  `json:"info"`
	State string  `json:"state"`
}

func getLiveURL(id string) (link string) {
	durl := "https://web.sinsyth.com/lxapi/douyujx.x?roomid="

	res, err := http.Get(durl + id)
	if err != nil {
		fmt.Println("get url failed:", err)
		return
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("read failed:", err)
		return
	}
	var dyres DYRes
	err = json.Unmarshal(data, &dyres)
	if err != nil {
		fmt.Println("json failed:", err)
		return
	}

	if dyres.State == "SUCCESS" {
		if dyres.Data.Link != "" {
			link = dyres.Data.Link
			return
		}
	} else if dyres.State == "NO" {
		fmt.Printf("closed")
	}

	return
}

func parseLiveURL(link string) (result string) {
	e := strings.LastIndex(link, "_")
	if e != -1 {
		link = link[:e]
	} else {
		e = strings.Index(link, "?")
		if e != -1 {
			link = link[:e]
		}

		e = strings.LastIndex(link, ".")
		if e != -1 {
			link = link[:e]
		}
	}

	s := strings.Index(link, ".")
	if s != -1 {
		link = link[s:]
	}

	if link != "" {
		result = "http://tx2play1" + link + ".flv"
	}

	return
}

func main() {
	var id string

	flag.StringVar(&id, "id", "9500,208114,93589,142488", "douyu id")
	flag.Parse()

	ids := strings.Split(id, ",")
	for _, v := range ids {
		link := getLiveURL(strings.TrimSpace(v))
		fmt.Println(parseLiveURL(link))
	}
}
