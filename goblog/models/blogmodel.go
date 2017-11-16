package models

import (
    "github.com/astaxie/beego/orm"
)

type User struct{
	Uid int  `PK`
	Name string 
	Password string
	Profile *Profile
	Blog []*Blog	`orm:"reverse(many)"` // 设置一对多的反向关系
}

type Profile struct{
	Id int
	Age int
	Introduce string
	User *User	`orm:"reverse(one)"` // 设置一对一反向关系(可选)
}

type Blog struct{
	Id int 
	Title string
	Content string
	
}