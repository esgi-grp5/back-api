package main

import (
	"go-micro/internal/config"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func (s *server) OAuth(c *gin.Context) {
	// Init var
	var (
		oauthReq  config.OAuthRequest
		oauthResp config.OAuthResponse
	)

	// Get JSON body
	if err := c.ShouldBindJSON(&oauthReq); err != nil {
		log.Err(err).Msg("Error in OAuth")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error with User API"})
		return
	}

	tokenAPI, err := s.db.SelectAuth(oauthReq.Secret)
	if err != nil {
		log.Err(err).Msg("Error in OAuth")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error with User API"})
		return
	}

	// Verify information with server
	if tokenAPI.ClientName != oauthReq.ID || tokenAPI.Token != oauthReq.Secret {
		log.Warn().Interface("OAuth request", oauthReq).Msg("Invalid ID or Secret")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID or Secret"})
		return
	}

	// Generate OAuthResponse
	oauthResp.AccessToken = RandomString(20)
	oauthResp.TokenType = "bearer_token"

	// Put on server
	s.oauth.OAuthResponse.AccessToken = oauthResp.AccessToken
	s.oauth.OAuthResponse.TokenType = oauthResp.TokenType

	// Return JSON reposne
	res := map[string]interface{}{
		"access_token": oauthResp.AccessToken,
		"token_type":   oauthResp.TokenType,
	}
	c.JSON(http.StatusOK, res)
}
