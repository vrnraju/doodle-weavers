package hub

import "github.com/gorilla/websocket"

type Client struct {
	conn *websocket.Conn
}
