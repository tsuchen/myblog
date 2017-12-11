package controllers

import (
	"fmt"
	"myblog/goblog/helper"
	"myblog/goblog/models"
	"strconv"
	"github.com/astaxie/beego"
)

type AdminCategoryController struct {
	beego.Controller
}

func (c *AdminCategoryController) Get() {
	cookie, _ := c.Ctx.Request.Cookie(sessionName)
	se := c.GetSession(cookie.Name)
	if se == nil {
		c.Data["URL"] = "http://localhost:8080"
		c.Data["Name"] = "xuchen"
		c.TplName = "login.html"
	} else {
		//通过用户名获取分类
		userInfo := helper.GlobalUserManager.GetUserInfo(se)
		if userInfo == nil {
			fmt.Println("用户信息不存在, 请重新登录。")
			c.TplName = "login.html"
			c.Data["UserName"] = "xuchen"
			c.Data["URL"] = "http://localhost:8080"
		} else {
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
			c.Data["UserName"] = userInfo.UserName
			c.Data["GroupListId"] = "CategoryList"
			c.Layout = "admin.html"
			c.TplName = "categorylist.html"
		}
	}

	c.Render()
}
