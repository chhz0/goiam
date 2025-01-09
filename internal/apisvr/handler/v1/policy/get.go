package policy

import (
	"github.com/chhz0/goiam/internal/pkg/httpcore"
	"github.com/chhz0/goiam/internal/pkg/logger"
	"github.com/chhz0/goiam/internal/pkg/middleware"
	"github.com/chhz0/goiam/pkg/log"
	"github.com/chhz0/goiam/pkg/meta"
	"github.com/gin-gonic/gin"
)

func (p *PolicyHandler) Get(ctx *gin.Context) {
	log.L(ctx, logger.UseKeys...).Info("get policy function called.")

	pol, err := p.srv.Policies().Get(ctx, ctx.GetString(middleware.KeyUsername),
		ctx.Param("name"), meta.GetOptions{})
	if err != nil {
		httpcore.WriteResponse(ctx, err, nil)
		return
	}

	httpcore.WriteResponse(ctx, nil, pol)
}
