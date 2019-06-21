package api

import (
	"fmt"
	"github.com/astaxie/beego"
)

type UploadApiController struct {
	beego.Controller
}

func (c *UploadApiController) Get() {
	c.ServeJSON()
}

func (c *UploadApiController) Post() {
	fmt.Println("post lz")
	c.ServeJSON()
}
