/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-08-31 00:49:27
 * @LastEditTime: 2019-09-03 17:14:20
 * @LastEditors: Please set LastEditors
 */
package login

import (
	"gin-web/dao"
	"gin-web/dao/service"
	"gin-web/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	LoginForm struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
)

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
	LoginViewUrl    = "/login"
	LoginHandlerUrl = "/login"
)

func LoginView(context *gin.Context) {
	context.HTML(http.StatusOK, "login.tmpl", gin.H{})
}

func LoginHandler(context *gin.Context) {
	var form LoginForm
	context.BindJSON(&form)

	if !form.check() {
		context.JSON(http.StatusOK, gin.H{
			"error": "您输入有误",
		})
	}

	println(form.Username)
	println(form.Password)

	user := &dao.User{Username: form.Username, Password: form.Password}

	if service.GetUserService().FindUser(user) {
		println(user.Username, user.Password)
		util.SetSession(context, "token", user.ID)
		context.JSON(http.StatusOK, gin.H{"success": "good!"})
	}

	context.JSON(http.StatusOK, gin.H{"error": "fuck"})
}
