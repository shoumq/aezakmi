package service

import (
	"aezakmi/internal/model"
	"time"

	"github.com/google/uuid"
)

func (s *Service) LogToken(
	userID, channel, tokenType, ip string,
	ttl int,
) error {
	return s.repo.CreateToken(&model.TokenLog{
		ID:        uuid.New().String(),
		UserID:    userID,
		Channel:   channel,
		TokenType: tokenType,
		IssuedAt:  time.Now(),
		ExpiresAt: time.Now().Add(time.Second * time.Duration(ttl)),
		IPAddress: ip,
	})
}
