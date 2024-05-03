package database

import (
	"fmt"
	"context"
	pgx "github.com/jackc/pgx/v5/pgxpool"
)


func DeleteUserById(pool *pgx.Pool, user_id string) error {
	query := `DELETE FROM public.user WHERE id = $1;`

	_, err := pool.Exec(context.Background(), query, user_id)
	if err != nil {
		return fmt.Errorf("delete user by id: %w", err)
	}

	return nil
}

func DeleteAuthByUserId(pool *pgx.Pool, user_id string) error {
	query := `DELETE FROM public.auth WHERE user_id = $1;`

	_, err := pool.Exec(context.Background(), query, user_id)
	if err != nil {
		return fmt.Errorf("delete auth by user id: %w", err)
	}

	return nil
}
