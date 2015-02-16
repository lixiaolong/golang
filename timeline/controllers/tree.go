package controllers

import (
	"github.com/astaxie/beego"
	//"github.com/lixiaolong/golang/timeline/models"
)

type TreeController struct {
	beego.Controller
}

func (c *TreeController) Get() {
	//dbname := beego.AppConfig.String("dbname")
	//apptitle := beego.AppConfig.String("apptitle")

	c.Data["page"] = "tree"

	c.Layout = "layout.html"
	c.TplNames = "tree/tree.tpl"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["HtmlHead"] = "tree/html_head.tpl"
	c.LayoutSections["Menu"] = "tree/menu.tpl"
	c.LayoutSections["Scripts"] = "tree/scripts.tpl"
}
