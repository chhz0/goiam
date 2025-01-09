package user

import (
	errcode "github.com/chhz0/goiam/internal/pkg/errorscore/code"
	"github.com/chhz0/goiam/internal/pkg/httpcore"
	"github.com/chhz0/goiam/internal/pkg/logger"
	"github.com/chhz0/goiam/internal/pkg/utils/authutil"
	"github.com/chhz0/goiam/pkg/errors"
	"github.com/chhz0/goiam/pkg/log"
	"github.com/chhz0/goiam/pkg/meta"
	"github.com/gin-gonic/gin"
)

type ChangePasswordRequest struct {
	OldPassword string `json:"oldPassword" binding:"omitempty"`

	NewPassword string `json:"newPassword" binding:"password"`
}

func (u *UserHandler) ChangePassword(ctx *gin.Context) {
	log.L(ctx, logger.UseKeys...).Info("user change password function called.")

	var r ChangePasswordRequest

	if err := ctx.ShouldBindJSON(&r); err != nil {
		httpcore.WriteResponse(ctx, errors.WithCode(errcode.ErrBind, err), nil)
		return
	}

	user, err := u.srv.Users().Get(ctx, ctx.Param("name"), meta.GetOptions{})
	if err != nil {
		httpcore.WriteResponse(ctx, err, nil)
		return
	}

	if err := user.Compare(r.OldPassword); err != nil {
		httpcore.WriteResponse(ctx, errors.WithCode(errcode.ErrPasswordIncorrect, err), nil)
		return
	}

	user.Password, _ = authutil.Encrypt(r.NewPassword)
	if err := u.srv.Users().ChangePassword(ctx, user); err != nil {
		httpcore.WriteResponse(ctx, err, nil)
		return
	}

	httpcore.WriteResponse(ctx, nil, nil)
}
