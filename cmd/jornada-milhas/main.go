// main starting point of the app
package main

import (
	"context"

	"github.com/google/uuid"
	"github.com/mateusmlo/jornada-milhas/cmd/api/controllers"
	"github.com/mateusmlo/jornada-milhas/cmd/api/routes"
	"github.com/mateusmlo/jornada-milhas/config"
	"github.com/mateusmlo/jornada-milhas/domain"
	repository "github.com/mateusmlo/jornada-milhas/internal/repositories"
	"github.com/spf13/viper"
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

func startServer(lc fx.Lifecycle, r *routes.Router, logger config.Logger, rh config.RequestHandler) {
	r.Setup()

	logger.Info("Staring server...")
	port := viper.GetString("SERVER_PORT")

	rh.Gin.Run(":" + port)

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			logger.Info("Starting server on port " + port)

			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Info("Stopping server...")

			return nil
		},
	})
}
