package app

import (
	"concurrency-chat/Logger"
	"concurrency-chat/app/handlers"
	"concurrency-chat/models"
	"github.com/gorilla/websocket"
	"net/http"
)

var (
	upgrader  = websocket.Upgrader{}
	broadcast = make(chan models.Message)
	clients   = make(map[*websocket.Conn]bool) // create a map to store connected clients
)

func HandleConnection(w http.ResponseWriter, r *http.Request) {
	// upgrade the http connection to a websocket connection
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		Logger.ErrorLogger().Printf("error upgrading connection: %v", err)
	}
	defer conn.Close()

	// add new client to map
	clients[conn] = true

	// loop to read for incoming messages from the client
	for {
		var message models.Message
		err := conn.ReadJSON(&message)
		if err != nil {
			Logger.ErrorLogger().Printf("error reading and parsing json: %v\n", err)
			delete(clients, conn)
			return
		}

		broadcast <- message
	}
}

func HandleMessages() {
	for {
		message := <-broadcast
		if handlers.ReceiveMessage(message) {
			for client := range clients {
				err := client.WriteJSON(message)
				if err != nil {
					Logger.ErrorLogger().Printf("error writing json to client: %v\n", err)
					delete(clients, client)
					return
				}
			}
		}
	}
}
