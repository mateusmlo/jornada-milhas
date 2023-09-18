package controllers

/*
import (
	"github.com/gin-gonic/gin"
	"github.com/mateusmlo/jornada-milhas/config"
	"github.com/mateusmlo/jornada-milhas/domain"
)

// JWTAuthController struct
type JWTAuthController struct {
	logger      config.Logger
	service     domain.AuthService
	userService domain.UserService
}

// NewJWTAuthController creates new controller
func NewJWTAuthController(
	logger config.Logger,
	service domain.AuthService,
	userService domain.UserService,
) JWTAuthController {
	return JWTAuthController{
		logger:      logger,
		service:     service,
		userService: userService,
	}
}

// SignIn signs in user
func (jwt JWTAuthController) SignIn(c *gin.Context) {
	jwt.logger.Info("SignIn route called")

	//TODO: validate email and password payload
	user, _ := jwt.userService.GetUserByEmail(_)

	token := jwt.service.GenerateJWT(user)
	c.JSON(200, gin.H{
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
*/
