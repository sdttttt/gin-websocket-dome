/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-08-31 00:49:27
 * @LastEditTime: 2019-08-31 17:48:10
 * @LastEditors: Please set LastEditors
 */
package login

import (
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

	context.JSON(http.StatusOK, gin.H{"success": "good!"})
}
