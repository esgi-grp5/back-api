package main

import (
	"go-micro/internal/config"
	"os"
	"strconv"

	_ "github.com/rizalgowandy/go-swag-sample/docs/ginsimple" // you need to update github.com/rizalgowandy/go-swag-sample with your own project path
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	log.Info().Msg("USER API Server is starting...")
	if err := run(); err != nil {
		log.Err(err).Msg("Fatal error")
		os.Exit(1)
	}
}

func run() error {
	/* Initialize config */
	c := config.Config("users")
	// Set zerolog level
	if c.Debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
	// Set gin port
	if c.Port == 0 {
		c.Port = 3000
	}
	/* Initialize server */
	s := newServer(c)
	/* Start server */
	if err := s.gin.Run(":" + strconv.Itoa(c.Port)); err != nil {
		return err
	}
	return nil
}
