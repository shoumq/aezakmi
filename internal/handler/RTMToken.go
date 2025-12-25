package handler

import (
	"aezakmi/internal/model"
	"net/http"
	"time"

	rtmtokenbuilder "github.com/AgoraIO/Tools/DynamicKey/AgoraDynamicKey/go/src/rtmtokenbuilder2"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GenerateRTMToken(c *gin.Context) {
	var req model.RtmTokenRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid body"})
		return
	}

	expire := uint32(time.Now().Unix() + int64(h.cfg.TokenTTL))

	token, err := rtmtokenbuilder.BuildToken(
		h.cfg.AgoraAppID,
		h.cfg.AgoraAppCert,
		req.UID,
		expire,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err = h.service.LogToken(req.UID, "", "RTC", c.ClientIP(), h.cfg.TokenTTL); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{
		"token":      token,
		"expires_in": h.cfg.TokenTTL,
	})
}
