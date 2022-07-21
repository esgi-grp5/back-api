package main

import (
	"go-micro/internal/database"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func (s *server) GetSerieWishList(c *gin.Context) {
	var userRequest User
	// Get JSON body
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		log.Err(err).Msg("Error in Login")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error with User API"})
		return
	}
	userID, err := strconv.Atoi(userRequest.ID)
	if err != nil {
		log.Err(err).Msg("Error in GetSerieWishList when Atoi")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error with User API"})
		return
	}
	// Get serie wishlist
	wishlist, err := s.db.SelectSerieWishList(userID)
	if err != nil {
		log.Err(err).Msg("Error in GetSerieWishList")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error with User API"})
		return
	}
	// Return serie wishlist
	c.JSON(http.StatusOK, wishlist)
}

func (s *server) AddSerieWishList(c *gin.Context) {
	var serieRequest database.Serie
	// Get JSON body
	if err := c.ShouldBindJSON(&serieRequest); err != nil {
		log.Err(err).Msg("Error in Login")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error with User API"})
		return
	}
	wishlist, err := s.db.SelectSerieWishList(serieRequest.UsernameID)
	if err != nil {
		log.Err(err).Msg("Error in AddSerieWishList")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error with User API"})
		return
	}
	for _, w := range wishlist {
		if w.SerieID == serieRequest.SerieID {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Serie already in wishlist"})
			return
		}
	}
	// Add serie to wishlist
	if err = s.db.InsertSerieWishList(serieRequest.UsernameID, serieRequest.SerieID); err != nil {
		log.Err(err).Msg("Error in AddSerieWishList")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error with User API"})
		return
	}
	// Return serie wishlist
	c.JSON(http.StatusOK, "success")
}

func (s *server) DeleteSerieWishList(c *gin.Context) {
	var serieRequest database.Serie
	// Get JSON body
	if err := c.ShouldBindJSON(&serieRequest); err != nil {
		log.Err(err).Msg("Error in Login")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error with User API"})
		return
	}
	// Delete serie from wishlist
	err := s.db.DeleteSerieWishList(serieRequest.UsernameID, serieRequest.SerieID)
	if err != nil {
		log.Err(err).Msg("Error in DeleteSerieWishList")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error with User API"})
		return
	}
	// Return serie wishlist
	c.JSON(http.StatusOK, "success")
}

func (s *server) GetSerieCount(c *gin.Context) {
	var serieRequest Serie
	// Get JSON body
	if err := c.ShouldBindJSON(&serieRequest); err != nil {
		log.Err(err).Msg("Error in Login")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error with User API"})
		return
	}
	serieID, err := strconv.Atoi(serieRequest.ID)
	if err != nil {
		log.Err(err).Msg("Error in GetSerieCount when Atoi")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error with User API"})
		return
	}
	// Get serie count
	count, err := s.db.SelectSerieCount(serieID)
	if err != nil {
		log.Err(err).Msg("Error in GetSerieCount")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error with User API"})
		return
	}
	res := map[string]interface{}{
		"count": count,
	}
	// Return serie count
	c.JSON(http.StatusOK, res)
}
