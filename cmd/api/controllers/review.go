package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	services "github.com/mateusmlo/jornada-milhas/domain/services"
	"github.com/mateusmlo/jornada-milhas/internal/dto"
	"github.com/mateusmlo/jornada-milhas/tools"
)

type reviewController struct {
	svc services.ReviewService
	us  services.UserService
	tu  tools.TokenUtils
}

// NewReviewController provides new review controller struct
func NewReviewController(rs services.ReviewService, tu tools.TokenUtils, us services.UserService) ReviewController {
	return &reviewController{
		svc: rs,
		tu:  tu,
		us:  us,
	}
}

// CreateReview creates a new location review
func (rc *reviewController) CreateReview(ctx *gin.Context) {
	var reviewPayload *dto.NewReviewDTO

	userID, err := rc.tu.ExtractTokenSub(ctx, false)
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

// GetReviewByUUID attempts to get a review by ID
func (rc *reviewController) GetReviewByUUID(ctx *gin.Context) {
	reviewID := ctx.Param("id")
	userID, err := rc.tu.ExtractTokenSub(ctx, false)
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

// GetUserReviews gets all reviews from a user
func (rc *reviewController) GetUserReviews(ctx *gin.Context) {
	userID, err := rc.tu.ExtractTokenSub(ctx, false)
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

// UpdateReview updates a review from a user
func (rc *reviewController) UpdateReview(ctx *gin.Context) {
	var reviewPayload *dto.UpdateReviewDTO

	userID, err := rc.tu.ExtractTokenSub(ctx, false)
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

// DeleteReview deletes a review by ID
func (rc *reviewController) DeleteReview(ctx *gin.Context) {
	reviewID := ctx.Param("id")
	userID, err := rc.tu.ExtractTokenSub(ctx, false)
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
