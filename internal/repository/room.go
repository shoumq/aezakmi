package repository

import (
	"aezakmi/internal/model"
)

func (r *Repository) CreateRoom(room *model.Room) error {
	return r.db.Create(room).Error
}
