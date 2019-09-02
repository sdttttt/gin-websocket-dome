/*
 * @Description: Application 入口
 * @Author: your name
 * @Date: 2019-08-30 19:11:42
 * @LastEditTime: 2019-09-02 22:03:32
 * @LastEditors: Please set LastEditors
 */
package main

import (
	"gin-web/router"
	"gin-web/socket"

	"github.com/gin-gonic/gin"
)

func main() {
	application := gin.Default()

	router := &router.Router{Gin: application}

	socket.TimeNoticeEnable()
	socket.ConnectCountNoticeEnable()

	go socket.GetConnectHub().RunAndListen()

	router.RegisterMiddleware()
	router.RigisterController()

	application.Run(":80")
}
