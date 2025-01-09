package apisvr

import (
	"context"
	"encoding/base64"
	"net/http"
	"strings"
	"time"

	ginjwt "github.com/appleboy/gin-jwt/v2"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/chhz0/goiam/internal/apisvr/dal"
	"github.com/chhz0/goiam/internal/pkg/logger"
	"github.com/chhz0/goiam/internal/pkg/middleware"
	"github.com/chhz0/goiam/internal/pkg/middleware/auth"
	"github.com/chhz0/goiam/internal/pkg/model"
	"github.com/chhz0/goiam/pkg/log"
	"github.com/chhz0/goiam/pkg/meta"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type loginInfo struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

func newBasicAuth() middleware.AuthStrategy {
	return auth.NewBasicStrategy(func(username, password string) bool {
		user, err := dal.Client().Users().Get(context.Background(), username, meta.GetOptions{})
		if err != nil {
			return false
		}

		if err := user.Compare(password); err != nil {
			return false
		}

		user.LoginedAt = time.Now()
		_ = dal.Client().Users().Update(context.Background(), user, meta.UpdateOptions{})

		return true
	})
}

func newJWTAuth() middleware.AuthStrategy {
	jwt, _ := ginjwt.New(&ginjwt.GinJWTMiddleware{
		Realm:            viper.GetString("jwt.realm"),
		SigningAlgorithm: "HS256",
		Key:              []byte(viper.GetString("jwt.key")),
		Timeout:          viper.GetDuration("jwt.timeout"),
		MaxRefresh:       viper.GetDuration("jwt.max_refresh"),
		Authenticator:    authenticator(),
		LoginResponse:    loginResponse(),
		LogoutResponse: func(c *gin.Context, code int) {
			c.JSON(http.StatusOK, gin.H{"msg": "ok"})
		},
		RefreshResponse: refreshResponse(),
		PayloadFunc:     payloadFunc(),
		IdentityKey:     middleware.KeyUsername,
		IdentityHandler: func(ctx *gin.Context) interface{} {
			claims := jwt.ExtractClaims(ctx)

			return claims[jwt.IdentityKey]
		},
		Authorizator: authorizator(),
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{"msg": message})
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		SendCookie:    true,
		TimeFunc:      time.Now,
		// TODO: HTTPStautsMessageFunc
	})

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
		var login loginInfo
		var err error

		if ctx.Request.Header.Get("Authorization") != "" {
			login, err = parseWithHeader(ctx)
		} else {
			login, err = parseWithBody(ctx)
		}

		if err != nil {
			// jwt.ErrFailedAuthentication
			return "", err
		}

		user, err := dal.Client().Users().Get(ctx, login.Username, meta.GetOptions{})
		if err != nil {
			log.Errorf("apisvr.authenticator get user failed: %s.", err.Error())

			return nil, jwt.ErrFailedAuthentication
		}

		if err := user.Compare(login.Password); err != nil {
			return "", jwt.ErrFailedAuthentication
		}

		user.LoginedAt = time.Now()
		_ = dal.Client().Users().Update(ctx, user, meta.UpdateOptions{})

		return user, nil
	}
}

func parseWithHeader(ctx *gin.Context) (loginInfo, error) {
	authSlice := strings.SplitN(ctx.Request.Header.Get("Authorization"), " ", 2)
	if len(authSlice) != 2 || authSlice[0] != "Basic" {
		log.Errorf("apisvr.parseWithHeader Authorization header format is wrong.")

		return loginInfo{}, jwt.ErrFailedAuthentication
	}

	payload, err := base64.StdEncoding.DecodeString(authSlice[1])
	if err != nil {
		log.Errorf("apisvr.parseWithHeader decode basic string: %s.", err.Error())

		return loginInfo{}, jwt.ErrFailedAuthentication
	}

	pair := strings.SplitN(string(payload), ":", 2)
	if len(pair) != 2 {
		log.Errorf("apisvr.parseWithHeader parse payload failed.")

		return loginInfo{}, jwt.ErrFailedAuthentication
	}

	return loginInfo{
		Username: pair[0],
		Password: pair[1],
	}, nil
}

func parseWithBody(ctx *gin.Context) (loginInfo, error) {
	var login loginInfo
	if err := ctx.ShouldBindJSON(&login); err != nil {
		log.Errorf("apisvr.parseWithBody bind json failed: %s.", err.Error())

		return loginInfo{}, jwt.ErrFailedAuthentication
	}

	return login, nil
}

func loginResponse() func(c *gin.Context, code int, token string, expire time.Time) {
	return tokenResp
}

func refreshResponse() func(ctx *gin.Context, code int, token string, expire time.Time) {
	return tokenResp
}

func tokenResp(ctx *gin.Context, code int, token string, expire time.Time) {
	ctx.JSON(http.StatusOK, gin.H{
		"token":  token,
		"expire": expire.Format(time.RFC3339),
	})
}

const (
	// APIServerAudience is the default audience for the APIServer jwt audience field.
	APIServerAudience = "iam.api.helloxx.cn"

	// APIServerIssuer is the default issuer for the APIServer jwt issuer field.
	APIServerIssuer = "iam-apisvr"
)

func payloadFunc() func(data interface{}) jwt.MapClaims {
	return func(data interface{}) jwt.MapClaims {
		claims := jwt.MapClaims{
			"iss": APIServerIssuer,
			"aud": APIServerAudience,
		}
		if u, ok := data.(*model.User); ok {
			claims[jwt.IdentityKey] = u.Name
			claims["sub"] = u.Name
		}

		return claims
	}
}

func authorizator() func(data interface{}, c *gin.Context) bool {
	return func(data interface{}, c *gin.Context) bool {
		if v, ok := data.(string); ok {
			log.L(c, logger.UseKeys...).Infof("apisvr.authorizator user: %s.", v)

			return true
		}

		return false
	}
}
