package database

import (
	"context"

	md "github.com/BMilkey/messenger/models"
	pgx "github.com/jackc/pgx/v5/pgxpool"
)

func UpdateAuthToken(pool *pgx.Pool, auth md.Auth) error {
	_, err := pool.Exec(context.Background(),
		`
		UPDATE public.auth
		SET auth_token=$1, auth_expires=$2
		WHERE user_id=$3;
		`,
		auth.Auth_token, auth.Auth_expires, auth.User_id)

	return err
}
