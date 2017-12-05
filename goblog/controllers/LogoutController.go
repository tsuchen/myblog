package controllers

import (
	"myblog/goblog/helper"

	"github.com/astaxie/beego"
)

type LogoutController struct {
	beego.Controller
}

func (c *LogoutController) Get() {
	resp := helper.NewResponse()
	defer resp.WriteRespByJson(c.Ctx.ResponseWriter)

	resp.RespMessage(helper.RS_success, helper.SUCCESS)
	resp.Data = "/login"

	cookie, _ := c.Ctx.Request.Cookie(sessionName)
	se := c.GetSession(cookie.Name)
	helper.GlobalUserManager.DeleteUserInfo(se)
	c.DelSession(cookie.Name)
}
