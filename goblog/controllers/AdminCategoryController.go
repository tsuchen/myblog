package controllers

import (
	"myblog/goblog/helper"
	"myblog/goblog/models"
	"strconv"
)

type AdminCategoryController struct {
	CommonController
}

func (c *AdminCategoryController) Get() {
	isLogin, se := c.checkUserStatus()
	if isLogin && se != nil {
		pageID, _ := c.GetInt(":page")
		totalPages, indexList, categoryList := models.GetCategoryByPageId(se, pageID)
		c.Data["TotalPages"] = totalPages
		c.Data["PageIndexList"] = indexList
		c.Data["Categorys"] = categoryList
		c.Data["GroupMenuId"] = "category-menu"
		c.Layout = "adminhome.html"
		c.TplName = "categorylist.html"
	}

	c.Render()
}

func (c *AdminCategoryController) Post() {
	resp := helper.NewResponse()
	defer resp.WriteRespByJson(c.Ctx.ResponseWriter)

	isLogin, se := c.checkUserStatus()
	if isLogin && se != nil {
		pageID, _ := c.GetInt(":page")
		oper := c.GetString("Type")
		categoryName := c.GetString("CategoryName")
		categoryID, _ := strconv.Atoi(c.GetString("CategoryId"))
		success := false
		if oper == "add" {
			//添加分类
			success = addCategory(se, categoryName)
		} else if oper == "delete" {
			//删除分类
			success = deleteCategory(se, categoryName)
		} else {
			//修改分类
			success = alterCategory(se, categoryID, categoryName)
		}
		if success {
			resp.RespMessage(helper.RS_success, helper.SUCCESS)
			resp.Data = "/admin/categorylist/p/" + strconv.Itoa(pageID)
		} else {
			resp.RespMessage(helper.RS_failed, helper.WARING)
		}
	} else {
		resp.RespMessage(helper.RS_failed, helper.WARING)
		c.Render()
	}
}

func addCategory(userName interface{}, name string) (success bool) {
	success = models.AddBlogCategory(userName, name)
	return
}

func deleteCategory(userName interface{}, name string) (success bool) {
	success = models.DeleteCategory(userName, name)
	return
}

func alterCategory(userName interface{}, id int, name string) (success bool) {
	success = models.AlterCategory(userName, id, name)
	return
}
