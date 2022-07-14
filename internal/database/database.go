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

/* Movie */

func (store *DatabasePostgres) SelectMovieWishList(idUser int) ([]Movie, error) {
	ctx := context.Background()
	var movies []Movie
	rows, err := store.db.Query(ctx, "SELECT * FROM movie WHERE username_id = $1", idUser)
	if err != nil {
		log.Error().Err(err).Msg("movie repository. cannot select movie")
	}
	defer rows.Close()
	for rows.Next() {
		var movie Movie
		err := rows.Scan(&movie.ID, &movie.UsernameID, &movie.MovieID)
		if err != nil {
			log.Error().Err(err).Msg("movie repository. cannot select movie")
		}
		movies = append(movies, movie)
	}
	return movies, nil
}

func (store *DatabasePostgres) InsertMovieWishList(usernameID, movieID int) error {
	ctx := context.Background()
	_, err := store.db.Exec(ctx, "INSERT INTO movie (username_id, movie_id) VALUES ($1, $2)", usernameID, movieID)
	if err != nil {
		log.Error().Err(err).Msg("movie repository. cannot insert movie")
	}
	return nil
}

func (store *DatabasePostgres) DeleteMovieWishList(usernameID, movieID int) error {
	ctx := context.Background()
	_, err := store.db.Exec(ctx, "DELETE FROM movie WHERE username_id = $1 AND movie_id = $2", usernameID, movieID)
	if err != nil {
		log.Error().Err(err).Msg("movie repository. cannot delete movie")
	}
	return nil
}
