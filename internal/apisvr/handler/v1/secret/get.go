package secret

import (
	"github.com/chhz0/goiam/internal/pkg/httpcore"
	"github.com/chhz0/goiam/internal/pkg/logger"
	"github.com/chhz0/goiam/internal/pkg/middleware"
	"github.com/chhz0/goiam/pkg/log"
	"github.com/chhz0/goiam/pkg/meta"
	"github.com/gin-gonic/gin"
)

func (s *SecretHandler) Get(ctx *gin.Context) {
	log.L(ctx, logger.UseKeys...).Info("get secret function called.")

	secret, err := s.srv.Secrets().Get(ctx, ctx.GetString(middleware.KeyUsername),
		ctx.Param("name"), meta.GetOptions{})
	if err != nil {
		httpcore.WriteResponse(ctx, err, nil)

		return
	}

	httpcore.WriteResponse(ctx, nil, secret)
}
