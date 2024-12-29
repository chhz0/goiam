package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// KeyXRequestID 请求ID
const KeyXRequestID = "X-Request-ID"

func RequestId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		rid := ctx.GetHeader(KeyXRequestID)

		if rid == "" {
			rid = uuid.Must(uuid.NewRandom()).String()
			ctx.Request.Header.Set(KeyXRequestID, rid)
			ctx.Set(KeyXRequestID, rid)
		}

		ctx.Writer.Header().Set(KeyXRequestID, rid)
		ctx.Next()
	}
}
