package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/vrnraju/doodle-weavers/internal/hub"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func serveWs(hubInstance *hub.Hub, w http.ResponseWriter, r *http.Request) {
	log.Println("Received connection attempt")
	roomId := "lobby"
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Failed to upgrade Connc for room %s", roomId)
		return
	}

	log.Printf("websocket connection opened for room %s", roomId)
	defer conn.Close()
	client := &hub.Client{}
	room := hubInstance.FindOrCreateRoom(roomId)

	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Failed to read message from client in room %s", roomId)
			break
		}
		log.Printf("TEMP: Message from %p in room %s: type=%d, msg=%s", client, room.Id, messageType, message)

	}

}

func main() {
	//first
	hubInstance := hub.NewHub()
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hubInstance, w, r)
	})
	serverAddress := ":8080"
	log.Printf("Starting Doodle Weavers server on %s\n", serverAddress)

	err := http.ListenAndServe(serverAddress, nil)
	if err != nil {
		log.Fatal("Error Starting Server: ", err)
	}

}
