package db

import (
	"context"
	"fmt"
	"os"

	"github.com/h00s-go/tiny-link-backend/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Database struct {
	config *config.Database
}

func NewDatabase(config *config.Database) *Database {
	return &Database{
		config: config,
	}
}

func (db *Database) Connect() error {
	dbpool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()
	return nil
}
