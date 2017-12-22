package controllers

type AdminBlogController struct {
	CommonController
}

func (c *AdminBlogController) Get() {
	if isLogin, _ := c.checkUserStatus(); isLogin {
	}

	c.Render()
}
