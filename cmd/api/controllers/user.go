package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	service "github.com/mateusmlo/jornada-milhas/domain/services"
	"github.com/mateusmlo/jornada-milhas/internal/dto"
	"github.com/mateusmlo/jornada-milhas/tools"
)

type userController struct {
	svc service.UserService
	ut  tools.TokenUtils
}

// NewUserController instantiates new user controller
func NewUserController(userService service.UserService, ut tools.TokenUtils) UserController {
	return &userController{
		svc: userService,
		ut:  ut,
	}
}

// CurrentUser returns informatio on current loged user
func (uc *userController) CurrentUser(ctx *gin.Context) {
	sub, err := uc.ut.ExtractTokenSub(ctx, false)
	if err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"error": err,
		})

		return
	}

	user, err := uc.svc.GetUserByUUID(sub.String())
	if err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{
			"error": err,
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": user})
}

// GetAllUsers returns all registered users
func (uc *userController) GetAllUsers(ctx *gin.Context) {
	users, err := uc.svc.GetAllUsers()

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
func (uc *userController) GetUserByUUID(ctx *gin.Context) {
	paramID := ctx.Param("id")

	user, err := uc.svc.GetUserByUUID(paramID)
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
func (uc *userController) CreateUser(ctx *gin.Context) {
	var userPayload *dto.NewUserDTO

	if err := ctx.BindJSON(&userPayload); err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})

		return
	}

	if err := uc.svc.CreateUser(userPayload); err != nil {
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
func (uc *userController) UpdateUser(ctx *gin.Context) {
	var updateUserPayload dto.UpdateUserDTO

	paramID := ctx.Param("id")
	if err := ctx.BindJSON(&updateUserPayload); err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})

		return
	}

	err := uc.svc.UpdateUser(paramID, updateUserPayload)
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
func (uc *userController) DeactivateUser(ctx *gin.Context) {
	paramID := ctx.Param("id")

	res, err := uc.svc.DeactivateUser(paramID)
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
