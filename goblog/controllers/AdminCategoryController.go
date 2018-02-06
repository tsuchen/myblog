package controllers

import (
	"fmt"
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
		pageId, _ := c.GetInt(":page")
		oper := c.GetString("Type")
		categoryName := c.GetString("CategoryName")
		categoryId := c.GetString("CategoryId")
		if oper == "add" {
			//添加分类
			success, message := addCategory(se, categoryName)
			fmt.Println(message)
			if success {
				resp.RespMessage(helper.RS_success, helper.SUCCESS)
				resp.Data = "/admin/categorylist/p/" + strconv.Itoa(pageId)
			} else {
				resp.RespMessage(helper.RS_failed, helper.WARING)
			}
		} else if oper == "delete" {
			//删除分类
			success, message := deleteCategory(se, categoryName)
			fmt.Println(message)
			if success {
				resp.RespMessage(helper.RS_success, helper.SUCCESS)
				resp.Data = "/admin/categorylist/p/" + strconv.Itoa(pageId)
			} else {
				resp.RespMessage(helper.RS_failed, helper.WARING)
			}
		} else {
			//修改分类
			success, message := alterCategory(se, categoryId, categoryName)
			fmt.Println(message)
			if success {
				resp.RespMessage(helper.RS_success, helper.SUCCESS)
				resp.Data = "/admin/categorylist/p/" + strconv.Itoa(pageId)
			} else {
				resp.RespMessage(helper.RS_failed, helper.WARING)
			}
		}
	} else {
		resp.RespMessage(helper.RS_failed, helper.WARING)
		c.Render()
	}
}

func addCategory(userName interface{}, name string) (success bool, message string) {
	success, message = models.AddBlogCategory(userName, name)
	return
}

func deleteCategory(userName interface{}, name string) (success bool, message string) {
	success, message = models.DeleteCategory(userName, name)
	return
}

func alterCategory(userName interface{}, id string, name string) (success bool, message string) {
	success, message = models.AlterCategory(userName, id, name)
	return
}
