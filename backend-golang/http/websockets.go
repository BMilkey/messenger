package http

import (
	"fmt"
	"net/http"

	cmap "github.com/orcaman/concurrent-map/v2"

	md "github.com/BMilkey/messenger/models"
	ws "github.com/gorilla/websocket"
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

var tokenToCreateMessageChannel = cmap.New[chan md.Message]()

// testWebSocketBroadCast   godoc
// @Summary                 Test websocket
// @Tags                    Test websocket
// @Description             НЕ РАБОТАЕТ ЧЕРЕЗ SWAGGER
// @ID                      sockets_test
// @Accept                  json-stream
// @Produce                 json-stream
// @Router                  /sockets/test [GET]
func SubscribeMessageCreated(w http.ResponseWriter, r *http.Request) {

	conn, err := wsupgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Warn("Failed to set websocket upgrade: %+v", err)
		return
	}

	if conn == nil {
		log.Warn("Connection is nil")
		return
	}

	_, requestToken, err := conn.ReadMessage()
	if err != nil {
		log.Warn("Cannot read start message in websocket")
		return
	}

	//	TODO get auth_token
	auth_token := "fake_auth_token"
LOOP:
	for {
		channel, ok := tokenToCreateMessageChannel.Get(auth_token)
		if !ok {
			log.Warn("Trying to get a message from tokenToCreateMessageChannel failed")
		}
		select {
		case <-r.Context().Done():
			break LOOP
		case msg := <-channel:

			serializedMsg, err := serializeToJSON(msg)
			if err != nil {
				log.Warn(err)
			}
			byteMsg := []byte(serializedMsg)
			conn.WriteMessage(ws.TextMessage, byteMsg)
		}
	}
	/*
			TokenToChannel.SetIfAbsent(auth_token, make(chan Message, 10))
			channel, ok := TokenToChannel.Get(auth_token)
			if !ok {
				return
			}

		BR:
			for {
				select {
				case isDone := <-isDoneChannel:
					if isDone {
						break BR
					}
					continue
				case msg := <-channel:
					println("auth_token: ", auth_token, ", text:", msg.Text)
				}

			}

			close(channel)
			TokenToChannel.Remove(auth_token)
	*/
	for i := 0; i < 10; i++ {
		msg := []byte(fmt.Sprintf("Message %d", i))
		conn.WriteMessage(ws.TextMessage, msg)
	}
	//conn.CloseHandler()
	conn.Close()
}
