/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-09-01 22:44:09
 * @LastEditTime: 2019-09-02 14:47:28
 * @LastEditors: Please set LastEditors
 */
package middleware

import (
	"gin-web/controller/login"
	"gin-web/socket"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var AWebSocketCallFilter = &WebSocketCallFilter{}

/** @description 这里的代码附上我写Java时的习惯,
 没有约定的代码是我非常不喜欢的
如果对你造成了困扰，我真的非常抱歉
*/
type (
	AbstractWebSocketCallFilter interface {
		MetmodAuthMiddleware(*gin.Context)
	}

	WebSocketCallFilter struct{}
)

/** @description 有些用户会直接访问 Provider of Websocket，
会在我们的程序中留下错误，这个Middleware就是为了解决这个问题
*/
func (filter *WebSocketCallFilter) MetmodAuthMiddleware(c *gin.Context) {

	if c.Request.URL.String() == socket.GinEchoUrl {
		log.Println("MetmodAuthMiddleware => Someone coming to Websocket service of us, happy!! ")
		if "websocket" == c.Request.Header.Get("Upgrade") {
			c.Next()
		}
		log.Println("MetmodAuthMiddleware => 原来你个憨憨，我送你一程。 ")

		// TIP : 当middleware想阻止访问某资源的时候请使用Abort  他会办好后面的一切.
		c.Abort()
		c.Redirect(http.StatusFound, login.LoginViewUrl)
	}

	c.Next()
}
