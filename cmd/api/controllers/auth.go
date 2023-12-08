package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	service "github.com/mateusmlo/jornada-milhas/domain/services"
	"github.com/mateusmlo/jornada-milhas/internal/dto"
	"github.com/mateusmlo/jornada-milhas/tools"
)

type authController struct {
	as service.AuthService
	rs service.RefreshService
	us service.UserService
	tu tools.TokenUtils
}

// NewAuthController creates new controller
func NewAuthController(
	as service.AuthService,
	rs service.RefreshService,
	tu tools.TokenUtils,
	us service.UserService,
) AuthController {
	return &authController{
		as: as,
		rs: rs,
		tu: tu,
		us: us,
	}
}

// SignIn signs in user
func (ac *authController) SignIn(ctx *gin.Context) {
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

// Logout logs out user
func (ac *authController) Logout(ctx *gin.Context) {
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

// RenewRefreshToken renews user token pair if refresh still valid
func (ac *authController) RenewTokenPair(ctx *gin.Context) {
	sub, err := ac.tu.ExtractTokenSub(ctx, true)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})

		return
	}

	ac.rs.DeleteRefreshToken(sub.String())

	newAccessTkn, newRefreshTkn, err := ac.as.GenerateTokenPair(sub)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})

		return
	}

	err = ac.rs.SetRefreshToken(newRefreshTkn, sub.String())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})

		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{
		"message":       "token renewed succesfully",
		"access_token":  newAccessTkn,
		"refresh_token": newRefreshTkn,
	})
}
