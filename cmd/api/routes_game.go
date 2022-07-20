package main

import (
	"go-micro/internal/database"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func (s *server) GetGameWishList(c *gin.Context) {
	var userRequest database.User
	// Get JSON body
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		log.Err(err).Msg("Error in Login")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error with User API"})
		return
	}
	// Get game wishlist
	wishlist, err := s.db.SelectGameWishList(userRequest.ID)
	if err != nil {
		log.Err(err).Msg("Error in GetGameWishList")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error with User API"})
		return
	}
	// Return game wishlist
	c.JSON(http.StatusOK, wishlist)
}

func (s *server) AddGameWishList(c *gin.Context) {
	var gameRequest database.Game
	// Get JSON body
	if err := c.ShouldBindJSON(&gameRequest); err != nil {
		log.Err(err).Msg("Error in Login")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error with User API"})
		return
	}
	wishlist, err := s.db.SelectGameWishList(gameRequest.UsernameID)
	if err != nil {
		log.Err(err).Msg("Error in AddGameWishList")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error with User API"})
		return
	}
	for _, w := range wishlist {
		if w.GameID == gameRequest.GameID {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Game already in wishlist"})
			return
		}
	}
	// Add game to wishlist
	if err = s.db.InsertGameWishList(gameRequest.UsernameID, gameRequest.GameID); err != nil {
		log.Err(err).Msg("Error in AddGameWishList")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error with User API"})
		return
	}
	// Return game wishlist
	c.JSON(http.StatusOK, "success")
}

func (s *server) DeleteGameWishList(c *gin.Context) {
	var gameRequest database.Game
	// Get JSON body
	if err := c.ShouldBindJSON(&gameRequest); err != nil {
		log.Err(err).Msg("Error in Login")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error with User API"})
		return
	}
	// Delete game from wishlist
	err := s.db.DeleteGameWishList(gameRequest.UsernameID, gameRequest.GameID)
	if err != nil {
		log.Err(err).Msg("Error in DeleteGameWishList")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error with User API"})
		return
	}
	// Return game wishlist
	c.JSON(http.StatusOK, "success")
}

func (s *server) GetGameCount(c *gin.Context) {
	var gameRequest database.Game
	// Get JSON body
	if err := c.ShouldBindJSON(&gameRequest); err != nil {
		log.Err(err).Msg("Error in Login")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error with User API"})
		return
	}
	// Get game count
	count, err := s.db.SelectGameCount(gameRequest.GameID)
	if err != nil {
		log.Err(err).Msg("Error in GetGameCount")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error with User API"})
		return
	}
	res := map[string]interface{}{
		"count": count,
	}
	// Return game count
	c.JSON(http.StatusOK, res)
}
