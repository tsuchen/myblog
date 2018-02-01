package controllers

import (
	"myblog/goblog/helper"
	"myblog/goblog/models"
)

type Resp struct {
	Data interface{}
}

type LoginController struct {
	CommonController
}

func (c *LoginController) Get() {
	c.Data["URL"] = defaultDomain
	c.Data["Name"] = defaultAdmin
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

	isFind, user := findUser(username, password)
	if isFind {
		//更新用户信息
		updateUserInfo(user)
		// 初始化session
		se := c.GetSession(sessionName)
		if se == nil {
			c.SetSession(sessionName, username)
		}
		resp.RespMessage(helper.RS_success, helper.SUCCESS)
		resp.Data = "/admin"
	} else {
		resp.RespMessage(helper.RS_password_error, helper.WARING)
	}
}

// 查找用户
func findUser(userName string, password string) (isFind bool, user models.User) {
	isFind, user = models.SelectUser(userName, password)

	return
}

// 更新用户信息
func updateUserInfo(user models.User) {
	newUserInfo := &helper.UserInfo{UserId: user.ID, UserName: user.Name, Password: user.Password}
	helper.GlobalUserManager.UpdateUserInfo(newUserInfo)
}
