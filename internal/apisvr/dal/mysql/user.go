package mysql

import (
	"context"

	"github.com/chhz0/goiam/internal/pkg/model"
	"gorm.io/gorm"
)

type users struct {
	db *gorm.DB
}

// Delete implements dal.UserSDal.
func (u *users) Delete(ctx context.Context, username string) error {
	panic("unimplemented")
}

// DeleteCollection implements dal.UserSDal.
func (u *users) DeleteCollection(ctx context.Context, username []string) error {
	panic("unimplemented")
}

// Get implements dal.UserSDal.
func (u *users) Get(ctx context.Context, username string) (*model.User, error) {
	panic("unimplemented")
}

// List implements dal.UserSDal.
func (u *users) List(ctx context.Context) ([]*model.User, error) {
	panic("unimplemented")
}

// Update implements dal.UserSDal.
func (u *users) Update(ctx context.Context, user *model.User) error {
	panic("unimplemented")
}

func newUsers(db *gorm.DB) *users {
	return &users{db: db}
}

func (u *users) Create(ctx context.Context, user *model.User) error {
	return u.db.Create(&user).Error
}
