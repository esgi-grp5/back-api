package main

import (
	"go-micro/internal/config"
	"go-micro/internal/database"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
)

// server struct
type server struct {
	gin    *gin.Engine
	config config.Configuration
	oauth  config.OAuthApp
	// DB
	db *database.DatabasePostgres
}

func newServer(c config.Configuration, db *database.DatabasePostgres) *server {
	// Initialize server
	s := &server{
		gin:    gin.New(),
		config: c,
		oauth:  c.OAuthApp,
		db:     db,
	}
	// Initialize router
	s.routes()
	return s
}

func RandomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	rand.Seed(time.Now().UTC().UnixNano())
	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}
