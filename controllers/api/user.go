package api

import (
	"fmt"
	"github.com/astaxie/beego"
	"go_project/models"
)

type UserController struct {
	beego.Controller
}

func (u *UserController) Register() {
	fmt.Println("register request...")
	response := make(map[string]interface{})
	errorMsg := ""
	nick := u.GetString("nick")
	if nick == "" {
		errorMsg = "require nick"
	}
	phoneNum := u.GetString("phoneNum")
	if phoneNum == "" {
		errorMsg = "require phoneNum"
	}
	pwd := u.GetString("pwd")
	if pwd == "" {
		errorMsg = "require pwd"
	}
	var id int
	var err error
	if errorMsg == "" {
		id, err = models.InsertUser(phoneNum, nick, pwd)
		if err != nil {
			errorMsg = err.Error()
			response["result"] = "error"
			response["msg"] = errorMsg
			fmt.Println("error: " + errorMsg)
		} else {
			response["result"] = "ok"
			response["msg"] = ""
			response["id"] = id
		}
	} else {
		response["result"] = "error"
		response["msg"] = errorMsg
	}
	u.Data["json"] = response
	u.ServeJSON()
}

func (u *UserController) Login() {
	response := make(map[string]interface{})
	errorMsg := ""
	phoneNum := u.GetString("phoneNum")
	if phoneNum == "" {
		errorMsg = "require phoneNum"
	}
	pwd := u.GetString("pwd")
	if pwd == "" {
		errorMsg = "require pwd"
	}
	if errorMsg == "" {

	} else {
		response["result"] = "error"
		response["msg"] = errorMsg
	}
	var user *models.User
	var err error
	user, err = models.Login(phoneNum, pwd)
	if err != nil {
		response["result"] = "error"
		response["msg"] = err.Error()
	} else {
		fmt.Println(fmt.Sprintf("id: %d", user.Id))
		response["result"] = "ok"
		response["msg"] = ""
		response["data"] = user
	}
	u.Data["json"] = response
	u.ServeJSON()
}

func (u *UserController) Delete() {

}
