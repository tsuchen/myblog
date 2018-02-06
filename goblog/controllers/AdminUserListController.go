package controllers

import (
	"myblog/goblog/helper"
	"myblog/goblog/models"
)

type AdminUserListController struct {
	CommonController
}

func (c *AdminUserListController) Get() {
	isLogin, se := c.checkUserStatus()
	if isLogin && se != nil {
		pageId, _ := c.GetInt(":page")
		totalPages, pageIndexList, users := models.GetUsersByPageId(pageId)
		c.Data["TotalPages"] = totalPages
		c.Data["Users"] = users
		c.Data["PageIndexList"] = pageIndexList
		c.Data["GroupMenuId"] = "user-menu"
		c.Layout = "adminhome.html"
		c.TplName = "userlist.html"
	}

	c.Render()
}

func (c *AdminUserListController) Post() {
	resp := helper.NewResponse()
	defer resp.WriteRespByJson(c.Ctx.ResponseWriter)

	isLogin, se := c.checkUserStatus()
	if isLogin && se != nil {
		oper := c.GetString("Type")
		if oper == "NewUser" {
			userName := c.GetString("UserName")
			password := c.GetString("Password")
			success := models.NewUser(userName, password)
			if success {
				resp.RespMessage(helper.RS_success, helper.SUCCESS)
				resp.Data = "/admin/userlist/p/1"
			} else {
				resp.RespMessage(helper.RS_failed, helper.WARING)
			}
		}
	} else {
		resp.RespMessage(helper.RS_failed, helper.WARING)
		c.Render()
	}
}
