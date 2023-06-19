package main

import (
	"log"
	c "rentals/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	startServer()
}

func startServer() {
	var server = "localhost:8080" // TODO: externalize
	log.Println("initiating server " + server)

	router := gin.Default()
	router.GET("/rentals", c.GetRentals)

	router.Run(server)
}
