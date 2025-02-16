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

func (p *PolicyHandler) Create(ctx *gin.Context) {
	log.L(ctx, logger.UseKeys...).Info("create policy function called.")

	var r model.Policy
	if err := ctx.ShouldBindJSON(&r); err != nil {
		httpcore.WriteResponse(ctx, errors.WithCode(errcode.ErrBind, err), nil)

		return
	}

	// TODO validate

	r.Username = ctx.GetString(middleware.KeyUsername)
	if err := p.srv.Policies().Create(ctx, &r, meta.CreateOptions{}); err != nil {
		httpcore.WriteResponse(ctx, err, nil)

		return
	}

	httpcore.WriteResponse(ctx, nil, r)
}
