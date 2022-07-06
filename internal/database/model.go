package database

import "time"

type Token struct {
	ID         int
	ClientName string
	Token      string
	Created    time.Time
}
