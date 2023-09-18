package config

import "github.com/gin-gonic/gin"

type RequestHandler struct {
	Gin *gin.Engine
}

// NewRequestHandler creates new req handler
func NewRequestHandler(logger Logger) RequestHandler {
	gin.DefaultWriter = logger.GetGinLogger()
	eng := gin.New()

	return RequestHandler{Gin: eng}
}
