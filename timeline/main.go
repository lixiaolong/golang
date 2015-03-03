package main

import (
	"github.com/astaxie/beego"
	_ "github.com/lixiaolong/golang/timeline/routers"
)

func main() {
	beego.AddFuncMap("markdown", Markdown)
	beego.Run()
}
