/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-08-30 23:57:34
 * @LastEditTime: 2019-09-02 01:19:55
 * @LastEditors: Please set LastEditors
 */
package router

import (
	"gin-web/controller/home"
	"gin-web/controller/login"
	"gin-web/middleware"
	"gin-web/socket"

	"github.com/gin-gonic/gin"
)

type Router struct {
	Gin *gin.Engine
}

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
	this.Gin.GET(home.HelloUrl, home.Hello)
	this.Gin.GET(login.LoginViewUrl, login.LoginView)
	this.Gin.POST(login.LoginHandlerUrl, login.LoginHandler)
	this.Gin.GET(socket.GinEchoUrl, socket.GinEcho)

}

func (this *Router) RegisterMiddleware() {

	this.Gin.Use(middleware.AWebSocketCallFilter.MetmodAuthMiddleware)

}
