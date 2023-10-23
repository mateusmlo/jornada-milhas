package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mateusmlo/jornada-milhas/config"
	service "github.com/mateusmlo/jornada-milhas/domain/services"
	"github.com/mateusmlo/jornada-milhas/internal/dto"
	"github.com/mateusmlo/jornada-milhas/tools"
)

type ReviewController struct {
	svc    *service.ReviewService
	us     *service.UserService
	logger config.GinLogger
}

func NewReviewController(reviewService *service.ReviewService, us *service.UserService, logger config.GinLogger) *ReviewController {
	return &ReviewController{
		svc:    reviewService,
		us:     us,
		logger: logger,
	}
}

func (rc *ReviewController) CreateReview(ctx *gin.Context) {
	var reviewPayload *dto.NewReviewDTO

	userID, err := tools.ExtractTokenSub(ctx)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": err,
		})
	}

	user, _ := rc.us.GetUserByUUID(userID.String())

	if err := ctx.BindJSON(&reviewPayload); err != nil {
		fmt.Println(err)
		rc.logger.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})

		return
	}

	reviewPayload.UserID = user.BaseModel.ID

	if err := rc.svc.CreateReview(reviewPayload); err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})

		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"data": reviewPayload,
	})
}
