package service

import (
	"aezakmi/internal/repository"

	"go.uber.org/zap"
)

type Service struct {
	repo *repository.Repository
	zap  *zap.SugaredLogger
}

func NewService(repo *repository.Repository, zap *zap.SugaredLogger) *Service {
	return &Service{repo: repo, zap: zap}
}
