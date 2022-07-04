package main

import (
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (s *server) routes() {
	// Routes
	s.gin.GET("/hello", HealthCheck)
	// Swagger
	url := ginSwagger.URL("/swagger/user/doc.json") // The url pointing to API definition
	s.gin.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}
