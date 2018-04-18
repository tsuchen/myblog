package controllers

import (
	"myblog/goblog/helper"
)

type AdminUploadController struct {
	CommonController
}

func (c *AdminUploadController) Get() {

}

func (c *AdminUploadController) Post() {
	resp := helper.NewUploadResponse()
	defer resp.WriteUploadRespByJson(c.Ctx.ResponseWriter)

	isLogin, se := c.checkUserStatus()
	if isLogin && se != nil {
		f, h, err := c.GetFile("editormd-image-file")
		if err != nil {
			resp.RespUploadMessage(0, "上传图片失败", "")
		} else {
			url := "/static/upload/img/" + h.Filename
			path := defaultUploadPath + "img/" + h.Filename
			isExist, _ := helper.PathExists(path)
			if isExist {
				f.Close()
				resp.RespUploadMessage(1, "图片已存在", url)
				return
			}
			f.Close()
			c.SaveToFile("editormd-image-file", path)

			resp.RespUploadMessage(1, "上传图片成功", url)
		}
	} else {
		resp.RespUploadMessage(0, "用户未登录", "")
		c.Render()
	}
}
