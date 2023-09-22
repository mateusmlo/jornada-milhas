package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mateusmlo/jornada-milhas/config"
	"github.com/mateusmlo/jornada-milhas/tools"
)

type IJWTMiddleware interface {
	ValidateToken(ctx *gin.Context) error
	ExtractTokenSub(ctx *gin.Context) (uuid.UUID, error)
}

type JWTMiddleware struct {
	env    config.Env
	logger config.GinLogger
}

func NewJWTAuthMiddleware(env config.Env, logger config.GinLogger) *JWTMiddleware {
	return &JWTMiddleware{
		env:    env,
		logger: logger,
	}
}

func (m *JWTMiddleware) JwtAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := tools.ValidateToken(ctx)
		if err != nil {
			m.logger.Error(err)
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized",
			})
			ctx.Abort()

			return
		}

		ctx.Next()
	}
}
