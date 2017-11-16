package models

import(
	// "database/sql"
	"github.com/astaxie/beego/orm"
	// "github.com/go-sql-driver/mysql"
)

func init(){
	//设置默认数据库
	orm.RegisterDataBase("default", "mysql", "root:123456@/myblog?charset=utf8", 30)
	// 开启 orm 调试模式：开发过程中建议打开，release时需要关闭
	orm.Debug = true
	// 自动建表
	orm.RunSyncdb("default", false, true)
}