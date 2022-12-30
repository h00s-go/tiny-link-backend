package db

import (
	"context"
	"errors"
	"fmt"

	"github.com/h00s-go/tiny-link-backend/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Database struct {
	config *config.Database
	conn   *pgxpool.Pool
}

func NewDatabase(config *config.Database) *Database {
	return &Database{
		config: config,
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
