package main

import (
	"go-micro/internal/config"
	"go-micro/internal/database"

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
