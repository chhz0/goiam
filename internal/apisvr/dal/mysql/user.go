package mysql

import (
	"context"

	"github.com/chhz0/goiam/internal/apisvr/dal"
	"github.com/chhz0/goiam/internal/pkg/model"
	"github.com/chhz0/goiam/pkg/meta"
	"gorm.io/gorm"
)

type users struct {
	db *gorm.DB
}

// Create implements dal.UserSDal.
func (u *users) Create(ctx context.Context, user *model.User, opts meta.CreateOptions) error {
	return u.db.Create(user).Error
}

// Delete implements dal.UserSDal.
func (u *users) Delete(ctx context.Context, username string, opts meta.DeleteOptions) error {
	panic("unimplemented")
}

// DeleteCollection implements dal.UserSDal.
func (u *users) DeleteCollection(ctx context.Context, username []string, opts meta.DeleteOptions) error {
	panic("unimplemented")
}

// Get implements dal.UserSDal.
func (u *users) Get(ctx context.Context, username string, opts meta.GetOptions) (*model.User, error) {
	panic("unimplemented")
}

// List implements dal.UserSDal.
func (u *users) List(ctx context.Context, opts meta.ListOptions) ([]*model.UserList, error) {
	panic("unimplemented")
}

// Update implements dal.UserSDal.
func (u *users) Update(ctx context.Context, user *model.User, opts meta.UpdateOptions) error {
	return u.db.Save(user).Error
}

var _ dal.UserSDal = (*users)(nil)

func newUsers(ds *dbStore) *users {
	return &users{ds.db}
}
