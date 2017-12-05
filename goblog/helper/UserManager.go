/*
*用户信息管理类
*/

package helper

import(
	"time"
	"sync"
)

type UserInfo struct {
	UserId		int
	UserName 	string
	Password	string
	/* 资料补充 */
	Sex 		string // 1:男， 2：女
	PNumber		int64  // 手机号
	Age			int
	Description string	
	/* 自动记录 */
	CreateTime 	time.Time // 创建时间
	LoginTime	time.Time // 登录时间
	LogoutTime 	time.Time // 登出时间
	LoginIp 	string    // 登录ip
}

var maxUserCount int = 100

var GlobalUserManager *UserManager

type UserManager struct {
	lock     sync.Mutex
	userList []*UserInfo
}

func NewUserManager(){
	GlobalUserManager = &UserManager{userList: make([]*UserInfo, 10, maxUserCount)} 
}

func (uManager *UserManager) GetUserInfo(username string) (info *UserInfo) {
	for _, user := range uManager.userList{
		if user.UserName == username {
			info = user
		}
	}
	return 
}

func (uManager *UserManager) insertUserInfo(info *UserInfo){
	uManager.userList = append(uManager.userList, info)
}

func (uManager *UserManager) UpdateUserInfo(info *UserInfo) (success bool){
	uManager.lock.Lock()
	defer uManager.lock.Unlock()

	for _, user := range uManager.userList{
		if user.UserId == info.UserId {		
			user = info 
			success = true
		}
	}
	//没有找到,则插入新用户信息
	if success == false{
		uManager.insertUserInfo(info)
	}

	return 
}

func (uManager *UserManager) DeleteUserInfo(username string) (success bool){
	for index, user := range uManager.userList{
		if user.UserName == username {
			uManager.userList = append(uManager.userList[:index], uManager.userList[index+1:]...)
			success = true
		}
	}
	return 
}