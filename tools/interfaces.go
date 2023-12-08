package tools

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// TokenUtils provides JWT utility functions such as validating, generating and extracting tokens from context
type TokenUtils interface {
	GenerateAccessToken(userID uuid.UUID) (string, error)
	GenerateRefreshToken(userID uuid.UUID) (string, error)
	ValidateAccessToken(ctx *gin.Context) error
	ValidateRefreshToken(ctx *gin.Context) error
	ExtractToken(ctx *gin.Context) string
	ExtractTokenSub(ctx *gin.Context, isRefresh bool) (uuid.UUID, error)
}
