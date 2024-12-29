package apisvr

import (
	"github.com/chhz0/goiam/internal/pkg/errorscore/errorno"
	"github.com/chhz0/goiam/internal/pkg/httpcore"
	"github.com/chhz0/goiam/internal/pkg/middleware/auth"
	"github.com/chhz0/goiam/pkg/errors"
	"github.com/gin-gonic/gin"
)

func initRouter(g *gin.Engine) {

}

func installMiddlewares(g *gin.Engine) {}

func installHandler(g *gin.Engine) {

	jwtStrategy, _ := newJWTAuth().(auth.JWTStrategy)
	g.POST("/login", jwtStrategy.LoginHandler)
	g.POST("/logout", jwtStrategy.LogoutHandler)
	g.POST("/refresh", jwtStrategy.RefreshHandler)

	auto := newAutoAuth()
	g.NoRoute(auto.AuthFunc(), func(ctx *gin.Context) {
		httpcore.WriteResponse(ctx,
			errors.WithCodef(errorno.ErrPageNotFound, "Page not found."),
			nil,
		)
	})

	v1 := g.Group("/v1")
	{
		// 用户接口
		userv1 := v1.Group("/users")
		{
			userv1.POST("")
			userv1.DELETE("")
			userv1.DELETE("/:name")
			userv1.PUT(":name/change-password")
			userv1.PUT(":name")
			userv1.GET("")
			userv1.GET(":name")
		}

		// 策略接口
		policyv1 := v1.Group("/policies")
		{
			policyv1.POST("")
		}

		// 密钥接口
		secretv1 := v1.Group("/secrets")
		{
			secretv1.POST("")
		}
	}

}
