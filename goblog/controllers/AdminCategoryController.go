package controllers

import (
	"fmt"
	"myblog/goblog/helper"
	"myblog/goblog/models"

	"github.com/astaxie/beego"
)

type AdminCategoryController struct {
	beego.Controller
}

func (c *AdminCategoryController) Get() {
	cookie, _ := c.Ctx.Request.Cookie(sessionName)
	se := c.GetSession(cookie.Name)
	if se == nil {
		c.Data["URL"] = "http://localhost:8080"
		c.Data["Name"] = "xuchen"
		c.TplName = "login.html"
	} else {
		//通过用户名获取分类
		userInfo := helper.GlobalUserManager.GetUserInfo(se)
		if userInfo == nil {
			fmt.Println("用户信息不存在, 请重新登录。")
			c.TplName = "login.html"
			c.Data["UserName"] = "xuchen"
			c.Data["URL"] = "http://localhost:8080"
		} else {
			categoryList := models.GetAllCategory(se)
			c.Data["Categorys"] = categoryList
			c.Data["UserName"] = userInfo.UserName
			c.Layout = "admin.html"
			c.TplName = "categorylist.html"
		}
	}

	c.Render()
}
