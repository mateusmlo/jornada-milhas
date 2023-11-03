package config

import "github.com/gin-gonic/gin"

type RequestHandler struct {
	Gin *gin.Engine
}

// NewRequestHandler creates new req handler
func NewRequestHandler() RequestHandler {
	eng := gin.New()
	eng.Use(gin.Logger(), gin.Recovery())

	return RequestHandler{Gin: eng}
}
