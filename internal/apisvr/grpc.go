package apisvr

import (
	"net"

	"github.com/chhz0/goiam/pkg/log"
	"google.golang.org/grpc"
)

type grpcServer struct {
	*grpc.Server
	address string
}

func (s *grpcServer) Run() {
	listen, err := net.Listen("tcp", s.address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	go func() {
		if err := s.Serve(listen); err != nil {
			log.Fatalf("failed to start grpc serve: %v", err)
		}
	}()

	log.Infof("grpc server started at %s", s.address)
}

func (s *grpcServer) Close() {
	s.GracefulStop()
	log.Infof("grpc server on %s stopped", s.address)
}
