package models

import (
	"fmt"

	"github.com/astaxie/beego"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	//读取配置
	sqlDB := beego.AppConfig.String("mysqldb")
	sqlUser := beego.AppConfig.String("mysqluser")
	sqlPassword := beego.AppConfig.String("mysqlpass")
	sqlURL := beego.AppConfig.String("mysqlurl")
	sqlInfo := sqlUser + ":" + sqlPassword + "@" + sqlURL + "/" + sqlDB + "?charset=utf8"
	fmt.Println(sqlInfo)

	//设置默认数据库
	orm.RegisterDataBase("default", "mysql", sqlInfo, 30)
	//注册自定义model
	orm.RegisterModel(new(User), new(Profile), new(Blog), new(Tag), new(Category))
	// 自动建表
	orm.RunSyncdb("default", false, true)
	// 开启 orm 调试模式：开发过程中建议打开，release时需要关闭
	orm.Debug = true
}

func NewUser() {
	o := orm.NewOrm()
	//创建用户基本信息
	profile := &Profile{}
	profile.Age = 25
	profile.PNumber = "13618076042"
	profile.Sex = "男"
	profile.Introduce = "这是我的博客，欢迎来访！"
	_, err := o.Insert(profile)
	if err != nil {
		fmt.Println("创建用户详情失败：", err)
	}

	user := User{Name: "admin", Password: "123456", Profile: profile}
	_, err = o.Insert(&user)
	if err != nil {
		fmt.Println("创建用户失败：", err)
	}
}

func SelectUser(userName string, password string) (isFind bool, u User) {
	o := orm.NewOrm()
	var user User
	qs := o.QueryTable("user")
	qs = qs.Filter("name", userName).Filter("password", password)
	err := qs.One(&user)
	if err == nil {
		isFind = true
	} else {
		isFind = false
	}
	u = user

	return
}

func GetAllUser() (userList []*User) {
	o := orm.NewOrm()
	num, err := o.QueryTable("user").All(&userList)
	fmt.Printf("Returned Rows Num: %d, %s", num, err)
	for _, user := range userList {
		if user.Profile != nil {
			o.Read(user.Profile)
		}
	}

	return
}

//获取用户所有的分类
func GetAllCategory(userName interface{}) (categoryList []*Category) {
	o := orm.NewOrm()
	o.QueryTable("category").Filter("Users__User__Name", userName).All(&categoryList)

	return
}
