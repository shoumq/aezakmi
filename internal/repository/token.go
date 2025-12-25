package repository

import (
	"aezakmi/internal/model"
)

func (r *Repository) CreateToken(log *model.TokenLog) error {
	return r.db.Create(log).Error
}
