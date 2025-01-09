package policy

import (
	"github.com/chhz0/goiam/internal/pkg/httpcore"
	"github.com/chhz0/goiam/internal/pkg/logger"
	"github.com/chhz0/goiam/internal/pkg/middleware"
	"github.com/chhz0/goiam/pkg/log"
	"github.com/chhz0/goiam/pkg/meta"
	"github.com/gin-gonic/gin"
)

func (p *PolicyHandler) DeleteCollection(ctx *gin.Context) {
	log.L(ctx, logger.UseKeys...).Info("delete policy collection function called.")

	if err := p.srv.Policies().DeleteCollection(ctx, ctx.GetString(middleware.KeyUsername),
		ctx.QueryArray("name"), meta.DeleteOptions{}); err != nil {
		httpcore.WriteResponse(ctx, err, nil)
		return
	}

	httpcore.WriteResponse(ctx, nil, nil)
}
