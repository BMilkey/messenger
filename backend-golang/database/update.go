package database

import (
	"context"
	"time"

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
		auth.Auth_token, time.Now().Add(time.Hour*24) /*auth.Auth_expires*/, auth.User_id)

	return err
}

func UpdateUser(pool *pgx.Pool, user md.User) error {
	_, err := pool.Exec(context.Background(),
		`
		UPDATE public.users
		SET name=$1, link=$2, about=$3, last_connection=$4, image_id=$5
		WHERE id=$6;
		`,
		user.Name, user.Link, user.About, time.Now(), user.Image_id, user.Id)
	return err
}
