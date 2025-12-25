package dto

type CreateRoomRequest struct {
	Name            string `json:"name"`
	IsPrivate       bool   `json:"is_private"`
	MaxParticipants int    `json:"max_participants"`
}
