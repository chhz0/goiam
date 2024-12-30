package dal

import (
	"context"

	"github.com/chhz0/goiam/internal/pkg/model"
	"github.com/chhz0/goiam/pkg/meta"
)

type PolicyDal interface {
	Create(ctx context.Context, policy *model.Policy, opts meta.CreateOptions) error
	Update(ctx context.Context, policy *model.Policy, opts meta.UpdateOptions) error
	Delete(ctx context.Context, username string, name string, opts meta.DeleteOptions) error
	DeleteCollection(ctx context.Context, username string, name []string, opts meta.DeleteOptions) error
	Get(ctx context.Context, username string, name string, opts meta.GetOptions) (*model.Policy, error)
	List(ctx context.Context, username string, opts meta.ListOptions) (*model.PolicyList, error)
}

type PolicyAuditDal interface {
	ClearOutdatedAudit(ctx context.Context, maxReserveDays int) (int64, error)
}
