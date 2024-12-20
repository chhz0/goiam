package config

import (
	"github.com/chhz0/goiam/internal/apisvr/options"
)

type Config struct {
	*options.APIServerOptions
}

func NewConfigWithOptions(opts *options.APIServerOptions) (*Config, error) {
	return &Config{opts}, nil
}
