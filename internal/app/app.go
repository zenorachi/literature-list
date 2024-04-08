package app

import (
	"fmt"
	"github.com/zenorachi/literature-list/api/rest"
	"github.com/zenorachi/literature-list/api/server"
	_ "github.com/zenorachi/literature-list/docs"
	"github.com/zenorachi/literature-list/internal/config"
	"github.com/zenorachi/literature-list/internal/service"
	"github.com/zenorachi/literature-list/pkg/logger"
	"os"
	"os/signal"
	"syscall"
)

func Run() {
	/* INIT CONFIG */
	cfg, err := config.New()
	if err != nil {
		logger.Fatal("error initializing config")
	}

	/* INIT SERVICES & DEPS */
	services := service.New(cfg)

	/* INIT HTTP HANDLER */
	handler := rest.NewHandler(services)

	/* INIT & RUN HTTP SERVER */
	srv := server.New(cfg, handler.InitRoutes(cfg))
	srv.Run()
	logger.Info(fmt.Sprintf("started listening server on port %s", cfg.HTTP.Port), "http server started")

	/* GRACEFUL SHUTDOWN */
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	/* WAITING FOR SYSCALL */
	select {
	case s := <-quit:
		logger.Info("app - Run - signal: " + s.String())
	case err = <-srv.Notify():
		logger.Error("server", err.Error())
	}

	/* SHUTTING DOWN */
	logger.Info("Shutting down...")
	if err = srv.Shutdown(); err != nil {
		logger.Error("server", err.Error())
	}
}
