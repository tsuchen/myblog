package controllers

import (
	"fmt"
	"myblog/goblog/helper"
	"myblog/goblog/models"
)

type AdminTagController struct {
	CommonController
}

func (c *AdminTagController) Get() {
	isLogin, se := c.checkUserStatus()
	if isLogin && se != nil {
		tagList := models.GetAllTags(se)
		c.Data["TagList"] = tagList
		c.Data["GroupListId"] = "TagList"
		c.Layout = "admin.html"
		c.TplName = "taglist.html"
	}

	c.Render()
}

func (c *AdminTagController) Post() {
	resp := helper.NewResponse()
	defer resp.WriteRespByJson(c.Ctx.ResponseWriter)

	isLogin, se := c.checkUserStatus()
	if isLogin && se != nil {
		oper := c.GetString("Type")
		tagName := c.GetString("TagName")
		if oper == "add" {
			if success, tips := addTag(se, tagName); success {
				resp.RespMessage(helper.RS_success, helper.SUCCESS)
				resp.Data = "/admin/tag"
			} else {
				fmt.Println(tips)
				resp.RespMessage(helper.RS_failed, helper.WARING)
			}
		} else if oper == "delete" {
			if success, tips := deleteTag(se, tagName); success {
				resp.RespMessage(helper.RS_success, helper.SUCCESS)
				resp.Data = "/admin/tag"
			} else {
				fmt.Println(tips)
				resp.RespMessage(helper.RS_failed, helper.WARING)
			}
		}
	} else {
		resp.RespMessage(helper.RS_failed, helper.WARING)
		c.Render()
	}
}

func addTag(userName interface{}, tagName string) (success bool, message string) {
	success, message = models.AddTag(userName, tagName)

	return
}

func deleteTag(userName interface{}, tagName string) (success bool, message string) {
	success, message = models.DeleteTag(userName, tagName)
	return
}
