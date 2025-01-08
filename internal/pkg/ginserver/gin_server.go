package ginserver

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/chhz0/goiam/internal/pkg/middleware"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

type Server struct {
	Conf *Config

	*gin.Engine

	ShutdownTimeout time.Duration

	secureServer   *http.Server
	insecureServer *http.Server
}

func (s *Server) init() {
	s.setup()
	s.installMiddlewares()
	s.installAPIs()
}

func (s *Server) setup() {

}

func (s *Server) installMiddlewares() {
	s.Use(middleware.RequestId())
	s.Use(middleware.Context())

	for _, m := range s.Conf.Middlewares {
		mw, ok := middleware.GinMiddlewares[m]
		if !ok {
			continue
		}
		s.Use(mw)
	}
}

func (s *Server) installAPIs() {
	if s.Conf.Healthz {
		s.GET("/healthz", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
		})
	}
	// TODO: enable prometheus and pprof handler

	s.GET("/version", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"version": "nil",
		})
	})
}

func (s *Server) Run() error {
	s.insecureServer = &http.Server{
		Addr:    s.Conf.InsecureServing.Address(),
		Handler: s,
	}

	s.secureServer = &http.Server{
		Addr:    s.Conf.SecureServing.Address(),
		Handler: s,
	}

	var eg errgroup.Group

	eg.Go(func() error {
		if err := s.insecureServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			return err
		}
		return nil
	})

	eg.Go(func() error {
		cert, key := s.Conf.SecureServing.CertKey.Cert, s.Conf.SecureServing.CertKey.Key
		if err := s.secureServer.ListenAndServeTLS(cert, key); err != nil && !errors.Is(err, http.ErrServerClosed) {
			return err
		}
		return nil
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if s.Conf.Healthz {
		if err := s.ping(ctx); err != nil {
			return err
		}
	}

	if err := eg.Wait(); err != nil {
		return err
	}

	return nil
}

// Shutdown 优雅关闭服务
func (s *Server) Shutdown() error {
	closer := []func(context.Context) error{
		s.insecureServer.Shutdown,
		s.secureServer.Shutdown,
	}

	return gracefulShutdown(closer)
}

func gracefulShutdown(closer []func(context.Context) error) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	for _, c := range closer {
		if err := c(ctx); err != nil {
			return err
		}
	}
	return nil
}

// ping 检查服务是否正常工作
func (s *Server) ping(ctx context.Context) error {
	url := fmt.Sprintf("http://%s/healthz", s.Conf.InsecureServing.BindAddress)

	for {
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
		if err != nil {
			return err
		}

		resp, err := http.DefaultClient.Do(req)
		if err == nil && resp.StatusCode == http.StatusOK {

			resp.Body.Close()

			return nil
		}

		time.Sleep(1 * time.Second)

		select {
		case <-ctx.Done():
			return nil
		default:
		}
	}
}
