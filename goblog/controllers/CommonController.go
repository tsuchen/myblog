package controllers

import (
	"fmt"
	"myblog/goblog/helper"
	"myblog/goblog/models"
	"strconv"

	"github.com/astaxie/beego"
)

var sessionName string = beego.AppConfig.String("SessionName")
var defaultAdmin string = "xuchen"
var defaultDomain string = "http://localhost:8080"

type CommonController struct {
	beego.Controller
}

func (c *CommonController) checkUserStatus() (hasLogin bool, session interface{}) {
	cookie, _ := c.Ctx.Request.Cookie(sessionName)
	se := c.GetSession(cookie.Name)
	if se == nil {
		fmt.Println("session不存在, 请先登录。")
		c.TplName = "login.html"
		c.Data["UserName"] = defaultAdmin
		c.Data["URL"] = defaultDomain
		hasLogin = false
	} else {
		userInfo := helper.GlobalUserManager.GetUserInfo(se)
		if userInfo == nil {
			c.Data["UserName"] = defaultAdmin
		} else {
			c.Data["UserName"] = userInfo.UserName
		}
		//添加url
		var categoryInfoList []*models.CategoryInfo
		list := models.GetAllCategory(se)
		for _, obj := range list {
			id := strconv.Itoa(obj.ID)
			url := "/admin/blogs/" + id
			info := &models.CategoryInfo{ID: obj.ID, Name: obj.Name, URL: url}
			categoryInfoList = append(categoryInfoList, info)
		}
		c.Data["CategoryInfos"] = categoryInfoList
		session = se
		hasLogin = true
	}

	return
}
