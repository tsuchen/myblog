package controllers

type CategoryInfo struct {
	ID   int
	Name string
	URL  string
}

type AdminController struct {
	CommonController
}

func (c *AdminController) Get() {
	isLogin, _ := c.checkUserStatus()
	if isLogin {
		c.TplName = "adminhome.html"
		// categoryList := models.GetAllCategory(se)
		// //添加url
		// var categoryInfoList []*CategoryInfo
		// for _, obj := range categoryList {
		// 	url := "/admin/category/" + strconv.Itoa(obj.ID)
		// 	info := &CategoryInfo{ID: obj.ID, Name: obj.Name, URL: url}
		// 	categoryInfoList = append(categoryInfoList, info)
		// }
		// users := models.GetAllUser()
		// c.Data["Users"] = users
		// c.Data["CategoryInfos"] = categoryInfoList
		// c.Data["GroupListId"] = "UserList"
		// c.Layout = "admin.html"
		// c.TplName = "userlist.html"
	}

	c.Render()
}
