package database

import (
	"context"

	pgx "github.com/jackc/pgx/v5/pgxpool"
)

func wrapRawExecQuerry(pool *pgx.Pool, querry string) error {
	_, err := pool.Exec(context.Background(), querry)

	if err != nil {
		return err
	}
	return nil
}
