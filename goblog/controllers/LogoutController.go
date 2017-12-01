package controllers

import (
	"github.com/astaxie/beego"
	"myblog/goblog/helper"
)

type LogoutController struct{
	beego.Controller
}

func (c *LogoutController) Get(){
	resp := helper.NewResponse()
	defer resp.WriteRespByJson(c.Ctx.ResponseWriter)
	
	resp.RespMessage(helper.RS_success, helper.SUCCESS)
	resp.Data = "/login"

	cookie, _:= c.Ctx.Request.Cookie(sessionName)
	c.DelSession(cookie.Name)
}