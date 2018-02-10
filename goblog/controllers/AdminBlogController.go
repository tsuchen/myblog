package controllers

import (
	"myblog/goblog/models"
)

type AdminBlogController struct {
	CommonController
}

func (c *AdminBlogController) Get() {
	if isLogin, se := c.checkUserStatus(); isLogin {
		cateID, _ := c.GetInt(":cateid")
		pageID, _ := c.GetInt(":page")
		cateName := models.GetCategoryNameById(cateID)
		c.Data["CategoryName"] = cateName
		c.Data["CateID"] = cateID
		totalPages, indexList, blogs := models.GetBlogs(se, cateID, pageID)
		c.Data["TotalPages"] = totalPages
		c.Data["PageIndexList"] = indexList
		c.Data["Blogs"] = blogs
		c.Data["GroupMenuId"] = "article-menu"
		c.Layout = "adminhome.html"
		c.TplName = "bloglist.html"
	}

	c.Render()
}

func (c *AdminBlogController) Post() {

}
