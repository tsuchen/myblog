/**
author: tsuchen
date: 2017.11.07
*/

package helper

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	RS_failed       = -1 // 操作失败
	RS_success      = 1  // 操作成功
	RS_params_error = 2  // 参数错误

	RS_user_exist       = 100 // 账号已存在
	RS_user_inexistence = 101 // 账号不存在
	RS_activate_failed  = 102 // 激活失败
	RS_password_error   = 103 // 密码错误
	RS_register_failed  = 104 // 注册失败

	RS_query_failed  = 200 // 查询失败
	RS_update_failed = 201 // 更新失败
	RS_create_failed = 202 // 创建失败
	RS_delete_failed = 203 // 删除失败
	RS_notin_trash   = 204 // 文章不在垃圾箱
	RS_undo_falied   = 205 // 撤销删除失败

	RS_user_not_activate = 300 // 用户暂未激活
	RS_user_not_login    = 301 // 用户没有登录

	RS_tag_exist = 400 // tag已存在
)

var descDict = map[int]string{
	RS_failed:            "操作失败",
	RS_success:           "操作成功",
	RS_params_error:      "参数错误",
	RS_user_exist:        "账号已存在",
	RS_user_inexistence:  "账号不存在",
	RS_activate_failed:   "激活失败",
	RS_password_error:    "密码错误",
	RS_register_failed:   "注册失败",
	RS_query_failed:      "查询失败",
	RS_update_failed:     "更新失败",
	RS_create_failed:     "创建失败",
	RS_delete_failed:     "删除失败",
	RS_notin_trash:       "文章不在垃圾箱",
	RS_undo_falied:       "撤销删除失败",
	RS_user_not_activate: "用户暂未激活",
	RS_user_not_login:    "用户没有登录",
	RS_tag_exist:         "tag已存在",
}

func Desc(code int) string {
	desc, found := descDict[code]
	if !found {
		return "未定义状态"
	}
	return desc
}

const (
	SUCCESS = "success"
	WARING  = "waring"
	ALTER   = "alter"
	INFO    = "info"
)

type Tips struct {
	Level string
	Msg   string
}

type Response struct {
	Status int
	Data   interface{}
	Tip    Tips
}

func NewResponse() (resp *Response) {
	return &Response{Status: RS_success}
}

func (resp *Response) RespMessage(status int, level string) {
	resp.Status = status
	if status == RS_success {
		resp.Tip = Tips{Level: SUCCESS, Msg: "操作成功！"}
	} else {
		resp.Tip = Tips{Level: level, Msg: "code:" + fmt.Sprint(status) + "" + Desc(status)}
	}
}

func (resp *Response) WriteRespByJson(w http.ResponseWriter) {
	obj, err := json.Marshal(resp)
	if err != nil {
		w.Write([]byte(`{status:-1,Tip:Tips{level:"alter",Msg:"code:-1|序列化失败"}`))
	} else {
		w.Write(obj)
	}
}

func GetNavigationPathStr(list []string) (path string) {
	for _, s := range list {
		path += "-" + s
	}

	return
}
