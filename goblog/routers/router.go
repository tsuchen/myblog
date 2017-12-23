package routers

import (
	"myblog/goblog/controllers"

	"github.com/astaxie/beego"
)

func init() {
	//后台路由
	beego.Router("/admin", &controllers.AdminController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/logout", &controllers.LogoutController{})
	beego.Router("/admin/category", &controllers.AdminCategoryController{})
	beego.Router("/admin/tag", &controllers.AdminTagController{})
	beego.Router("/admin/blogs/:cateid([0-9]+)", &controllers.AdminBlogController{})
}
