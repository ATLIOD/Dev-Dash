package db

import (
	"DevDash/internal/config"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Database struct {
	// field for pgx pool
}

func New() (*Database, error) {
	// initialize pgx pool and return database instance
	return &Database{}, nil
}

func OpenDB(dbConfig config.DBConfig) (*pgxpool.Pool, error) {
	// Parse the connection string into a pgxpool.Config
	poolConfig, err := pgxpool.ParseConfig(dbConfig.Dsn)
	if err != nil {
		fmt.Printf("Error parsing DSN: %v\n", err)
		return nil, err
	}

	poolConfig.MaxConns = int32(dbConfig.MaxConns)
	poolConfig.MaxConnIdleTime = dbConfig.MaxConnIdleTime
	poolConfig.MinConns = int32(dbConfig.MinConns)

	pool, err := pgxpool.New(context.Background(), dbConfig.Dsn)
	if err != nil {
		fmt.Printf("Unable to create connection pool: %v\n", err)
		return nil, err
	}

	// Test the connection
	err = pool.Ping(context.Background())
	if err != nil {
		return nil, err
	}

	return pool, nil
}

func (db *Database) Close() error {
	// close database connection
	return nil
}
