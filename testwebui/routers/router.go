package routers

import (
	"github.com/lixiaolong/golang/testwebui/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
