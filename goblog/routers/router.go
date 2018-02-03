package routers

import (
	"myblog/goblog/controllers"

	"github.com/astaxie/beego"
)

func init() {
	//后台路由
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/logout", &controllers.LogoutController{})
	beego.Router("/admin", &controllers.AdminController{})
	beego.Router("/admin/userlist/p/:page([0-9]+)", &controllers.AdminUserListController{})
	beego.Router("/admin/categorylist/p/:page([0-9]+)", &controllers.AdminCategoryController{})
	beego.Router("/admin/tag", &controllers.AdminTagController{})
	beego.Router("/admin/blogs/:cateid([0-9]+)", &controllers.AdminBlogController{})
	beego.Router("/admin/editblog", &controllers.AdminEditBlogController{})
}
