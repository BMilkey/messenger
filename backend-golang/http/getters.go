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

func userByAuthHandler(c *gin.Context, pool *pgx.Pool) {
	var request struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	login_hash := db.Sha256Hash(request.Login)
	password_hash := db.Sha256Hash(request.Password)
	auth, err := db.SelectAuthByLoginHash(pool, login_hash)
	if err != nil {
		log.Info(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	auth.Auth_token = db.GenerateUUID()

	if auth.Password_hash != password_hash {
		log.Error("Incorrect password or login")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect password or login"})
		return
	}

	user, err := db.SelectUserById(pool, auth.User_id)
	if err != nil {
		log.Info(err)
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if err := db.UpdateAuthToken(pool, auth); err != nil {
		log.Info(err)
		db.DeleteAuthByUserId(pool, user.Id)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert auth"})
		return
	}

	if !prolongToken(c, pool, auth.Auth_token) {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"auth_token":  auth.Auth_token,
		"name":        user.Name,
		"link":        user.Link,
		"about":       user.About,
		"last_online": user.Last_connection,
		"image_id":    user.Image_id,
	})
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
	authToken := db.GenerateUUID()

	auth := md.Auth{
		Login_hash:    hashedLogin,
		Password_hash: hashedPassword,
		Email:         db.GenerateUUID(),
		User_id:       newUserId,
		Auth_token:    authToken,
		Auth_expires:  time.Now().Add(time.Minute * 1),
	}
	user := md.User{
		Id:              newUserId,
		Name:            request.Name,
		Link:            "@" + newUserId,
		About:           "There's should be description for " + request.Name,
		Last_connection: time.Now(),
		Image_id:        "fake",
	}

	if err := db.InsertUser(pool, user); err != nil {
		log.Info(err)
		// Rollback auth insertion

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert user"})
		return
	}

	if err := db.InsertAuth(pool, auth); err != nil {
		log.Info(err)
		db.DeleteAuthByUserId(pool, user.Id)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert auth"})
		return
	}

	if !prolongToken(c, pool, auth.Auth_token) {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"auth_token":  authToken,
		"name":        user.Name,
		"link":        user.Link,
		"about":       user.About,
		"last_online": user.Last_connection,
		"image_id":    user.Image_id,
	})
}

/*
	func userByIdHandler(c *gin.Context, pool *pgx.Pool) {
		var request struct {
			UserID     string `json:"user_id"`
			Auth_token string `json:"auth_token"`
		}
		if err := c.BindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if !validateAuthToken(pool, request.Auth_token) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid auth token"})
			return
		}

		user, err := db.SelectUserById(pool, request.UserID)
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
*/
func createChatReturnUsersHandler(c *gin.Context, pool *pgx.Pool) {
	var request struct {
		Auth_token  string   `json:"auth_token"`
		Title       string   `json:"title"`
		Users_links []string `json:"users_links"`
	}

	if err := c.BindJSON(&request); err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if !validateAuthToken(pool, request.Auth_token) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid auth token"})
		return
	}

	auth, err := db.SelectAuthByToken(pool, request.Auth_token)
	if err != nil {
		log.Info(err)
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	chat := md.Chat{
		Id:          db.GenerateUUID(),
		Link:        "@" + db.GenerateUUID(),
		Title:       request.Title,
		User_id:     auth.User_id,
		Create_time: time.Now(),
		About:       "There's should be description for " + request.Title,
		Image_id:    "fake",
	}

	if err := db.InsertChat(pool, chat); err != nil {
		log.Info(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert chat"})
		return
	}

	users := make([]md.User, 0)
	for i := 0; i < len(request.Users_links); i++ {
		user, err := db.SelectUserByLink(pool, request.Users_links[i])
		if err != nil {
			log.Info(err)
			c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("User %s not found", request.Users_links[i])})
			return
		}
		err = db.InsertChatParticipant(pool, chat.Id, user.Id)
		if err != nil {
			log.Info(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert chat participant"})
			return
		}
		users = append(users, user)
	}
	users = emptyUsersId(users)
	// Serialize the users array into JSON
	jsonData, err := serializeToJSON(users)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to serialize users"})
		return
	}

	if !prolongToken(c, pool, request.Auth_token) {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"users": string(jsonData),
	})

}

func chatsByTokenHandler(c *gin.Context, pool *pgx.Pool) {
	var request struct {
		Auth_token string `json:"auth_token"`
	}
	if err := c.BindJSON(&request); err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if !validateAuthToken(pool, request.Auth_token) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid auth token"})
		return
	}

	auth, err := db.SelectAuthByToken(pool, request.Auth_token)
	if err != nil {
		log.Info(err)
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	userID := auth.User_id

	chat_ids, err := db.SelectChatIdsByUserId(pool, userID)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to select chats"})
		return
	}
	chats := make([]md.Chat, 0)
	for i := 0; i < len(chat_ids); i++ {
		chat, err := db.SelectChatById(pool, chat_ids[i])
		if err != nil {
			log.Error(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to select chats"})
			return
		}
		user, err := db.SelectUserById(pool, chat.User_id)
		if err != nil {
			log.Error(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to select chats"})
			return
		}
		chat.User_id = user.Link
		chats = append(chats, chat)
	}

	jsonData, err := serializeToJSON(chats)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to serialize chats"})
		return
	}

	if !prolongToken(c, pool, request.Auth_token) {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"chats": string(jsonData),
	})

}

func usersByNameHandler(c *gin.Context, pool *pgx.Pool) {
	var request struct {
		Name       string `json:"name"`
		Auth_token string `json:"auth_token"`
	}
	if err := c.BindJSON(&request); err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if !validateAuthToken(pool, request.Auth_token) {
		log.Error("Invalid auth token")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid auth token"})
		return
	}

	users, err := db.SelectUsersByName(pool, request.Name)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to select users"})
		return
	}

	users = emptyUsersId(users)

	jsonData, err := serializeToJSON(users)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to serialize users"})
		return
	}

	if !prolongToken(c, pool, request.Auth_token) {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"users": string(jsonData),
	})
}

func usersByChatIdHandler(c *gin.Context, pool *pgx.Pool) {
	var request struct {
		ChatID     string `json:"chat_id"`
		Auth_token string `json:"auth_token"`
	}
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !isTokenHaveAccessToChat(c, pool, request.ChatID, request.Auth_token) {
		return
	}

	userIDs, err := db.SelectUserIdsByChatId(pool, request.ChatID)
	if err != nil {
		log.Info(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to select users"})
		return
	}

	users := make([]md.User, 0)
	for i := 0; i < len(userIDs); i++ {
		user, err := db.SelectUserById(pool, userIDs[i])
		if err != nil {
			log.Info(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to select users"})
			return
		}
		users = append(users, user)
	}

	users = emptyUsersId(users)

	jsonData, err := serializeToJSON(users)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to serialize users"})
		return
	}

	if !prolongToken(c, pool, request.Auth_token) {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"users": string(jsonData),
	})
}

func messagesByChatId(c *gin.Context, pool *pgx.Pool) {
	var request struct {
		ChatID     string `json:"chat_id"`
		Auth_token string `json:"auth_token"`
	}

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !isTokenHaveAccessToChat(c, pool, request.ChatID, request.Auth_token) {
		return
	}

	messages, err := db.SelectMessagesByChatId(pool, request.ChatID)
	if err != nil {
		log.Info(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to select messages"})
		return
	}

	jsonData, err := serializeToJSON(messages)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to serialize messages"})
		return
	}

	if !prolongToken(c, pool, request.Auth_token) {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"messages": string(jsonData),
	})

}

func createMessageHandler(c *gin.Context, pool *pgx.Pool) {
	var request struct {
		Chat_id      string `json:"chat_id"`
		Text         string `json:"text"`
		Auth_token   string `json:"auth_token"`
		Reply_msg_id string `json:"reply_msg_id"`
	}

	if err := c.BindJSON(&request); err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !isTokenHaveAccessToChat(c, pool, request.Chat_id, request.Auth_token) {
		return
	}

	auth, err := db.SelectAuthByToken(pool, request.Auth_token)
	if err != nil {
		log.Info(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to select user"})
		return
	}
	reply_msg_id := "fake"
	reply_msg, err := db.SelectMessageById(pool, request.Reply_msg_id)

	if err == nil {
		reply_msg_id = reply_msg.Id
	}

	message := md.Message{
		Id:           db.GenerateUUID(),
		Chat_id:      request.Chat_id,
		User_id:      auth.User_id,
		Create_time:  time.Now(),
		Text:         request.Text,
		Reply_msg_id: reply_msg_id,
	}

	err = db.InsertMessage(pool, message)
	if err != nil {
		log.Error(err)
	}

	if !prolongToken(c, pool, request.Auth_token) {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":   message,
		"reply_msg": reply_msg,
	})
}

func addUserToChatHandler(c *gin.Context, pool *pgx.Pool) {
	var request struct {
		Chat_id    string `json:"chat_id"`
		User_link  string `json:"user_link"`
		Auth_token string `json:"auth_token"`
	}

	if err := c.BindJSON(&request); err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !isTokenHaveAccessToChat(c, pool, request.Chat_id, request.Auth_token) {
		return
	}

	user, err := db.SelectUserByLink(pool, request.User_link)
	if err != nil {
		log.Info(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to select user"})
		return
	}

	err = db.InsertChatParticipant(pool, request.Chat_id, user.Id)
	if err != nil {
		log.Info(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add user to chat"})
		return
	}

	if !prolongToken(c, pool, request.Auth_token) {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})

}

func changeUserInfoHandler(c *gin.Context, pool *pgx.Pool) {
	var request struct {
		Auth_token string `json:"auth_token"`
		New_name   string `json:"new_name"`
		New_link   string `json:"new_link"`
		New_about  string `json:"new_about"`
		New_image  string `json:"new_image"`
	}

	if err := c.BindJSON(&request); err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	auth, err := db.SelectAuthByToken(pool, request.Auth_token)
	if err != nil {
		log.Info(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user := md.User{
		Id:              auth.User_id,
		Name:            request.New_name,
		Link:            request.New_link,
		About:           request.New_about,
		Last_connection: time.Now(),
		Image_id:        request.New_image,
	}
	err = db.UpdateUser(pool, user)
	if err != nil {
		log.Info(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	prolongToken(c, pool, request.Auth_token)

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

// test shit

func testLoginWoHashHandler(c *gin.Context, pool *pgx.Pool) {
	var request struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//login_hash := db.Sha256Hash(request.Login)
	//password_hash := db.Sha256Hash(request.Password)
	auth, err := db.SelectAuthByLoginHash(pool, request.Login)
	if err != nil {
		log.Info(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if request.Password != auth.Password_hash {
		log.Error("Incorrect password or login")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect password or login"})
		return
	}

	user, err := db.SelectUserById(pool, auth.User_id)
	if err != nil {
		log.Info(err)
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"name":        user.Name,
		"link":        user.Link,
		"about":       user.About,
		"last_online": user.Last_connection,
		"image_id":    user.Image_id,
	})
}
