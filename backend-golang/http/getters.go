package http

import (
	"time"
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	md "github.com/BMilkey/messenger/models"
)



func getChatsHandler(c *gin.Context) {
    // Implement logic to get list of chats for given user_id
    userID := c.Param("user_id")

    // Your logic here
	log.Trace(fmt.Sprintf("getChatsHandler %v", userID))
    // Example response
    chats := []md.Chat{
        {Id: "1", User_id: userID, Create_time: time.Now(), About: "Chat about something", Image_id: "123"},
        {Id: "2", User_id: userID, Create_time: time.Now(), About: "Another chat", Image_id: "456"},
    }
    
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