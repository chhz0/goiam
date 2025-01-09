package secret

import (
	"github.com/AlekSi/pointer"
	errcode "github.com/chhz0/goiam/internal/pkg/errorscore/code"
	"github.com/chhz0/goiam/internal/pkg/httpcore"
	"github.com/chhz0/goiam/internal/pkg/logger"
	"github.com/chhz0/goiam/internal/pkg/middleware"
	"github.com/chhz0/goiam/internal/pkg/model"
	"github.com/chhz0/goiam/internal/pkg/utils/idutil"
	"github.com/chhz0/goiam/pkg/errors"
	"github.com/chhz0/goiam/pkg/log"
	"github.com/chhz0/goiam/pkg/meta"
	"github.com/gin-gonic/gin"
)

const maxSecretCount = 10

func (s *SecretHandler) Create(ctx *gin.Context) {
	log.L(ctx, logger.UseKeys...).Info("create secret function called.")

	var r model.Secret
	if err := ctx.ShouldBindJSON(&r); err != nil {
		httpcore.WriteResponse(ctx, errors.WithCode(errcode.ErrBind, err), nil)
		return
	}

	// TODO: validata

	username := ctx.GetString(middleware.KeyUsername)
	secrets, err := s.srv.Secrets().List(ctx, username, meta.ListOptions{
		Offset: pointer.ToInt64(0),
		Limit:  pointer.ToInt64(-1),
	})
	if err != nil {
		httpcore.WriteResponse(ctx, err, nil)
		return
	}

	if secrets.TotalCount >= maxSecretCount {
		httpcore.WriteResponse(ctx, errors.WithCodef(errcode.ErrReachMaxCount, "secret count: %d", secrets.TotalCount), nil)
		return
	}

	r.Username = username

	r.SecretID = idutil.NewSecretID()
	r.SecretKey = idutil.NewSecretKey()

	if err := s.srv.Secrets().Create(ctx, &r, meta.CreateOptions{}); err != nil {
		httpcore.WriteResponse(ctx, err, nil)
		return
	}

	httpcore.WriteResponse(ctx, nil, r)
}
