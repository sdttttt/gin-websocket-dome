package router

import (
	"gin-web/controller/login"
	"gin-web/controller/register"
	"gin-web/middleware"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

type Router struct {
	Gin *gin.Engine
}

/**
 * @description rigister All Controller
 */
func (this *Router) RigisterController() {

	/**
	* Resource List
	 */
	this.Gin.Static("static", "./static")

	this.Gin.POST(login.LoginHandlerUrl, login.LoginHandler)

	this.Gin.POST(register.RegisterHandlerUrl, register.RegisterHandler)

}

/**
 * @description rigister All Middleware
 */
func (this *Router) RegisterMiddleware() {

	store := cookie.NewStore([]byte("secret"))

	this.Gin.Use(sessions.Sessions("mysession", store))
	this.Gin.Use(middleware.OriginAcceptMiddleware)

}
