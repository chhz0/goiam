package auth

import (
	ginjwt "github.com/appleboy/gin-jwt/v2"
	"github.com/chhz0/goiam/internal/pkg/middleware"
	"github.com/gin-gonic/gin"
)

const AuthzAudience = "iam.authz"

type JWTStrategy struct {
	ginjwt.GinJWTMiddleware
}

// // AuthFunc implements middleware.AuthStrategy.
// func (j JWTStrategy) AuthFunc() gin.HandlerFunc {
// 	panic("unimplemented")
// }

func NewJWTStrategy(gjwt ginjwt.GinJWTMiddleware) JWTStrategy {
	return JWTStrategy{gjwt}
}

// AuthFunc implements middleware.AuthStrategy.
func (j JWTStrategy) AuthFunc() gin.HandlerFunc {
	return j.MiddlewareFunc()
}

var _ middleware.AuthStrategy = (*JWTStrategy)(nil)
