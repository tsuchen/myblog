package controllers

import (
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
		pageID, _ := c.GetInt(":page")
		totalPages, indexList, categoryList := models.GetTagByPageId(se, pageID)
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
		tagID, _ := strconv.Atoi(c.GetString("TagId"))
		tagName := c.GetString("TagName")
		success := false
		if oper == "add" {
			success = addTag(se, tagName)
		} else if oper == "delete" {
			success = deleteTag(se, tagName)
		} else if oper == "alter" {
			success = alterTag(se, tagID, tagName)
		}
		if success {
			resp.RespMessage(helper.RS_success, helper.SUCCESS)
			resp.Data = "/admin/taglist/p/" + strconv.Itoa(pageID)
		} else {
			resp.RespMessage(helper.RS_failed, helper.WARING)
		}
	} else {
		resp.RespMessage(helper.RS_failed, helper.WARING)
		c.Render()
	}
}

func addTag(userName interface{}, name string) (success bool) {
	success = models.AddTag(userName, name)
	return
}

func deleteTag(userName interface{}, name string) (success bool) {
	success = models.DeleteTag(userName, name)
	return
}

func alterTag(userName interface{}, id int, name string) (success bool) {
	success = models.AlterTag(userName, id, name)
	return
}
