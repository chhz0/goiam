package policy

import (
	"github.com/chhz0/goiam/internal/pkg/httpcore"
	"github.com/chhz0/goiam/internal/pkg/logger"
	"github.com/chhz0/goiam/internal/pkg/middleware"
	"github.com/chhz0/goiam/pkg/log"
	"github.com/chhz0/goiam/pkg/meta"
	"github.com/gin-gonic/gin"
)

func (p *PolicyHandler) Delete(ctx *gin.Context) {
	log.L(ctx, logger.UseKeys...).Info("delete policy function called.")
	if err := p.srv.Policies().Delete(ctx, ctx.GetString(middleware.KeyUsername),
		ctx.Param("name"), meta.DeleteOptions{}); err != nil {
		httpcore.WriteResponse(ctx, err, nil)
		return
	}

	httpcore.WriteResponse(ctx, nil, nil)
}
