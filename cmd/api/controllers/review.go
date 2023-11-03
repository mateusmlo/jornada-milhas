package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	service "github.com/mateusmlo/jornada-milhas/domain/services"
	"github.com/mateusmlo/jornada-milhas/internal/dto"
	"github.com/mateusmlo/jornada-milhas/tools"
)

type ReviewController struct {
	svc *service.ReviewService
	us  *service.UserService
}

func NewReviewController(reviewService *service.ReviewService, us *service.UserService) *ReviewController {
	return &ReviewController{
		svc: reviewService,
		us:  us,
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

	if err := ctx.BindJSON(&reviewPayload); err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})

		return
	}

	if err := rc.svc.CreateReview(reviewPayload, userID); err != nil {
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

func (rc *ReviewController) GetReviewByUUID(ctx *gin.Context) {
	reviewID := ctx.Param("id")
	userID, err := tools.ExtractTokenSub(ctx)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": err,
		})
	}

	r, err := rc.svc.FindByUUID(reviewID, userID)
	if err != nil {
		fmt.Println(err)

		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusFound, gin.H{
		"data": &r,
	})
}

func (rc *ReviewController) GetUserReviews(ctx *gin.Context) {
	userID, err := tools.ExtractTokenSub(ctx)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": err,
		})
	}

	revs, err := rc.svc.GetUserReviews(userID)
	if err != nil {
		fmt.Println(err)

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusFound, gin.H{
		"data": revs,
	})
}

func (rc *ReviewController) UpdateReview(ctx *gin.Context) {
	var reviewPayload *dto.UpdateReviewDTO

	userID, err := tools.ExtractTokenSub(ctx)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": err,
		})
	}

	if err := ctx.BindJSON(&reviewPayload); err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})

		return
	}

	if err := rc.svc.UpdateReview(reviewPayload.ID.String(), *reviewPayload, userID); err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})

		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{
		"data": 1,
	})
}

func (rc *ReviewController) DeleteReview(ctx *gin.Context) {
	reviewID := ctx.Param("id")
	userID, err := tools.ExtractTokenSub(ctx)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": err,
		})
	}

	res, err := rc.svc.DeleteReview(reviewID, userID)
	if err != nil {
		fmt.Println(err)

		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{
		"data": res,
	})
}
