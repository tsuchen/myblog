package controllers

import (
	"myblog/goblog/models"
)

type AdminUserListController struct {
	CommonController
}

func (c *AdminUserListController) Get() {
	isLogin, se := c.checkUserStatus()
	if isLogin && se != nil {
		pageId, _ := c.GetInt(":page")
		pageIndexList, users := models.GetUsersByPageId(pageId)
		c.Data["Users"] = users
		c.Data["PageIndexList"] = pageIndexList
		c.Data["GroupMenuId"] = "user-menu"
		c.Layout = "adminhome.html"
		c.TplName = "userlist.html"
	}

	c.Render()
}
