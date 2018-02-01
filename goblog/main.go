package main

import (
	_ "myblog/goblog/routers"

	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/session/mysql"
)

func main() {
	beego.Run()
}
