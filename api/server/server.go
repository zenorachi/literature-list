package server

import (
	"context"
	"fmt"
	"github.com/zenorachi/literature-list/internal/config"
	"net/http"
	"time"
)

type Server struct {
	server          *http.Server
	notify          chan error
	shutdownTimeout time.Duration
}

func New(cfg *config.Config, handler http.Handler) *Server {
	httpServer := &http.Server{
		Handler:      handler,
		ReadTimeout:  cfg.HTTP.ReadTimeout,
		WriteTimeout: cfg.HTTP.WriteTimeout,
		Addr:         fmt.Sprintf("%s:%s", cfg.HTTP.Host, cfg.HTTP.Port),
	}

	return &Server{
		server:          httpServer,
		notify:          make(chan error, 1),
		shutdownTimeout: cfg.HTTP.ShutdownTimeout,
	}
}

func (s *Server) Run() {
	go func() {
		s.notify <- s.server.ListenAndServe()
		close(s.notify)
	}()
}

func (s *Server) Notify() <-chan error {
	return s.notify
}

func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()

	return s.server.Shutdown(ctx)
}
