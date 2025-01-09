package apisvr

import (
	"github.com/chhz0/goiam/internal/apisvr/dal/mysql"
	"github.com/chhz0/goiam/internal/apisvr/handler/v1/policy"
	"github.com/chhz0/goiam/internal/apisvr/handler/v1/secret"
	"github.com/chhz0/goiam/internal/apisvr/handler/v1/user"
	errcode "github.com/chhz0/goiam/internal/pkg/errorscore/code"
	"github.com/chhz0/goiam/internal/pkg/httpcore"
	"github.com/chhz0/goiam/internal/pkg/middleware/auth"
	"github.com/chhz0/goiam/pkg/errors"
	"github.com/gin-gonic/gin"
)

func initRouter(g *gin.Engine) {
	initMiddlewares(g)
	initHandler(g)
}

func initMiddlewares(g *gin.Engine) {}

func initHandler(g *gin.Engine) {

	jwtStrategy, _ := newJWTAuth().(auth.JWTStrategy)
	g.POST("/login", jwtStrategy.LoginHandler)
	g.POST("/logout", jwtStrategy.LogoutHandler)
	g.POST("/refresh", jwtStrategy.RefreshHandler)

	auto := newAutoAuth()
	g.NoRoute(auto.AuthFunc(), func(ctx *gin.Context) {
		httpcore.WriteResponse(ctx,
			errors.WithCodef(errcode.ErrPageNotFound, "Page not found."),
			nil,
		)
	})

	mysqlIns, err := mysql.GetMysqlFactoryOr(nil)
	if err != nil {
		panic(err)
	}

	v1 := g.Group("/v1")
	{

		// 用户接口
		userv1 := v1.Group("/users")
		{
			userHandler := user.NewUserHandler(mysqlIns)

			userv1.POST("", userHandler.Create)
			userv1.Use(auto.AuthFunc())
			userv1.DELETE("", userHandler.Delete)
			userv1.DELETE("/:name", userHandler.DeleteCollection)
			userv1.PUT(":name/change-password", userHandler.ChangePassword)
			userv1.PUT(":name", userHandler.Update)
			userv1.GET("", userHandler.List)
			userv1.GET(":name", userHandler.Get)
		}

		v1.Use(auto.AuthFunc())
		// 策略接口
		policyv1 := v1.Group("/policies")
		{
			policyHandler := policy.NewPolicyHandler(mysqlIns)

			policyv1.POST("", policyHandler.Create)
			policyv1.DELETE("", policyHandler.Delete)
			policyv1.DELETE("/:name", policyHandler.DeleteCollection)
			policyv1.PUT(":name", policyHandler.Update)
			policyv1.GET("", policyHandler.List)
			policyv1.GET(":name", policyHandler.Get)
		}

		// 密钥接口
		secretv1 := v1.Group("/secrets")
		{
			secretHandler := secret.NewSecretHandler(mysqlIns)

			secretv1.POST("", secretHandler.Create)
			secretv1.DELETE("", secretHandler.Delete)
			secretv1.PUT(":name", secretHandler.Update)
			secretv1.GET("", secretHandler.List)
			secretv1.GET(":name", secretHandler.Get)
		}
	}

}
