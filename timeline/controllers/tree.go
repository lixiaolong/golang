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
	var page = "tree"

	c.Data["page"] = "tree"

	c.Layout = "layout.html"
	c.TplNames = page + "/mainpage.tpl"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["HtmlHead"] = page + "/html_head.tpl"
	c.LayoutSections["Scripts"] = page + "/scripts.tpl"
}
