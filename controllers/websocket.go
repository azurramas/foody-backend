package controllers

import (
	"bytes"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]bool) // connected clients

//WsHandler ->
func WsHandler(w http.ResponseWriter, r *http.Request) {

	var upgrader = websocket.Upgrader{}
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	// upgrade this connection to a WebSocket
	// connection
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrade:", err)
		return
	}

	defer c.Close()

	clients[c] = true

	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			fmt.Println(err)
			break
		}

		if bytes.Equal(message, []byte("2")) {
			for client := range clients {
				err := client.WriteMessage(mt, message)
				if err != nil {
					log.Printf("error: %v", err)
					client.Close()
					delete(clients, client)
				}
			}
		}
	}
}
