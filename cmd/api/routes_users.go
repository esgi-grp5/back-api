package main

import (
	"go-micro/internal/database"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func (s *server) Login(c *gin.Context) {
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

func (s *server) Register(c *gin.Context) {
	var userRequest database.User
	// Get JSON body
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		log.Err(err).Msg("Error in Login")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error with User API"})
		return
	}

	// Verify information with database
	if checkMail(userRequest.Mail, s) {
		log.Warn().Str("User request", userRequest.Mail).Msg("Mail already exist")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Mail already exist"})
		return
	}
	if checkUsername(userRequest.Username, s) {
		log.Warn().Str("User request", userRequest.Mail).Msg("Username already exist")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username already exist"})
		return
	}

	// Put on database
	err := s.db.InsertUser(userRequest)
	if err != nil {
		log.Err(err).Msg("Error in Login")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error with User API"})
		return
	}

	// Return JSON reposne
	res := map[string]interface{}{
		"data": "success",
	}
	c.JSON(http.StatusOK, res)
}

func checkMail(mail string, s *server) bool {
	user, err := s.db.GetMailFromUser(mail)
	if err != nil {
		log.Err(err).Msg("")
	}
	return user.Mail == mail
}

func checkUsername(username string, s *server) bool {
	user, err := s.db.GetUsernameFromUser(username)
	if err != nil {
		log.Err(err).Msg("")
	}
	return user.Username == username
}
