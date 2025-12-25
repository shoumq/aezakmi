package handler

import (
	"aezakmi/internal/config"
	"aezakmi/internal/service"

	"go.uber.org/zap"
)

type Handler struct {
	service *service.Service
	cfg     *config.Config
	zap     *zap.SugaredLogger
}

func NewHandler(s *service.Service, cfg *config.Config, zap *zap.SugaredLogger) *Handler {
	return &Handler{service: s, cfg: cfg, zap: zap}
}
