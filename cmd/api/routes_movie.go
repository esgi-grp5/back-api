package main

import (
	"go-micro/internal/database"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func (s *server) GetMovieWishList(c *gin.Context) {
	userRequest := c.Param("id")
	userID, err := strconv.Atoi(userRequest)
	if err != nil {
		log.Err(err).Msg("Error in GetMovieWishList when Atoi")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error with User API"})
		return
	}
	// Get movie wishlist
	wishlist, err := s.db.SelectMovieWishList(userID)
	if err != nil {
		log.Err(err).Msg("Error in GetMovieWishList")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error with User API"})
		return
	}
	// Return movie wishlist
	c.JSON(http.StatusOK, wishlist)
}

func (s *server) AddMovieWishList(c *gin.Context) {
	var movieRequest database.Movie
	// Get JSON body
	if err := c.ShouldBindJSON(&movieRequest); err != nil {
		log.Err(err).Msg("Error in Login")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error with User API"})
		return
	}
	wishlist, err := s.db.SelectMovieWishList(movieRequest.UsernameID)
	if err != nil {
		log.Err(err).Msg("Error in AddMovieWishList")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error with User API"})
		return
	}
	for _, w := range wishlist {
		if w.MovieID == movieRequest.MovieID {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Movie already in wishlist"})
			return
		}
	}
	// Add movie to wishlist
	if err = s.db.InsertMovieWishList(movieRequest.UsernameID, movieRequest.MovieID); err != nil {
		log.Err(err).Msg("Error in AddMovieWishList")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error with User API"})
		return
	}
	// Return movie wishlist
	c.JSON(http.StatusOK, "success")
}

func (s *server) DeleteMovieWishList(c *gin.Context) {
	var movieRequest database.Movie
	// Get JSON body
	if err := c.ShouldBindJSON(&movieRequest); err != nil {
		log.Err(err).Msg("Error in Login")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error with User API"})
		return
	}
	// Delete movie from wishlist
	err := s.db.DeleteMovieWishList(movieRequest.UsernameID, movieRequest.MovieID)
	if err != nil {
		log.Err(err).Msg("Error in DeleteMovieWishList")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error with User API"})
		return
	}
	// Return movie wishlist
	c.JSON(http.StatusOK, "success")
}

func (s *server) GetMovieCount(c *gin.Context) {
	movieRequest := c.Param("id")
	movieID, err := strconv.Atoi(movieRequest)
	if err != nil {
		log.Err(err).Msg("Error in GetMovieCount when Atoi")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error with User API"})
		return
	}
	// Get movie count
	count, err := s.db.SelectMovieCount(movieID)
	if err != nil {
		log.Err(err).Msg("Error in GetMovieCount")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error with User API"})
		return
	}
	res := map[string]interface{}{
		"count": count,
	}
	// Return movie count
	c.JSON(http.StatusOK, res)
}
