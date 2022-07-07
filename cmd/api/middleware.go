package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func (s *server) OAuthMiddleware(c *gin.Context) {
	log.Info().Msg("API used")
	tokenBearer := c.Request.Header.Get("Authorization")
	token := strings.TrimPrefix(tokenBearer, "Bearer ")
	if strings.HasPrefix(tokenBearer, "Bearer ") && s.oauth.OAuthResponse.AccessToken == token {
		c.Next()
	} else {
		c.String(http.StatusUnauthorized, "Unauthorized")
		c.Abort()
	}
}
