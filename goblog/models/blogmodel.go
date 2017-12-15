package models

type User struct {
	ID        int `orm:"auto"`
	Name      string
	Password  string
	Profile   *Profile    `orm:"rel(one)"`      //设置一对一关系
	Blogs     []*Blog     `orm:"reverse(many)"` // 设置一对多的反向关系
	Categorys []*Category `orm:"rel(m2m)"`
}

type Profile struct {
	ID        int `orm:"auto"`
	Age       int
	Introduce string
	Sex       string
	PNumber   string
	User      *User `orm:"reverse(one)"` // 设置一对一反向关系(可选)
}

type Blog struct {
	ID      int `orm:"auto"`
	Title   string
	Content string
	User    *User  `orm:"rel(fk)"` //设置一对多关系
	Tags    []*Tag `orm:"rel(m2m)"`
	Category *Category `orm:"rel(fk)"` 
}

type Tag struct {
	ID    int `orm:"auto"`
	Name  string
	Blogs []*Blog `orm:"reverse(many)"`
}

type Category struct {
	ID    int `orm:"auto"`
	Name  string
	Users []*User `orm:"reverse(many)"`
	Blogs []*Blog `orm:"reverse(many)"`
}
