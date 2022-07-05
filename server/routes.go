package main

func (s *server) routes() {
	// Routes
	s.gin.POST("/oauth", s.OAuth)
	s.gin.GET("/hello", s.HealthCheck)
	// Access with OAuth
	s.gin.GET("/flutter", s.FlutterAccess)
}
