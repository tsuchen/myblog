package controllers

import (
	"fmt"
	"myblog/goblog/helper"
	"myblog/goblog/models"
	"strconv"
)

type AdminTagController struct {
	CommonController
}

func (c *AdminTagController) Get() {
	isLogin, se := c.checkUserStatus()
	if isLogin && se != nil {
		pageId, _ := c.GetInt(":page")
		totalPages, indexList, categoryList := models.GetTagByPageId(se, pageId)
		c.Data["TotalPages"] = totalPages
		c.Data["PageIndexList"] = indexList
		c.Data["Tags"] = categoryList
		c.Data["GroupMenuId"] = "tag-menu"
		c.Layout = "adminhome.html"
		c.TplName = "taglist.html"
	}

	c.Render()
}

func (c *AdminTagController) Post() {
	resp := helper.NewResponse()
	defer resp.WriteRespByJson(c.Ctx.ResponseWriter)

	isLogin, se := c.checkUserStatus()
	if isLogin && se != nil {
		pageID, _ := c.GetInt(":page")
		oper := c.GetString("Type")
		tagID := c.GetString("TagId")
		tagName := c.GetString("TagName")
		if oper == "add" {
			if success, tips := addTag(se, tagName); success {
				resp.RespMessage(helper.RS_success, helper.SUCCESS)
				resp.Data = "/admin/taglist/p/" + strconv.Itoa(pageID)
			} else {
				fmt.Println(tips)
				resp.RespMessage(helper.RS_failed, helper.WARING)
			}
		} else if oper == "delete" {
			if success, tips := deleteTag(se, tagName); success {
				resp.RespMessage(helper.RS_success, helper.SUCCESS)
				resp.Data = "/admin/taglist/p/" + strconv.Itoa(pageID)
			} else {
				fmt.Println(tips)
				resp.RespMessage(helper.RS_failed, helper.WARING)
			}
		} else if oper == "alter" {
			if success, tips := alterTag(se, tagID, tagName); success {
				resp.RespMessage(helper.RS_success, helper.SUCCESS)
				resp.Data = "/admin/taglist/p/" + strconv.Itoa(pageID)
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

func addTag(userName interface{}, name string) (success bool, message string) {
	success, message = models.AddTag(userName, name)

	return
}

func deleteTag(userName interface{}, name string) (success bool, message string) {
	success, message = models.DeleteTag(userName, name)
	return
}

func alterTag(userName interface{}, id string, name string) (success bool, message string) {
	success, message = models.AlterTag(userName, id, name)
	return
}
