package user

import (
	errcode "github.com/chhz0/goiam/internal/pkg/errorscore/code"
	"github.com/chhz0/goiam/internal/pkg/httpcore"
	"github.com/chhz0/goiam/pkg/errors"
	"github.com/chhz0/goiam/pkg/log"
	"github.com/chhz0/goiam/pkg/meta"
	"github.com/gin-gonic/gin"
)

func (u *UserHandler) List(ctx *gin.Context) {
	log.L(ctx).Info("user list function called.")

	var r meta.ListOptions
	if err := ctx.ShouldBindJSON(&r); err != nil {
		httpcore.WriteResponse(ctx, errors.WithCode(errcode.ErrBind, err), nil)

		return
	}

	users, err := u.srv.Users().List(ctx, r)
	if err != nil {
		httpcore.WriteResponse(ctx, err, nil)

		return
	}

	httpcore.WriteResponse(ctx, nil, users)
}
