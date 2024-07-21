package http

import (
	md "github.com/BMilkey/messenger/models"

	pgx "github.com/jackc/pgx/v5/pgxpool"

	log "github.com/sirupsen/logrus"

	db "github.com/BMilkey/messenger/database"
)

func NotifyCreateMessageSubscribers(pool *pgx.Pool, message md.Message) error {
	userIds, err := db.SelectUserIdsByChatId(pool, message.Chat_id)
	if err != nil {
		log.Error(err)
		return err
	}

	for _, userId := range userIds {
		channel, ok := userIdToCreateMessageChannel.Get(userId)
		if !ok {
			continue
		}
		channel <- message
	}
	return nil
}
