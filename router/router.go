/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-08-30 23:57:34
 * @LastEditTime: 2019-09-05 17:25:20
 * @LastEditors: Please set LastEditors
 */
package router

import (
	"gin-web/controller/login"
	"gin-web/controller/register"
	"gin-web/middleware"
	"gin-web/socket"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

type Router struct {
	Gin *gin.Engine
}

/*
*****************************************************
* Warning => 这个Project已经改成了前后端分离
*				视图控制器应该应该已经不会再启动了
******************************************************
 */

/**
 * @description rigister All Controller
 */
func (this *Router) RigisterController() {

	/**
	* Resource List
	 */
	this.Gin.LoadHTMLGlob("template/*")
	this.Gin.Static("static", "./static")

	/**
	*
	* Router List
	 */
	//this.Gin.GET(home.HelloUrl, home.Hello)

	//this.Gin.GET(login.LoginViewUrl, login.LoginView)
	this.Gin.POST(login.LoginHandlerUrl, login.LoginHandler)

	//this.Gin.GET(register.RegisterViewUrl, register.RegisterView)
	this.Gin.POST(register.RegisterHandlerUrl, register.RegisterHandler)

	this.Gin.GET(socket.GinEchoUrl, socket.GinEcho)

}

/**
 * @description rigister All Middleware
 */
func (this *Router) RegisterMiddleware() {

	store := cookie.NewStore([]byte("secret"))
	this.Gin.Use(sessions.Sessions("mysession", store))
	this.Gin.Use(middleware.MetmodAuthMiddleware)

}
