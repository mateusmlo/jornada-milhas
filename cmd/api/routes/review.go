package routes

import (
	"github.com/mateusmlo/jornada-milhas/cmd/api/controllers"
	"github.com/mateusmlo/jornada-milhas/cmd/api/middlewares"
	"github.com/mateusmlo/jornada-milhas/config"
)

type ReviewRouter struct {
	rh     config.RequestHandler
	rc     *controllers.ReviewController
	md     *middlewares.JWTMiddleware
	logger config.GinLogger
}

func NewReviewRouter(rc *controllers.ReviewController, logger config.GinLogger, rh config.RequestHandler, md *middlewares.JWTMiddleware) *ReviewRouter {
	return &ReviewRouter{
		rh:     rh,
		rc:     rc,
		logger: logger,
		md:     md,
	}
}

func (r *ReviewRouter) Setup() {
	r.logger.Info("Setting up review routes...")

	private := r.rh.Gin.Group("/api")
	private.Use(r.md.JwtAuthMiddleware())
	private.POST("/review", r.rc.CreateReview)
}
