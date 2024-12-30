package model

import (
	"encoding/json"
	"fmt"

	"github.com/chhz0/goiam/internal/pkg/utils/idutil"
	"github.com/chhz0/goiam/pkg/meta"
	"github.com/ory/ladon"
	"gorm.io/gorm"
)

// AuthzPolicy is the policy for authorization
type AuthzPolicy struct {
	ladon.DefaultPolicy
}

func (ap *AuthzPolicy) String() string {
	buf, _ := json.Marshal(ap)

	return string(buf)
}

type Policy struct {
	meta.ObjectMeta `json:"metadata,omitempty"`

	Username string `json:"username" gorm:"column:username;type:varchar(255);" validata:"omitempty"`

	Policy AuthzPolicy `json:"policy,omitempty" gorm:"-" validata:"omitempty"`

	PolicyShadow string `json:"-" gorm:"column:policy_shadow;type:longtext" validata:"omitempty"`
}

type PolicyList struct {
	*meta.ListMeta `json:",inline"`

	Items []*Policy `json:"items"`
}

func (p *Policy) TableName() string {
	return "policy"
}

func (p *Policy) BeforeCreate(tx *gorm.DB) (err error) {
	if err := p.ObjectMeta.BeforeCreate(tx); err != nil {
		return fmt.Errorf("failed to run ObjectMeta.BeforeCreate in Policy hook :%w.", err)
	}

	p.Policy.ID = p.Name
	p.PolicyShadow = p.Policy.String()

	return nil
}

func (p *Policy) AfterCreate(tx *gorm.DB) (err error) {
	p.InstanceID = idutil.GenerateInstanceID(p.TableName(), p.ID, "policy-")

	return tx.Save(p).Error
}

func (p *Policy) BeforeUpdate(tx *gorm.DB) (err error) {
	if err := p.ObjectMeta.BeforeUpdate(tx); err != nil {
		return fmt.Errorf("failed to run ObjectMeta.BeforeUpdate in Policy hook :%w.", err)
	}

	p.Policy.ID = p.Name
	p.PolicyShadow = p.Policy.String()

	return nil
}

func (p *Policy) AfterFind(tx *gorm.DB) (err error) {
	if err := p.ObjectMeta.AfterFind(tx); err != nil {
		return fmt.Errorf("failed to run ObjectMeta.AfterFind in Policy hook :%w.", err)
	}

	if err := json.Unmarshal([]byte(p.PolicyShadow), &p.Policy); err != nil {
		return fmt.Errorf("failed to unmarshal policyShadow in Policy hook: %w", err)
	}

	return nil
}
