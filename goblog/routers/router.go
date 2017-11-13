package routers

import (
	"myblog/goblog/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/homepage", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
}