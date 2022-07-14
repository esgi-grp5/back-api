package main

import (
	"go-micro/internal/database"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func (s *server) GetSerieWishList(c *gin.Context) {
	var userRequest database.User
	// Get JSON body
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		log.Err(err).Msg("Error in Login")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error with User API"})
		return
	}
	// Get movie wishlist
	wishlist, err := s.db.SelectSerieWishList(userRequest.ID)
	if err != nil {
		log.Err(err).Msg("Error in GetMovieWishList")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error with User API"})
		return
	}
	// Return movie wishlist
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
	// Add movie to wishlist
	if err = s.db.InsertSerieWishList(serieRequest.UsernameID, serieRequest.SerieID); err != nil {
		log.Err(err).Msg("Error in AddSerieWishList")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error with User API"})
		return
	}
	// Return movie wishlist
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
	// Delete movie from wishlist
	err := s.db.DeleteSerieWishList(serieRequest.UsernameID, serieRequest.SerieID)
	if err != nil {
		log.Err(err).Msg("Error in DeleteSerieWishList")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error with User API"})
		return
	}
	// Return movie wishlist
	c.JSON(http.StatusOK, "success")
}
