package main

import (
	"myblog/goblog/helper"
	_ "myblog/goblog/routers"

	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/session/mysql"
)

func getRowIndex(i int) (out int) {
	out = i + 1
	return
}

func init() {
	helper.NewUserManager()
	beego.AddFuncMap("getRowIndex", getRowIndex)
}

func main() {
	beego.Run()
}
