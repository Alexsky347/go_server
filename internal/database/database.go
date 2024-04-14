package database

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

func New(dataBaseURL string) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(context.Background(), dataBaseURL)
	if err != nil {
		return nil, err
	}

	var greeting string
	err = pool.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)
	if err != nil {
		return nil, err
	}
	return pool, nil
}
