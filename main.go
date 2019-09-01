/*
 * @Description: Application 入口
 * @Author: your name
 * @Date: 2019-08-30 19:11:42
 * @LastEditTime: 2019-08-31 18:58:24
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

	/**
	*	Socket Server Start.
	 */
	go socket.Run()

	router := router.Router{Gin: application}
	router.Rigister()

	application.Run(":80")
}
