package controllers

import (
	"myblog/goblog/helper"
	"myblog/goblog/models"

	"github.com/astaxie/beego"
)

type Resp struct {
	Data interface{}
}

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	c.Data["URL"] = "http://localhost:8080"
	c.Data["Name"] = "tsuchen"
	c.TplName = "login.html"
}

func (c *LoginController) Post() {
	resp := helper.NewResponse()

	defer resp.WriteRespByJson(c.Ctx.ResponseWriter)

	username := c.GetString("username")
	password := c.GetString("password")
	if username == "" && password == "" {
		resp.RespMessage(helper.RS_params_error, helper.WARING)
		models.NewUser()
		return
	}

	isFind := FindUser(username, password)
	if isFind {
		resp.RespMessage(helper.RS_success, helper.SUCCESS)
		resp.Data = "/homepage"
	} else {
		resp.RespMessage(helper.RS_password_error, helper.WARING)
	}
}

// 查找用户
func FindUser(userName string, password string) (isFind bool) {
	isFind = models.SelectUser(userName, password)

	return
}
