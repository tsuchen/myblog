package controllers

import (
	"myblog/goblog/models"
)

type AdminBlogController struct {
	CommonController
}

func (c *AdminBlogController) Get() {
	if isLogin, se := c.checkUserStatus(); isLogin {
		categoryID, _ := c.GetInt(":cateid")
		blogs := models.GetBlogsByCategoryId(se, categoryID)
		c.Data["Blogs"] = blogs
		c.Data["GroupListId"] = "BlogList"
		c.Layout = "admin.html"
		c.TplName = "bloglist.html"
	}

	c.Render()
}
