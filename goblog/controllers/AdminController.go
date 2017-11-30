
package controllers

import (
	"github.com/astaxie/beego"
)

type AdminController struct {
	beego.Controller
}

func (c *AdminController) Get() {
	cookie, _:= c.Ctx.Request.Cookie(sessionName)
	se := c.GetSession(cookie.Name)
	if se != nil{
		c.TplName = "admin.html"
		c.Data["UserName"] = "xuchen"
	}else{
		c.Data["URL"] = "http://localhost:8080"
		c.Data["UserName"] = "xuchen"
		c.TplName = "login.html"
	}
}
