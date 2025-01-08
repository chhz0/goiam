package dal

import (
	"context"

	"github.com/chhz0/goiam/internal/pkg/model"
	"github.com/chhz0/goiam/pkg/meta"
)

type UserSDal interface {
	Create(ctx context.Context, user *model.User, opts meta.CreateOptions) error
	Update(ctx context.Context, user *model.User, opts meta.UpdateOptions) error
	Delete(ctx context.Context, username string, opts meta.DeleteOptions) error
	DeleteCollection(ctx context.Context, username []string, opts meta.DeleteOptions) error
	Get(ctx context.Context, username string, opts meta.GetOptions) (*model.User, error)
	List(ctx context.Context, opts meta.ListOptions) (*model.UserList, error)
}
