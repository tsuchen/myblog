package controllers

import (
	"fmt"
	"myblog/goblog/helper"
	"myblog/goblog/models"
)

type AdminCategoryController struct {
	CommonController
}

func (c *AdminCategoryController) Get() {
	isLogin, se := c.checkUserStatus()
	if isLogin && se != nil {
		categoryList := models.GetAllCategory(se)
		//添加url
		var categoryInfoList []*CategoryInfo
		for _, obj := range categoryList {
			url := "/admin/category/" + obj.Name
			info := &CategoryInfo{ID: obj.ID, Name: obj.Name, URL: url}
			categoryInfoList = append(categoryInfoList, info)
		}
		c.Data["CategoryInfos"] = categoryInfoList
		c.Data["Categorys"] = categoryList
		c.Data["GroupListId"] = "CategoryList"
		c.Layout = "admin.html"
		c.TplName = "categorylist.html"
	}

	c.Render()
}

func (c *AdminCategoryController) Post() {
	resp := helper.NewResponse()
	defer resp.WriteRespByJson(c.Ctx.ResponseWriter)

	isLogin, se := c.checkUserStatus()
	if isLogin && se != nil {
		oper := c.GetString("Type")
		categoryName := c.GetString("CategoryName")
		categoryId := c.GetString("CategoryId")
		if oper == "add" {
			//添加分类
			success, message := addCategory(se, categoryName)
			fmt.Println(message)
			if success {
				resp.RespMessage(helper.RS_success, helper.SUCCESS)
				// categoryList := models.GetAllCategory(se)
				resp.Data = "/admin/category"
			} else {
				resp.RespMessage(helper.RS_failed, helper.WARING)
			}
		} else if oper == "delete" {
			//删除分类
			success, message := deleteCategory(se, categoryName)
			fmt.Println(message)
			if success {
				resp.RespMessage(helper.RS_success, helper.SUCCESS)
				resp.Data = "/admin/category"
			} else {
				resp.RespMessage(helper.RS_failed, helper.WARING)
			}
		} else {
			//修改分类
			success, message := alterCategory(se, categoryId, categoryName)
			fmt.Println(message)
			if success {
				resp.RespMessage(helper.RS_success, helper.SUCCESS)
				resp.Data = "/admin/category"
			} else {
				resp.RespMessage(helper.RS_failed, helper.WARING)
			}
		}

	} else {
		resp.RespMessage(helper.RS_failed, helper.WARING)
		c.Render()
	}
}

func addCategory(userName interface{}, categoryName string) (success bool, message string) {
	success, message = models.AddBlogCategory(userName, categoryName)
	return
}

func deleteCategory(userName interface{}, categoryName string) (success bool, message string) {
	success, message = models.DeleteCategory(userName, categoryName)
	return
}

func alterCategory(userName interface{}, categoryId string, categoryName string) (success bool, message string) {
	success, message = models.AlterCategory(userName, categoryId, categoryName)
	return
}
