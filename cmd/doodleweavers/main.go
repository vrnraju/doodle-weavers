package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	log.Println("Received connection attempt...")
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Failed to upgrade Connc")
		return
	}
	defer ws.Close()
	log.Println("Client Connected!!")

	for {
		messageType, message, err := ws.ReadMessage()
		if err != nil {
			log.Println("Read error:", err)
			break
		}
		err = ws.WriteMessage(messageType, message)
		if err != nil {
			log.Println("Write error:", err)
			break
		}
	}
	log.Println("Websocket connection closed. ")

}
func main() {
	//first
	http.HandleFunc("/ws", handleConnections)
	serverAddress := ":8080"
	log.Printf("Starting Doodle Weavers server on %s\n", serverAddress)

	err := http.ListenAndServe(serverAddress, nil)
	if err != nil {
		log.Fatal("Error Starting Server: ", err)
	}

}
