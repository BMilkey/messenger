package http

import (
	"encoding/json"
	"net/http"
	"time"
	db "github.com/BMilkey/messenger/database"
	md "github.com/BMilkey/messenger/models"
	"github.com/gin-gonic/gin"
	pgx "github.com/jackc/pgx/v5/pgxpool"
	log "github.com/sirupsen/logrus"
)

func prolongToken(c *gin.Context, pool *pgx.Pool, token string) bool {
	auth, err := db.SelectAuthByToken(pool, token)
	if err != nil {
		log.Info(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to select user"})
		return false
	}
	err = db.UpdateAuthToken(pool, auth)
	if err != nil {
		log.Info(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update auth token"})
		return false
	}
	user, err := db.SelectUserById(pool, auth.User_id)
	if err != nil {
		log.Info(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to select user"})
		return false
	}
	user.Last_connection = time.Now()
	err = db.UpdateUser(pool, user)
	if err != nil {
		log.Info(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return false
	}
	
	return true
}

func validateAuthToken(pool *pgx.Pool, token string) bool {

	savedToken, err := db.SelectAuthByToken(pool, token)
	if err != nil {
		return false
	}
	if token != savedToken.Auth_token {
		return false
	}
	return true

}

func serializeToJSON(data interface{}) ([]byte, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}

func emptyUsersId(users []md.User) []md.User {
	for i := 0; i < len(users); i++ {
		users[i].Id = ""
	}
	return users
}

func isTokenHaveAccessToChat(c *gin.Context, pool *pgx.Pool, chat_id string, token string) bool {
	if !validateAuthToken(pool, token) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid auth token"})
		return false
	}

	userIDs, err := db.SelectUserIdsByChatId(pool, chat_id)
	if err != nil {
		log.Info(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to select users"})
		return false
	}

	auth, err := db.SelectAuthByToken(pool, token)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return false
	}

	flag := false
	for i := 0; i < len(userIDs); i++ {
		if userIDs[i] == auth.User_id {
			flag = true
		}
	}

	if !flag {
		log.Error("You're not in this chat group")
		c.JSON(http.StatusForbidden, gin.H{"error": "You're not in this chat group"})
		return false
	}
	return true
}
