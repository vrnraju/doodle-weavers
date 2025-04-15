package hub

import "log"

type Hub struct {
	rooms map[string]*Room
}

func NewHub() *Hub {
	return &Hub{
		rooms: make(map[string]*Room),
	}
}

func (h *Hub) FindOrCreateRoom(id string) *Room {
	room, ok := h.rooms[id]
	if !ok {
		log.Printf("Hub: Creating new room: %s", id)
		room = NewRoom(id)
		h.rooms[id] = room
	}
	return room
}
