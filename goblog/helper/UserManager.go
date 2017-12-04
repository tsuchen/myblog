/*
*用户信息管理类
*/

package helper

import(
	"time"
	"sync"
)

type UserInfo struct {
	UserId	int
	UserName string
	URL	string
	/* 资料补充 */
	Sex string // 1:男， 2：女
	PNumber int64  // 手机号
	Age	int
	Description string	
	/* 自动记录 */
	CreateTime time.Time // 创建时间
	LoginTime time.Time // 登录时间
	LoginIp string    // 登录ip
	LogoutTime time.Time // 登出时间 
}


var GlobalUserManager *UserManager

type UserManager struct {
	lock     sync.Mutex
	userInfo *UserInfo
}

func NewUserManager(){
	GlobalUserManager = &UserManager{userInfo: &UserInfo{UserId: 0}} 
}

func (uManager *UserManager) GetUserInfo() (info *UserInfo) {
	info = uManager.userInfo
	return 
}

func (uManager *UserManager) SetUserInfo(info *UserInfo) {
	uManager.lock.Lock()
	defer uManager.lock.Unlock()
	uManager.userInfo = info
}