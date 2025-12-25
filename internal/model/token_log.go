package model

import "time"

type TokenLog struct {
	ID        string `gorm:"primaryKey"`
	UserID    string
	Channel   string
	TokenType string
	IssuedAt  time.Time
	ExpiresAt time.Time
	IPAddress string
}
