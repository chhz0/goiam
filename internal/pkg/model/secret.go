package model

import "github.com/chhz0/goiam/pkg/meta"

type Secret struct {
	meta.ObjectMeta `json:"metadata,omitempty"`
	Username        string `json:"username" gorm:"column:username;type:varchar(64);" validata:"omitempty"`

	SecretID  string `json:"secretId" gorm:"column:secret_id;" validata:"omitempty"`
	SecretKey string `json:"secretKey" gorm:"column:secret_key;" validata:"omitempty"`

	Expires     int64  `json:"expires" gorm:"column:expires;default:0" validata:"omitempty"`
	Description string `json:"description" gorm:"column:description;type:varchar(64);" validata:"omitempty"`
}
