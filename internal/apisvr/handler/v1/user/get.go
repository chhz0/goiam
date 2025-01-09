package user

import (
	"github.com/chhz0/goiam/internal/pkg/httpcore"
	"github.com/chhz0/goiam/internal/pkg/logger"
	"github.com/chhz0/goiam/pkg/log"
	"github.com/chhz0/goiam/pkg/meta"
	"github.com/gin-gonic/gin"
)

func (u *UserHandler) Get(ctx *gin.Context) {
	log.L(ctx, logger.UseKeys...).Info("user get function called.")

	user, err := u.srv.Users().Get(ctx, ctx.Param("name"), meta.GetOptions{})
	if err != nil {
		httpcore.WriteResponse(ctx, err, nil)
		return
	}

	httpcore.WriteResponse(ctx, nil, user)
}
