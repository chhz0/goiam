package user

import (
	"github.com/chhz0/goiam/internal/apisvr/dal"
	"github.com/chhz0/goiam/internal/apisvr/service"
)

type UserHandler struct {
	srv service.Service
}

func NewUserHandler(f dal.Factory) *UserHandler {
	return &UserHandler{
		srv: service.NewService(f),
	}
}
