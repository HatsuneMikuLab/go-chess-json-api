package main

import (
	"go_chess/httpd/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/api/ping", handlers.Ping())
	router.Run()
}
