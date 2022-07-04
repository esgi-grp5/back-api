package main

import (
	"os"

	_ "github.com/rizalgowandy/go-swag-sample/docs/ginsimple" // you need to update github.com/rizalgowandy/go-swag-sample with your own project path
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	log.Info().Msg("GAMES API Server is starting...")
	if err := run(); err != nil {
		log.Err(err).Msg("Fatal error")
		os.Exit(1)
	}
}

func run() error {
	// Initialize server
	s := newServer()
	// Start server
	if err := s.gin.Run(":3001"); err != nil {
		return err
	}
	return nil
}
