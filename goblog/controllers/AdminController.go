package controllers

import "myblog/goblog/models"

type AdminController struct {
	CommonController
}

func (c *AdminController) Get() {
	isLogin, _ := c.checkUserStatus()
	if isLogin {
		// c.TplName = "adminhome.html"
		// categoryList := models.GetAllCategory(se)
		// //添加url
		// var categoryInfoList []*CategoryInfo
		// for _, obj := range categoryList {
		// 	url := "/admin/category/" + strconv.Itoa(obj.ID)
		// 	info := &CategoryInfo{ID: obj.ID, Name: obj.Name, URL: url}
		// 	categoryInfoList = append(categoryInfoList, info)
		// }
		pathList := []string{"主页", "用户列表"}
		users := models.GetAllUser()
		c.Data["PathList"] = pathList
		c.Data["Users"] = users
		c.Data["TitleName"] = "用户列表"
		// c.Data["CategoryInfos"] = categoryInfoList
		// c.Data["GroupListId"] = "UserList"
		c.Layout = "adminhome.html"
		c.TplName = "userprofile.html"
	}

	c.Render()
}
