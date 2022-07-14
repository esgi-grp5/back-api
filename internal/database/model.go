package database

import "time"

type Token struct {
	ID         int
	ClientName string
	Token      string
	Created    time.Time
}

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Mail     string `json:"mail"`
	Password string `json:"password"`
}

type Movie struct {
	ID         int `json:"id"`
	UsernameID int `json:"username_id"`
	MovieID    int `json:"movie_id"`
}
