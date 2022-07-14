package main

func (s *server) routes() {
	// Access without OAuth
	s.gin.POST("/oauth", s.OAuth)
	s.gin.GET("/health", s.HealthCheck)
	// Access with OAuth
	oauth := s.gin.Use(s.OAuthMiddleware)
	oauth.GET("/hello", s.FlutterAccess)
	oauth.POST("/login", s.Login)
	oauth.POST("/register", s.Register)
	oauth.GET("/movie/wishlist", s.GetMovieWishList)
	oauth.POST("/movie/wishlist", s.AddMovieWishList)
	oauth.DELETE("/movie/wishlist", s.DeleteMovieWishList)
	// oauth.GET("/serie/wishlist", s.GetMovie)
	// oauth.GET("/game/wishlist", s.GetMovie)
	// oauth.GET("/music/wishlist", s.GetMovie)
}
