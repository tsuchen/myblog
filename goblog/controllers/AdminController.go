package controllers

import (
	"myblog/goblog/helper"
	"myblog/goblog/models"
	"time"
)

type AdminController struct {
	CommonController
}

func (c *AdminController) Get() {
	isLogin, se := c.checkUserStatus()
	if isLogin && se != nil {
		user := models.GetUserByName(se)
		c.Data["User"] = user
		c.Data["Birth"] = (user.Profile.Birth).Format("2006-01-02")
		local, _ := time.LoadLocation("Local")
		created := user.Created.In(local)
		c.Data["CreateTime"] = created.Format("2006-01-02 15:04:05")
		updated := user.Updated.In(local)
		c.Data["UpdateTime"] = updated.Format("2006-01-02 15:04:05")
		c.Data["GroupMenuId"] = "user-menu"
		c.Layout = "adminhome.html"
		c.TplName = "userprofile.html"
	}

	c.Render()
}

func (c *AdminController) Post() {
	resp := helper.NewResponse()
	defer resp.WriteRespByJson(c.Ctx.ResponseWriter)

	isLogin, se := c.checkUserStatus()
	if isLogin && se != nil {
		if oper, err := c.GetInt("Type"); err == nil {
			if oper == 1 {
				//更新用户信息
				nickName := c.GetString("NickName")
				sex := c.GetString("Sex")
				birthStr := c.GetString("Birth")
				birth, _ := time.Parse("2006-01-02", birthStr)
				pNumber := c.GetString("PhoneNumber")
				email := c.GetString("Email")
				desc := c.GetString("Desc")
				userProfile := models.UserProfile{nickName, sex, pNumber, email, desc, birth}
				success := updateProfile(se, userProfile)
				if success {
					resp.RespMessage(helper.RS_success, helper.WARING)
					resp.Data = "/admin"
				} else {
					resp.RespMessage(helper.RS_failed, helper.WARING)
				}
			} else if oper == 2 {
				//更新用户密码
				oldPass := c.GetString("OldPassword")
				newPass := c.GetString("Password")
				success := updatePassword(se, oldPass, newPass)
				if success {
					resp.RespMessage(helper.RS_success, helper.WARING)
					resp.Data = "/admin"
				} else {
					resp.RespMessage(helper.RS_failed, helper.WARING)
				}
			}
		}
	} else {
		resp.RespMessage(helper.RS_failed, helper.WARING)
		c.Render()
	}
}

func updateProfile(userName interface{}, info models.UserProfile) (success bool) {
	success = models.UpdateUserProfile(userName, info)

	return
}

func updatePassword(userName interface{}, oldPass string, newPass string) (success bool) {
	success = models.UpdatePassword(userName, oldPass, newPass)

	return
}
