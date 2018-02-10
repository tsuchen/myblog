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
	beego.Router("/admin/taglist/p/:page([0-9]+)", &controllers.AdminTagController{})
	beego.Router("/admin/bloglist/cate/:cateid([0-9]+)/p/:page([0-9]+)", &controllers.AdminBlogController{})
	beego.Router("/admin/editblog/blog/:blogid([0-9])", &controllers.AdminEditBlogController{})
}
