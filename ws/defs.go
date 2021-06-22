package ws

import (
	"github.com/gorilla/websocket"
)

type (
	Command byte
	ID int64
	ClientRequest struct {
		Command Command
		Args []interface{}
	}
	Room struct {
		ID ID
		Players [2]*websocket.Conn
		Spectators []*websocket.Conn
	}
	Server struct {
		Rooms []*Room
		Guests []*websocket.Conn
	}

)

const (
	create Command = 0
	join Command = 1
	watch Command = 2
	move Command = 3
	resign Command = 4
)

var upgrader = websocket.Upgrader{}