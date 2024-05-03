package http

import (
	"github.com/gin-gonic/gin"
	pgx "github.com/jackc/pgx/v5/pgxpool"

	"github.com/BMilkey/messenger/hlp"
)

func setupRouter(dbpool *pgx.Pool) *gin.Engine {
	r := gin.Default()

	// Define routes
	r.POST("/post/auth/user_id_by_auth/", func(c *gin.Context) {
		userIdByAuthHandler(c, dbpool)
	})
	r.POST("/post/auth/register_user/", func(c *gin.Context) {
		registerUserHandler(c, dbpool)
	})
	r.POST("/post/auth/login_user/", func(c *gin.Context) {
		userByIdHandler(c, dbpool)
	})
	r.POST("/post/auth/logout_user/", func(c *gin.Context) {
		chatIdsByUserId(c, dbpool)
	})
	r.POST("/post/chat/chat_ids_by_user_id/", func(c *gin.Context) {
		chatIdsByUserId(c, dbpool)
	})
	r.POST("/post/chat/user_ids_by_chat_id/", func(c *gin.Context) {
		userIdsByChatId(c, dbpool)
	})

	return r
}

func StartServer(cfg hlp.HttpConfig, dbpool *pgx.Pool) {
	r := setupRouter(dbpool)
	r.Run(":" + cfg.Port)
}

func Test() {
	router := gin.Default()

	router.GET("/get_chats/:user_id", getChatsHandler)
	router.GET("/get_messages/:chat_id", getMessagesHandler)
	router.GET("/get_chat_participants/:chat_id", getChatParticipantsHandler)
	router.POST("/get_images", getImagesHandler)
	router.POST("/get_files", getFilesHandler)

	// Start server
	router.Run(":8080")
}
