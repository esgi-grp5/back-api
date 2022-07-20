package main

func (s *server) routes() {
	/* Access without OAuth */
	s.gin.POST("/oauth", s.OAuth)
	s.gin.GET("/health", s.HealthCheck)
	/* Access with OAuth */
	oauth := s.gin.Use(s.OAuthMiddleware)
	oauth.GET("/hello", s.FlutterAccess)
	oauth.POST("/login", s.Login)
	oauth.POST("/register", s.Register)
	// Movie
	oauth.GET("/movie/wishlist", s.GetMovieWishList)
	oauth.POST("/movie/wishlist", s.AddMovieWishList)
	oauth.DELETE("/movie/wishlist", s.DeleteMovieWishList)
	oauth.GET("/movie/count", s.GetMovieCount)
	// Serie
	oauth.GET("/serie/wishlist", s.GetSerieWishList)
	oauth.POST("/serie/wishlist", s.AddSerieWishList)
	oauth.DELETE("/serie/wishlist", s.DeleteSerieWishList)
	oauth.GET("/serie/count", s.GetSerieCount)
	// Game
	oauth.GET("/game/wishlist", s.GetGameWishList)
	oauth.POST("/game/wishlist", s.AddGameWishList)
	oauth.DELETE("/game/wishlist", s.DeleteGameWishList)
	oauth.GET("/game/count", s.GetGameCount)
	// Music
	oauth.GET("/music/wishlist", s.GetMusicWishList)
	oauth.POST("/music/wishlist", s.AddMusicWishList)
	oauth.DELETE("/music/wishlist", s.DeleteMusicWishList)
	oauth.GET("/music/count", s.GetMusicCount)
}
