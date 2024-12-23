package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// XRequestIdKey 请求ID
const XRequestIdKey = "X-Request-ID"

func RequestId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		rid := ctx.GetHeader(XRequestIdKey)

		if rid == "" {
			rid = uuid.Must(uuid.NewRandom()).String()
			ctx.Request.Header.Set(XRequestIdKey, rid)
			ctx.Set(XRequestIdKey, rid)
		}

		ctx.Writer.Header().Set(XRequestIdKey, rid)
		ctx.Next()
	}
}
