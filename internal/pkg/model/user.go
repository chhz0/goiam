package model

import (
	"fmt"
	"time"

	"github.com/chhz0/goiam/internal/pkg/utils/authutil"
	"github.com/chhz0/goiam/internal/pkg/utils/idutil"
	"github.com/chhz0/goiam/pkg/meta"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type User struct {
	*meta.ObjectMeta `json:"metadata,omitempty"`

	Status int `json:"status" gorm:"column:status;default:1" validata:"omitempty"`

	Nickname string `json:"nickname" gorm:"column:nickname;type:varchar(64);not null" validata:"required,min=1,max=64"`

	Password string `json:"password,omitempty" gorm:"column:password;type:varchar(64);not null" validata:"required,min=1,max=64"`

	Email string `json:"email" gorm:"column:email;type:varchar(64);not null" validata:"required,email,min=1,max=64"`

	Phone string `json:"phone" gorm:"column:phone;type:varchar(64);not null" validata:"omitempty"`

	IsAdmin int `json:"IsAdmin,omitempty" gorm:"column:is_admin;default:0" validata:"omitempty"`

	TotalPolicy int64 `json:"totalPolicy" gorm:"-" validata:"omitempty"`

	LoginedAt time.Time `json:"loginedAt,omitempty" gorm:"column:logined_ad;default:null;"`
}

func (u *User) TableName() string {
	return "user"
}

func (u *User) Compare(pwd string) error {
	if err := authutil.Compare(u.Password, pwd); err != nil {
		return fmt.Errorf("failed to compare password: %w", err)
	}
	return nil
}

func (u *User) Validata() error {
	// TODO: 支持参数验证
	validator.New(validator.WithRequiredStructEnabled())
	return nil
}

func (u *User) AfterCreate(tx *gorm.DB) (err error) {
	u.InstanceID = idutil.GenerateInstanceID(u.TableName(), u.ID, "user-")

	return tx.Save(u).Error
}

type UserList struct {
	meta.ListMeta `json:",inline"`

	Items []*User `json:"items"`
}
