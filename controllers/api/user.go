package api

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

type User struct {
	id       string
	phoneNum string
	nick     string
}

func (u *UserController) Register() {
	fmt.Println("register request...")
	var user User
	data := u.Ctx.Input.RequestBody
	fmt.Println(data)
	fmt.Println("info: " + u.GetString("jsoninfo"))
	err := json.Unmarshal(data, &user)
	if err != nil {
		fmt.Println("json.Unmarshal is err:", err.Error())
	}
	fmt.Println(user)
	response := make(map[string]interface{})
	response["result"] = "ok"
	response["msg"] = "success"
	u.Data["json"] = response
	u.ServeJSON()
}

func (u *UserController) Login() {

}

func (u *UserController) Delete() {

}
