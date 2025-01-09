package secret

import (
	errcode "github.com/chhz0/goiam/internal/pkg/errorscore/code"
	"github.com/chhz0/goiam/internal/pkg/httpcore"
	"github.com/chhz0/goiam/internal/pkg/logger"
	"github.com/chhz0/goiam/internal/pkg/middleware"
	"github.com/chhz0/goiam/pkg/errors"
	"github.com/chhz0/goiam/pkg/log"
	"github.com/chhz0/goiam/pkg/meta"
	"github.com/gin-gonic/gin"
)

func (s *SecretHandler) List(ctx *gin.Context) {
	log.L(ctx, logger.UseKeys...).Info("list secret function called.")

	var r meta.ListOptions
	if err := ctx.ShouldBindQuery(&r); err != nil {
		httpcore.WriteResponse(ctx, errors.WithCode(errcode.ErrBind, err), nil)
		return
	}

	secrets, err := s.srv.Secrets().List(ctx, ctx.GetString(middleware.KeyUsername), r)
	if err != nil {
		httpcore.WriteResponse(ctx, err, nil)
		return
	}

	httpcore.WriteResponse(ctx, nil, secrets)
}
