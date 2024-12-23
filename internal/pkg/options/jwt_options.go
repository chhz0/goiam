package options

import (
	"time"

	"github.com/chhz0/goiam/internal/pkg/ginserver"
	"github.com/chhz0/goiam/pkg/simplecobra"
	"github.com/spf13/pflag"
)

const (
	defaultJwtRealm      = "iam jwt"
	defaultJwtKey        = ""
	defaultJwtTimeout    = "24h"
	defaultJwtMaxRefresh = "24h"
)

type JwtOptions struct {
	Realm      string        `json:"realm" mapstructure:"realm"`
	Key        string        `json:"key" mapstructure:"key"`
	Timeout    time.Duration `json:"timeout" mapstructure:"timeout"`
	MaxRefresh time.Duration `json:"max-refresh" mapstructure:"max-refresh"`
}

// LocalFlagsAndRequired implements simplecobra.Flags.
func (j *JwtOptions) LocalFlagsAndRequired() (fs *pflag.FlagSet, required []string) {
	fs = pflag.NewFlagSet("jwt", pflag.ExitOnError)

	fs.StringVar(&j.Realm, "jwt.realm", j.Realm, "Realm name to display to the user.")
	fs.StringVar(&j.Key, "jwt.key", j.Key, "Private key used to sign jwt token.")
	fs.DurationVar(&j.Timeout, "jwt.timeout", j.Timeout, "JWT token timeout.")

	fs.DurationVar(&j.MaxRefresh, "jwt.max-refresh", j.MaxRefresh, ""+
		"This field allows clients to refresh their token until MaxRefresh has passed.")

	return
}

// PersistentFlagsAndRequired implements simplecobra.Flags.
func (j *JwtOptions) PersistentFlagsAndRequired() (fs *pflag.FlagSet, required []string) {
	return
}

var _ simplecobra.Flags = (*JwtOptions)(nil)

func (j *JwtOptions) AppendToServer(c *ginserver.Config) error {
	c.JWT.Realm = j.Realm
	c.JWT.Key = j.Key
	c.JWT.Timeout = j.Timeout
	c.JWT.MaxRefresh = j.MaxRefresh

	return nil
}

func NewDefaultJwtOptions() *JwtOptions {
	dTimeout, _ := time.ParseDuration(defaultJwtTimeout)
	dRefresh, _ := time.ParseDuration(defaultJwtMaxRefresh)

	return &JwtOptions{
		Realm:      defaultJwtRealm,
		Key:        defaultJwtKey,
		Timeout:    dTimeout,
		MaxRefresh: dRefresh,
	}
}
