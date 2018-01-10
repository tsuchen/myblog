package controllers

import (
	"fmt"
	"myblog/goblog/helper"
	"myblog/goblog/models"
	"strconv"

	"github.com/astaxie/beego"
)

var sessionName string = beego.AppConfig.String("SessionName")

type CategoryInfo struct {
	ID   int
	Name string
	URL  string
}

type CommonController struct {
	beego.Controller
}

func (c *CommonController) checkUserStatus() (hasLogin bool, session interface{}) {
	cookie, _ := c.Ctx.Request.Cookie(sessionName)
	se := c.GetSession(cookie.Name)
	if se == nil {
		fmt.Println("session不存在, 请先登录。")
		c.TplName = "login.html"
		c.Data["UserName"] = "xuchen"
		c.Data["URL"] = "http://localhost:8080"
		hasLogin = false
	} else {
		userInfo := helper.GlobalUserManager.GetUserInfo(se)
		if userInfo == nil {
			fmt.Println("用户信息不存在, 请重新登录。")
			c.TplName = "login.html"
			c.Data["UserName"] = "xuchen"
			c.Data["URL"] = "http://localhost:8080"
			hasLogin = false
		} else {
			c.Data["UserName"] = userInfo.UserName
			//添加url
			var categoryInfoList []*CategoryInfo
			list := models.GetAllCategory(se)
			for _, obj := range list {
				id := strconv.Itoa(obj.ID)
				url := "/admin/blogs/" + id
				info := &CategoryInfo{ID: obj.ID, Name: obj.Name, URL: url}
				categoryInfoList = append(categoryInfoList, info)
			}
			c.Data["CategoryInfos"] = categoryInfoList
			session = se
			hasLogin = true
		}
	}

	return
}
