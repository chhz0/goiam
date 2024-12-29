package model

import (
	"time"

	"github.com/chhz0/goiam/pkg/meta"
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
