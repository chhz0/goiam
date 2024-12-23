package auth

import (
	"encoding/base64"
	"strings"

	"github.com/chhz0/goiam/internal/pkg/constants/errorsno"
	"github.com/chhz0/goiam/internal/pkg/httpcore"
	"github.com/chhz0/goiam/internal/pkg/middleware"
	"github.com/chhz0/goiam/pkg/errors"
	"github.com/gin-gonic/gin"
)

type BasicStrategy struct {
	compare func(username, password string) bool
}

func NewBasicStrategy(compare func(username, password string) bool) *BasicStrategy {
	return &BasicStrategy{compare: compare}
}

// AuthFunc implements middleware.AuthStrategy.
func (b BasicStrategy) AuthFunc() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		basic := strings.SplitN(ctx.Request.Header.Get("Authorization"), " ", 2)

		if len(basic) != 2 || basic[0] != "Basic" {
			httpcore.WriteResponse(ctx,
				errors.WithCodef(errorsno.ErrSignatureInvalid, "Authorization header format is wrong."),
				nil,
			)

			ctx.Abort()
			return
		}

		payload, _ := base64.StdEncoding.DecodeString(basic[1])
		pair := strings.SplitN(string(payload), ":", 2)

		if len(pair) != 2 || !b.compare(pair[0], pair[1]) {
			httpcore.WriteResponse(ctx,
				errors.WithCodef(errorsno.ErrSignatureInvalid, "Authorization header format is wrong."),
				nil,
			)
			ctx.Abort()

			return
		}

		ctx.Set(middleware.UsernameKey, pair[0])
		ctx.Next()
	}
}

var _ middleware.AuthStrategy = (*BasicStrategy)(nil)
