package routes

import (
	"fmt"

	"github.com/mateusmlo/jornada-milhas/cmd/api/controllers"
	"github.com/mateusmlo/jornada-milhas/config"
)

type AuthRouter struct {
	rh config.RequestHandler
	ac controllers.JWTAuthController
}

func NewAuthRouter(ac controllers.JWTAuthController, rh config.RequestHandler) *AuthRouter {
	return &AuthRouter{
		rh: rh,
		ac: ac,
	}
}

func (r *AuthRouter) Setup() {
	fmt.Println("\nSetting up auth routes...")

	api := r.rh.Gin.Group("/api/auth")

	api.POST("/login", r.ac.SignIn)
}
