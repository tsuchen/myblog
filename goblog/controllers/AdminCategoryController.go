package controllers

import (
	"myblog/goblog/models"
	"strconv"
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
			url := "/admin/category/" + strconv.Itoa(obj.ID)
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
