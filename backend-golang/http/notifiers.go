package http

import (
	md "github.com/BMilkey/messenger/models"

	pgx "github.com/jackc/pgx/v5/pgxpool"

	log "github.com/sirupsen/logrus"

	db "github.com/BMilkey/messenger/database"

	cmap "github.com/orcaman/concurrent-map/v2"
)

var userIdToMessageBroker = cmap.New[Broker[md.Message]]()

// setup for multiple listeners (multiple clients with same user_id)
func NotifyCreateMessageSubscribers(pool *pgx.Pool, message md.Message) error {
	userIds, err := db.SelectUserIdsByChatId(pool, message.Chat_id)
	if err != nil {
		log.Error(err)
		return err
	}

	for _, userId := range userIds {
		broker, ok := userIdToMessageBroker.Get(userId)
		if !ok {
			log.Warn("User ", userId, " not found message broker")
			continue
		}
		broker.Publish(message)
	}
	return nil
}

var userIdToChatBroker = cmap.New[Broker[md.ChatInfo]]()

// setup for multiple listeners (multiple clients with same user_id)
func NotifyCreateChatSubscribers(pool *pgx.Pool, chat md.ChatInfo) error {
	userIds, err := db.SelectUserIdsByChatId(pool, chat.Id)
	if err != nil {
		log.Error(err)
		return err
	}
	for _, userId := range userIds {
		broker, ok := userIdToChatBroker.Get(userId)
		if !ok {
			log.Warn("User ", userId, " not found chat broker")
			continue
		}
		broker.Publish(chat)
	}
	return nil
}
