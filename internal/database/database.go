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

/* Token */

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

/* User */

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

// Get mail from user
func (store *DatabasePostgres) GetMailFromUser(mail string) (User, error) {
	ctx := context.Background()
	var user User
	err := store.db.QueryRow(ctx, "SELECT mail FROM users WHERE mail = $1", mail).Scan(&user.Mail)
	if err != nil {
		log.Error().Err(err).Msg("user repository. cannot select user")
	}
	return user, nil
}

// Get username from user
func (store *DatabasePostgres) GetUsernameFromUser(username string) (User, error) {
	ctx := context.Background()
	var user User
	err := store.db.QueryRow(ctx, "SELECT username FROM users WHERE username = $1", username).Scan(&user.Username)
	if err != nil {
		log.Error().Err(err).Msg("user repository. cannot select user")
	}
	return user, nil
}

// Insert user
func (store *DatabasePostgres) InsertUser(user User) error {
	ctx := context.Background()
	_, err := store.db.Exec(ctx, "INSERT INTO users (username, mail, password) VALUES ($1, $2, $3)", user.Username, user.Mail, user.Password)
	if err != nil {
		log.Error().Err(err).Msg("user repository. cannot insert user")
	}
	return nil
}
