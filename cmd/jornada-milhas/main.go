// main starting point of the app
package main

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/google/uuid"
	"github.com/mateusmlo/jornada-milhas/cmd/api/controllers"
	"github.com/mateusmlo/jornada-milhas/cmd/api/middlewares"
	"github.com/mateusmlo/jornada-milhas/cmd/api/routes"
	"github.com/mateusmlo/jornada-milhas/config"
	service "github.com/mateusmlo/jornada-milhas/domain/services"
	repository "github.com/mateusmlo/jornada-milhas/internal/repositories"
	"github.com/mateusmlo/jornada-milhas/utils"
	"go.uber.org/fx"
)

func main() {
	uuid.EnableRandPool()

	app := fx.New(
		config.Module,
		middlewares.Module,
		controllers.Module,
		repository.Module,
		service.Module,
		routes.Module,
		utils.Module,
		fx.Invoke(startServer),
	)

	app.Run()
}

func startServer(
	lc fx.Lifecycle,
	ur *routes.UserRouter,
	ar *routes.AuthRouter,
	rr *routes.ReviewRouter,
	rh *config.RequestHandler,
	env *config.Env,
) {
	ur.SetupRoutes()
	ar.SetupRoutes()
	rr.SetupRoutes()

	fmt.Println("\nStaring server...")
	port := env.ServerPort

	srv := &http.Server{Addr: ":" + port, Handler: rh.Gin}

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			ln, err := net.Listen("tcp", srv.Addr)

			if err != nil {
				fmt.Println("Failed to start HTTP Server at", srv.Addr)
				return err
			}

			go func() {
				err := srv.Serve(ln)
				if err != nil {
					fmt.Println(err)
				}
			}() // process an incoming request in a go routine

			return nil
		},
		OnStop: func(ctx context.Context) error {
			fmt.Println("Stopping server...")

			err := srv.Shutdown(ctx)
			if err != nil {
				fmt.Println(err)
				return err
			}

			return nil
		},
	})
}
