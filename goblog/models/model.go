package models

import (
	"fmt"
	"math"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

var CountOfOnePage float64 = 10

type CategoryInfo struct {
	ID   int
	Name string
	URL  string
}

type UserProfile struct {
	NickName    string
	Sex         string
	PhoneNumber string
	Email       string
	Desc        string
	Birth       time.Time
}

func init() {
	//读取配置
	sqlDB := beego.AppConfig.String("mysqldb")
	sqlUser := beego.AppConfig.String("mysqluser")
	sqlPassword := beego.AppConfig.String("mysqlpass")
	sqlURL := beego.AppConfig.String("mysqlurl")
	sqlInfo := sqlUser + ":" + sqlPassword + "@" + sqlURL + "/" + sqlDB + "?charset=utf8&loc=Local"

	//设置默认数据库
	orm.RegisterDataBase("default", "mysql", sqlInfo, 30)
	//注册自定义model
	orm.RegisterModel(new(User), new(Profile), new(Blog), new(Tag), new(Category))
	// 自动建表
	orm.RunSyncdb("default", false, true)
	// 设置为 UTC 时间
	orm.DefaultTimeLoc = time.UTC
	// 开启 orm 调试模式：开发过程中建议打开，release时需要关闭
	orm.Debug = true
}

func NewUser() {
	o := orm.NewOrm()
	profile := &Profile{}
	profile.NickName = "admin"
	profile.Sex = "man"
	if _, err := o.Insert(profile); err == nil {
		user := User{Name: "admin", Password: "123456", Profile: profile}
		_, err = o.Insert(&user)
		if err != nil {
			fmt.Println("创建用户失败：", err)
		}
	}
}

func SelectUser(userName string, password string) (isFind bool, user User) {
	o := orm.NewOrm()
	err := o.QueryTable("user").Filter("name", userName).Filter("password", password).One(&user)
	isFind = (err == nil)

	return
}

func GetUserByName(userName interface{}) (user User) {
	o := orm.NewOrm()
	err := o.QueryTable("user").Filter("name", userName).One(&user)
	if err == nil {
		if user.Profile != nil {
			o.Read(user.Profile)
		}
	}

	return
}

func UpdateUserProfile(userName interface{}, info UserProfile) bool {
	success := false
	o := orm.NewOrm()
	var profile Profile
	err := o.QueryTable("profile").Filter("User__Name", userName).One(&profile)
	if err == nil {
		profile.Birth = info.Birth
		profile.NickName = info.NickName
		profile.Sex = info.Sex
		profile.PNumber = info.PhoneNumber
		profile.Introduce = info.Desc
		profile.Email = info.Email
		if _, error := o.Update(&profile); error == nil {
			success = true
		}
	}

	var user User
	err = o.QueryTable("user").Filter("Name", userName).One(&user)
	if err == nil {
		now := time.Now()
		user.Updated = now
		o.Update(&user)
	}

	return success
}

func UpdatePassword(userName interface{}, oldPass string, newPass string) bool {
	success := false
	o := orm.NewOrm()
	var user User
	err := o.QueryTable("user").Filter("Name", userName).One(&user)
	if err == nil {
		pass := user.Password
		if pass == oldPass && pass != newPass {
			user.Password = newPass
			if _, error := o.Update(&user); error == nil {
				success = true
			}
		}
	}

	return success
}

func GetUsersByPageId(pageIndex int) (cur float64, total float64, userList []*User) {
	o := orm.NewOrm()
	qs := o.QueryTable("user")
	if num, err := qs.All(&userList); err == nil {
		if num == 0 {
			total = 0
			cur = 1
			return
		}
		total = math.Ceil(float64(num) / CountOfOnePage)
		var start float64
		if float64(pageIndex) > total {
			cur = 1
			start = 0
		} else {
			start = (float64(pageIndex) - 1) * CountOfOnePage
		}
		qs.Limit(int64(start), int64(CountOfOnePage)).All(&userList)
	}

	return
}

//获取用户所有的分类
func GetAllCategory(userName interface{}) (categoryList []*Category) {
	o := orm.NewOrm()
	o.QueryTable("category").Filter("Users__User__Name", userName).All(&categoryList)

	return
}

func GetCategoryNameById(id int) (name string) {
	o := orm.NewOrm()
	var category Category
	err := o.QueryTable("category").Filter("ID", id).One(&category)
	if err == nil {
		name = category.Name
	} else {
		name = "所有博客"
	}

	return
}

func existCategory(userName interface{}, categoryName string) (exist bool) {
	list := GetAllCategory(userName)
	for _, obj := range list {
		if obj.Name == categoryName {
			exist = true
			break
		}
	}
	exist = false

	return
}

//添加博客分类
func AddBlogCategory(userName interface{}, categoryName string) (success bool, message string) {
	if exist := existCategory(userName, categoryName); exist {
		success = false
		message = "已存在此分类，添加失败"
		return
	}

	var user User
	o := orm.NewOrm()
	o.QueryTable("user").Filter("Name", userName).One(&user)
	m2m := o.QueryM2M(&user, "Categorys")
	category := Category{Name: categoryName}
	o.Insert(&category)
	if _, err := m2m.Add(&category); err == nil {
		success = true
		message = "添加分类成功。"
	} else {
		success = false
		message = "添加分类失败。"
	}

	return
}

//删除博客分类
func DeleteCategory(userName interface{}, categoryName string) (success bool, message string) {
	var user User
	o := orm.NewOrm()
	o.QueryTable("user").Filter("Name", userName).One(&user)
	m2m := o.QueryM2M(&user, "Categorys")
	var category Category
	o.QueryTable("category").Filter("Name", categoryName).One(&category)

	if _, err := m2m.Remove(category); err == nil {
		success = true
		message = "删除分类成功。"
	} else {
		success = false
		message = "删除分类失败。"
	}

	return
}

//修改博客分类
func AlterCategory(userName interface{}, categoryId string, categoryName string) (success bool, message string) {
	id, _ := strconv.Atoi(categoryId)
	o := orm.NewOrm()
	var category Category
	o.QueryTable("category").Filter("ID", id).One(&category)
	category.Name = categoryName
	if _, err := o.Update(&category); err != nil {
		success = false
		message = "修改分类失败。"
	} else {
		success = true
		message = "修改分类成功"
	}

	return
}

//获取所有标签
func GetAllTags(userName interface{}) (tagList []*Tag) {
	o := orm.NewOrm()
	o.QueryTable("tag").Filter("Users__User__Name", userName).All(&tagList)

	return
}

func existTag(userName interface{}, tagName string) (exist bool) {
	list := GetAllTags(userName)
	for _, obj := range list {
		if obj.Name == tagName {
			exist = true
			break
		}
	}
	exist = false

	return
}

//添加标签
func AddTag(userName interface{}, tagName string) (success bool, message string) {
	if exist := existTag(userName, tagName); exist {
		success = false
		message = "已存在此标签，添加失败"
		return
	}

	var user User
	o := orm.NewOrm()
	o.QueryTable("user").Filter("Name", userName).One(&user)
	m2m := o.QueryM2M(&user, "Tags")
	tag := Tag{Name: tagName}
	o.Insert(&tag)
	if _, err := m2m.Add(&tag); err == nil {
		success = true
		message = "添加标签成功。"
	} else {
		success = false
		message = "添加分类失败。"
	}

	return
}

func DeleteTag(userName interface{}, tagName string) (success bool, message string) {
	var user User
	o := orm.NewOrm()
	o.QueryTable("user").Filter("Name", userName).One(&user)
	m2m := o.QueryM2M(&user, "Tags")
	var tag Tag
	o.QueryTable("tag").Filter("Name", tagName).One(&tag)

	if _, err := m2m.Remove(tag); err == nil {
		success = true
		message = "删除标签成功。"
	} else {
		success = false
		message = "删除标签失败。"
	}

	return
}

func AlterTag(userName interface{}, tagId string, tagName string) (success bool, message string) {
	id, _ := strconv.Atoi(tagId)

	o := orm.NewOrm()
	var tag Tag
	o.QueryTable("tag").Filter("ID", id).One(&tag)
	tag.Name = tagName
	if _, err := o.Update(&tag); err != nil {
		success = false
		message = "修改标签失败。"
	} else {
		success = true
		message = "修改标签成功"
	}

	return
}

func getAllBlogs(userName interface{}) (blogs []*Blog) {
	o := orm.NewOrm()

	var user User
	err := o.QueryTable("user").Filter("Name", userName).One(&user)
	if err == nil {

	}

	return
}

func GetBlogsByCategoryId(userName interface{}, categoryId int) (blogs []*Blog) {
	o := orm.NewOrm()
	var user User
	err := o.QueryTable("user").Filter("Name", userName).One(&user)
	if err == nil {
		if categoryId == 0 {
			o.QueryTable("blog").Filter("User", user.ID).RelatedSel().All(&blogs)
		} else {
			o.QueryTable("blog").Filter("User", user.ID).Filter("Category", categoryId).All(&blogs)
		}
	}

	return
}
