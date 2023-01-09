package db

import (
	"context"
	"errors"
	"fmt"

	"github.com/h00s-go/tiny-link-backend/config"
	"github.com/h00s-go/tiny-link-backend/db/migrations"
	"github.com/h00s-go/tiny-link-backend/db/sql"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Database struct {
	config     *config.Database
	conn       *pgxpool.Pool
	migrations map[int]string
}

func NewDatabase(config *config.Database) *Database {
	return &Database{
		config: config,
		migrations: map[int]string{
			1: migrations.CreateLinks,
		},
	}
}

func (db *Database) Connect() error {
	conn, err := pgxpool.New(context.Background(), db.ConnString())
	if err != nil {
		return errors.New("Unable to create connection pool: " + err.Error())
	}
	db.conn = conn

	if err := db.conn.Ping(context.Background()); err != nil {
		return errors.New("Unable to ping database: " + err.Error())
	}
	return nil
}

func (db *Database) Close() {
	db.conn.Close()
}

func (db *Database) ConnString() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s pool_max_conns=10 pool_min_conns=2",
		db.config.Host,
		db.config.Port,
		db.config.User,
		db.config.Password,
		db.config.Name,
	)
}

func (db *Database) Migrate() error {
	if _, err := db.conn.Exec(context.Background(), sql.CreateSchema); err != nil {
		return err
	}

	var version int
	err := db.conn.QueryRow(context.Background(), sql.SelectSchema).Scan(&version)
	if err != nil {
		if err == pgx.ErrNoRows {
			if _, err = db.conn.Exec(context.Background(), sql.InsertSchema); err != nil {
				return err
			}
		} else {
			return errors.New("Unable to scan row: " + err.Error())
		}
	}

	for i := version + 1; i <= len(db.migrations); i++ {
		if _, err := db.conn.Exec(context.Background(), db.migrations[i]); err != nil {
			return err
		}
	}

	return nil
}
