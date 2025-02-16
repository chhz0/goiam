package service

import (
	"github.com/chhz0/goiam/internal/apisvr/dal"
	v1 "github.com/chhz0/goiam/internal/apisvr/service/v1"
)

//go:generate mockgen -source=./service.go -destination=./mock/mock_service.go -package=mock Service
type Service interface {
	Users() v1.UserSrv
	Policies() v1.PolicySrv
	Secrets() v1.SecretSrv
}

type service struct {
	factory dal.Factory
}

// Users implements Service.
func (s *service) Users() v1.UserSrv {
	return v1.NewUsers(s.factory)
}

// Policies implements Service.
func (s *service) Policies() v1.PolicySrv {
	return v1.NewPolicies(s.factory)
}

// Secrets implements Service.
func (s *service) Secrets() v1.SecretSrv {
	return v1.NewSecrets(s.factory)
}

func NewService(factory dal.Factory) Service {
	return &service{
		factory: factory,
	}
}
