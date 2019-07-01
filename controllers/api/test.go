package api

import "github.com/astaxie/beego"

type TestController struct {
	beego.Controller
}

func (c *TestController) Get() {
	response := make(map[string]interface{})
	response["result"] = "ok"
	c.Data["json"] = response
	c.ServeJSON()
}
