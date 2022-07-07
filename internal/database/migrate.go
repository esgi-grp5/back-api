package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
	"github.com/rs/zerolog/log"
)

var schemaVersion = len(migrations)

func Migrate(username, password, host, bdd string) (int, error) {
	var currentVersion int
	var newVersion int

	ctx := context.Background()
	u := fmt.Sprintf(
		"postgres://%s:%s@%s/%s",
		username,
		password,
		host,
		bdd,
	)
	conn, err := pgx.Connect(ctx, u)
	if err != nil {
		return 999, err
	}

	err = conn.QueryRow(ctx, "SELECT api FROM schema").Scan(&currentVersion)
	if err != nil {
		return 999, err
	}

	log.Info().Int("current", currentVersion).Int("latest", schemaVersion).Msg("Schema version")

	for version := currentVersion; version < schemaVersion; version++ {
		newVersion = version + 1
		log.Info().Msg(fmt.Sprintf("* Migrating to version: %v", newVersion))

		tx, err := conn.Begin(ctx)
		if err != nil {
			return newVersion, fmt.Errorf("cannot start transaction. %v", err)
		}

		if err := migrations[version](tx.Conn()); err != nil {
			_ = tx.Rollback(ctx)
			return newVersion, fmt.Errorf("cannot execute sql %v", err)
		}

		if _, err := tx.Exec(ctx, `UPDATE schema SET api = $1`, newVersion); err != nil {
			_ = tx.Rollback(ctx)
			return newVersion, fmt.Errorf("cannot update schema. %s", err)
		}

		if err := tx.Commit(ctx); err != nil {
			return newVersion, fmt.Errorf("cannot commit. %v", err)
		}
	}
	return newVersion, nil
}

var migrations = []func(conn *pgx.Conn) error{
	// v1
	// create token table
	func(conn *pgx.Conn) (err error) {
		sql := `
		CREATE SEQUENCE token_id_seq START 1;
		CREATE TABLE "public"."token" (
			"id" integer DEFAULT nextval('token_id_seq') NOT NULL,
			"client_name" text NOT NULL,
			"token" text NOT NULL,
			"created" timestamp NOT NULL,
			CONSTRAINT "token_pkey" PRIMARY KEY ("id")
		) WITH (oids = false);`

		_, err = conn.Exec(context.Background(), sql)
		return err
	},
	// v2
	// create user table
	func(conn *pgx.Conn) (err error) {
		sql := `
		CREATE SEQUENCE users_id_seq START 1;
		CREATE TABLE "public"."users" (
			"id" integer DEFAULT nextval('users_id_seq') NOT NULL,
			"username" text NOT NULL,
			"mail" text NOT NULL,
			"password" text NOT NULL,
			CONSTRAINT "users_pkey" PRIMARY KEY ("id")
		) WITH (oids = false);`

		_, err = conn.Exec(context.Background(), sql)
		return err
	},
}
