
package controllers

import (
	"fmt"
	"myblog/goblog/helper"
	"github.com/astaxie/beego"
)

type AdminController struct {
	beego.Controller
}

func (c *AdminController) Get() {
	cookie, _:= c.Ctx.Request.Cookie(sessionName)
	se := c.GetSession(cookie.Name)
	if se == nil {
		c.TplName = "login.html"
		c.Data["UserName"] = "xuchen"
		c.Data["URL"] = "http://localhost:8080"
	}else{
		userInfo := helper.GlobalUserManager.GetUserInfo()
		fmt.Println(userInfo)
		if userInfo.UserId != 0 {
			c.TplName = "admin.html"
			c.Data["UserName"] = "xuchen"
		}else{
			c.TplName = "login.html"
			c.Data["UserName"] = "xuchen"
			c.Data["URL"] = "http://localhost:8080"
		}
	}
	c.Render()
}
