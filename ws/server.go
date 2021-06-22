package ws

import (
	"github.com/gorilla/websocket"
	"net/http"
	"fmt"
)

func NewSocketServer() *Server {
	return &Server{
		Rooms: make([]*Room, 0, 100),
		Guests: make([]*websocket.Conn, 100),
	}
}

func (s *Server) Accept(res http.ResponseWriter, req *http.Request) {
	client, err := upgrader.Upgrade(res, req, nil)
	if err != nil {
		println(err)
	}
	defer client.Close()
	//s.Guests = append(s.Guests, client)
	s.Serve(client)
}

func (s *Server) Serve(client *websocket.Conn) {
	for {
		clientRequest := &ClientRequest{}
		err := client.ReadJSON(clientRequest)
		if err != nil {
			fmt.Printf("Wrong format %v", err)
			continue
		}
		fmt.Printf("Request from client %v", clientRequest)
		switch clientRequest.Command {
		case create: 
			if len(clientRequest.Args) != 2 {
				fmt.Println("This command should have 2 arguments.")
			}
			side := clientRequest.Args[0]
			time := clientRequest.Args[1]
		}
	}
}

