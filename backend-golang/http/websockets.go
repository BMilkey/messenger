package http

import (
	"fmt"
	"net/http"
	"time"

	db "github.com/BMilkey/messenger/database"
	md "github.com/BMilkey/messenger/models"
	ws "github.com/gorilla/websocket"
	pgx "github.com/jackc/pgx/v5/pgxpool"
	cmap "github.com/orcaman/concurrent-map/v2"
	log "github.com/sirupsen/logrus"
)

var wsupgrader = ws.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// testWebSocketBroadCast   godoc
// @Summary                 Test websocket
// @Tags                    Test websocket
// @Description             НЕ РАБОТАЕТ ЧЕРЕЗ SWAGGER
// @Description Код, которым тестил
// @Description let socket = new WebSocket("ws://147.45.70.245:80/sockets/test");
// @Description
// @Description socket.onopen = function(e) {
// @Description     console.log("[open] Connection established");
// @Description     socket.send("Hello server!");
// @Description };
// @Description
// @Description socket.onmessage = function(event) {
// @Description     console.log(`[message] Data received from server: ${event.data}`);
// @Description };
// @Description
// @Description socket.onclose = function(event) {
// @Description     if (event.wasClean) {
// @Description         console.log(`[close] Connection closed cleanly, code=${event.code} reason=${event.reason}`);
// @Description     } else {
// @Description         console.log('[close] Connection died');
// @Description     }
// @Description };
// @Description
// @Description socket.onerror = function(error) {
// @Description     console.log(`[error] ${error.message}`);
// @Description };
// @ID                      sockets_test
// @Accept                  json-stream
// @Produce                 json-stream
// @Router                  /sockets/test [GET]
func testWebSocketBroadCast(w http.ResponseWriter, r *http.Request) {
	conn, err := wsupgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Failed to set websocket upgrade: %+v", err)
		return
	}

	if conn == nil {
		log.Println("Connection is nil")
		return
	}

	for i := 0; i < 10; i++ {
		msg := []byte(fmt.Sprintf("Message %d", i))
		conn.WriteMessage(ws.TextMessage, msg)
	}
	//conn.CloseHandler()
	conn.Close()
}

//var messagesChannel := make(chan md.Message, 100)

var userIdToCreateMessageChannel = cmap.New[chan md.Message]()
var sameUserMultipleListeners = cmap.New[chan bool]()

//var auth

// subscribeMessageCreated 	godoc
// @Summary 			Subscribe to all messages, that you should receive
// @Tags 				WebSocket API для подписок
// @Description 		Подписка на получаемые сообщения
// @ID 					subscribe_message_created
// @Accept  			json
// @Produce  			json
// @Param 				input body string true "credentials"
// @Success 			200 {object} md.Message "data"
// @Router 				/sockets/subscribe_message_created [get]
func subscribeMessageCreated(w http.ResponseWriter, r *http.Request, pool *pgx.Pool) {

	conn, err := wsupgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Warn("Failed to set websocket upgrade: %+v", err)
		return
	}

	if conn == nil {
		log.Warn("Connection is nil")
		return
	}
	defer conn.Close()

	_, requestAuthToken, err := conn.ReadMessage()
	if err != nil {
		log.Warn("Cannot read start message in websocket")
		return
	}

	userId, err := db.SelectUserIdByAuthToken(pool, string(requestAuthToken))
	if err != nil {
		log.Warn("Cannot get userId by auth token")
		return
	}
	// userId := ...
	b, ok := sameUserMultipleListeners.Get(userId)
	if ok {
		b <- true
	}

	sameUserMultipleListeners.SetIfAbsent(userId, make(chan bool, 1))

	userIdToCreateMessageChannel.SetIfAbsent(userId, make(chan md.Message, 10))
	defer func() { userIdToCreateMessageChannel.Remove(userId) }()

LOOP:
	for {
		channel, ok := userIdToCreateMessageChannel.Get(userId)
		if !ok {
			log.Warn("Trying to get a message from tokenToCreateMessageChannel failed")
			return
		}
		sameUser, ok := sameUserMultipleListeners.Get(userId)
		if !ok {
			log.Warn("Trying to get a message from sameUserMultipleListeners failed")
			return
		}

		select {
		case <-sameUser:
			sameUserMultipleListeners.Remove(userId)
			break LOOP
		case <-time.After(time.Second * 120):
			break LOOP
		case <-r.Context().Done():
			break LOOP
		case msg := <-channel:

			serializedMsg, err := serializeToJSON(msg)
			if err != nil {
				log.Warn(err)
			}
			byteMsg := []byte(serializedMsg)
			conn.WriteMessage(ws.TextMessage, byteMsg)
			log.Info("Sent to client ", conn.RemoteAddr(), "with token", string(requestAuthToken), ":", string(serializedMsg))
		}
	}

}
