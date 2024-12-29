package middleware

import (
	"github.com/chhz0/goiam/internal/pkg/logger"
	"github.com/gin-gonic/gin"
)

const (
	KeyUsername = "username"
	KeyUserID   = "userID"
)

// Context 是一个中间件，向 gin.Context 注入 日志需要的字段信息，以实现调用链追踪
func Context() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set(logger.KeyRequestID, ctx.GetString(KeyXRequestID))
		ctx.Set(logger.KeyUsername, ctx.GetString(KeyUsername))
		ctx.Next()
	}
}
