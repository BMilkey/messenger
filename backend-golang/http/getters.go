package http

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	pgx "github.com/jackc/pgx/v5/pgxpool"
	log "github.com/sirupsen/logrus"

	db "github.com/BMilkey/messenger/database"
	md "github.com/BMilkey/messenger/models"
)

func userIdsByChatId(c *gin.Context, pool *pgx.Pool) {
	chatID := c.Param("chat_id")
	userIDs, err := db.SelectUserIdsByChatId(pool, chatID)
	if err != nil {
		log.Info(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to select users"})
		return
	}

	c.JSON(http.StatusOK, userIDs)
}


func chatIdsByUserId(c *gin.Context, pool *pgx.Pool) {
	userID := c.Param("user_id")
	chats, err := db.SelectChatIdsByUserId(pool, userID)
	if err != nil {
		log.Info(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to select chats"})
		return
	}

	c.JSON(http.StatusOK, chats)
}

func userByIdHandler(c *gin.Context, pool *pgx.Pool) {
	userID := c.Param("user_id")

	user, err := db.SelectUserById(pool, userID)
	if err != nil {
		log.Info(err)
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":          user.Id,
		"name":        user.Name,
		"link":        user.Link,
		"about":       user.About,
		"last_online": user.Last_connection,
		"image_id":    user.Image_id,
	})
}

func userIdByAuthHandler(c *gin.Context, pool *pgx.Pool) {
	login := c.Param("login")
	password := c.Param("password")

	log.Trace(fmt.Sprintf("getUserIdByAuthHandler %v %v", login, password))
	login_hash := db.Sha256Hash(login)
	password_hash := db.Sha256Hash(password)
	auth, err := db.SelectAuthByLoginHash(pool, login_hash)
	if err != nil {
		log.Info(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if auth.Password_hash != password_hash {
		log.Error("Incorrect password")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect password"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user_id": auth.User_id})
}

func registerUserHandler(c *gin.Context, pool *pgx.Pool) {
	var request struct {
		Login    string `json:"login"`
		Password string `json:"password"`
		Name     string `json:"name"`
	}
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedLogin := db.Sha256Hash(request.Login)
	hashedPassword := db.Sha256Hash(request.Password)
	newUserId := db.GenerateUUID()

	auth := md.Auth{
		Login_hash:    hashedLogin,
		Password_hash: hashedPassword,
		User_id:       newUserId,
	}
	user := md.User{
		Id:              newUserId,
		Name:            request.Name,
		Link:            "@" + request.Name,
		About:           "There's should be description for " + request.Name,
		Last_connection: time.Now(),
		Image_id:        nil,
	}

	err := db.InsertAuth(pool, auth)
	if err != nil {
		log.Info(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert auth"})
		db.DeleteAuthByUserId(pool, newUserId)
		return
	}

	err = db.InsertAuth(pool, auth)
	if err != nil {
		log.Info(err)
		db.DeleteAuthByUserId(pool, newUserId)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert user"})

		return
	}

	if err := db.InsertUser(pool, user); err != nil {
		log.Info(err)
		db.DeleteAuthByUserId(pool, newUserId)
		db.DeleteUserById(pool, newUserId)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to set user auth"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user_id": newUserId})
}
func getChatsHandler(c *gin.Context) {
	// Implement logic to get list of chats for given user_id
	userID := c.Param("user_id")

	// Your logic here
	log.Trace(fmt.Sprintf("getChatsHandler %v", userID))
	// Example response
	chats := []md.Chat{
		{Id: "1", User_id: userID, Create_time: time.Now(), About: "Chat about something", Image_id: new(string)},
		{Id: "2", User_id: userID, Create_time: time.Now(), About: "Another chat", Image_id: new(string)},
	}
	*chats[0].Image_id = "123"
	*chats[1].Image_id = "456"

	c.JSON(http.StatusOK, chats)
}

func getMessagesHandler(c *gin.Context) {
	// Implement logic to get list of messages for given chat_id
	chatID := c.Param("chat_id")

	// Your logic here
	log.Trace(fmt.Sprintf("getMessagesHandler %v", chatID))
	// Example response
	messages := []md.Message{
		{Id: "1", Chat_id: chatID, User_id: "user1", Create_time: time.Now(), Text: "Hello"},
		{Id: "2", Chat_id: chatID, User_id: "user2", Create_time: time.Now(), Text: "Hi there"},
	}

	c.JSON(http.StatusOK, messages)
}

func getChatParticipantsHandler(c *gin.Context) {
	// Implement logic to get list of chat participants for given chat_id
	chatID := c.Param("chat_id")

	// Your logic here
	log.Trace(fmt.Sprintf("getChatParticipantsHandler %v", chatID))
	// Example response
	participants := []md.User{
		{Id: "user1", Name: "John", Link: "@user1", About: "About John", Last_connection: time.Now()},
		{Id: "user2", Name: "Jane", Link: "@user2", About: "About Jane", Last_connection: time.Now()},
	}

	c.JSON(http.StatusOK, participants)
}

func getImagesHandler(c *gin.Context) {
	// Implement logic to get list of paths to images for given list of image_ids
	// Example request body: {"image_ids": ["1", "2", "3"]}
	var request struct {
		ImageIDs []string `json:"image_ids"`
	}
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Your logic here
	log.WithFields(log.Fields{"paths": request.ImageIDs}).Trace("getImagesHandler")
	// Example response
	images := []md.Image{
		{Id: "1", Name: "krasivaya-ava-devushke-50.jpg", Size: 15000, Path: "/https://lifeo.ru/wp-content/uploads/krasivaya-ava-devushke-50.jpg"},
		{Id: "2", Name: "image2.jpg", Size: 10000, Path: "/path/to/image2.jpg"},
	}

	c.JSON(http.StatusOK, images)
}

func getFilesHandler(c *gin.Context) {
	// Implement logic to get list of paths to files for given list of file_ids
	// Example request body: {"file_ids": ["1", "2", "3"]}
	var request struct {
		FileIDs []string `json:"file_ids"`
	}
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Your logic here
	log.WithFields(log.Fields{"paths": request.FileIDs}).Trace("getImagesHandler")
	// Example response
	files := []md.File{
		{Id: "1", Name: "peach-cat-yay.gif", Size: 66666, Path: "https://media.tenor.com/nLKPVCtMbioAAAAi/peach-cat-yay.gif"},
		{Id: "2", Name: "file2.txt", Size: 12345, Path: "/path/to/file2.txt"},
	}

	c.JSON(http.StatusOK, files)
}
