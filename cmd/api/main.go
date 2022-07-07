package main

import (
	"context"
	"errors"
	"fmt"
	"go-micro/internal/config"
	"go-micro/internal/database"
	"os"
	"strconv"

	"github.com/jackc/pgx/v4/pgxpool"
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
	c := config.Config()
	// Set zerolog level
	if c.Debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
	// Set gin port
	if c.Port == 0 {
		c.Port = 3000
	}
	// Check DB config
	if c.DB.Host == "" {
		return errors.New("missing db hostname")
	}
	if c.DB.Username == "" {
		return errors.New("missing db username")
	}
	if c.DB.Password == "" {
		return errors.New("missing db password")
	}
	if c.DB.Name == "" {
		return errors.New("missing db dbName")
	}
	/* Set up db */
	u := fmt.Sprintf(
		"postgres://%s:%s@%s/%s",
		c.DB.Username,
		c.DB.Password,
		c.DB.Host,
		c.DB.Name,
	)
	pool, err := pgxpool.Connect(context.Background(), u)
	if err != nil {
		return fmt.Errorf("cannot connect to pgxpool. %s", err)
	}
	defer pool.Close()
	dbPool := database.NewDatabasePostgres(pool)
	// Migration DB
	v, err := database.Migrate(c.DB.Username, c.DB.Password, c.DB.Host, c.DB.Name)
	if err != nil {
		return fmt.Errorf("cannot migrate to version %d. %s", v, err)
	}
	/* Initialize server */
	s := newServer(c, dbPool)
	/* Start server */
	if err := s.gin.Run(":" + strconv.Itoa(c.Port)); err != nil {
		return err
	}
	return nil
}
