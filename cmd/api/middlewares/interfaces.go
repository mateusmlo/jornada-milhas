package middlewares

import "github.com/gin-gonic/gin"

type AuthMiddleware interface {
	ValidateAccessToken() gin.HandlerFunc
	ValidateRefreshToken() gin.HandlerFunc
}
