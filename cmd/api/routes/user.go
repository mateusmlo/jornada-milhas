package routes

import (
	"fmt"

	"github.com/mateusmlo/jornada-milhas/cmd/api/controllers"
	"github.com/mateusmlo/jornada-milhas/cmd/api/middlewares"
	"github.com/mateusmlo/jornada-milhas/config"
)

type UserRouter struct {
	rh *config.RequestHandler
	uc controllers.UserController
	md middlewares.AuthMiddleware
}

func NewUserRouter(
	uc controllers.UserController,
	rh *config.RequestHandler,
	md middlewares.AuthMiddleware,
) *UserRouter {
	return &UserRouter{
		uc: uc,
		rh: rh,
		md: md,
	}
}

func (r *UserRouter) SetupRoutes() {
	fmt.Println("\nSetting up user routes...")

	public := r.rh.Gin.Group("/v1/user")
	public.POST("/register", r.uc.CreateUser)
	public.GET("/whoami", r.md.ValidateAccessToken(), r.uc.CurrentUser)
	public.PATCH("/:id", r.md.ValidateAccessToken(), r.uc.UpdateUser)

	private := r.rh.Gin.Group("/v1/admin")
	private.Use(r.md.ValidateAccessToken())
	private.GET("/user/:id", r.uc.GetUserByUUID)
	private.GET("/user", r.uc.GetAllUsers)
	private.DELETE("/user/:id", r.uc.DeactivateUser)
}
