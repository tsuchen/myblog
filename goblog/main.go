package main

import (
	"myblog/goblog/helper"
	_ "myblog/goblog/routers"

	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/session/mysql"
)

func init() {
	helper.NewUserManager()
}

func main() {
	beego.Run()
}
