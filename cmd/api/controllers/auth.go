package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	service "github.com/mateusmlo/jornada-milhas/domain/services"
	"github.com/mateusmlo/jornada-milhas/internal/dto"
	"github.com/mateusmlo/jornada-milhas/tools"
)

// AuthController struct
type AuthController struct {
	as *service.AuthService
	rs *service.RefreshService
	us *service.UserService
	tu *tools.TokenUtils
}

// NewAuthController creates new controller
func NewAuthController(
	as *service.AuthService,
	rs *service.RefreshService,
	tu *tools.TokenUtils,
	us *service.UserService,
) *AuthController {
	return &AuthController{
		as: as,
		rs: rs,
		tu: tu,
		us: us,
	}
}

// SignIn signs in user
func (ac *AuthController) SignIn(ctx *gin.Context) {
	var payload dto.AuthDTO

	if err := ctx.BindJSON(&payload); err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})

		return
	}

	user, err := ac.us.FindByEmail(payload.Email)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err,
		})

		return
	}

	token, err := ac.as.CreateSession(payload, user)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": err,
		})

		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{
		"message": "logged in successfully",
		"data":    token,
	})
}

func (ac *AuthController) Logout(ctx *gin.Context) {
	sub, err := ac.tu.ExtractTokenSub(ctx, true)
	if err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"error": err,
		})

		return
	}

	res := ac.rs.DeleteRefreshToken(sub.String())

	ctx.JSON(http.StatusAccepted, gin.H{
		"message": "logged out succesfully",
		"status":  res,
	})
}

func (ac *AuthController) RenewRefreshToken(ctx *gin.Context) {
	sub, err := ac.tu.ExtractTokenSub(ctx, true)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})

		return
	}

	ac.rs.DeleteRefreshToken(sub.String())

	newTkn, err := ac.tu.GenerateRefreshToken(sub)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})

		return
	}

	err = ac.rs.SetRefreshToken(newTkn, sub.String())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})

		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{
		"message":       "refresh token renewed succesfully",
		"refresh_token": newTkn,
	})
}
