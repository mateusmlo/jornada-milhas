package routes

import (
	"fmt"

	"github.com/mateusmlo/jornada-milhas/cmd/api/controllers"
	"github.com/mateusmlo/jornada-milhas/cmd/api/middlewares"
	"github.com/mateusmlo/jornada-milhas/config"
)

type ReviewRouter struct {
	rh *config.RequestHandler
	rc controllers.ReviewController
	md middlewares.AuthMiddleware
}

func NewReviewRouter(
	rc controllers.ReviewController,
	rh *config.RequestHandler,
	md middlewares.AuthMiddleware,
) *ReviewRouter {
	return &ReviewRouter{
		rh: rh,
		rc: rc,
		md: md,
	}
}

func (r *ReviewRouter) SetupRoutes() {
	fmt.Println("\nSetting up review routes...")

	private := r.rh.Gin.Group("/v1/review")
	private.Use(r.md.ValidateAccessToken())
	private.POST("/", r.rc.CreateReview)
	private.GET("/user", r.rc.GetUserReviews)
	private.PUT("/:id", r.rc.UpdateReview)
	private.DELETE("/:id", r.rc.DeleteReview)
}
