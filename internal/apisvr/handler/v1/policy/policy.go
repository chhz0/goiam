package policy

import (
	"github.com/chhz0/goiam/internal/apisvr/dal"
	"github.com/chhz0/goiam/internal/apisvr/service"
)

type PolicyHandler struct {
	srv service.Service
}

func NewPolicyHandler(f dal.Factory) *PolicyHandler {
	return &PolicyHandler{
		srv: service.NewService(f),
	}
}
