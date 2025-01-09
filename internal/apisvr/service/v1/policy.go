package v1

import (
	"context"

	"github.com/chhz0/goiam/internal/apisvr/dal"
	errcode "github.com/chhz0/goiam/internal/pkg/errorscore/code"
	"github.com/chhz0/goiam/internal/pkg/model"
	"github.com/chhz0/goiam/pkg/errors"
	"github.com/chhz0/goiam/pkg/meta"
)

type PolicySrv interface {
	Create(ctx context.Context, policy *model.Policy, opts meta.CreateOptions) error
	Update(ctx context.Context, policy *model.Policy, opts meta.UpdateOptions) error
	Delete(ctx context.Context, username string, name string, opts meta.DeleteOptions) error
	DeleteCollection(ctx context.Context, username string, name []string, opts meta.DeleteOptions) error
	Get(ctx context.Context, username string, name string, opts meta.GetOptions) (*model.Policy, error)
	List(ctx context.Context, username string, opts meta.ListOptions) (*model.PolicyList, error)
}

type policyService struct {
	dal dal.Factory
}

// Create implements PolicySrv.
func (p *policyService) Create(ctx context.Context, policy *model.Policy, opts meta.CreateOptions) error {
	if err := p.dal.Policies().Create(ctx, policy, opts); err != nil {
		return errors.WithCode(errcode.ErrDatabase, err)
	}

	return nil
}

// Delete implements PolicySrv.
func (p *policyService) Delete(ctx context.Context, username string, name string, opts meta.DeleteOptions) error {
	if err := p.dal.Policies().Delete(ctx, username, name, opts); err != nil {
		return errors.WithCode(errcode.ErrDatabase, err)
	}

	return nil
}

// DeleteCollection implements PolicySrv.
func (p *policyService) DeleteCollection(ctx context.Context, username string, name []string, opts meta.DeleteOptions) error {
	if err := p.dal.Policies().DeleteCollection(ctx, username, name, opts); err != nil {
		return errors.WithCode(errcode.ErrDatabase, err)
	}

	return nil
}

// Get implements PolicySrv.
func (p *policyService) Get(ctx context.Context, username string, name string, opts meta.GetOptions) (*model.Policy, error) {
	policy, err := p.dal.Policies().Get(ctx, username, name, opts)
	if err != nil {
		return nil, errors.WithCode(errcode.ErrDatabase, err)
	}

	return policy, nil
}

// List implements PolicySrv.
func (p *policyService) List(ctx context.Context, username string, opts meta.ListOptions) (*model.PolicyList, error) {
	policies, err := p.dal.Policies().List(ctx, username, opts)
	if err != nil {
		return nil, errors.WithCode(errcode.ErrDatabase, err)
	}

	return policies, nil
}

// Update implements PolicySrv.
func (p *policyService) Update(ctx context.Context, policy *model.Policy, opts meta.UpdateOptions) error {
	if err := p.dal.Policies().Update(ctx, policy, opts); err != nil {
		return errors.WithCode(errcode.ErrDatabase, err)
	}

	return nil
}

func NewPolicies(f dal.Factory) *policyService {
	return &policyService{
		dal: f,
	}
}
