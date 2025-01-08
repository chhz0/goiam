package v1

import (
	"context"
	"regexp"

	"github.com/chhz0/goiam/internal/apisvr/dal"
	errcode "github.com/chhz0/goiam/internal/pkg/errorscore/code"
	"github.com/chhz0/goiam/internal/pkg/model"
	"github.com/chhz0/goiam/pkg/errors"
	"github.com/chhz0/goiam/pkg/meta"
)

type UserSrv interface {
	Create(ctx context.Context, user *model.User, opts meta.CreateOptions) error
	Update(ctx context.Context, user *model.User, opts meta.UpdateOptions) error
	Delete(ctx context.Context, username string, opts meta.DeleteOptions) error
	DeleteCollection(ctx context.Context, username []string, opts meta.DeleteOptions) error
	Get(ctx context.Context, username string, opts meta.GetOptions) (*model.User, error)
	List(ctx context.Context, opts meta.ListOptions) (*model.UserList, error)
	ListWithBadPerformance(ctx context.Context, opts meta.ListOptions) (*model.UserList, error)
	ChangePassword(ctx context.Context, user *model.User) error
}

type userService struct {
	dal dal.Factory
}

// ChangePassword implements UserSrv.
func (u *userService) ChangePassword(ctx context.Context, user *model.User) error {
	if err := u.dal.Users().Update(ctx, user, meta.UpdateOptions{}); err != nil {
		return errors.WithCode(errcode.ErrDatabase, err)
	}

	return nil
}

// Create implements UserSrv.
func (u *userService) Create(ctx context.Context, user *model.User, opts meta.CreateOptions) error {
	if err := u.dal.Users().Create(ctx, user, opts); err != nil {
		if match, _ := regexp.MatchString("Duplicate entry '.*' for key 'idx_name'", err.Error()); match {
			return errors.WithCode(errcode.ErrUserAlreadyExist, err)
		}

		return errors.WithCode(errcode.ErrDatabase, err)
	}

	return nil
}

// Delete implements UserSrv.
func (u *userService) Delete(ctx context.Context, username string, opts meta.DeleteOptions) error {
	panic("unimplemented")
}

// DeleteCollection implements UserSrv.
func (u *userService) DeleteCollection(ctx context.Context, username []string, opts meta.DeleteOptions) error {
	panic("unimplemented")
}

// Get implements UserSrv.
func (u *userService) Get(ctx context.Context, username string, opts meta.GetOptions) (*model.User, error) {
	panic("unimplemented")
}

// List implements UserSrv.
func (u *userService) List(ctx context.Context, opts meta.ListOptions) (*model.UserList, error) {
	panic("unimplemented")
}

// ListWithBadPerformance implements UserSrv.
func (u *userService) ListWithBadPerformance(ctx context.Context, opts meta.ListOptions) (*model.UserList, error) {
	panic("unimplemented")
}

// Update implements UserSrv.
func (u *userService) Update(ctx context.Context, user *model.User, opts meta.UpdateOptions) error {
	panic("unimplemented")
}

func NewUsers(f dal.Factory) *userService {
	return &userService{
		dal: f,
	}
}
