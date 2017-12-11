package controllers

import(
	"github.com/astaxie/beego"
	"myblog/goblog/helper"
	"fmt"
)

var sessionName string = beego.AppConfig.String("SessionName")

type CommonController struct{
	beego.Controller
}

func (c *CommonController) checkUserStatus() (hasLogin bool, session interface{}) {
	cookie, _ := c.Ctx.Request.Cookie(sessionName)
	se := c.GetSession(cookie.Name)
	if se == nil {
		fmt.Println("session不存在, 请先登录。")
		c.TplName = "login.html"
		c.Data["UserName"] = "xuchen"
		c.Data["URL"] = "http://localhost:8080"
		hasLogin = false
	} else {
		userInfo := helper.GlobalUserManager.GetUserInfo(se)
		if userInfo == nil {
			fmt.Println("用户信息不存在, 请重新登录。")
			c.TplName = "login.html"
			c.Data["UserName"] = "xuchen"
			c.Data["URL"] = "http://localhost:8080"
			hasLogin = false
		} else {
			c.Data["UserName"] = userInfo.UserName
			session = se
			hasLogin = true
		}
	}

	return 
}