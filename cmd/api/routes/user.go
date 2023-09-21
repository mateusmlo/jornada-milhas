package routes

import (
	"github.com/mateusmlo/jornada-milhas/cmd/api/controllers"
	"github.com/mateusmlo/jornada-milhas/cmd/api/middlewares"
	"github.com/mateusmlo/jornada-milhas/config"
)

type UserRouter struct {
	rh     config.RequestHandler
	uc     *controllers.UserController
	md     *middlewares.JWTMiddleware
	logger config.Logger
}

func NewUserRouter(uc *controllers.UserController, logger config.Logger, rh config.RequestHandler, md *middlewares.JWTMiddleware) *UserRouter {
	return &UserRouter{
		uc:     uc,
		rh:     rh,
		logger: logger,
		md:     md,
	}
}

func (r *UserRouter) Setup() {
	r.logger.Info("Setting up user routes...")

	public := r.rh.Gin.Group("/api")
	public.POST("/user", r.uc.CreateUser)
	public.GET("/user/whoami", r.md.JwtAuthMiddleware(), r.uc.CurrentUser)
	public.PATCH("/user/:id", r.md.JwtAuthMiddleware(), r.uc.UpdateUser)

	private := r.rh.Gin.Group("/api/admin")
	private.Use(r.md.JwtAuthMiddleware())
	private.GET("/user/:id", r.uc.GetUserByUUID)
	private.GET("/user", r.uc.GetAllUsers)
	private.DELETE("/user/:id", r.uc.DeactivateUser)

}
