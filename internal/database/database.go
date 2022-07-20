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
		return Token{}, err
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
		return User{}, err
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
		return User{}, err
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
		return User{}, err
	}
	return user, nil
}

// Insert user
func (store *DatabasePostgres) InsertUser(user User) error {
	ctx := context.Background()
	_, err := store.db.Exec(ctx, "INSERT INTO users (username, mail, password) VALUES ($1, $2, $3)", user.Username, user.Mail, user.Password)
	if err != nil {
		log.Error().Err(err).Msg("user repository. cannot insert user")
		return err
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
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var movie Movie
		err := rows.Scan(&movie.ID, &movie.UsernameID, &movie.MovieID)
		if err != nil {
			log.Error().Err(err).Msg("movie repository. cannot select movie")
			return nil, err
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
		return err
	}
	return nil
}

func (store *DatabasePostgres) DeleteMovieWishList(usernameID, movieID int) error {
	ctx := context.Background()
	_, err := store.db.Exec(ctx, "DELETE FROM movie WHERE username_id = $1 AND movie_id = $2", usernameID, movieID)
	if err != nil {
		log.Error().Err(err).Msg("movie repository. cannot delete movie")
		return err
	}
	return nil
}

func (store *DatabasePostgres) SelectMovieCount(movieID int) (int, error) {
	ctx := context.Background()
	var count int
	err := store.db.QueryRow(ctx, "SELECT count(*) FROM movie WHERE movie_id = $1", movieID).Scan(&count)
	if err != nil {
		log.Error().Err(err).Msg("movie repository. cannot count movie")
		return 0, err
	}
	return count, nil
}

/* Serie */

func (store *DatabasePostgres) SelectSerieWishList(idUser int) ([]Serie, error) {
	ctx := context.Background()
	var series []Serie
	rows, err := store.db.Query(ctx, "SELECT * FROM serie WHERE username_id = $1", idUser)
	if err != nil {
		log.Error().Err(err).Msg("serie repository. cannot select serie")
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var serie Serie
		err := rows.Scan(&serie.ID, &serie.UsernameID, &serie.SerieID)
		if err != nil {
			log.Error().Err(err).Msg("serie repository. cannot select serie")
			return nil, err
		}
		series = append(series, serie)
	}
	return series, nil
}

func (store *DatabasePostgres) InsertSerieWishList(usernameID, serieID int) error {
	ctx := context.Background()
	_, err := store.db.Exec(ctx, "INSERT INTO serie (username_id, serie_id) VALUES ($1, $2)", usernameID, serieID)
	if err != nil {
		log.Error().Err(err).Msg("serie repository. cannot insert serie")
	}
	return nil
}

func (store *DatabasePostgres) DeleteSerieWishList(usernameID, serieID int) error {
	ctx := context.Background()
	_, err := store.db.Exec(ctx, "DELETE FROM serie WHERE username_id = $1 AND serie_id = $2", usernameID, serieID)
	if err != nil {
		log.Error().Err(err).Msg("serie repository. cannot delete serie")
	}
	return nil
}

func (store *DatabasePostgres) SelectSerieCount(serieID int) (int, error) {
	ctx := context.Background()
	var count int
	err := store.db.QueryRow(ctx, "SELECT count(*) FROM serie WHERE serie_id = $1", serieID).Scan(&count)
	if err != nil {
		log.Error().Err(err).Msg("serie repository. cannot count serie")
		return 0, err
	}
	return count, nil
}

/* Game */

func (store *DatabasePostgres) SelectGameWishList(idUser int) ([]Game, error) {
	ctx := context.Background()
	var games []Game
	rows, err := store.db.Query(ctx, "SELECT * FROM game WHERE username_id = $1", idUser)
	if err != nil {
		log.Error().Err(err).Msg("game repository. cannot select game")
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var game Game
		err := rows.Scan(&game.ID, &game.UsernameID, &game.GameID)
		if err != nil {
			log.Error().Err(err).Msg("game repository. cannot select game")
			return nil, err
		}
		games = append(games, game)
	}
	return games, nil
}

func (store *DatabasePostgres) InsertGameWishList(usernameID, gameID int) error {
	ctx := context.Background()
	_, err := store.db.Exec(ctx, "INSERT INTO game (username_id, game_id) VALUES ($1, $2)", usernameID, gameID)
	if err != nil {
		log.Error().Err(err).Msg("game repository. cannot insert game")
	}
	return nil
}

func (store *DatabasePostgres) DeleteGameWishList(usernameID, gameID int) error {
	ctx := context.Background()
	_, err := store.db.Exec(ctx, "DELETE FROM game WHERE username_id = $1 AND game_id = $2", usernameID, gameID)
	if err != nil {
		log.Error().Err(err).Msg("game repository. cannot delete game")
	}
	return nil
}

func (store *DatabasePostgres) SelectGameCount(gameID int) (int, error) {
	ctx := context.Background()
	var count int
	err := store.db.QueryRow(ctx, "SELECT count(*) FROM game WHERE game_id = $1", gameID).Scan(&count)
	if err != nil {
		log.Error().Err(err).Msg("game repository. cannot count game")
		return 0, err
	}
	return count, nil
}

/* Music */

func (store *DatabasePostgres) SelectMusicWishList(idUser int) ([]Music, error) {
	ctx := context.Background()
	var musics []Music
	rows, err := store.db.Query(ctx, "SELECT * FROM music WHERE username_id = $1", idUser)
	if err != nil {
		log.Error().Err(err).Msg("music repository. cannot select music")
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var music Music
		err := rows.Scan(&music.ID, &music.UsernameID, &music.MusicID)
		if err != nil {
			log.Error().Err(err).Msg("music repository. cannot select music")
			return nil, err
		}
		musics = append(musics, music)
	}
	return musics, nil
}

func (store *DatabasePostgres) InsertMusicWishList(usernameID, musicID int) error {
	ctx := context.Background()
	_, err := store.db.Exec(ctx, "INSERT INTO music (username_id, music_id) VALUES ($1, $2)", usernameID, musicID)
	if err != nil {
		log.Error().Err(err).Msg("music repository. cannot insert music")
	}
	return nil
}

func (store *DatabasePostgres) DeleteMusicWishList(usernameID, musicID int) error {
	ctx := context.Background()
	_, err := store.db.Exec(ctx, "DELETE FROM music WHERE username_id = $1 AND music_id = $2", usernameID, musicID)
	if err != nil {
		log.Error().Err(err).Msg("music repository. cannot delete music")
	}
	return nil
}

func (store *DatabasePostgres) SelectMusicCount(musicID int) (int, error) {
	ctx := context.Background()
	var count int
	err := store.db.QueryRow(ctx, "SELECT count(*) FROM music WHERE music_id = $1", musicID).Scan(&count)
	if err != nil {
		log.Error().Err(err).Msg("music repository. cannot count music")
		return 0, err
	}
	return count, nil
}
