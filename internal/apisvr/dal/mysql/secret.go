package mysql

import (
	"context"

	errcode "github.com/chhz0/goiam/internal/pkg/errorscore/code"
	"github.com/chhz0/goiam/internal/pkg/model"
	"github.com/chhz0/goiam/internal/pkg/utils/gormutil"
	"github.com/chhz0/goiam/pkg/errors"
	"github.com/chhz0/goiam/pkg/meta"
	"github.com/chhz0/goiam/pkg/meta/fields"
	"gorm.io/gorm"
)

type secrets struct {
	db *gorm.DB
}

func (s *secrets) Create(ctx context.Context, secret *model.Secret, opts meta.CreateOptions) error {
	return s.db.Create(secret).Error
}

func (s *secrets) Update(ctx context.Context, secret *model.Secret, opts meta.UpdateOptions) error {
	return s.db.Save(&secret).Error
}

func (s *secrets) Delete(ctx context.Context, username, secretID string, opts meta.DeleteOptions) error {
	if opts.Unscoped {
		s.db = s.db.Unscoped()
	}

	err := s.db.Where("username = ? and secret_id = ?", username, secretID).Delete(&model.Secret{}).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.WithCode(errcode.ErrDatabase, err)
	}

	return nil
}

func (s *secrets) DeleteCollection(ctx context.Context, username string, secretIDs []string, opts meta.DeleteOptions) error {
	if opts.Unscoped {
		s.db = s.db.Unscoped()
	}

	return s.db.Where("username = ? and secret_id in (?)", username, secretIDs).Delete(&model.Secret{}).Error
}

func (s *secrets) Get(ctx context.Context, username, secretID string, opts meta.GetOptions) (*model.Secret, error) {
	secret := &model.Secret{}

	err := s.db.Where("username =? and secret_id =?", username, secretID).First(&secret).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.WithCode(errcode.ErrSecretNotFound, err)
		}
		return nil, errors.WithCode(errcode.ErrDatabase, err)
	}

	return secret, nil
}

func (s *secrets) List(ctx context.Context, username string, opts meta.ListOptions) (*model.SecretList, error) {
	ret := &model.SecretList{}
	ol := gormutil.UnpointerLO(opts.Limit, opts.Offset)

	if username != "" {
		s.db = s.db.Where("username =?", username)
	}

	selector, _ := fields.ParseSelector(opts.FieldSelector)
	name, _ := selector.RequiresExactMatch("name")

	d := s.db.Where("name like?", "%"+name+"%").
		Offset(ol.Offset).Limit(ol.Limit).Order("id desc").Find(&ret.Items).
		Offset(-1).Limit(-1).Count(&ret.TotalCount)

	return ret, d.Error
}

func newSecrets(ds *dbStore) *secrets {
	return &secrets{db: ds.db}
}
