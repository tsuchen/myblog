package controllers

import (
	"myblog/goblog/helper"
	"myblog/goblog/models"
	"strconv"
)

type AdminEditBlogController struct {
	CommonController
}

func (c *AdminEditBlogController) Get() {
	if isLogin, se := c.checkUserStatus(); isLogin {
		blogID, _ := c.GetInt(":blogid")
		tags := models.GetAllTags(se)
		c.Data["Tags"] = tags
		blog := models.GetArticleByID(blogID)
		if blog != nil {
			article := blog.(models.Blog)
			c.Data["IsNew"] = false
			c.Data["SelectedCate"] = article.Category.Name
			c.Data["Title"] = article.Title
			c.Data["Content"] = article.Content
			var tagStr string
			selectTags := article.Tags
			for _, tag := range selectTags {
				tagStr += tag.Name + ";"
			}
			c.Data["SelectTags"] = tagStr
		} else {
			c.Data["IsNew"] = true
		}
		c.Data["BlogID"] = blogID
		c.Data["GroupMenuId"] = "editblog-menu"
		c.Layout = "adminhome.html"
		c.TplName = "editblog.html"
	}

	c.Render()
}

func (c *AdminEditBlogController) Post() {
	resp := helper.NewResponse()
	defer resp.WriteRespByJson(c.Ctx.ResponseWriter)
	if isLogin, se := c.checkUserStatus(); isLogin {
		blogID, _ := c.GetInt(":blogid")
		model := c.GetString("Type")
		var success bool
		if model == "save" {

		} else if model == "delete" {
			success := deleteArticle(se, blogID)
			if success {
				resp.RespMessage(helper.RS_success, helper.SUCCESS)
				resp.Data = "/admin/bloglist/cate/1/p/1"
			} else {
				resp.RespMessage(helper.RS_failed, helper.WARING)
			}
		} else if model == "send" {
			//发表文章
			id := strconv.Itoa(blogID)
			title := c.GetString("Title")
			cate := c.GetString("Cate")
			tags := c.GetString("Tags")
			content := c.GetString("Content")
			success = sendArticle(se, id, title, cate, tags, content)
			if success {
				resp.RespMessage(helper.RS_success, helper.SUCCESS)
				resp.Data = "/admin/bloglist/cate/1/p/1"
			} else {
				resp.RespMessage(helper.RS_failed, helper.WARING)
			}
		}
	} else {
		resp.RespMessage(helper.RS_failed, helper.WARING)
		c.Render()
	}
}

// 发表文章
func sendArticle(userName interface{}, id string, title string, cate string, tags string, content string) (success bool) {
	args := make(map[string]string)
	args["title"] = title
	args["blogid"] = id
	args["category"] = cate
	args["tags"] = tags
	args["content"] = content
	success = models.SendArticleByID(userName, args)

	return
}

//删除文章
func deleteArticle(userName interface{}, id int) (success bool) {
	success = models.DeleteArticle(userName, id)
	return
}
