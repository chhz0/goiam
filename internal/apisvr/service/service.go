package service

import (
	"github.com/chhz0/goiam/internal/apisvr/dal"
	v1 "github.com/chhz0/goiam/internal/apisvr/service/v1"
)

type Service interface {
	Policies() v1.PolicySrv
}

type service struct {
	factory dal.Factory
}

// Policies implements Service.
func (s *service) Policies() v1.PolicySrv {
	return v1.NewPolicies(s.factory)
}

func NewService(factory dal.Factory) Service {
	return &service{
		factory: factory,
	}
}
