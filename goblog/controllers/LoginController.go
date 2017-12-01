package controllers

import (
	"myblog/goblog/helper"
	"myblog/goblog/models"

	"github.com/astaxie/beego"
)

var sessionName = beego.AppConfig.String("SessionName")

type Resp struct {
	Data interface{}
}

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	c.Data["URL"] = "http://localhost:8080"
	c.Data["Name"] = "xuchen"
	c.TplName = "login.html"

	c.Render()
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

	isFind, user := FindUser(username, password)
	if isFind {
		userInfo := helper.GlobalUserManager.GetUserInfo()
		if userInfo.UserId != 0 {
			return 
		}
		userInfo.UserId = user.ID
		userInfo.UserName = user.Name
		userInfo.Age = user.Profile.Age

		// 初始化session 
		se := c.GetSession(sessionName)
		if se == nil {
			c.SetSession(sessionName, username)
		} else {
			c.SetSession(sessionName, username)
		}
		resp.RespMessage(helper.RS_success, helper.SUCCESS)
		resp.Data = "/admin"
	} else {
		resp.RespMessage(helper.RS_password_error, helper.WARING)
	}
}

// 查找用户
func FindUser(userName string, password string) (isFind bool, user models.User) {
	isFind, user = models.SelectUser(userName, password)

	return
}
