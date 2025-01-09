package secret

import (
	"github.com/chhz0/goiam/internal/pkg/httpcore"
	"github.com/chhz0/goiam/internal/pkg/logger"
	"github.com/chhz0/goiam/internal/pkg/middleware"
	"github.com/chhz0/goiam/pkg/log"
	"github.com/chhz0/goiam/pkg/meta"
	"github.com/gin-gonic/gin"
)

func (s *SecretHandler) DeleteCollection(ctx *gin.Context) {
	log.L(ctx, logger.KeyUsername).Info("delete secret collection function called.")

	if err := s.srv.Secrets().DeleteCollection(ctx, ctx.GetString(middleware.KeyUsername),
		ctx.QueryArray("name"), meta.DeleteOptions{}); err != nil {
		httpcore.WriteResponse(ctx, err, nil)
		return
	}

	httpcore.WriteResponse(ctx, nil, nil)
}
