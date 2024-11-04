package test

import (
	"net/http"
)

// @Summary 账号管理模块
// @Router /account
// @Tags Account
type Account struct {
}

// @Summary 注册用户
// @Produce json
// @Param tag_id body string false "标签ID"
// @Param title body string false "文章标题"
// @Param desc body string false "文章简述"
// @Param cover_image_url body string false "封面图片地址"
// @Param content body string false "文章内容"
// @Param modified_by body string true "修改者"
// @Success 200 {object} string "注册成功"
// @Failure 400 {object} string "请求错误"
// @Router /account/register [GET]
func (ctrl *Account) Register(w http.ResponseWriter, req *http.Request) {
	return
}

// 绑定微信OPENID
// @Router /account/bindwxuser [POST]
func (ctrl *Account) Bindwxuser(w http.ResponseWriter, req *http.Request) {
	return
}

// 修改基础资料
// @Router /account/updateprofile POST
func (ctrl *Account) Updateprofile(w http.ResponseWriter, req *http.Request) {
	return
}

// 删除用户信息
// @Router /account/delete DELETE
func (ctrl *Account) Delete(w http.ResponseWriter, req *http.Request) {
	return
}

// 使用用户名登录登录
// @Router /account/login [POST,GET]
func (ctrl *Account) Login(w http.ResponseWriter, req *http.Request) {
	return
}

// 自动登录
func (ctrl *Account) AutoLogin(w http.ResponseWriter, req *http.Request) {
	return
}
