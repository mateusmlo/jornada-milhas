package routes

import (
	"fmt"

	"github.com/mateusmlo/jornada-milhas/cmd/api/controllers"
	"github.com/mateusmlo/jornada-milhas/cmd/api/middlewares"
	"github.com/mateusmlo/jornada-milhas/config"
)

type UserRouter struct {
	rh *config.RequestHandler
	uc *controllers.UserController
	md *middlewares.JWTMiddleware
}

func NewUserRouter(uc *controllers.UserController, rh *config.RequestHandler, md *middlewares.JWTMiddleware) *UserRouter {
	return &UserRouter{
		uc: uc,
		rh: rh,
		md: md,
	}
}

func (r *UserRouter) Setup() {
	fmt.Println("\nSetting up user routes...")

	public := r.rh.Gin.Group("/api")
	public.POST("/user", r.uc.CreateUser)
	public.GET("/user/whoami", r.md.ValidateAccessToken(), r.uc.CurrentUser)
	public.PATCH("/user/:id", r.md.ValidateAccessToken(), r.uc.UpdateUser)

	private := r.rh.Gin.Group("/api/admin")
	private.Use(r.md.ValidateAccessToken())
	private.GET("/user/:id", r.uc.GetUserByUUID)
	private.GET("/user", r.uc.GetAllUsers)
	private.DELETE("/user/:id", r.uc.DeactivateUser)

}
