package http

import (
	docs "github.com/BMilkey/messenger/docs"
	"github.com/BMilkey/messenger/hlp"
	"github.com/gin-gonic/gin"
	pgx "github.com/jackc/pgx/v5/pgxpool"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func setupRouter(dbpool *pgx.Pool) *gin.Engine {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/post"
	post := r.Group("/post")
	{

		auth := post.Group("/auth")
		{
			auth.POST("/user_by_auth", func(c *gin.Context) {
				userByAuthHandler(c, dbpool)
			})
			auth.POST("/register_user", func(c *gin.Context) {
				registerUserHandler(c, dbpool)
			})
		}
		chat := post.Group("/chat")
		{
			chat.POST("/create_chat_return_users", func(c *gin.Context) {
				createChatReturnUsersHandler(c, dbpool)
			})
			chat.POST("/chats_by_token", func(c *gin.Context) {
				chatsByTokenHandler(c, dbpool)
			})
			chat.POST("/users_by_name", func(c *gin.Context) {
				usersByNameHandler(c, dbpool)
			})
			chat.POST("/users_by_chat_id", func(c *gin.Context) {
				usersByChatIdHandler(c, dbpool)
			})
			chat.POST("/messages_by_chat_id", func(c *gin.Context) {
				messagesByChatId(c, dbpool)
			})
			chat.POST("/create_message", func(c *gin.Context) {
				createMessageHandler(c, dbpool)
			})
			chat.POST("/add_user_to_chat", func(c *gin.Context) {
				addUserToChatHandler(c, dbpool)
			})
			chat.POST("/change_user_info", func(c *gin.Context) {
				changeUserInfoHandler(c, dbpool)
			})
		}
		test := post.Group("/test")
		{
			test.POST("/ping", func(c *gin.Context) {
				pingHandler(c, dbpool)
			})
		}
	}

	chat_sockets := r.Group("/sockets")
	{
		chat_sockets.GET("/test", func(c *gin.Context) {
			//c.Done()
			testWebSocketBroadCast(c.Writer, c.Request)
		})
		chat_sockets.GET("/subscribe_message_created", func(c *gin.Context) {
			subscribeMessageCreated(c.Writer, c.Request, dbpool)
		})
		chat_sockets.GET("/subscribe_сhat_сreated", func(c *gin.Context) {
			subscribeChatCreated(c.Writer, c.Request, dbpool)
		})
	}

	// Define routes

	// test shit
	r.POST("/testLoginWoHashHandler/", func(c *gin.Context) {
		testLoginWoHashHandler(c, dbpool)
	})

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return r
}

func StartServer(cfg hlp.HttpConfig, dbpool *pgx.Pool) {
	r := setupRouter(dbpool)

	r.Run(":" + cfg.Port)
	//r.RunTLS("0.0.0.0:"+cfg.Port, cfg.Certificate, cfg.Key)
}
