package models

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"
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
	URL    string
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
	orm.RegisterModel(new(User), new(Profile), new(Blog), new(Tag), new(Category), new(TempBlog))
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

func getPageIndexList(curPageIndex int, totalPages float64, url string) (pageIndexList []*PageIndexInfo) {
	pageIndexList = append(pageIndexList, &PageIndexInfo{curPageIndex, "", true})
	startIndex, endIndex := curPageIndex, curPageIndex
	for {
		if startIndex <= 1 && endIndex >= int(totalPages) {
			break
		}

		if startIndex > 1 {
			startIndex--
			urlStr := fmt.Sprintf(url, startIndex)
			pageIndexList = append(pageIndexList, &PageIndexInfo{startIndex, urlStr, false})
		}

		if endIndex < int(totalPages) {
			endIndex++
			urlStr := fmt.Sprintf(url, startIndex)
			pageIndexList = append(pageIndexList, &PageIndexInfo{endIndex, urlStr, false})
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
	if num, err := qs.Count(); err == nil {
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

		url := "admin/userlist/p/%d"
		pageIndexList = getPageIndexList(pageIndex, totalPage, url)

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

func getCategoryByName(userName interface{}, cateName string) interface{} {
	o := orm.NewOrm()
	var cate Category
	err := o.QueryTable("category").Filter("Users__User__Name", userName).Filter("Name", cateName).One(&cate)
	if err != nil {
		return nil
	}

	return &cate
}

func GetCategoryByPageId(userName interface{}, pageIndex int) (totalPage float64, indexList []*PageIndexInfo, cats []*Category) {
	o := orm.NewOrm()
	qs := o.QueryTable("category").Filter("Users__User__Name", userName)
	if num, err := qs.Count(); err == nil {
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

		url := "admin/categorylist/p/%d"
		indexList = getPageIndexList(pageIndex, totalPage, url)

		offset := (pageIndex - 1) * int(CountOfOnePage)
		qs.Limit(int64(CountOfOnePage), int64(offset)).All(&cats)
	}

	return
}

func GetCategoryNameById(cateID int) (cateName string) {
	o := orm.NewOrm()
	var cate Category
	err := o.QueryTable("category").Filter("ID", cateID).One(&cate)
	if err == nil {
		cateName = cate.Name
	} else {
		cateName = "所有博客"
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

	if category.ID == 1 {
		fmt.Println("不能删除此分类")
		return false
	}

	var user User
	err = o.QueryTable("user").Filter("Name", userName).One(&user)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	m2m := o.QueryM2M(&user, "Categorys")
	if _, err = m2m.Remove(category); err != nil {
		fmt.Println(err.Error())
		return false
	}

	var bloglist []*Blog
	_, err = o.QueryTable("blog").Filter("Category", category.ID).RelatedSel().All(&bloglist)
	if err == nil {
		for _, blog := range bloglist {
			if blog.Category != nil {
				//默认分类
				blog.Category.ID = 1
				o.Update(blog)
			}
		}
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

func getTagByName(userName interface{}, tagName string) interface{} {
	o := orm.NewOrm()
	var tag Tag
	err := o.QueryTable("tag").Filter("Users__User__Name", userName).Filter("Name", tagName).One(&tag)
	if err != nil {
		return nil
	}

	return &tag
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

		url := "admin/taglist/p/%d"
		indexList = getPageIndexList(pageIndex, totalPage, url)

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

	var blogs []*Blog
	_, err = o.QueryTable("blog").Filter("Tags__Tag__Name", tagName).All(&blogs)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	for _, blog := range blogs {
		m2m = o.QueryM2M(blog, "Tags")
		m2m.Remove(&tag)
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

func GetBlogs(userName interface{}, cateID int, pageID int) (totalPages float64, indexList []*PageIndexInfo, blogs []*Blog) {
	totalPages = 1

	o := orm.NewOrm()
	var user User
	err := o.QueryTable("user").Filter("Name", userName).One(&user)
	if err != nil {
		println(err.Error())
		return
	}

	var num int64 //文章数量
	if cateID == 1 {
		//所有文章
		num, err = o.QueryTable("blog").Filter("User", user.ID).Count()
	} else {
		num, err = o.QueryTable("blog").Filter("User", user.ID).Filter("Category", cateID).Count()
	}

	if err != nil {
		println(err.Error())
		return
	}

	if num != 0 {
		totalPages = math.Ceil(float64(num) / CountOfOnePage)
	}

	if float64(pageID) > totalPages {
		pageID = int(totalPages)
	} else if float64(pageID) < 1 {
		pageID = 1
	}

	url := "admin/bloglist/cate/" + strconv.Itoa(cateID) + "/p/%d"
	indexList = getPageIndexList(pageID, totalPages, url)

	offset := (pageID - 1) * int(CountOfOnePage)

	if cateID == 1 {
		//所有文章
		qs := o.QueryTable("blog").Filter("User", user.ID)
		qs.All(&blogs)
	} else {
		qs := o.QueryTable("blog").Filter("User", user.ID).Filter("Category", cateID)
		qs.Limit(int64(CountOfOnePage), int64(offset)).All(&blogs)
	}

	return
}

func GetTempArticleByID(id int) interface{} {
	o := orm.NewOrm()
	var tempBlog TempBlog
	err := o.QueryTable("temp_blog").Filter("Blog__ID", id).One(&tempBlog)
	if err == nil {
		return tempBlog
	}

	return nil
}

func GetArticleByID(id int) interface{} {
	o := orm.NewOrm()
	var blog Blog
	err := o.QueryTable("blog").Filter("ID", id).One(&blog)
	if err != nil {
		return nil
	}

	if blog.Category != nil {
		o.Read(blog.Category)
	}
	_, err = o.QueryTable("tag").Filter("Blogs__Blog__ID", id).All(&blog.Tags)
	if err != nil {
		fmt.Println(err.Error())
	}

	return blog
}

//暂存博客
func SaveArticle(userName interface{}, args map[string]string) bool {
	o := orm.NewOrm()
	var user User
	err := o.QueryTable("user").Filter("Name", userName).One(&user)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	blogID, _ := strconv.Atoi(args["blogid"])
	title := args["title"]
	content := args["content"]
	cateName := args["category"]
	tags := args["tags"]

	var tempBlog TempBlog
	if blogID != 0 {
		if err = o.QueryTable("temp_blog").Filter("Blog__ID", blogID).One(&tempBlog); err == nil {
			tempBlog.Title = title
			tempBlog.Content = content
			tempBlog.Category = cateName
			tempBlog.Tags = tags
			if _, err = o.Update(&tempBlog); err != nil {
				fmt.Println(err.Error())
				return false
			}
		} else {
			var blog Blog
			err = o.QueryTable("blog").Filter("ID", blogID).One(&blog)
			if err != nil {
				fmt.Println(err.Error())
				return false
			}
			tempBlog.Blog = &blog
			tempBlog.User = &user
			tempBlog.Title = title
			tempBlog.Content = content
			tempBlog.Category = cateName
			tempBlog.Tags = tags
			if _, err = o.Insert(&tempBlog); err != nil {
				fmt.Println(err.Error())
				return false
			}
		}
	}

	return true
}

//发表博客
func SendArticleByID(userName interface{}, args map[string]string) bool {
	o := orm.NewOrm()
	var user User
	err := o.QueryTable("user").Filter("Name", userName).One(&user)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	blogID, _ := strconv.Atoi(args["blogid"])
	title := args["title"]
	content := args["content"]
	cateName := args["category"]
	AddBlogCategory(userName, cateName)
	var tagStrs []string
	tagStrs = strings.Split(args["tags"], ";")
	for _, tagName := range tagStrs {
		AddTag(userName, tagName)
	}
	//tag
	var tags []*Tag
	for _, tagStr := range tagStrs {
		tag := getTagByName(userName, tagStr)
		if reflect.TypeOf(tag) != nil {
			tags = append(tags, tag.(*Tag))
		}
	}

	var blog Blog

	if blogID == 0 {
		//新文章
		blog.User = &user
		blog.Title = title
		blog.Content = content
		//category
		cate := getCategoryByName(userName, cateName)
		if reflect.TypeOf(cate) != nil {
			blog.Category = cate.(*Category)
		}
		if _, err = o.Insert(&blog); err != nil {
			//插入新文章失败
			fmt.Println(err.Error())
			return false
		}
		//tag
		m2m := o.QueryM2M(&blog, "Tags")
		// m2m.Clear()
		for _, tag := range tags {
			//添加m2m关系失败
			if _, err = m2m.Add(tag); err != nil {
				//添加标签失败
				fmt.Println(err.Error())
			}
		}
	} else {
		b := GetArticleByID(blogID)
		if reflect.TypeOf(b) == nil {
			return false
		}
		blog = b.(Blog)
		blog.Title = title
		blog.Content = content
		//category
		cate := getCategoryByName(userName, cateName)
		if reflect.TypeOf(cate) != nil {
			blog.Category = cate.(*Category)
		}
		if _, err = o.Update(&blog); err != nil {
			fmt.Println(err.Error())
			return false
		}
		m2m := o.QueryM2M(&blog, "Tags")
		m2m.Clear()
		for _, tag := range tags {
			//添加m2m关系失败
			if _, err = m2m.Add(tag); err != nil {
				//添加标签失败
				fmt.Println(err.Error())
			}
		}

		o.QueryTable("temp_blog").Filter("Blog__ID", blogID).Delete()
	}

	return true
}

func DeleteArticle(userName interface{}, blogID int) bool {
	o := orm.NewOrm()
	var user User
	err := o.QueryTable("user").Filter("Name", userName).One(&user)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	userID := user.ID
	var blog Blog
	qs := o.QueryTable("blog").Filter("ID", blogID).Filter("User", userID).RelatedSel()
	err = qs.One(&blog)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	m2m := o.QueryM2M(&blog, "Tags")
	_, err = m2m.Remove(blog)
	if err != nil {
		return false
	}

	if _, err = qs.Delete(); err == nil {
		o.QueryTable("temp_blog").Filter("Blog__ID", blogID).Delete()
	}

	return true
}
