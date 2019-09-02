/*
 * @Description: Application 入口
 * @Author: your name
 * @Date: 2019-08-30 19:11:42
 * @LastEditTime: 2019-09-02 00:18:17
 * @LastEditors: Please set LastEditors
 */
package main

import (
	"gin-web/router"

	"github.com/gin-gonic/gin"
)

func main() {
	application := gin.Default()

	/**
	*	Socket Server Start.
	*	現在我提供了Gin 框架的快速啟動
	*	這種啟動方法不是我所推薦的
	 */
	//go socket.Run()

	router := &router.Router{Gin: application}

	router.RegisterMiddleware()
	router.RigisterController()

	application.Run(":80")
}
