package handler

import (
	"aezakmi/internal/model"
	"net/http"
	"time"

	rtctokenbuilder "github.com/AgoraIO/Tools/DynamicKey/AgoraDynamicKey/go/src/rtctokenbuilder2"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GenerateRTCToken(c *gin.Context) {
	var req model.RtcTokenRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body"})
		return
	}

	var role rtctokenbuilder.Role
	switch req.Role {
	case "host":
		role = rtctokenbuilder.RolePublisher
	case "audience":
		role = rtctokenbuilder.RoleSubscriber
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid role"})
		return
	}

	expire := uint32(time.Now().Unix() + int64(h.cfg.TokenTTL))

	token, err := rtctokenbuilder.BuildTokenWithUserAccount(
		h.cfg.AgoraAppID,
		h.cfg.AgoraAppCert,
		req.Channel,
		req.UID,
		role,
		expire,
		expire,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err = h.service.LogToken(req.UID, req.Channel, "RTC", c.ClientIP(), h.cfg.TokenTTL); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{
		"token":      token,
		"expires_in": h.cfg.TokenTTL,
	})
}
