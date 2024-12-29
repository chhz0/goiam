package dal

import (
	"context"

	"github.com/chhz0/goiam/internal/pkg/model"
)

type SecretDal interface {
	Create(ctx context.Context, secret *model.Secret)
}
