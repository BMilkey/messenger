package http

import (
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	return r
}

func StartServer() {
	r := setupRouter()
	r.Run(":8080")

}

func Test() {
	router := gin.Default()

    // Define routes
    router.GET("/get_chats/:user_id", getChatsHandler)
    router.GET("/get_messages/:chat_id", getMessagesHandler)
    router.GET("/get_chat_participants/:chat_id", getChatParticipantsHandler)
    router.POST("/get_images", getImagesHandler)
    router.POST("/get_files", getFilesHandler)

    // Start server
    router.Run(":8080")
}
