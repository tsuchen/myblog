package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	//设置默认数据库
	orm.RegisterDataBase("default", "mysql", "root:smx10221102@/myblog?charset=utf8", 30)
	//注册自定义model
	orm.RegisterModel(new(User), new(Profile), new(Blog), new(Tag))
	// 自动建表
	orm.RunSyncdb("default", false, true)
	// 开启 orm 调试模式：开发过程中建议打开，release时需要关闭
	orm.Debug = true
}

func NewUser() {
	o := orm.NewOrm()
	profile := &Profile{}
	user := User{ID: 1, Name: "xuchen", Password: "smx10221102", Profile: profile}
	_, err := o.Insert(&user)
	if err != nil {
		fmt.Println("创建用户失败：", err)
	}

	profile.ID = 1
	profile.Age = 25
	profile.Introduce = "这是我的博客，欢迎来访！"
	_, err = o.Insert(profile)
	if err != nil {
		fmt.Println("创建用户详情：", err)
	}
}

func SelectUser(userName string, password string) (isFind bool) {
	o := orm.NewOrm()
	var user User
	qs := o.QueryTable("user")
	qs = qs.Filter("name", userName).Filter("password", password)
	err := qs.One(&user)
	fmt.Printf("Returned Rows Num: %s\n", err)
	if err == nil {
		isFind = true
	} else {
		isFind = false
	}

	return
}
