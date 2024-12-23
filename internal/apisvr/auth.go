package apisvr

import (
	ginjwt "github.com/appleboy/gin-jwt/v2"
	"github.com/chhz0/goiam/internal/pkg/middleware"
	"github.com/chhz0/goiam/internal/pkg/middleware/auth"
	"github.com/gin-gonic/gin"
)

const (
	// APIServerAudience is the default audience for the APIServer jwt audience field.
	APIServerAudience = "iam.api.helloxx.cn"

	// APIServerIssuer is the default issuer for the APIServer jwt issuer field.
	APIServerIssuer = "iam-apisvr"
)

type loginInfo struct {
	Username string `json:"username" form:"username" query:"username" binding:"required,username"`
	Password string `json:"password" form:"password" query:"password" binding:"required,password"`
}

func newBasicAuth() middleware.AuthStrategy {
	return auth.NewBasicStrategy(func(username, password string) bool {

		return true
	})
}

func newJWTAuth() middleware.AuthStrategy {
	jwt, _ := ginjwt.New(&ginjwt.GinJWTMiddleware{})

	return auth.NewJWTStrategy(*jwt)
}

func newAutoAuth() middleware.AuthStrategy {
	return auth.NewAutoStrategyWith(
		auth.WithBasicStrategy(newBasicAuth().(auth.BasicStrategy)),
		auth.WithJWTStrategy(newJWTAuth().(auth.JWTStrategy)),
	)
}

func authenticator() func(ctx *gin.Context) (any, error) {
	return func(ctx *gin.Context) (any, error) {

		return nil, nil
	}
}
