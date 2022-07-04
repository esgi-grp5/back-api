package main

import (
	"github.com/gin-gonic/gin"
)

// server struct
type server struct {
	gin *gin.Engine
}

// @title Gin Swagger Example API
// @version 2.0
// @description This is a sample server server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:3000
// @BasePath /
// @schemes http
func newServer() *server {
	// Initialize server
	s := &server{
		gin: gin.New(),
	}
	// Initialize router
	s.routes()
	return s
}
