package controllers

import "myblog/goblog/models"

type AdminEditBlogController struct {
	CommonController
}

func (c *AdminEditBlogController) Get() {
	if isLogin, se := c.checkUserStatus(); isLogin {
		tagList := models.GetAllTags(se)
		c.Data["Tags"] = tagList
		c.TplName = "editblog.html"
	}

	c.Render()
}

func (c *AdminEditBlogController) Post() {

}
