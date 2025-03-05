package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"os"
)

func (pg *Postgres) NewPool() *pgxpool.Pool {
	pool, err := pgxpool.New(context.Background(), pg.connString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	return pool
}
