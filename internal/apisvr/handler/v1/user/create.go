package user

import (
	"time"

	errcode "github.com/chhz0/goiam/internal/pkg/errorscore/code"
	"github.com/chhz0/goiam/internal/pkg/httpcore"
	"github.com/chhz0/goiam/internal/pkg/model"
	"github.com/chhz0/goiam/internal/pkg/utils/authutil"
	"github.com/chhz0/goiam/pkg/errors"
	"github.com/chhz0/goiam/pkg/log"
	"github.com/chhz0/goiam/pkg/meta"
	"github.com/gin-gonic/gin"
)

func (u *UserHandler) Create(ctx *gin.Context) {
	log.L(ctx).Info("user create function called.")

	var r model.User

	if err := ctx.ShouldBindJSON(&r); err != nil {
		httpcore.WriteResponse(ctx, errors.WithCode(errcode.ErrBind, err), nil)

		return
	}

	if err := r.Validata(); err != nil {
		httpcore.WriteResponse(ctx, errors.WithCode(errcode.ErrValidation, err), nil)

		return
	}

	r.Password, _ = authutil.Encrypt(r.Password)
	r.Status = 1
	r.LoginedAt = time.Now()

	if err := u.srv.Users().Create(ctx, &r, meta.CreateOptions{}); err != nil {
		httpcore.WriteResponse(ctx, err, nil)

		return
	}

	httpcore.WriteResponse(ctx, nil, r)
}
