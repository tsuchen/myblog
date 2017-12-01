package main

import (
	_ "myblog/goblog/routers"
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/session/mysql"
	"myblog/goblog/helper"
)

func init(){
	helper.NewUserManager()
}

func main() {
	beego.Run()
}

