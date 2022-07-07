package database

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/rs/zerolog/log"
)

type DatabasePostgres struct {
	db *pgxpool.Pool
}

func NewDatabasePostgres(db *pgxpool.Pool) *DatabasePostgres {
	return &DatabasePostgres{
		db: db,
	}
}

// Select token config of client
func (store *DatabasePostgres) SelectAuth(apiToken string) (Token, error) {
	ctx := context.Background()
	var token Token
	err := store.db.QueryRow(ctx, "SELECT id, client_name, token, created FROM token WHERE token = $1", apiToken).Scan(&token.ID, &token.ClientName, &token.Token, &token.Created)
	if err != nil {
		log.Error().Err(err).Msg("token repository. cannot select token")
	}
	return token, nil
}

// Select user where mail is equal to mail
func (store *DatabasePostgres) SelectUser(mailUser string) (User, error) {
	ctx := context.Background()
	var user User
	err := store.db.QueryRow(ctx, "SELECT id, username, mail, password FROM users WHERE mail = $1", mailUser).Scan(&user.ID, &user.Username, &user.Mail, &user.Password)
	if err != nil {
		log.Error().Err(err).Msg("user repository. cannot select user")
	}
	return user, nil
}
