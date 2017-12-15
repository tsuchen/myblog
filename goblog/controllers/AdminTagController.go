package controllers

import(

)

type AdminTagController struct {
	CommonController
}

func (c *AdminTagController) Get() {
	isLogin, se := c.checkUserStatus()
	if isLogin && se != nil {
		
	}

	c.Render()
}