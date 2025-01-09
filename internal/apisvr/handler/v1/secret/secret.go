package secret

import (
	"github.com/chhz0/goiam/internal/apisvr/dal"
	"github.com/chhz0/goiam/internal/apisvr/service"
)

type SecretHandler struct {
	srv service.Service
}

func NewSecretHandler(f dal.Factory) *SecretHandler {
	return &SecretHandler{
		srv: service.NewService(f),
	}
}
