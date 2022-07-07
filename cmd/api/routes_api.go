package main

import (
	"go-micro/internal/database"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

/* Access without OAuth */

func (s *server) HealthCheck(c *gin.Context) {
	res := map[string]interface{}{
		"data": "Server is up and running",
	}

	c.JSON(http.StatusOK, res)
}

/* Access with OAuth */

func (s *server) FlutterAccess(c *gin.Context) {
	if verify(c, s) {
		res := map[string]interface{}{
			"data": "Hello world!",
		}
		c.JSON(http.StatusOK, res)
	}
}

func (s *server) Login(c *gin.Context) {
	if verify(c, s) {
		var userRequest database.User
		// Get JSON body
		if err := c.ShouldBindJSON(&userRequest); err != nil {
			log.Err(err).Msg("Error in OAuth")
			c.JSON(http.StatusBadRequest, gin.H{"error": "Error with User API"})
			return
		}

		user, err := s.db.SelectUser(userRequest.Mail)
		if err != nil {
			log.Err(err).Msg("Error in OAuth")
			c.JSON(http.StatusBadRequest, gin.H{"error": "Error with User API"})
			return
		}

		if user.Password == userRequest.Password {
			res := map[string]interface{}{
				"id":       user.ID,
				"username": user.Username,
				"mail":     user.Mail,
				"password": user.Password,
			}
			c.JSON(http.StatusOK, res)
		} else {
			log.Warn().Interface("User request", userRequest).Msg("Invalid password")
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid password"})
		}
	}
}
