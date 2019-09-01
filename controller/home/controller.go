/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-08-31 00:48:59
 * @LastEditTime: 2019-08-31 18:21:13
 * @LastEditors: Please set LastEditors
 */
package home

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	HelloUrl = "/"
)

func Hello(context *gin.Context) {
	context.HTML(http.StatusOK, "home.tmpl", nil)
}
