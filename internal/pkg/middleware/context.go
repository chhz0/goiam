package middleware

import "github.com/gin-gonic/gin"

const UsernameKey = "username"

// Context 是一个中间件，向 gin.Context 注入一些公共的前缀字段
func Context() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set("requestId", ctx.GetString(XRequestIdKey))
		ctx.Set("username", ctx.GetString(UsernameKey))
		ctx.Next()
	}
}
