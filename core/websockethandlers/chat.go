package websockethandlers

import (
	"context"
	"net/http"
	"log"

	"github.com/gorilla/websocket"

	"github.com/Tauhid-UAP/golang-sample-web-app/core/chat"
	"github.com/Tauhid-UAP/golang-sample-web-app/core/middleware"
	"github.com/Tauhid-UAP/golang-sample-web-app/core/redisclient"
)

func ChatHandler(websocketUpgrader websocket.Upgrader, hub *chat.Hub) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		roomName := r.URL.Query().Get("roomName")
		if roomName == "" {
			http.Error(w, "roomName required", http.StatusBadRequest)
			return
		}

		websocketConnection, err := websocketUpgrader.Upgrade(w, r, nil)
		if err != nil {
			return
		}

		userID := r.Context().Value(middleware.UserIDKey).(string)
		client := &chat.Client{
			Conn: websocketConnection,
			Receiver: make(chan []byte, 256),
			UserID: userID,
			RoomName: roomName,
		}

		ctx := context.Background()
		room := hub.GetOrCreateRoom(ctx, roomName)
		room.Register <- client

		go client.ReceiveMessages()

		for {
			_, message, err := websocketConnection.ReadMessage()
			log.Printf("Message: %s", string(message))
			if err != nil {
				log.Printf("Error reading websocket message: %v", err)
				break
			}

			payload := []byte(userID + ": " + string(message))
			log.Printf("Publishing to room: %s", payload)
			redisclient.PublishToRoom(ctx, roomName, payload)
		}
		
		log.Printf("Unregistering client")
		room.Unregister <- client
	}
}
