package http

import (
	"github.com/gin-gonic/gin"
	pgx "github.com/jackc/pgx/v5/pgxpool"

	"github.com/BMilkey/messenger/hlp"
)

func setupRouter(dbpool *pgx.Pool) *gin.Engine {
	r := gin.Default()

	// Define routes
	r.POST("/post/auth/user_by_auth/", func(c *gin.Context) {
		userByAuthHandler(c, dbpool)
	})
	r.POST("/post/auth/register_user/", func(c *gin.Context) {
		registerUserHandler(c, dbpool)
	})
	r.POST("/post/chat/chats_by_token/", func(c *gin.Context) {
		chatsByTokenHandler(c, dbpool)
	})
	/*
		r.POST("/post/chat/user_ids_by_chat_id/", func(c *gin.Context) {
			userIdsByChatIdHandler(c, dbpool)
		})
	*/
	r.POST("/post/chat/create_chat_return_users/", func(c *gin.Context) {
		createChatReturnUsersHandler(c, dbpool)
	})
	r.POST("/post/chat/users_by_name/", func(c *gin.Context) {
		usersByNameHandler(c, dbpool)
	})
	r.POST("/post/chat/users_by_chat_id/", func(c *gin.Context) {
		usersByChatIdHandler(c, dbpool)
	})
	r.POST("/post/chat/messages_by_chat_id/", func(c *gin.Context) {
		messagesByChatId(c, dbpool)
	})
	r.POST("/post/chat/create_message/", func(c *gin.Context) {
		createMessageHandler(c, dbpool)
	})
	r.POST("/post/chat/add_user_to_chat/", func(c *gin.Context) {
		addUserToChatHandler(c, dbpool)
	})
	r.POST("/post/chat/change_user_info/", func(c *gin.Context) {
		changeUserInfoHandler(c, dbpool)
	})

	
	// test shit
	r.POST("/testLoginWoHashHandler/", func(c *gin.Context) {
		testLoginWoHashHandler(c, dbpool)
	})

	return r
}

func StartServer(cfg hlp.HttpConfig, dbpool *pgx.Pool) {
	r := setupRouter(dbpool)
	r.Run(":" + cfg.Port)
}
