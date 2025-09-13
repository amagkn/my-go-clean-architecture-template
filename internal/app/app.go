package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/amagkn/my-go-clean-architecture-template/config"
	"github.com/amagkn/my-go-clean-architecture-template/pkg/base_errors"
	"github.com/amagkn/my-go-clean-architecture-template/pkg/http_server"
	"github.com/amagkn/my-go-clean-architecture-template/pkg/logger"
	"github.com/amagkn/my-go-clean-architecture-template/pkg/postgres"
	"github.com/amagkn/my-go-clean-architecture-template/pkg/router"
	"github.com/go-chi/chi/v5"
)

type Dependences struct {
	RouterHTTP *chi.Mux
	Postgres   *postgres.Pool
}

func Run(ctx context.Context, c config.Config) (err error) {
	var deps Dependences

	deps.Postgres, err = postgres.New(ctx, c.Postgres)
	if err != nil {
		return base_errors.WithPath("postgres.New", err)
	}
	defer deps.Postgres.Close()

	deps.RouterHTTP = router.New()

	ProductDomain(deps)

	httpServer := http_server.New(deps.RouterHTTP, c.HTTP.Port)
	defer httpServer.Close()

	waiting(httpServer)

	return nil
}

func waiting(httpServer *http_server.Server) {
	logger.Info("App started")

	wait := make(chan os.Signal, 1)
	signal.Notify(wait, os.Interrupt, syscall.SIGTERM)

	select {
	case i := <-wait:
		logger.Info("App got signal: " + i.String())
	case err := <-httpServer.Notify():
		logger.Error(err, "App got notify: httpServer.Notify")
	}

	logger.Info("App is stopping")
}
