package routers

import (
	"github.com/astaxie/beego"
	"github.com/lixiaolong/golang/timeline/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/tree", &controllers.TreeController{})
}
