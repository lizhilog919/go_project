package routers

import (
	"github.com/astaxie/beego"
	"go_project/controllers/api"
)

func init() {
	beego.Router("/api/upload", &api.UploadApiController{})
	beego.Router("/api/user/register", &api.UserController{}, "*:Register")
	beego.Router("/api/user/login", &api.UserController{}, "*:Login")
	beego.Router("/api/user/delete", &api.UserController{}, "POST:Delete")
}
