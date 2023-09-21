package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mateusmlo/jornada-milhas/config"
	"github.com/mateusmlo/jornada-milhas/domain"
	"github.com/mateusmlo/jornada-milhas/internal/dto"
)

// JWTAuthController struct
type JWTAuthController struct {
	logger      config.Logger
	authService *domain.AuthService
	userService *domain.UserService
}

// NewJWTAuthController creates new controller
func NewJWTAuthController(
	logger config.Logger,
	authService *domain.AuthService,
	userService *domain.UserService,
) JWTAuthController {
	return JWTAuthController{
		logger:      logger,
		authService: authService,
		userService: userService,
	}
}

// SignIn signs in user
func (jwt JWTAuthController) SignIn(ctx *gin.Context) {
	var payload dto.AuthDTO

	if err := ctx.BindJSON(&payload); err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})

		return
	}

	user, err := jwt.userService.FindByEmail(payload.Email)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err,
		})

		return
	}

	token, err := jwt.authService.CreateSession(payload, user)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": err,
		})

		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{
		"message": "logged in successfully",
		"token":   token,
	})
}

// Register registers user
func (jwt JWTAuthController) Register(c *gin.Context) {
	jwt.logger.Info("Register route called")
	c.JSON(200, gin.H{
		"message": "register route",
	})
}
