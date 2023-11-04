package config

import (
	"go.uber.org/zap"
)

type EchoHandler struct {
	log *zap.Logger
}

func NewEchoHandler(log *zap.Logger) *EchoHandler {
	return &EchoHandler{
		log: log,
	}
}
