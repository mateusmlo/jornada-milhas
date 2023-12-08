package routes

import (
	"fmt"

	"github.com/mateusmlo/jornada-milhas/cmd/api/controllers"
	"github.com/mateusmlo/jornada-milhas/cmd/api/middlewares"
	"github.com/mateusmlo/jornada-milhas/config"
)

type AuthRouter struct {
	rh *config.RequestHandler
	ac controllers.AuthController
	md middlewares.AuthMiddleware
}

func NewAuthRouter(
	ac controllers.AuthController,
	md middlewares.AuthMiddleware,
	rh *config.RequestHandler,
) *AuthRouter {
	return &AuthRouter{
		rh: rh,
		md: md,
		ac: ac,
	}
}

func (r *AuthRouter) SetupRoutes() {
	fmt.Println("\nSetting up auth routes...")

	api := r.rh.Gin.Group("/v1/auth")

	api.POST("/login", r.ac.SignIn)
	api.GET("/logout", r.md.ValidateRefreshToken(), r.ac.Logout)
	api.GET("/refresh", r.md.ValidateRefreshToken(), r.ac.RenewTokenPair)
}
