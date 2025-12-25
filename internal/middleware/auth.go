package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SimpleAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetHeader("X-User-Id")
		if userID == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "missing X-User-Id",
			})
			return
		}
		c.Set("userID", userID)
		c.Next()
	}
}
