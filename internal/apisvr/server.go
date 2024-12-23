package apisvr

import (
	"github.com/chhz0/goiam/internal/apisvr/config"
	"github.com/chhz0/goiam/internal/pkg/ginserver"
	"github.com/chhz0/goiam/pkg/graceful"
	"github.com/chhz0/goiam/pkg/graceful/shutdownmanagers/posixsignal"
)

type apiServer struct {
	gracefulShutdown *graceful.GracefulShutdown
	ginServer        *ginserver.Server
	// todo redisServer gRPCServer
}

// todo 添加额外配置
type ExtraConfig struct {
}

func newAPIServer(cfg *config.Config) (*apiServer, error) {
	// 创建优雅服务
	gs := graceful.New()
	gs.AddShutdownManager(posixsignal.NewPosixSignalManager())

	ginServerConf, err := buildGinServerConfig(cfg)
	if err != nil {
		return nil, err
	}

	ginServer, err := ginServerConf.Complete().Server()
	if err != nil {
		return nil, err
	}

	return &apiServer{
		gracefulShutdown: gs,
		ginServer:        ginServer,
	}, nil
}

func buildGinServerConfig(cfg *config.Config) (gConf *ginserver.Config, lastErr error) {
	gConf = ginserver.NewZeroConfig()

	if lastErr = cfg.Server.AppendToServer(gConf); lastErr != nil {
		return
	}

	if lastErr = cfg.SecureServing.AppendToServer(gConf); lastErr != nil {
		return
	}

	if lastErr = cfg.InsecureServing.AppendToServer(gConf); lastErr != nil {
		return
	}

	if lastErr = cfg.Fearure.AppendToServer(gConf); lastErr != nil {
		return
	}

	return
}

func (s *apiServer) PreRun() *preApiServer {
	initRouter(s.ginServer.Engine)

	s.gracefulShutdown.AddShutdownCallback(graceful.OnShutdownFunc(func(str string) error {
		// todo 关闭mysql、grpc

		s.ginServer.Shutdown()

		return nil
	}))

	return &preApiServer{s}
}

type preApiServer struct {
	*apiServer
}

func (s *preApiServer) Run() error {

	if err := s.gracefulShutdown.Start(); err != nil {
		panic(err)
	}

	return s.ginServer.Run()
}
