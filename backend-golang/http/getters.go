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

// pingHandler		 	godoc
// @Summary 			Ping
// @Tags 				Test
// @Description 		Пинг
// @ID 					ping
// @Accept  			json
// @Produce  			json
// @Success 			200
// @Router 				/test/ping [post]
func pingHandler(c *gin.Context, pool *pgx.Pool) {
	c.JSON(http.StatusOK, "")
}

// userByAuthHandler 	godoc
// @Summary 			Auth user
// @Tags 				API для авторизации и регистрации пользователя
// @Description 		Авторизация пользователя
// @ID 					user_by_auth
// @Accept  			json
// @Produce  			json
// @Param 				input body md.SignInRequest true "credentials"
// @Success 			200 {object} md.SignInResponse "data"
// @Router 				/auth/user_by_auth [post]
func userByAuthHandler(c *gin.Context, pool *pgx.Pool) {
	var request md.SignInRequest

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

	c.JSON(http.StatusOK, md.SignInResponse{
		Auth_token:  auth.Auth_token,
		Name:        user.Name,
		Link:        user.Link,
		About:       user.About,
		Last_online: user.Last_connection,
		Image_id:    user.Image_id,
	})
}

// registerUserHandler 	godoc
// @Summary 			Register user
// @Tags 				API для авторизации и регистрации пользователя
// @Description 		Регистрация пользователя
// @ID 					register_user
// @Accept  			json
// @Produce  			json
// @Param 				input body md.SignUpRequest true "credentials"
// @Success 			200 {object} md.SignInResponse "data"
// @Router 				/auth/register_user [post]
func registerUserHandler(c *gin.Context, pool *pgx.Pool) {
	var request md.SignUpRequest

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

	c.JSON(http.StatusOK, md.SignInResponse{
		Auth_token:  auth.Auth_token,
		Name:        user.Name,
		Link:        user.Link,
		About:       user.About,
		Last_online: user.Last_connection,
		Image_id:    user.Image_id,
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

// createChatReturnUsersHandler 	godoc
// @Summary 						Create chat
// @Tags 							API для работы с чатами и сообщениями
// @Description 					Создать чат
// @ID 								create_chat_return_users
// @Accept  						json
// @Produce  						json
// @Param 							input body md.CreateChatRequest true "credentials"
// @Success 						200 {object} md.ChatUsers "data"
// @Router 							/chat/create_chat_return_users [post]
func createChatReturnUsersHandler(c *gin.Context, pool *pgx.Pool) {
	var request md.CreateChatRequest

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
	create_user, err := db.SelectUserById(pool, auth.User_id)
	if err != nil {
		log.Info(err)
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("User_id for chat create %s not found", auth.User_id)})
		return
	}
	err = db.InsertChatParticipant(pool, chat.Id, create_user.Id)
	if err != nil {
		log.Info(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert chat participant"})
		return
	}
	users = append(users, create_user)

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
	if !prolongToken(c, pool, request.Auth_token) {
		return
	}

	c.JSON(http.StatusOK, md.ChatUsers{
		Chat_id: chat.Id,
		Users:   users,
	})

}

// chatsByTokenHandler 	godoc
// @Summary 			Get chats by auth token
// @Tags 				API для работы с чатами и сообщениями
// @Description 		Чаты по токену авторизации
// @ID 					chats_by_token
// @Accept  			json
// @Produce  			json
// @Param 				input body md.ByTokenRequest true "credentials"
// @Success 			200 {object} md.Chats "data"
// @Router 				/chat/chats_by_token [post]
func chatsByTokenHandler(c *gin.Context, pool *pgx.Pool) {
	var request md.ByTokenRequest
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

	if !prolongToken(c, pool, request.Auth_token) {
		return
	}

	c.JSON(http.StatusOK, md.Chats{
		Chats: chats,
	})

}

// usersByNameHandler 	godoc
// @Summary 			Get users by name
// @Tags 				API для работы с чатами и сообщениями
// @Description 		Юзеры по имени
// @ID 					users_by_name
// @Accept  			json
// @Produce  			json
// @Param 				input body md.UsersByNameRequest true "credentials"
// @Success 			200 {object} md.Users "data"
// @Router 				/chat/users_by_name [post]
func usersByNameHandler(c *gin.Context, pool *pgx.Pool) {
	var request md.UsersByNameRequest

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

	if !prolongToken(c, pool, request.Auth_token) {
		return
	}

	c.JSON(http.StatusOK, md.Users{
		Users: users,
	})
}

// usersByChatIdHandler 	godoc
// @Summary 				Get users by chat_id
// @Tags 					API для работы с чатами и сообщениями
// @Description 			Юзеры по ИД чата
// @ID 						users_by_chat_id
// @Accept  				json
// @Produce  				json
// @Param 					input body md.ByChatIdRequest true "credentials"
// @Success 				200 {object} md.Users "data"
// @Router 					/chat/users_by_chat_id [post]
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

	if !prolongToken(c, pool, request.Auth_token) {
		return
	}

	c.JSON(http.StatusOK, md.Users{
		Users: users,
	})
}

// messagesByChatId		 	godoc
// @Summary 				Get messages by chat_id
// @Tags 					API для работы с чатами и сообщениями
// @Description 			Сообщения по ИД чата
// @ID 						messages_by_chat_id
// @Accept  				json
// @Produce  				json
// @Param 					input body md.MessagesByChatIdRequest true "credentials"
// @Success 				200 {object} md.Messages "data"
// @Router 					/chat/messages_by_chat_id [post]
func messagesByChatId(c *gin.Context, pool *pgx.Pool) {
	var request md.MessagesByChatIdRequest

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !isTokenHaveAccessToChat(c, pool, request.Chat_id, request.Auth_token) {
		return
	}

	messages, err := db.SelectMessagesByChatId(pool, request)
	if err != nil {
		log.Info(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to select messages"})
		return
	}

	for i := 0; i < len(messages); i++ {
		user, err := db.SelectUserById(pool, messages[i].User_id)
		if err != nil {
			log.Info(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to select user"})
			return
		}
		messages[i].User_id = user.Link
	}

	if !prolongToken(c, pool, request.Auth_token) {
		return
	}

	c.JSON(http.StatusOK, md.Messages{
		Messages: messages,
	})

}

// createMessageHandler		godoc
// @Summary 				Create message
// @Tags 					API для работы с чатами и сообщениями
// @Description 			Создать сообщение
// @ID 						create_message
// @Accept  				json
// @Produce  				json
// @Param 					input body md.CreateMessageRequest true "credentials"
// @Success 				200 {object} md.CreateMessageResponse "data"
// @Router 					/chat/create_message [post]
func createMessageHandler(c *gin.Context, pool *pgx.Pool) {
	var request md.CreateMessageRequest

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

	c.JSON(http.StatusOK, md.CreateMessageResponse{
		Message:   message,
		Reply_msg: reply_msg,
	})
}

// addUserToChatHandler		godoc
// @Summary 				Add user to chat
// @Tags 					API для работы с чатами и сообщениями
// @Description 			Добавить юзера в чат
// @ID 						add_user_to_chat
// @Accept  				json
// @Produce  				json
// @Param 					input body md.AddUserToChatRequest true "credentials"
// @Success 				200 {object} md.User "data"
// @Router 					/chat/add_user_to_chat [post]
func addUserToChatHandler(c *gin.Context, pool *pgx.Pool) {
	var request md.AddUserToChatRequest

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

	c.JSON(http.StatusOK, md.User(user))

}

// changeUserInfoHandler	godoc
// @Summary 				Change user info
// @Tags 					API для работы с чатами и сообщениями
// @Description 			Изменить информацию о юзере
// @ID 						change_user_info
// @Accept  				json
// @Produce  				json
// @Param 					input body md.ChangeUserInfoRequest true "credentials"
// @Success 				200 {object} md.User "data"
// @Router 					/chat/change_user_info [post]
func changeUserInfoHandler(c *gin.Context, pool *pgx.Pool) {
	var request md.ChangeUserInfoRequest

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

	c.JSON(http.StatusOK, md.User(user))
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
