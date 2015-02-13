package routers

import (
	"github.com/lixiaolong/golang/timeline/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
