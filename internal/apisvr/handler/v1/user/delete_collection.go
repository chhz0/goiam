package user

import (
	"github.com/chhz0/goiam/internal/pkg/httpcore"
	"github.com/chhz0/goiam/internal/pkg/logger"
	"github.com/chhz0/goiam/pkg/log"
	"github.com/chhz0/goiam/pkg/meta"
	"github.com/gin-gonic/gin"
)

func (u *UserHandler) DeleteCollection(ctx *gin.Context) {
	log.L(ctx, logger.UseKeys...).Info("user delete collection function called.")

	username := ctx.QueryArray("name")

	if err := u.srv.Users().DeleteCollection(ctx, username, meta.DeleteOptions{}); err != nil {
		httpcore.WriteResponse(ctx, err, nil)
		return
	}

	httpcore.WriteResponse(ctx, nil, nil)
}
