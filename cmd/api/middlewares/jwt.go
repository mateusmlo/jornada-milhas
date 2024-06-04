package middlewares

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mateusmlo/jornada-milhas/config"
	"github.com/mateusmlo/jornada-milhas/utils"
)

type JWTMiddleware struct {
	env *config.Env
	tu  utils.TokenUtils
}

func NewJWTAuthMiddleware(env *config.Env, tu utils.TokenUtils) AuthMiddleware {
	return &JWTMiddleware{
		env: env,
		tu:  tu,
	}
}

func (m *JWTMiddleware) ValidateAccessToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := m.tu.ValidateAccessToken(ctx)
		if err != nil {
			fmt.Println(err)
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized",
			})
			ctx.Abort()

			return
		}

		ctx.Next()
	}
}

func (m *JWTMiddleware) ValidateRefreshToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := m.tu.ValidateRefreshToken(ctx)
		if err != nil {
			fmt.Println(err)
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized",
			})
			ctx.Abort()

			return
		}

		ctx.Next()
	}
}
