package controllers

import (
	"net/http"

	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
		uc.logger.Error(err)
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

	id, err := uuid.Parse(paramID)
	if err != nil {
		uc.logger.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	user, err := uc.service.GetUserByUUID(id)

	if err != nil {
		uc.logger.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
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
	var userPayload models.User

	if err := ctx.BindJSON(userPayload); err != nil {
		uc.logger.Error(err)
		spew.Dump(userPayload)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})

		return
	}

	if err := uc.service.CreateUser(userPayload); err != nil {
		uc.logger.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})

		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"data": userPayload,
	})
}
