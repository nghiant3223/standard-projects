package httpserver

import (
	"context"
	"errors"
	"net"
	"net/http"
	"strconv"
)

type Server struct {
	Port int

	server *http.Server
}

func NewServer(port int, handler http.Handler) *Server {
	return &Server{
		Port:   port,
		server: &http.Server{Handler: handler},
	}
}

func (s *Server) Start(ctx context.Context) error {
	portString := strconv.Itoa(s.Port)
	s.server.Addr = net.JoinHostPort("", portString)
	err := s.server.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		return nil
	}
	return err
}

func (s *Server) Stop(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
