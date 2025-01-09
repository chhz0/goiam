package secret

import (
	errcode "github.com/chhz0/goiam/internal/pkg/errorscore/code"
	"github.com/chhz0/goiam/internal/pkg/httpcore"
	"github.com/chhz0/goiam/internal/pkg/logger"
	"github.com/chhz0/goiam/internal/pkg/middleware"
	"github.com/chhz0/goiam/internal/pkg/model"
	"github.com/chhz0/goiam/pkg/errors"
	"github.com/chhz0/goiam/pkg/log"
	"github.com/chhz0/goiam/pkg/meta"
	"github.com/gin-gonic/gin"
)

func (s *SecretHandler) Update(ctx *gin.Context) {
	log.L(ctx, logger.UseKeys...).Info("update secret function called.")

	var r model.Secret
	if err := ctx.ShouldBindJSON(&r); err != nil {
		httpcore.WriteResponse(ctx, errors.WithCode(errcode.ErrBind, err), nil)
		return
	}

	username := ctx.GetString(middleware.KeyUsername)
	name := ctx.Param("name")

	secret, err := s.srv.Secrets().Get(ctx, username, name, meta.GetOptions{})
	if err != nil {
		httpcore.WriteResponse(ctx, errors.WithCode(errcode.ErrBind, err), nil)
		return
	}

	secret.Expires = r.Expires
	secret.Description = r.Description
	secret.ExtenAttrs = r.ExtenAttrs

	// TODO: validate

	if err := s.srv.Secrets().Update(ctx, secret, meta.UpdateOptions{}); err != nil {
		httpcore.WriteResponse(ctx, err, nil)
		return
	}

	httpcore.WriteResponse(ctx, nil, secret)
}
