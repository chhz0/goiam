package policy

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

func (p *PolicyHandler) Update(ctx *gin.Context) {
	log.L(ctx, logger.UseKeys...).Info("update policy function called.")

	var r model.Policy
	if err := ctx.ShouldBindJSON(&r); err != nil {
		httpcore.WriteResponse(ctx, errors.WithCode(errcode.ErrBind, err), nil)
		return
	}

	pol, err := p.srv.Policies().Get(ctx, ctx.GetString(middleware.KeyUsername),
		ctx.Param("name"), meta.GetOptions{})
	if err != nil {
		httpcore.WriteResponse(ctx, err, nil)
		return
	}

	pol.Policy = r.Policy
	pol.ExtenAttrs = r.ExtenAttrs

	// TODO: validate

	if err := p.srv.Policies().Update(ctx, pol, meta.UpdateOptions{}); err != nil {
		httpcore.WriteResponse(ctx, err, nil)
		return
	}

	httpcore.WriteResponse(ctx, nil, pol)
}
