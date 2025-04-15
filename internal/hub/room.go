package hub

import "log"

type Room struct {
	Id      string
	Clients map[*Client]bool
}

func NewRoom(id string) *Room {
	return &Room{
		Id:      id,
		Clients: make(map[*Client]bool),
	}
}

func (r *Room) Broadcast(messageType int, message []byte) {
	log.Printf("Attempting to broadcast to add clients in room %s", r.Id)

	for client := range r.Clients {
		err := client.conn.WriteMessage(messageType, message)
		if err != nil {
			log.Printf("Failed to broadcast message to client with in room %s: %v", r.Id, err)
		}
	}
}
