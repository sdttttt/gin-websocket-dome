package middleware

import "github.com/gin-gonic/gin"

func OriginAcceptMiddleware(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
}
