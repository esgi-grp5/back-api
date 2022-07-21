package main

import (
	"go-micro/internal/database"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func (s *server) GetMusicWishList(c *gin.Context) {
	var userRequest User
	// Get JSON body
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		log.Err(err).Msg("Error in Login")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error with User API"})
		return
	}
	userID, err := strconv.Atoi(userRequest.ID)
	if err != nil {
		log.Err(err).Msg("Error in GetMusicWishList when Atoi")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error with User API"})
		return
	}
	// Get music wishlist
	wishlist, err := s.db.SelectMusicWishList(userID)
	if err != nil {
		log.Err(err).Msg("Error in GetMusicWishList")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error with User API"})
		return
	}
	// Return music wishlist
	c.JSON(http.StatusOK, wishlist)
}

func (s *server) AddMusicWishList(c *gin.Context) {
	var musicRequest database.Music
	// Get JSON body
	if err := c.ShouldBindJSON(&musicRequest); err != nil {
		log.Err(err).Msg("Error in Login")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error with User API"})
		return
	}
	wishlist, err := s.db.SelectMusicWishList(musicRequest.UsernameID)
	if err != nil {
		log.Err(err).Msg("Error in AddMusicWishList")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error with User API"})
		return
	}
	for _, w := range wishlist {
		if w.MusicID == musicRequest.MusicID {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Serie already in wishlist"})
			return
		}
	}
	// Add music to wishlist
	if err = s.db.InsertMusicWishList(musicRequest.UsernameID, musicRequest.MusicID); err != nil {
		log.Err(err).Msg("Error in AddMusicWishList")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error with User API"})
		return
	}
	// Return music wishlist
	c.JSON(http.StatusOK, "success")
}

func (s *server) DeleteMusicWishList(c *gin.Context) {
	var musicRequest database.Music
	// Get JSON body
	if err := c.ShouldBindJSON(&musicRequest); err != nil {
		log.Err(err).Msg("Error in Login")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error with User API"})
		return
	}
	// Delete music from wishlist
	err := s.db.DeleteMusicWishList(musicRequest.UsernameID, musicRequest.MusicID)
	if err != nil {
		log.Err(err).Msg("Error in DeleteMusicWishList")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error with User API"})
		return
	}
	// Return music wishlist
	c.JSON(http.StatusOK, "success")
}

func (s *server) GetMusicCount(c *gin.Context) {
	var musicRequest Music
	// Get JSON body
	if err := c.ShouldBindJSON(&musicRequest); err != nil {
		log.Err(err).Msg("Error in Login")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error with User API"})
		return
	}
	musicID, err := strconv.Atoi(musicRequest.ID)
	if err != nil {
		log.Err(err).Msg("Error in GetMusicCount when Atoi")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error with User API"})
		return
	}
	// Get music count
	count, err := s.db.SelectMusicCount(musicID)
	if err != nil {
		log.Err(err).Msg("Error in GetMusicCount")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error with User API"})
		return
	}
	res := map[string]interface{}{
		"count": count,
	}
	// Return music count
	c.JSON(http.StatusOK, res)
}
