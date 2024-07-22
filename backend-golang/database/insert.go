package database

import (
	"context"

	md "github.com/BMilkey/messenger/models"
	pgx "github.com/jackc/pgx/v5/pgxpool"
)

func InsertUser(pool *pgx.Pool, user md.User) error {
	_, err := pool.Exec(context.Background(),
		`
		INSERT INTO public.users
		(id, name, link, about, last_connection, image_id)
		VALUES
		($1, $2, $3, $4, $5, $6)
		`,
		user.Id, user.Name, user.Link, user.About, user.Last_online, user.Image_id)
	return err
}

func InsertAuth(pool *pgx.Pool, auth md.Auth) error {
	_, err := pool.Exec(context.Background(),
		`
			INSERT INTO public.auth(
			login_hash, password_hash, user_id, auth_token, auth_expires, email)
			VALUES
			($1, $2, $3, $4, $5, $6)
		`,
		auth.Login_hash, auth.Password_hash, auth.User_id, auth.Auth_token, auth.Auth_expires, auth.Email)
	return err
}

func InsertChat(pool *pgx.Pool, chat md.ChatInfo) error {
	_, err := pool.Exec(context.Background(),
		`
		INSERT INTO public.chats(
			id, link, title, created_by_user_id, create_time, about, image_id)
			VALUES ($1, $2, $3, $4, $5, $6, $7);
		`,
		chat.Id, chat.Link, chat.Title, chat.User_id, chat.Create_time, chat.About, chat.Image_id)
	return err
}

func InsertChatParticipant(pool *pgx.Pool, chat_id string, user_id string) error {
	_, err := pool.Exec(context.Background(),
		`
		INSERT INTO public.chat_participants(
			chat_id, user_id)
			VALUES ($1, $2);
		`,
		chat_id, user_id)
	return err
}

func InsertMessage(pool *pgx.Pool, message md.Message) error {

	_, err := pool.Exec(context.Background(),
		`
		INSERT INTO public.messages(
			id, chat_id, user_id, create_time, text, reply_message_id)
			VALUES ($1, $2, $3, $4, $5, $6);
		`,
		message.Id, message.Chat_id, message.User_id, message.Create_time, message.Text, message.Reply_msg_id)
	return err
}
