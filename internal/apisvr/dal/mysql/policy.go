package mysql

import (
	"context"

	"github.com/chhz0/goiam/internal/apisvr/dal"
	"github.com/chhz0/goiam/internal/pkg/errorscore/errorno"
	"github.com/chhz0/goiam/internal/pkg/model"
	"github.com/chhz0/goiam/internal/pkg/utils/gormutil"
	"github.com/chhz0/goiam/pkg/errors"
	"github.com/chhz0/goiam/pkg/meta"
	"gorm.io/gorm"
)

type policies struct {
	db *gorm.DB
}

// Create implements dal.PolicyDal.
func (p *policies) Create(ctx context.Context, policy *model.Policy, opts meta.CreateOptions) error {
	return p.db.Create(&policy).Error
}

// Delete implements dal.PolicyDal.
func (p *policies) Delete(ctx context.Context, username string, name string, opts meta.DeleteOptions) error {
	if opts.Unscoped {
		p.db = p.db.Unscoped()
	}

	err := p.db.Where("username = ? AND name = ?", username, name).Delete(&model.Policy{}).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.WithCode(errorno.ErrDatabase, err)
	}

	return nil
}

func (p *policies) DeleteByUser(ctx context.Context, username string, opts meta.DeleteOptions) error {
	if opts.Unscoped {
		p.db = p.db.Unscoped()
	}

	return p.db.Where("username = ?", username).Delete(&model.Policy{}).Error
}

// DeleteCollection implements dal.PolicyDal.
func (p *policies) DeleteCollection(ctx context.Context, username string, name []string, opts meta.DeleteOptions) error {
	if opts.Unscoped {
		p.db = p.db.Unscoped()
	}

	return p.db.Where("username = ? AND name IN (?)", username, name).Delete(&model.Policy{}).Error
}

func (p *policies) DeleteCollectionByUser(ctx context.Context, username []string, opts meta.DeleteOptions) error {
	if opts.Unscoped {
		p.db = p.db.Unscoped()
	}

	return p.db.Where("username IN (?)", username).Delete(&model.Policy{}).Error
}

// Get implements dal.PolicyDal.
func (p *policies) Get(ctx context.Context, username string, name string, opts meta.GetOptions) (*model.Policy, error) {
	policy := &model.Policy{}
	err := p.db.Where("username = ? AND name = ?", username, name).First(&policy).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.WithCode(errorno.ErrPolicyNotFound, err)
		}
		return nil, errors.WithCode(errorno.ErrDatabase, err)
	}

	return policy, nil
}

// List implements dal.PolicyDal.
func (p *policies) List(ctx context.Context, username string, opts meta.ListOptions) (*model.PolicyList, error) {
	ret := &model.PolicyList{}
	ol := gormutil.UnpointerLO(opts.Limit, opts.Offset)

	if username != "" {
		p.db = p.db.Where("username = ?", username)
	}

	// todo : selector

	d := p.db.Where("name like ?", "%"+opts.FieldSelector+"%").
		Offset(ol.Offset).Limit(ol.Limit).Order("id desc").Find(&ret.Items).
		Offset(-1).Limit(-1).Count(&ret.TotalCount)

	return ret, d.Error
}

// Update implements dal.PolicyDal.
func (p *policies) Update(ctx context.Context, policy *model.Policy, opts meta.UpdateOptions) error {
	return p.db.Save(&policy).Error
}

var _ dal.PolicyDal = (*policies)(nil)

func newPolicies(ds *dbStore) *policies {
	return &policies{db: ds.db}
}
