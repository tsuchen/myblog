package models

type User struct {
	ID       int	  `orm:"auto"`
	Name     string
	Password string
	Profile  *Profile `orm:"rel(one)"`      //设置一对一关系
	Blog     []*Blog  `orm:"reverse(many)"` // 设置一对多的反向关系
	Category []*Category `orm:"rel(m2m)"`
}

type Profile struct {
	ID        int	`orm:"auto"`
	Age       int
	Introduce string
	Sex       string
	PNumber	  string
	User      *User `orm:"reverse(one)"` // 设置一对一反向关系(可选)
}

type Blog struct {
	ID      int		`orm:"auto"`
	Title   string
	Content string
	User    *User  `orm:"rel(fk)"` //设置一对多关系
	Tag     []*Tag `orm:"rel(m2m)"`
}

type Tag struct {
	ID   int	`orm:"auto"`
	Name string
	Blog []*Blog `orm:"reverse(many)"`
}

type Category struct {
	ID		int		`orm:"auto"`
	Name	string 
	User	[]*User  `orm:"reverse(many)"`
}