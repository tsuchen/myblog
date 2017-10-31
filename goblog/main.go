package main

import (
	_ "myblog/goblog/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

