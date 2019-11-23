/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-08-31 00:49:46
 * @LastEditTime: 2019-09-12 19:10:52
 * @LastEditors: Please set LastEditors
 */
package register

import (
	"gin-web/dao"
	"gin-web/dao/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterForm struct {
	Username string `json:"username"`

	Password string `json:"password"`

	Repassword string `json:"repassword"`
}

func (form *RegisterForm) check() bool {
	if form.Username == "" || form.Password == "" || form.Repassword == "" {
		return false
	}
	if len(form.Username) == 0 || len(form.Password) == 0 || len(form.Repassword) == 0 {
		return false
	}

	return form.Password == form.Repassword
}

const (
	RegisterViewUrl = "/register"

	RegisterHandlerUrl = "/register"
)

/**
 * @description return Register View
 */
func RegisterView(context *gin.Context) {

	context.HTML(http.StatusOK, "register.tmpl", nil)

}

/**
 * @description Create User
	TODO:
	- [x] Albe work ?
*/
func RegisterHandler(context *gin.Context) {

	var form RegisterForm
	context.BindJSON(&form)

	if !form.check() {
		context.JSON(http.StatusOK, gin.H{
			"error": "输入信息有问题",
		})
		return
	}

	userInfo := &dao.User{Username: form.Username, Password: form.Repassword}

	if service.GetUserService().CreateUser(userInfo) {
		context.JSON(http.StatusOK, gin.H{
			"success": "good",
		})
		return
	}

	if userInfo.Username == "" {
		context.JSON(http.StatusOK, gin.H{
			"error": "用户名已经被那啥了",
		})
	} else {
		context.JSON(http.StatusOK, gin.H{
			"error": "不知道为啥，注册失败了",
		})
	}

}
