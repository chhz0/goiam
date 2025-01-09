package apisvr

import (
	"log"

	"github.com/chhz0/goiam/internal/apisvr/config"
	"github.com/chhz0/goiam/internal/apisvr/dal"
	"github.com/chhz0/goiam/internal/apisvr/dal/mysql"
	"github.com/chhz0/goiam/internal/pkg/ginserver"
	"github.com/chhz0/goiam/pkg/graceful"
	"github.com/chhz0/goiam/pkg/graceful/shutdownmanagers/posixsignal"
)

type apiServer struct {
	gracefulShutdown *graceful.GracefulShutdown
	ginServer        *ginserver.Server
	grpcServer       *grpcServer
	// TODO redisServer
}

func newAPIServer(cfg *config.Config) (*apiServer, error) {
	// 创建优雅服务
	gs := graceful.New()
	gs.AddShutdownManager(posixsignal.NewPosixSignalManager())

	// 创建mysql instance
	mysqlFactory, err := mysql.GetMysqlFactoryOr(cfg.MySQL)
	if err != nil {
		return nil, err
	}
	dal.SetClient(mysqlFactory)

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
		mysqlStore, _ := mysql.GetMysqlFactoryOr(nil)
		if mysqlStore != nil {
			_ = mysqlStore.Close()
		}

		// TODO: grpc
		// s.grpcServer.Close()

		_ = s.ginServer.Shutdown()

		return nil
	}))

	return &preApiServer{s}
}

type preApiServer struct {
	*apiServer
}

func (s *preApiServer) Run() error {
	// go s.grpcServer.Run()

	if err := s.gracefulShutdown.Start(); err != nil {
		log.Fatalf("start shutdown manager failed: %v", err)
	}

	return s.ginServer.Run()
}
