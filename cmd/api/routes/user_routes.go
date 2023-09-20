package routes

import (
	"github.com/mateusmlo/jornada-milhas/cmd/api/controllers"
	"github.com/mateusmlo/jornada-milhas/config"
)

type Router struct {
	rh     config.RequestHandler
	uc     *controllers.UserController
	logger config.Logger
}

func NewUserRouter(uc *controllers.UserController, logger config.Logger, rh config.RequestHandler) *Router {
	return &Router{
		uc:     uc,
		rh:     rh,
		logger: logger,
	}
}

func (r *Router) Setup() {
	r.logger.Info("Setting up user routes...")

	api := r.rh.Gin.Group("/api")

	api.GET("/user", r.uc.GetAllUsers)
	api.GET("/user/:id", r.uc.GetUserByUUID)
	api.POST("/user", r.uc.CreateUser)
	api.DELETE("/user/:id", r.uc.DeactivateUser)
	api.PATCH("/user/:id", r.uc.UpdateUser)

}
