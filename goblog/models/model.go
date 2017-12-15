package models

import (
	"fmt"

	"github.com/astaxie/beego"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"

	"strconv"
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


//添加博客分类
func AddBlogCategory(userName interface{}, categoryName string) (success bool, message string) {
	fmt.Println("CategoryName: ",categoryName)
	o := orm.NewOrm()
	var categoryList []*Category
	_, err := o.QueryTable("category").Filter("Users__User__Name", userName).All(&categoryList)
	if err != nil {
		success = false
		message = "查询分类失败。"
		return 
	}

	//查询分类名称
	isFind := false
	for _, obj := range categoryList {
		if obj.Name == categoryName {
			isFind = true
			break
		}
	}

	if isFind {
		success = false
		message = "添加分类失败，重复添加。"
		return
	}

	var user User
	o.QueryTable("user").Filter("Name", userName).One(&user)
	m2m := o.QueryM2M(&user, "Categorys")
	category := &Category{Name: categoryName}
	o.Insert(category)
	_, err = m2m.Add(category)
	if err != nil {
		success = false
		message = "添加分类失败。"
		return
	}

	success = true
	message = "添加分类成功。"

	return 
}

//删除博客分类
func DeleteCategory(userName interface{}, categoryName string) (success bool, message string) {
	o := orm.NewOrm()
	var categoryList []*Category
	_, err := o.QueryTable("category").Filter("Users__User__Name", userName).All(&categoryList)
	if err != nil {
		success = false
		message = "查询分类失败。"
		return 
	}

	//查询分类名称
	isFind := false
	for _, obj := range categoryList {
		if obj.Name == categoryName {
			isFind = true
			break
		}
	}

	if !isFind {
		success = false
		message = "删除分类失败，不存在此分类。"
		return
	}

	var user User
	o.QueryTable("user").Filter("Name", userName).One(&user)
	m2m := o.QueryM2M(&user, "Categorys")
	var category Category
	o.QueryTable("category").Filter("Name", categoryName).One(&category)
	_, err = m2m.Remove(category)
	if err != nil {
		success = false
		message = "删除分类失败。"
		return
	}

	if !m2m.Exist(&category) {
		o.Delete(&category)
	}

	success = true
	message = "删除分类成功。"

	return 
}

//修改博客分类
func AlterCategory(userName interface{}, categoryId string, categoryName string) (success bool, message string) {
	id, _ := strconv.Atoi(categoryId)

	o := orm.NewOrm()
	var categoryList []*Category
	_, err := o.QueryTable("category").Filter("Users__User__Name", userName).All(&categoryList)
	if err != nil {
		success = false
		message = "查询分类失败。"
		return 
	}

	//查询分类
	isFind := false
	for _, obj := range categoryList {
		if obj.ID == id {
			isFind = true
		}else{
			if obj.Name == categoryName {
				success = false
				message = "修改分类失败, 已存在此分类名称。"
				return 
			}
		}
	}

	if !isFind {
		success = false
		message = "修改分类失败，不存在此分类。"
		return
	}

	var category Category
	err = o.QueryTable("category").Filter("ID", id).One(&category)
	category.Name = categoryName
	if _, err := o.Update(&category); err != nil {
		success = false
		message = "修改分类失败。"
	}else{
		success = true
		message = "修改分类成功"
	}

	return 
}