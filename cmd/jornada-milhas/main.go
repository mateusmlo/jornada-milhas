// main starting point of the app
package main

import (
	"context"
	"net"
	"net/http"

	"github.com/google/uuid"
	"github.com/mateusmlo/jornada-milhas/cmd/api/controllers"
	"github.com/mateusmlo/jornada-milhas/cmd/api/routes"
	"github.com/mateusmlo/jornada-milhas/config"
	"github.com/mateusmlo/jornada-milhas/domain"
	repository "github.com/mateusmlo/jornada-milhas/internal/repositories"
	"go.uber.org/fx"
)

func main() {
	uuid.EnableRandPool()
	config.LoadEnvs()

	app := fx.New(
		config.Module,
		controllers.Module,
		repository.Module,
		domain.Module,
		routes.Module,
		fx.Invoke(startServer),
	)

	app.Run()
}

func startServer(lc fx.Lifecycle, r *routes.Router, logger config.Logger, rh config.RequestHandler, env config.Env) {
	r.Setup()

	logger.Info("Staring server...")
	port := env.ServerPort

	srv := &http.Server{Addr: ":" + port, Handler: rh.Gin}

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			ln, err := net.Listen("tcp", srv.Addr)

			if err != nil {
				logger.Error("Failed to start HTTP Server at", srv.Addr)
				return err
			}

			go srv.Serve(ln) // process an incoming request in a go routine

			logger.Info("Succeeded to start HTTP Server at", srv.Addr)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Info("Stopping server...")

			srv.Shutdown(ctx)

			return nil
		},
	})
}
