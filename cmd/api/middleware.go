package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (s *server) OAuthMiddleware(c *gin.Context) {
	const prefix = "Bearer "
	header := c.GetHeader("Authorization")
	token := header[len(prefix):]

	if header == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Authorization header missing or empty"})
	}

	if !strings.HasPrefix(header, "Bearer ") {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Bearer prefix missing or mismatch with RFC 6750"})
	}

	if s.oauth.OAuthResponse.AccessToken != token {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
	}

	c.Next()
}
