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
		cur, total, users := models.GetUsersByPageId(pageId)
		var pageIndexList []int
		for i := 1; i <= int(total); i++ {
			pageIndexList = append(pageIndexList, i)
		}
		c.Data["Users"] = users
		c.Data["CurPageIndex"] = cur
		c.Data["PageIndexList"] = pageIndexList
		c.Data["GroupMenuId"] = "user-menu"
		c.Layout = "adminhome.html"
		c.TplName = "userlist.html"
	}

	c.Render()
}
