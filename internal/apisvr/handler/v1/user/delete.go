package user

import (
	"github.com/chhz0/goiam/internal/pkg/httpcore"
	"github.com/chhz0/goiam/internal/pkg/logger"
	"github.com/chhz0/goiam/pkg/log"
	"github.com/chhz0/goiam/pkg/meta"
	"github.com/gin-gonic/gin"
)

func (u *UserHandler) Delete(ctx *gin.Context) {
	log.L(ctx, logger.UseKeys...).Info("user delete function called.")

	if err := u.srv.Users().Delete(ctx, ctx.Param("name"), meta.DeleteOptions{Unscoped: true}); err != nil {
		httpcore.WriteResponse(ctx, err, nil)
		return
	}

	httpcore.WriteResponse(ctx, nil, nil)
}
