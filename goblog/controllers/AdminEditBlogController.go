package controllers

import (
	"myblog/goblog/models"
)

type AdminEditBlogController struct {
	CommonController
}

func (c *AdminEditBlogController) Get() {
	if isLogin, _ := c.checkUserStatus(); isLogin {
		blogID, _ := c.GetInt(":blogid")
		article, err := models.GetArticleByID(blogID)
		if err == nil {
			c.Data["IsNew"] = false
		} else {
			c.Data["IsNew"] = true
		}
		c.Data["Article"] = article
		c.Data["GroupMenuId"] = "editblog-menu"
		c.Layout = "adminhome.html"
		c.TplName = "editblog.html"
	}

	c.Render()
}
