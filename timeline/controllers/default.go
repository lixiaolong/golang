package controllers

import (
	"github.com/astaxie/beego"
	"github.com/lixiaolong/golang/timeline/models"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	//dbname, _ := beego.Config("string", "dbname", "models/timeline.db")
	//apptitle, _ := beego.Config("string", "apptitle", "记录")
	//apptitle := "版本修改记录"
	//dbname := "models/timeline.db"

	dbname := beego.AppConfig.String("dbname")
	apptitle := beego.AppConfig.String("apptitle")

	tl := models.NewTimeline()
	tl.SetTitle(apptitle)

	tl.LoadDB(dbname, 0, 10)

	c.Data["ver"] = &tl

	c.TplNames = "index.tpl"
}

func (c *MainController) Post() {
	//apptitle, _ := beego.Config("string", "apptitle", "记录")
	dbname := beego.AppConfig.String("dbname")
	apptitle := beego.AppConfig.String("apptitle")

	title := c.GetString("title")
	time := c.GetString("time")
	details := c.GetString("details")

	tl := models.NewTimeline()
	tl.SetTitle(apptitle)
	tl.AppendLine(time, title, details)
	tl.SaveDB(dbname)
	c.Redirect("/", 302)
}
