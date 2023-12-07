package routes

import (
	"fmt"

	"github.com/mateusmlo/jornada-milhas/cmd/api/controllers"
	"github.com/mateusmlo/jornada-milhas/cmd/api/middlewares"
	"github.com/mateusmlo/jornada-milhas/config"
)

type AuthRouter struct {
	rh *config.RequestHandler
	ac *controllers.AuthController
	md *middlewares.JWTMiddleware
}

func NewAuthRouter(ac *controllers.AuthController, md *middlewares.JWTMiddleware, rh *config.RequestHandler) *AuthRouter {
	return &AuthRouter{
		rh: rh,
		md: md,
		ac: ac,
	}
}

func (r *AuthRouter) Setup() {
	fmt.Println("\nSetting up auth routes...")

	api := r.rh.Gin.Group("/api/auth")

	api.POST("/login", r.ac.SignIn)
	api.GET("/logout", r.md.ValidateRefreshToken(), r.ac.Logout)
	api.GET("/refresh", r.md.ValidateRefreshToken(), r.ac.RenewRefreshToken)
}
