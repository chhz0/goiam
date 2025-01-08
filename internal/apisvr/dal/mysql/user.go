package mysql

import (
	"context"

	"github.com/chhz0/goiam/internal/apisvr/dal"
	errcode "github.com/chhz0/goiam/internal/pkg/errorscore/code"
	"github.com/chhz0/goiam/internal/pkg/model"
	"github.com/chhz0/goiam/internal/pkg/utils/gormutil"
	"github.com/chhz0/goiam/pkg/errors"
	"github.com/chhz0/goiam/pkg/meta"
	"github.com/chhz0/goiam/pkg/meta/fields"
	"gorm.io/gorm"
)

type users struct {
	db *gorm.DB
}

// Create implements dal.UserSDal.
func (u *users) Create(ctx context.Context, user *model.User, opts meta.CreateOptions) error {
	return u.db.Create(&user).Error
}

// Delete implements dal.UserSDal.
func (u *users) Delete(ctx context.Context, username string, opts meta.DeleteOptions) error {
	pol := newPolicies(&dbStore{u.db})
	if err := pol.DeleteByUser(ctx, username, opts); err != nil {
		return err
	}

	if opts.Unscoped {
		u.db = u.db.Unscoped()
	}
	err := u.db.Where("name = ?", username).Delete(&model.User{}).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.WithCode(errcode.ErrDatabase, err)
	}

	return nil
}

// DeleteCollection implements dal.UserSDal.
func (u *users) DeleteCollection(ctx context.Context, username []string, opts meta.DeleteOptions) error {
	pol := newPolicies(&dbStore{u.db})
	if err := pol.DeleteCollectionByUser(ctx, username, opts); err != nil {
		return err
	}

	if opts.Unscoped {
		u.db = u.db.Unscoped()
	}

	return u.db.Where("name IN (?)", username).Delete(&model.User{}).Error
}

// Get implements dal.UserSDal.
func (u *users) Get(ctx context.Context, username string, opts meta.GetOptions) (*model.User, error) {
	user := &model.User{}
	err := u.db.Where("name = ?", username).First(user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.WithCode(errcode.ErrUserNotFound, err)
		}

		return nil, errors.WithCode(errcode.ErrDatabase, err)
	}

	return user, nil
}

// List implements dal.UserSDal.
func (u *users) List(ctx context.Context, opts meta.ListOptions) (*model.UserList, error) {
	ret := &model.UserList{}
	ol := gormutil.UnpointerLO(opts.Limit, opts.Offset)

	selector, _ := fields.ParseSelector(opts.FieldSelector)
	username, _ := selector.RequiresExactMatch("name")
	d := u.db.Where("name like ? and status = 1", "%"+username+"%").
		Offset(ol.Offset).Limit(ol.Limit).Order("id desc").Find(&ret.Items).
		Offset(-1).Limit(-1).Count(&ret.TotalCount)

	return ret, d.Error
}

func (u *users) ListOptional(ctx context.Context, opts meta.ListOptions) (*model.UserList, error) {
	ret := &model.UserList{}
	ol := gormutil.UnpointerLO(opts.Limit, opts.Offset)

	where := model.User{}
	whereNot := model.User{
		IsAdmin: 0,
	}

	selector, _ := fields.ParseSelector(opts.FieldSelector)
	username, found := selector.RequiresExactMatch("name")
	if found {
		where.Name = username
	}
	d := u.db.Where(where).Not(whereNot).
		Offset(ol.Offset).Limit(ol.Limit).Order("id desc").Find(&ret.Items).
		Offset(-1).Limit(-1).Count(&ret.TotalCount)

	return ret, d.Error
}

// Update implements dal.UserSDal.
func (u *users) Update(ctx context.Context, user *model.User, opts meta.UpdateOptions) error {
	return u.db.Save(user).Error
}

var _ dal.UserSDal = (*users)(nil)

func newUsers(ds *dbStore) *users {
	return &users{ds.db}
}
