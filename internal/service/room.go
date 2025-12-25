package service

import (
	"aezakmi/internal/model"
	"time"

	"github.com/google/uuid"
)

func (s *Service) CreateRoom(
	name string,
	isPrivate bool,
	max int,
	userID string,
) (*model.Room, error) {

	room := &model.Room{
		ID:              uuid.New().String(),
		Name:            name,
		AgoraChannel:    "agora_" + uuid.New().String(),
		IsPrivate:       isPrivate,
		MaxParticipants: max,
		CreatedBy:       userID,
		CreatedAt:       time.Now(),
	}

	return room, s.repo.CreateRoom(room)
}
