package main

import (
	"net"
	"go_chess/httpd/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	server, err := net.Listen("tcp", ":777")
	router := gin.Default()

	router.GET("/api/ping", handlers.Ping())
	router.Run()
}
