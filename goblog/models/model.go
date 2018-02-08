package models

import (
	"fmt"
	"math"
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

type PageIndexInfo struct {
	Index  int
	Active bool
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

func NewUser(name string, pass string) (success bool) {
	o := orm.NewOrm()
	profile := &Profile{}
	profile.NickName = name
	if _, err := o.Insert(profile); err == nil {
		user := User{Name: name, Password: pass, Profile: profile}
		_, err = o.Insert(&user)
		success = true
		if err != nil {
			fmt.Println("创建新用户失败：", err)
			success = false
		}
	}

	return
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

func getPageIndexList(curPageIndex int, totalPages float64) (pageIndexList []*PageIndexInfo) {
	pageIndexList = append(pageIndexList, &PageIndexInfo{curPageIndex, true})
	startIndex, endIndex := curPageIndex, curPageIndex
	for {
		if startIndex <= 1 && endIndex >= int(totalPages) {
			break
		}

		if startIndex > 1 {
			startIndex--
			pageIndexList = append(pageIndexList, &PageIndexInfo{startIndex, false})
		}

		if endIndex < int(totalPages) {
			endIndex++
			pageIndexList = append(pageIndexList, &PageIndexInfo{endIndex, false})
		}

		//最多显示5个页码
		if len(pageIndexList) > 5 {
			break
		}
	}

	//按照升序排序
	for i, _ := range pageIndexList {
		isSwap := false
		for j := len(pageIndexList) - 1; j > i; j-- {
			if pageIndexList[j].Index < pageIndexList[j-1].Index {
				temp := pageIndexList[j]
				pageIndexList[j] = pageIndexList[j-1]
				pageIndexList[j-1] = temp
				isSwap = true
			}
		}
		if !isSwap {
			break
		}
	}

	return
}

func GetUsersByPageId(pageIndex int) (totalPage float64, pageIndexList []*PageIndexInfo, userList []*User) {
	o := orm.NewOrm()
	qs := o.QueryTable("user")
	if num, err := qs.All(&userList); err == nil {
		if num == 0 {
			totalPage = 1
		} else {
			totalPage = math.Ceil(float64(num) / CountOfOnePage)
		}

		if float64(pageIndex) > totalPage {
			pageIndex = int(totalPage)
		} else if float64(pageIndex) < 1 {
			pageIndex = 1
		}
		pageIndexList = getPageIndexList(pageIndex, totalPage)

		offset := (pageIndex - 1) * int(CountOfOnePage)
		qs.Limit(int64(CountOfOnePage), int64(offset)).All(&userList)
	}

	return
}

//获取用户所有的分类
func getAllCategory() (categoryList []*Category) {
	o := orm.NewOrm()
	o.QueryTable("category").All(&categoryList)

	return
}

func GetCategoryByPageId(userName interface{}, pageIndex int) (totalPage float64, indexList []*PageIndexInfo, cats []*Category) {
	o := orm.NewOrm()
	qs := o.QueryTable("category").Filter("Users__User__Name", userName)
	if num, err := qs.All(&cats); err == nil {
		if num == 0 {
			totalPage = 1
		} else {
			totalPage = math.Ceil(float64(num) / CountOfOnePage)
		}

		if float64(pageIndex) > totalPage {
			pageIndex = int(totalPage)
		} else if float64(pageIndex) < 1 {
			pageIndex = 1
		}

		indexList = getPageIndexList(pageIndex, totalPage)

		offset := (pageIndex - 1) * int(CountOfOnePage)
		qs.Limit(int64(CountOfOnePage), int64(offset)).All(&cats)
	}

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

func GetUserAllCategory(userName interface{}) (list []*Category) {
	o := orm.NewOrm()
	o.QueryTable("category").Filter("Users__User__Name", userName).All(&list)
	return
}

//添加博客分类
func AddBlogCategory(userName interface{}, categoryName string) bool {
	var cat Category
	o := orm.NewOrm()
	err := o.QueryTable("category").Filter("Name", categoryName).One(&cat)
	//如果数据库不存在此分类，则插入该分类
	if err != nil {
		println(err.Error())
		cat = Category{Name: categoryName}
		o.Insert(&cat)
	}

	var user User
	o.QueryTable("user").Filter("Name", userName).One(&user)
	m2m := o.QueryM2M(&user, "Categorys")
	//是否存在m2m关系
	if m2m.Exist(&cat) {
		println("存在m2m关系，添加分类失败")
		return false
	}
	//添加m2m关系失败
	if _, err = m2m.Add(&cat); err != nil {
		println(err.Error())
		return false
	}

	return true
}

//删除分类
func DeleteCategory(userName interface{}, categoryName string) bool {
	o := orm.NewOrm()
	var category Category
	err := o.QueryTable("category").Filter("Name", categoryName).One(&category)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	var user User
	o.QueryTable("user").Filter("Name", userName).One(&user)
	m2m := o.QueryM2M(&user, "Categorys")

	if _, err = m2m.Remove(category); err != nil {
		fmt.Println(err.Error())
		return false
	}

	relDeleteCategoryByName(categoryName)

	return true
}

//修改博客分类
func AlterCategory(userName interface{}, id int, catName string) bool {
	result := AddBlogCategory(userName, catName)
	relDeleteCategoryByID(userName, id)

	return result
}

func relDeleteCategoryByID(userName interface{}, categoryID int) {
	o := orm.NewOrm()
	var cate Category
	err := o.QueryTable("category").Filter("ID", categoryID).One(&cate)
	if err != nil {
		println(err.Error())
		return
	}

	var user User
	o.QueryTable("user").Filter("Name", userName).One(&user)
	m2m := o.QueryM2M(&user, "Categorys")
	if _, err = m2m.Remove(cate); err != nil {
		fmt.Println(err.Error())
		return
	}

	//检查是否存在m2m关系
	m2m = o.QueryM2M(&cate, "Users")
	if num, error := m2m.Count(); error == nil {
		fmt.Println("Total nums:", num)
		if num == 0 {
			o.QueryTable("category").Filter("ID", categoryID).Delete()
		}
	}
}

func relDeleteCategoryByName(cateName string) {
	o := orm.NewOrm()
	var cate Category
	err := o.QueryTable("category").Filter("Name", cateName).One(&cate)
	if err != nil {
		println(err.Error())
		return
	}
	//检查是否存在m2m关系
	m2m := o.QueryM2M(&cate, "Users")
	num, _ := m2m.Count()
	if num == 0 {
		o.QueryTable("category").Filter("Name", cateName).Delete()
	}
}

//获取所有标签
func GetAllTags(userName interface{}) (tagList []*Tag) {
	o := orm.NewOrm()
	o.QueryTable("tag").Filter("Users__User__Name", userName).All(&tagList)

	return
}

func GetTagByPageId(userName interface{}, pageIndex int) (totalPage float64, indexList []*PageIndexInfo, tags []*Tag) {
	o := orm.NewOrm()
	qs := o.QueryTable("tag").Filter("Users__User__Name", userName)
	if num, err := qs.All(&tags); err == nil {
		if num == 0 {
			totalPage = 1
		} else {
			totalPage = math.Ceil(float64(num) / CountOfOnePage)
		}

		if float64(pageIndex) > totalPage {
			pageIndex = int(totalPage)
		} else if float64(pageIndex) < 1 {
			pageIndex = 1
		}

		indexList = getPageIndexList(pageIndex, totalPage)

		offset := (pageIndex - 1) * int(CountOfOnePage)
		qs.Limit(int64(CountOfOnePage), int64(offset)).All(&tags)
	}

	return
}

//添加标签
func AddTag(userName interface{}, tagName string) bool {
	var tag Tag
	o := orm.NewOrm()
	err := o.QueryTable("tag").Filter("Name", tagName).One(&tag)
	//如果数据库不存在此分类，则插入该分类
	if err != nil {
		println(err.Error())
		tag = Tag{Name: tagName}
		o.Insert(&tag)
	}

	var user User
	o.QueryTable("user").Filter("Name", userName).One(&user)
	m2m := o.QueryM2M(&user, "Tags")
	//是否存在m2m关系
	if m2m.Exist(&tag) {
		println("存在m2m关系，添加标签失败")
		return false
	}
	//添加m2m关系失败
	if _, err = m2m.Add(&tag); err != nil {
		println(err.Error())
		return false
	}

	return true
}

func DeleteTag(userName interface{}, tagName string) bool {
	o := orm.NewOrm()
	var tag Tag
	err := o.QueryTable("tag").Filter("Name", tagName).One(&tag)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	var user User
	o.QueryTable("user").Filter("Name", userName).One(&user)
	m2m := o.QueryM2M(&user, "Tags")
	if _, err = m2m.Remove(tag); err != nil {
		fmt.Println(err.Error())
		return false
	}

	relDeleteTagByName(tagName)

	return true
}

func AlterTag(userName interface{}, tagID int, tagName string) bool {
	result := AddTag(userName, tagName)
	relDeleteTagByID(userName, tagID)

	return result
}

func relDeleteTagByName(tagName string) {
	o := orm.NewOrm()
	var tag Tag
	err := o.QueryTable("tag").Filter("Name", tagName).One(&tag)
	if err != nil {
		println(err.Error())
		return
	}
	//检查是否存在m2m关系
	m2m := o.QueryM2M(&tag, "Users")
	num, _ := m2m.Count()
	if num == 0 {
		o.QueryTable("tag").Filter("Name", tagName).Delete()
	}
}

func relDeleteTagByID(userName interface{}, tagID int) {
	o := orm.NewOrm()
	var tag Tag
	err := o.QueryTable("tag").Filter("ID", tagID).One(&tag)
	if err != nil {
		println(err.Error())
		return
	}

	var user User
	o.QueryTable("user").Filter("Name", userName).One(&user)
	m2m := o.QueryM2M(&user, "Tags")
	if _, err = m2m.Remove(tag); err != nil {
		fmt.Println(err.Error())
		return
	}

	//检查是否存在m2m关系
	m2m = o.QueryM2M(&tag, "Users")
	if num, error := m2m.Count(); error == nil {
		fmt.Println("Total nums:", num)
		if num == 0 {
			o.QueryTable("tag").Filter("ID", tagID).Delete()
		}
	}
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
