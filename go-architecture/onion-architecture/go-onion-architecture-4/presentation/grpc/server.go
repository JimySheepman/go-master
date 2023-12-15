package grpc

import (
	"fmt"
	"log"
	"net"

	"auth/internal/config"
	pb "auth/pkg/grpc/go/auth"
	"auth/presentation"

	"github.com/google/wire"
	"google.golang.org/grpc"
)

type Server struct {
	*grpc.Server
}

var _ presentation.Server = (*Server)(nil)

var Set = wire.NewSet(NewServer, NewHandler)

func NewServer(h *Handler) *Server {
	grpcServer := grpc.NewServer()
	pb.RegisterAuthServer(grpcServer, h)
	return &Server{grpcServer}
}

func (s *Server) Start() error {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", config.Config.GRPCServerHost(), config.Config.GRPCServerPort()))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	return s.Serve(lis)
}

func (s *Server) GracefulShutdown() error {
	s.GracefulStop()
	return nil
}
