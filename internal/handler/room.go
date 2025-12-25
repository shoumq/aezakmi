package handler

import (
	"aezakmi/internal/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateRoom(c *gin.Context) {
	var req dto.CreateRoomRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.zap.Warnw("invalid body", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body"})
		return
	}

	userID := c.GetString("userID")
	h.zap.Infow("creating room", "userID", userID, "name", req.Name, "isPrivate", req.IsPrivate, "maxParticipants", req.MaxParticipants)

	room, err := h.service.CreateRoom(
		req.Name,
		req.IsPrivate,
		req.MaxParticipants,
		userID,
	)
	if err != nil {
		h.zap.Errorw("failed to create room", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	h.zap.Infow("room created successfully", "roomID", room.ID, "channelName", room.AgoraChannel)
	c.JSON(http.StatusOK, gin.H{
		"room_id":      room.ID,
		"channel_name": room.AgoraChannel,
	})
}
