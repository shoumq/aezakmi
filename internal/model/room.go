package model

import "time"

type Room struct {
	ID              string    `gorm:"primaryKey" json:"room_id"`
	Name            string    `json:"name"`
	AgoraChannel    string    `json:"agora_channel_name"`
	IsPrivate       bool      `json:"is_private"`
	MaxParticipants int       `json:"max_participants"`
	CreatedBy       string    `json:"created_by"`
	CreatedAt       time.Time `json:"created_at"`
}
