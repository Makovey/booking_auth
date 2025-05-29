package app

import (
	"context"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Makovey/booking_utils/pkg/config"
	"github.com/Makovey/booking_utils/pkg/logger"
)

type App struct {
	log logger.Logger
	cfg config.AuthConfig
}

func NewApp(
	log logger.Logger,
	cfg config.AuthConfig,
) *App {
	return &App{
		log: log,
		cfg: cfg,
	}
}

func (a *App) Run() error {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill, syscall.SIGTERM)
	defer stop()

	a.startHTTPServer(ctx)

	<-ctx.Done()
	stop()

	return nil
}

func (a *App) startHTTPServer(ctx context.Context) {
	fn := "app.startHTTPServer"

	addr := net.JoinHostPort("localhost", a.cfg.AuthPort()) // TODO: move to cfg

	srv := &http.Server{
		Addr:    addr,
		Handler: a.initRouter(),
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			a.log.Infof("[%s]: server stopped, cause: %v", fn, err)
		}
	}()

	a.log.Infof("[%s]: server started at %s", fn, addr)

	<-ctx.Done()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		a.log.Infof("[%s]: can't shutdown server, cause: %v", fn, err)
	}
}
