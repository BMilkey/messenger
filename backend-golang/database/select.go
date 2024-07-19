package database

import (
	"context"

	md "github.com/BMilkey/messenger/models"
	pgx "github.com/jackc/pgx/v5/pgxpool"
)

func SelectUserById(pool *pgx.Pool, user_id string) (md.User, error) {
	var user md.User

	err := pool.QueryRow(context.Background(), `
		SELECT id, name, link, about, last_connection, image_id 
		FROM public.users 
		WHERE users.id = $1
		`,
		user_id).
		Scan(&user.Id, &user.Name, &user.Link, &user.About, &user.Last_online, &user.Image_id)

	if err != nil {
		return user, err
	}

	return user, nil
}

func SelectUsersByName(pool *pgx.Pool, name string) ([]md.User, error) {
	var users []md.User

	rows, err := pool.Query(context.Background(), `
        SELECT id, name, link, about, last_connection, image_id 
        FROM public.users 
        WHERE users.name like $1
		`,
		"%"+name+"%")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user md.User
		err := rows.Scan(&user.Id, &user.Name, &user.Link, &user.About, &user.Last_online, &user.Image_id)
		if err != nil {
			return users, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func SelectUserByLink(pool *pgx.Pool, link string) (md.User, error) {
	var user md.User

	err := pool.QueryRow(context.Background(), `
	SELECT id, name, link, about, last_connection, image_id
	FROM public.users
	where link = $1
		`,
		link).
		Scan(&user.Id, &user.Name, &user.Link, &user.About, &user.Last_online, &user.Image_id)

	if err != nil {
		return user, err
	}

	return user, nil
}

func SelectChatById(pool *pgx.Pool, chat_id string) (md.Chat, error) {
	var chat md.Chat

	err := pool.QueryRow(context.Background(), `
		SELECT id, link, title, created_by_user_id, create_time, about, image_id 
		FROM public.chats 
		WHERE chats.id = $1
		`,
		chat_id).
		Scan(&chat.Id, &chat.Link, &chat.Title, &chat.User_id, &chat.Create_time, &chat.About, &chat.Image_id)

	if err != nil {
		return chat, err
	}

	return chat, nil
}

func SelectChatByLink(pool *pgx.Pool, chat_link string) (md.Chat, error) {
	var chat md.Chat

	err := pool.QueryRow(context.Background(), `
		SELECT id, link, title, created_by_user_id, create_time, about, image_id 
		FROM public.chats 
		WHERE chats.link = $1
		`,
		chat_link).
		Scan(&chat.Id, &chat.Link, &chat.Title, &chat.User_id, &chat.Create_time, &chat.About, &chat.Image_id)

	if err != nil {
		return chat, err
	}

	return chat, nil
}

func SelectChatsByTitle(pool *pgx.Pool, title string) ([]md.Chat, error) {
	var chats []md.Chat

	rows, err := pool.Query(context.Background(), `
        SELECT id, link, title, created_by_user_id, create_time, about, image_id 
        FROM public.chats 
        WHERE chats.title = $1
		`,
		title)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var chat md.Chat
		err := rows.Scan(&chat.Id, &chat.Link, &chat.Title, &chat.User_id, &chat.Create_time, &chat.About, &chat.Image_id)
		if err != nil {
			return chats, err
		}
		chats = append(chats, chat)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return chats, nil
}

func SelectMessageById(pool *pgx.Pool, msg_id string) (md.Message, error) {
	var msg md.Message

	err := pool.QueryRow(context.Background(), `
		SELECT id, chat_id, user_id, create_time, text, reply_message_id
		FROM public.messages
		WHERE messages.id = $1
		`, msg_id).
		Scan(&msg.Id, &msg.Chat_id, &msg.User_id, &msg.Create_time, &msg.Text, &msg.Reply_msg_id)

	if err != nil {
		return msg, err
	}

	return msg, nil
}

func SelectMessagesByChatId(pool *pgx.Pool, request md.MessagesByChatIdRequest) ([]md.Message, error) {
	var msgs []md.Message

	query := `
		SELECT id, chat_id, user_id, create_time, text, reply_message_id
		FROM public.messages
		WHERE messages.chat_id = $1`

	// Add conditional clauses for FromDate and ToDate
	if request.From_date != nil {
		query += ` AND messages.create_time >= $2`
	}
	if request.To_date != nil {
		query += ` AND messages.create_time <= $3`
	}

	rows, err := pool.Query(context.Background(), query,
		request.Chat_id,
		request.From_date,
		request.To_date)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var msg md.Message
		err := rows.Scan(&msg.Id, &msg.Chat_id, &msg.User_id, &msg.Create_time, &msg.Text, &msg.Reply_msg_id)
		if err != nil {
			return msgs, err
		}
		msgs = append(msgs, msg)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return msgs, nil
}

func SelectMessagesByText(pool *pgx.Pool, text string) ([]md.Message, error) {
	var msgs []md.Message

	rows, err := pool.Query(context.Background(), `
		SELECT id, chat_id, user_id, create_time, text, reply_message_id
		FROM public.messages
		WHERE messages.text LIKE $1
		`,
		"%"+text+"%")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var msg md.Message
		err := rows.Scan(&msg.Id, &msg.Chat_id, &msg.User_id, &msg.Create_time, &msg.Text, &msg.Reply_msg_id)
		if err != nil {
			return msgs, err
		}
		msgs = append(msgs, msg)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return msgs, nil
}

func SelectMessagesByChatAndText(pool *pgx.Pool, chat_id string, text string) ([]md.Message, error) {
	var msgs []md.Message

	rows, err := pool.Query(context.Background(), `
		SELECT id, chat_id, user_id, create_time, text, reply_message_id
		FROM public.messages
		WHERE messages.chat_id = $1
			AND messages.text LIKE $2
		`,
		chat_id,
		"%"+text+"%")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var msg md.Message
		err := rows.Scan(&msg.Id, &msg.Chat_id, &msg.User_id, &msg.Create_time, &msg.Text, &msg.Reply_msg_id)
		if err != nil {
			return msgs, err
		}
		msgs = append(msgs, msg)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return msgs, nil
}

func SelectAuthByLoginHash(pool *pgx.Pool, login_hash string) (md.Auth, error) {
	var auth md.Auth

	err := pool.QueryRow(context.Background(), `
		SELECT login_hash, password_hash, email, user_id, auth_token, auth_expires
		FROM public.auth
		WHERE auth.login_hash = $1
		`, login_hash).
		Scan(&auth.Login_hash, &auth.Password_hash, &auth.Email, &auth.User_id, &auth.Auth_token, &auth.Auth_expires)

	if err != nil {
		return auth, err
	}

	return auth, nil
}

func SelectChatIdsByUserId(pool *pgx.Pool, user_id string) ([]string, error) {
	query := `
		SELECT chat_id
		FROM public.chat_participants
		WHERE user_id = $1
	`

	rows, err := pool.Query(context.Background(), query, user_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	chat_ids := make([]string, 0)
	for rows.Next() {
		var chat_id string
		err := rows.Scan(&chat_id)
		if err != nil {
			return chat_ids, err
		}
		chat_ids = append(chat_ids, chat_id)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return chat_ids, nil
}

func SelectUserIdsByChatId(pool *pgx.Pool, chat_id string) ([]string, error) {
	query := `
		SELECT user_id
		FROM public.chat_participants
		WHERE chat_id = $1
	`

	rows, err := pool.Query(context.Background(), query, chat_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users_id := make([]string, 0)
	for rows.Next() {
		var user_id string
		var _ string
		err := rows.Scan(&user_id)
		if err != nil {
			return users_id, err
		}
		users_id = append(users_id, user_id)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users_id, nil
}

func SelectAuthByToken(pool *pgx.Pool, token string) (md.Auth, error) {
	auth := md.Auth{}
	err := pool.QueryRow(context.Background(), `
		SELECT login_hash, password_hash, email, user_id, auth_token, auth_expires
		FROM public.auth
		WHERE auth.auth_token = $1
	`, token).
		Scan(&auth.Login_hash, &auth.Password_hash, &auth.Email, &auth.User_id, &auth.Auth_token, &auth.Auth_expires)
	if err != nil {
		return auth, err
	}
	return auth, nil
}
