/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-08-31 00:48:59
 * @LastEditTime: 2019-09-04 17:52:14
 * @LastEditors: Please set LastEditors
 */
package home

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/**
 * @description 控制器的声明请按照处理器对应路由路径（ ** 对应 **Url ）
 */
const (
	HelloUrl = "/"
)

func Hello(context *gin.Context) {
	context.HTML(http.StatusOK, "home.tmpl", nil)
}
