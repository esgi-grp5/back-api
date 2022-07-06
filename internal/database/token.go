package database

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/rs/zerolog/log"
)

type TokenPostgres struct {
	db *pgxpool.Pool
}

func NewTokenPostGres(db *pgxpool.Pool) *TokenPostgres {
	return &TokenPostgres{
		db: db,
	}
}

// Select token config of client
func (store *TokenPostgres) SelectAuth(apiToken string) (Token, error) {
	ctx := context.Background()
	var token Token
	err := store.db.QueryRow(ctx, "SELECT id, client_name, token, created FROM token WHERE token = $1", apiToken).Scan(&token.ID, &token.ClientName, &token.Token, &token.Created)
	if err != nil {
		log.Error().Err(err).Msg("token repository. cannot select token")
	}
	return token, nil
}
