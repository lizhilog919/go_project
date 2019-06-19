package routers

import (
	"github.com/astaxie/beego"
	"go_project/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
