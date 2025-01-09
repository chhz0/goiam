package apisvr

import (
	"context"

	"github.com/chhz0/goiam/internal/apisvr/config"
	"github.com/chhz0/goiam/internal/apisvr/options"
	"github.com/chhz0/goiam/internal/pkg/logger"
	"github.com/chhz0/goiam/pkg/log"
	"github.com/chhz0/goiam/pkg/simplecobra"
)

func New() *simplecobra.Executor {
	opts := options.NewNilAPIOptions()

	x := simplecobra.NewRootCmd("iamapi",
		simplecobra.WithRootLong(`The IAM API is a web service that provides identity authentication and access management,
		built with the Go language.`),
		simplecobra.WithVersion("0.0.1 Snapshot"),
		simplecobra.WithConfig(true, "../../config/iam-api-server.yaml"),
		simplecobra.WithFlagSets(opts),
		simplecobra.WithRunFunc(
			run(opts),
		),
	)

	return x
}

func run(opts *options.APIServerOptions) func(ctx context.Context, args []string) error {
	return func(ctx context.Context, args []string) error {
		logger.NewLogger()
		defer log.Sync()

		cfg, _ := config.NewConfigWithOptions(opts)

		return apiRun(cfg)
	}
}
