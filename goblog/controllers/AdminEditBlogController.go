package controllers

type AdminEditBlogController struct {
	CommonController
}

func (c *AdminEditBlogController) Get() {
	if isLogin, _ := c.checkUserStatus(); isLogin {
		c.Data["GroupMenuId"] = "editblog-menu"
		c.Layout = "adminhome.html"
		c.TplName = "editblog.html"
	}

	c.Render()
}

func (c *AdminEditBlogController) Post() {

}
