package database

import (
	"fmt"

	md "github.com/BMilkey/messenger/models"
	pgx "github.com/jackc/pgx/v5/pgxpool"
)

func InsertUser(pool *pgx.Pool, user md.User) error {
	querry := fmt.Sprintf(
		`
		INSERT INTO public.users
		(id, name, link, about, last_connection, image_id)
		VALUES
		(%s, %s, %s, %s, %s)
		`,
		user.Id, user.Name, user.Link, user.About, user.Last_connection)
	return wrapRawExecQuerry(pool, querry)
}

func InsertAuth(pool *pgx.Pool, auth md.Auth) error {
	querry := fmt.Sprintf(
		`
		INSERT INTO public.auth
		(user_id, login_hash, password_hash)
		VALUES
		(%s, %s, %s)
		`,
		auth.User_id, auth.Login_hash, auth.Password_hash)
	return wrapRawExecQuerry(pool, querry)
}
