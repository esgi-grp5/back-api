package main

import (
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (s *server) routes() {
	// Routes
	s.gin.POST("/oauth", s.OAuth)
	s.gin.GET("/hello", s.HealthCheck)
	// Access with OAuth
	s.gin.GET("/flutter", s.FlutterAccess)
	// Swagger
	url := ginSwagger.URL("/swagger/user/doc.json") // The url pointing to API definition
	s.gin.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}
