package v1

import (
	"context"

	"github.com/chhz0/goiam/internal/apisvr/dal"
	errcode "github.com/chhz0/goiam/internal/pkg/errorscore/code"
	"github.com/chhz0/goiam/internal/pkg/model"
	"github.com/chhz0/goiam/pkg/errors"
	"github.com/chhz0/goiam/pkg/meta"
)

type SecretSrv interface {
	Create(ctx context.Context, secret *model.Secret, opts meta.CreateOptions) error
	Update(ctx context.Context, secret *model.Secret, opts meta.UpdateOptions) error
	Delete(ctx context.Context, username, secretID string, opts meta.DeleteOptions) error
	DeleteCollection(ctx context.Context, username string, secretID []string, opts meta.DeleteOptions) error
	Get(ctx context.Context, username, secretID string, opts meta.GetOptions) (*model.Secret, error)
	List(ctx context.Context, username string, opts meta.ListOptions) (*model.SecretList, error)
}

type secretService struct {
	dal dal.Factory
}

// Create implements SecretSrv.
func (s *secretService) Create(ctx context.Context, secret *model.Secret, opts meta.CreateOptions) error {
	if err := s.dal.Secrets().Create(ctx, secret, opts); err != nil {
		return errors.WithCode(errcode.ErrDatabase, err)
	}

	return nil
}

// Delete implements SecretSrv.
func (s *secretService) Delete(ctx context.Context, username string, secretID string, opts meta.DeleteOptions) error {
	if err := s.dal.Secrets().Delete(ctx, username, secretID, opts); err != nil {
		return err
	}

	return nil
}

// DeleteCollection implements SecretSrv.
func (s *secretService) DeleteCollection(ctx context.Context, username string, secretID []string, opts meta.DeleteOptions) error {
	if err := s.dal.Secrets().DeleteCollection(ctx, username, secretID, opts); err != nil {
		return err
	}

	return nil
}

// Get implements SecretSrv.
func (s *secretService) Get(ctx context.Context, username string, secretID string, opts meta.GetOptions) (*model.Secret, error) {
	secret, err := s.dal.Secrets().Get(ctx, username, secretID, opts)
	if err != nil {
		return nil, err
	}

	return secret, nil
}

// List implements SecretSrv.
func (s *secretService) List(ctx context.Context, username string, opts meta.ListOptions) (*model.SecretList, error) {
	secrets, err := s.dal.Secrets().List(ctx, username, opts)
	if err != nil {
		return nil, errors.WithCode(errcode.ErrDatabase, err)
	}

	return secrets, nil
}

// Update implements SecretSrv.
func (s *secretService) Update(ctx context.Context, secret *model.Secret, opts meta.UpdateOptions) error {
	if err := s.dal.Secrets().Update(ctx, secret, opts); err != nil {
		return errors.WithCode(errcode.ErrDatabase, err)
	}

	return nil
}

func NewSecrets(f dal.Factory) *secretService {
	return &secretService{
		dal: f,
	}
}
