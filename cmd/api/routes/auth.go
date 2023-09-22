package routes

import (
	"github.com/mateusmlo/jornada-milhas/cmd/api/controllers"
	"github.com/mateusmlo/jornada-milhas/config"
)

type AuthRouter struct {
	rh     config.RequestHandler
	ac     controllers.JWTAuthController
	logger config.GinLogger
}

func NewAuthRouter(ac controllers.JWTAuthController, logger config.GinLogger, rh config.RequestHandler) *AuthRouter {
	return &AuthRouter{
		rh:     rh,
		ac:     ac,
		logger: logger,
	}
}

func (r *AuthRouter) Setup() {
	r.logger.Info("Setting up auth routes...")

	api := r.rh.Gin.Group("/api/auth")

	api.POST("/login", r.ac.SignIn)
}
