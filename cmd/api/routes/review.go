package routes

import (
	"fmt"

	"github.com/mateusmlo/jornada-milhas/cmd/api/controllers"
	"github.com/mateusmlo/jornada-milhas/cmd/api/middlewares"
	"github.com/mateusmlo/jornada-milhas/config"
)

type ReviewRouter struct {
	rh config.RequestHandler
	rc *controllers.ReviewController
	md *middlewares.JWTMiddleware
}

func NewReviewRouter(rc *controllers.ReviewController, rh config.RequestHandler, md *middlewares.JWTMiddleware) *ReviewRouter {
	return &ReviewRouter{
		rh: rh,
		rc: rc,
		md: md,
	}
}

func (r *ReviewRouter) Setup() {
	fmt.Println("\nSetting up review routes...")

	private := r.rh.Gin.Group("/api")
	private.Use(r.md.JwtAuthMiddleware())
	private.POST("/review", r.rc.CreateReview)
	private.GET("/review/user", r.rc.GetUserReviews)
	private.PUT("/review/:id", r.rc.UpdateReview)
	private.DELETE("/review/:id", r.rc.DeleteReview)
}
