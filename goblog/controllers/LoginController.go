package controllers

import (
	"myblog/goblog/helper"

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
	}else if username == "tsuchen" && password == "123456" {
		resp.RespMessage(helper.RS_success, helper.SUCCESS)
		resp.Data = "/homepage"
	}else{
		
	}
}
