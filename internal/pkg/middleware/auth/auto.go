package auth

import (
	"strings"

	errcode "github.com/chhz0/goiam/internal/pkg/errorscore/code"
	"github.com/chhz0/goiam/internal/pkg/httpcore"
	"github.com/chhz0/goiam/internal/pkg/middleware"
	"github.com/chhz0/goiam/pkg/errors"
	"github.com/gin-gonic/gin"
)

type AutoStrategy struct {
	basic middleware.AuthStrategy
	jwt   middleware.AuthStrategy
}

type withStrategy func(*AutoStrategy)

func WithBasicStrategy(basic middleware.AuthStrategy) withStrategy {
	return func(as *AutoStrategy) {
		as.basic = basic
	}
}

func WithJWTStrategy(jwt middleware.AuthStrategy) withStrategy {
	return func(as *AutoStrategy) {
		as.jwt = jwt
	}
}

// NewAutoStrategyWith returns a new AutoStrategy with given strategies.
func NewAutoStrategyWith(strategies ...withStrategy) *AutoStrategy {
	auto := &AutoStrategy{}

	for _, sg := range strategies {
		sg(auto)
	}

	return auto
}

const (
	basicHeader = "Basic"
	jwtHeader   = "Bearer"
)

// AuthFunc implements middleware.AuthStrategy.
func (a AutoStrategy) AuthFunc() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		operation := middleware.AuthOperator{}
		authHeader := strings.SplitN(ctx.Request.Header.Get("Authorization"), " ", 2)

		if len(authHeader) != 2 {
			httpcore.WriteResponse(ctx,
				errors.WithCodef(errcode.ErrInvalidAuthHeader, "Authorization header format is wrong."),
				nil,
			)
			ctx.Abort()

			return
		}

		switch authHeader[0] {
		case basicHeader:
			operation.SetStrategy(a.basic)
		case jwtHeader:
			operation.SetStrategy(a.jwt)
		default:
			httpcore.WriteResponse(ctx,
				errors.WithCodef(errcode.ErrSignatureInvalid, "unrecognized Authorization header."),
				nil,
			)
			ctx.Abort()
			return
		}

		operation.AuthFunc()(ctx)

		ctx.Next()
	}
}

var _ middleware.AuthStrategy = (*AutoStrategy)(nil)
