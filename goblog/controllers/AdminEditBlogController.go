package controllers

import (
	"myblog/goblog/helper"
	"myblog/goblog/models"
)

type AdminEditBlogController struct {
	CommonController
}

func (c *AdminEditBlogController) Get() {
	if isLogin, se := c.checkUserStatus(); isLogin {
		blogID, _ := c.GetInt(":blogid")
		tags := models.GetAllTags(se)
		c.Data["Tags"] = tags
		article, err := models.GetArticleByID(blogID)
		if err == nil {
			c.Data["IsNew"] = false
			c.Data["SelectCate"] = article.Category.Name
			c.Data["Title"] = article.Title
			var tagStr string
			selectTags := article.Tags
			for _, tag := range selectTags {
				tagStr += tag.Name + ";"
			}
			c.Data["SelectTags"] = tagStr
			c.Data["Content"] = article.Content
			c.Data["BlogID"] = blogID
		} else {
			c.Data["IsNew"] = true
		}
		c.Data["GroupMenuId"] = "editblog-menu"
		c.Layout = "adminhome.html"
		c.TplName = "editblog.html"
	}

	c.Render()
}

func (c *AdminEditBlogController) Post() {
	resp := helper.NewResponse()
	defer resp.WriteRespByJson(c.Ctx.ResponseWriter)
	if isLogin, _ := c.checkUserStatus(); isLogin {
		blogID, _ := c.GetInt("blogid")
		model := c.GetString("Type")
		if model == "save" {

		} else if model == "delete" {

		} else if model == "send" {

		}
	} else {
		c.Render()
	}
}

// 发表博客
func sendAtricle(userName interface{}, id string, title string, cate string, tags string, content string) (success bool) {
	args := make(map[string]string)
	args["title"] = title
	args["blogid"] = id
	args["category"] = cate
	args["tags"] = tags
	args["content"] = content
	success = models.SendArticleByID(userName, args)

	return
}
