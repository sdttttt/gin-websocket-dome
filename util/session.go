/*
 * @Description: In User Settings Edit
 * @Author: your name
 * @Date: 2019-09-03 12:39:37
 * @LastEditTime: 2019-09-05 17:30:48
 * @LastEditors: Please set LastEditors
 */
package util

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func SetSession(context *gin.Context, key string, value interface{}) {
	session := sessions.Default(context)
	session.Set(key, value)
	session.Save()
}

func GetSession(context *gin.Context, key string) interface{} {
	session := sessions.Default(context)
	return session.Get(key)
}
