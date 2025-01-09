package user

import (
	errcode "github.com/chhz0/goiam/internal/pkg/errorscore/code"
	"github.com/chhz0/goiam/internal/pkg/httpcore"
	"github.com/chhz0/goiam/internal/pkg/logger"
	"github.com/chhz0/goiam/internal/pkg/model"
	"github.com/chhz0/goiam/pkg/errors"
	"github.com/chhz0/goiam/pkg/log"
	"github.com/chhz0/goiam/pkg/meta"
	"github.com/gin-gonic/gin"
)

func (u *UserHandler) Update(ctx *gin.Context) {
	log.L(ctx, logger.UseKeys...).Info("user update function called.")

	var r model.User
	if err := ctx.ShouldBindJSON(&r); err != nil {
		httpcore.WriteResponse(ctx, errors.WithCode(errcode.ErrBind, err), nil)
		return
	}

	user, err := u.srv.Users().Get(ctx, ctx.Param("name"), meta.GetOptions{})
	if err != nil {
		httpcore.WriteResponse(ctx, err, nil)
		return
	}

	user.Nickname = r.Nickname
	user.Email = r.Email
	user.Phone = r.Phone
	user.ExtenAttrs = r.ExtenAttrs
	if err := u.srv.Users().Update(ctx, user, meta.UpdateOptions{}); err != nil {
		httpcore.WriteResponse(ctx, err, nil)
		return
	}

	httpcore.WriteResponse(ctx, nil, user)
}
