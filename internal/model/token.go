package model

type RtcTokenRequest struct {
	Channel string `json:"channel"`
	UID     string `json:"uid"`
	Role    string `json:"role"`
}

type RtmTokenRequest struct {
	UID string `json:"uid"`
}
