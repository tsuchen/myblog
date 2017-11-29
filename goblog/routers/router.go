package routers

import (
	"myblog/goblog/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/admin", &controllers.AdminController{})
	beego.Router("/login", &controllers.LoginController{})
}