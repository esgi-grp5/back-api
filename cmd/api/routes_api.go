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
			log.Err(err).Msg("Error in Login")
			c.JSON(http.StatusBadRequest, gin.H{"error": "Error with User API"})
			return
		}

		// Verify information with database
		userDB, err := s.db.SelectUser(userRequest.Mail)
		if err != nil {
			log.Err(err).Msg("Error in Login")
			c.JSON(http.StatusBadRequest, gin.H{"error": "Error with User API"})
			return
		}

		// Verify information with request and database
		if userDB.Password == userRequest.Password {
			res := map[string]interface{}{
				"id":       userDB.ID,
				"username": userDB.Username,
				"mail":     userDB.Mail,
				"password": userDB.Password,
			}
			c.JSON(http.StatusOK, res)
		} else {
			log.Warn().Str("User request", userRequest.Mail).Msg("Invalid password")
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid password"})
		}
	}
}
