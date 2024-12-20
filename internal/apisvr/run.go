package apisvr

import "github.com/chhz0/goiam/internal/apisvr/config"

func apiRun(cfg *config.Config) error {
	server, err := newAPIServer(cfg)
	if err != nil {
		return nil
	}

	return server.PreRun().Run()
}
