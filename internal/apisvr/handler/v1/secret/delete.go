package secret

import (
	"github.com/chhz0/goiam/internal/pkg/httpcore"
	"github.com/chhz0/goiam/internal/pkg/logger"
	"github.com/chhz0/goiam/internal/pkg/middleware"
	"github.com/chhz0/goiam/pkg/log"
	"github.com/chhz0/goiam/pkg/meta"
	"github.com/gin-gonic/gin"
)

func (s *SecretHandler) Delete(ctx *gin.Context) {
	log.L(ctx, logger.UseKeys...).Info("delete secret function called.")

	if err := s.srv.Secrets().Delete(ctx, ctx.GetString(middleware.KeyUsername),
		ctx.Param("name"), meta.DeleteOptions{Unscoped: true}); err != nil {
		httpcore.WriteResponse(ctx, err, nil)
		return
	}

	httpcore.WriteResponse(ctx, nil, nil)
}
