package main

func (s *server) routes() {
	// Access without OAuth
	s.gin.POST("/oauth", s.OAuth)
	s.gin.GET("/health", s.HealthCheck)
	// Access with OAuth
	s.gin.GET("/flutter", s.FlutterAccess)
}
