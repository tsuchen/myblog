package controllers

import (
	"fmt"
	"myblog/goblog/helper"
	"myblog/goblog/models"

	"github.com/astaxie/beego"
)

type AdminController struct {
	beego.Controller
}

func (c *AdminController) Get() {
	cookie, _ := c.Ctx.Request.Cookie(sessionName)
	se := c.GetSession(cookie.Name)
	if se == nil {
		fmt.Println("session不存在, 请先登录。")
		c.TplName = "login.html"
		c.Data["UserName"] = "xuchen"
		c.Data["URL"] = "http://localhost:8080"
	} else {
		userInfo := helper.GlobalUserManager.GetUserInfo(se)
		if userInfo == nil {
			fmt.Println("用户不存在, 请重新登录。")
			c.TplName = "login.html"
			c.Data["UserName"] = "xuchen"
			c.Data["URL"] = "http://localhost:8080"
		} else {
			users := models.GetAllUser()
			c.Data["Users"] = users
			c.Data["UserName"] = userInfo.UserName
			c.Layout = "admin.html"
			c.TplName = "userlist.html"
		}
	}
	c.Render()
}
