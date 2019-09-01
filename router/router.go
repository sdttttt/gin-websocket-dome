/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-08-30 23:57:34
 * @LastEditTime: 2019-08-31 17:25:24
 * @LastEditors: Please set LastEditors
 */
package router

import (
	"gin-web/controller/home"
	"gin-web/controller/login"

	"github.com/gin-gonic/gin"
)

type Router struct {
	Gin *gin.Engine
}

func (router *Router) Rigister() {

	router.Gin.LoadHTMLGlob("template/*")
	router.Gin.Static("static", "./static")

	/**
	*
	* Router List
	 */
	router.Gin.GET(home.HelloUrl, home.Hello)
	router.Gin.GET(login.LoginViewUrl, login.LoginView)
	router.Gin.POST(login.LoginHandlerUrl, login.LoginHandler)

}
