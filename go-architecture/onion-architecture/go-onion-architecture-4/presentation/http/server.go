package http

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"auth/internal/config"
	"auth/presentation"
	"auth/presentation/http/handler"

	"github.com/google/wire"
)

type Server struct {
	*http.Server
}

var _ presentation.Server = (*Server)(nil)

var Set = wire.NewSet(NewServer, handler.New)

func NewServer(h http.Handler) *Server {
	s := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", config.Config.HTTPServerHost(), config.Config.HTTPServerPort()),
		Handler: h,
	}
	return &Server{s}
}

func (s *Server) Start() error {
	log.Printf("HTTP server running and listen %s ...\n", s.Addr)
	return s.ListenAndServe()
}

func (s *Server) GracefulShutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return s.Shutdown(ctx)
}
