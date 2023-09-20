package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mateusmlo/jornada-milhas/config"
	"github.com/mateusmlo/jornada-milhas/domain"
	"github.com/mateusmlo/jornada-milhas/internal/models"
)

// UserController data
type UserController struct {
	service *domain.UserService
	logger  config.Logger
}

// NewUserController instantiates new user controller
func NewUserController(userService *domain.UserService, logger config.Logger) *UserController {
	return &UserController{
		service: userService,
		logger:  logger,
	}
}

// GetAllUsers returns all registered users
func (uc *UserController) GetAllUsers(ctx *gin.Context) {
	users, err := uc.service.GetAllUsers()

	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	}

	ctx.JSON(http.StatusFound, gin.H{
		"data": &users,
	})
}

// GetUserByUUID gets one user by UUID
func (uc *UserController) GetUserByUUID(ctx *gin.Context) {
	paramID := ctx.Param("id")

	user, err := uc.service.GetUserByUUID(paramID)
	if err != nil && user == nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"data": &user,
	})
}

// CreateUser creates new user
func (uc *UserController) CreateUser(ctx *gin.Context) {
	var userPayload *models.User

	if err := ctx.BindJSON(&userPayload); err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})

		return
	}

	if err := uc.service.CreateUser(userPayload); err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})

		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"data": userPayload,
	})
}

// UpdateUser updates user info
func (uc *UserController) UpdateUser(ctx *gin.Context) {
	var updateUserPayload models.UpdateUserDTO

	paramID := ctx.Param("id")
	if err := ctx.BindJSON(&updateUserPayload); err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})

		return
	}

	err := uc.service.UpdateUser(paramID, updateUserPayload)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})

		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{
		"data": 1,
	})
}

// DeactivateUser deactivates a user
func (uc *UserController) DeactivateUser(ctx *gin.Context) {
	paramID := ctx.Param("id")

	res, err := uc.service.DeactivateUser(paramID)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})

		return

	}

	ctx.JSON(http.StatusAccepted, gin.H{
		"data": res,
	})
}
