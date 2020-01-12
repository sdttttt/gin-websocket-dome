package middleware

import "github.com/gin-gonic/gin"

/**
	跨域
*/
func OriginAcceptMiddleware(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
}
