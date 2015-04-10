package main

import (
	"github.com/astaxie/beego"
	_ "github.com/lixiaolong/golang/timeline/routers"
)

func Add(a, b int) int {
	return a + b
}

func Sub(a, b int) int {
	return a - b
}

func main() {
	beego.AddFuncMap("markdown", Markdown)
	beego.AddFuncMap("add", Add)
	beego.AddFuncMap("sub", Sub)
	beego.Run()
}
