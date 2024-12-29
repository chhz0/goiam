package dal

import (
	"context"

	"github.com/chhz0/goiam/internal/pkg/model"
)

type UserSDal interface {
	Create(ctx context.Context, user *model.User) error
	Update(ctx context.Context, user *model.User) error
	Delete(ctx context.Context, username string) error
	DeleteCollection(ctx context.Context, username []string) error
	Get(ctx context.Context, username string) (*model.User, error)
	List(ctx context.Context) ([]*model.User, error)
}
