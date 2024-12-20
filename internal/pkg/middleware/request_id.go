package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const XRequestIdKey = "X-Request-Id"

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
