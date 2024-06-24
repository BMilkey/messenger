package http

import (
	"fmt"
	"net/http"

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
