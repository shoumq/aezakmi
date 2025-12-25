package repository

import (
	"aezakmi/internal/config"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Repository struct {
	db  *gorm.DB
	cfg *config.Config
	zap *zap.SugaredLogger
}

func NewRepository(db *gorm.DB, cfg *config.Config, zap *zap.SugaredLogger) *Repository {
	return &Repository{db: db, cfg: cfg, zap: zap}
}
