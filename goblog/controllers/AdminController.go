package controllers

import (
	"fmt"
	"myblog/goblog/helper"
	"myblog/goblog/models"
	"strconv" 
	"github.com/astaxie/beego"
)

type CategoryInfo struct{
	ID int
	Name string
	URL string
}

type AdminController struct {
	beego.Controller
}

func (c *AdminController) Get() {
	cookie, _ := c.Ctx.Request.Cookie(sessionName)
	se := c.GetSession(cookie.Name)
	if se == nil {
		fmt.Println("session不存在, 请先登录。")
		c.TplName = "login.html"
		c.Data["UserName"] = "xuchen"
		c.Data["URL"] = "http://localhost:8080"
	} else {
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
			users := models.GetAllUser()
			c.Data["Users"] = users
			c.Data["UserName"] = userInfo.UserName
			c.Data["CategoryInfos"] = categoryInfoList
			c.Data["GroupListId"] = "UserList"
			c.Layout = "admin.html"
			c.TplName = "userlist.html"
		}
	}
	c.Render()
}
