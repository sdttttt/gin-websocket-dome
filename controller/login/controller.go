package login

import (
	"gin-web/dao"
	"gin-web/dao/service"
	"gin-web/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

/*
	表单检查
*/
func (form *LoginForm) check() bool {

	if len(form.Username) == 0 || form.Username == "" {
		return false
	}

	if len(form.Password) == 0 || form.Password == "" {
		return false
	}

	return true
}

const (
	LoginHandlerUrl = "/login"
)

/*
	View Controller
*/
func LoginView(context *gin.Context) {
	context.HTML(http.StatusOK, "login.tmpl", gin.H{})
}

/*
	登陆逻辑处理
*/
func LoginHandler(context *gin.Context) {
	var form LoginForm
	context.BindJSON(&form)

	if !form.check() {
		context.JSON(http.StatusOK, gin.H{
			"error": "您输入有误",
		})
		return
	}

	println(form.Username)
	println(form.Password)

	user := &dao.User{Username: form.Username, Password: form.Password}

	if service.GetUserService().FindUser(user) {
		println(user.Username, user.Password)
		util.SetSession(context, "token", user.Username)
		context.JSON(http.StatusOK, gin.H{"success": "good!"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"error": "用户名或密码不正确"})
}
