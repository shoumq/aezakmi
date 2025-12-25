package dto

type RTCTokenRequest struct {
	Channel string `json:"channel"`
	UID     string `json:"uid"`
	Role    string `json:"role"`
}
