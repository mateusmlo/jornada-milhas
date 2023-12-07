package controllers

import "github.com/gin-gonic/gin"

// AuthController provides authentication resources
type AuthController interface {
	SignIn(ctx *gin.Context)
	Logout(ctx *gin.Context)
	RenewRefreshToken(ctx *gin.Context)
}

// UserController provides user resources
type UserController interface {
	CurrentUser(ctx *gin.Context)
	GetAllUsers(ctx *gin.Context)
	GetUserByUUID(ctx *gin.Context)
	CreateUser(ctx *gin.Context)
	UpdateUser(ctx *gin.Context)
	DeactivateUser(ctx *gin.Context)
}

// ReviewController provides location review resources
type ReviewController interface {
	CreateReview(ctx *gin.Context)
	GetReviewByUUID(ctx *gin.Context)
	GetUserReviews(ctx *gin.Context)
	UpdateReview(ctx *gin.Context)
	DeleteReview(ctx *gin.Context)
}
