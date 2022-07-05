package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func verify(c *gin.Context, s *server) bool {
	log.Info().Msg("API used")
	tokenBearer := c.Request.Header.Get("Authorization")
	token := strings.TrimPrefix(tokenBearer, "Bearer ")
	if strings.HasPrefix(tokenBearer, "Bearer ") && s.oauth.OAuthResponse.AccessToken == token {
		return true
	} else {
		c.String(http.StatusUnauthorized, "Unauthorized")
		return false
	}
}
