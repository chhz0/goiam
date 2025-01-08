package dal

import (
	"context"

	"github.com/chhz0/goiam/internal/pkg/model"
	"github.com/chhz0/goiam/pkg/meta"
)

type SecretDal interface {
	Create(ctx context.Context, secret *model.Secret, opts meta.CreateOptions) error
	Update(ctx context.Context, secret *model.Secret, opts meta.UpdateOptions) error
	Delete(ctx context.Context, username, secretID string, opts meta.DeleteOptions) error
	DeleteCollection(ctx context.Context, username string, secretID []string, opts meta.DeleteOptions) error
	Get(ctx context.Context, username, secretID string, opts meta.GetOptions) (*model.Secret, error)
	List(ctx context.Context, username string, opts meta.ListOptions) (*model.SecretList, error)
}
