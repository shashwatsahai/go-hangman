package main

import (
	"log"

	"github.com/shashwatsahai/go-hangman/server"
)

func main() {
	serverInstance := &server.ServerStart{
		Hosturl: "localhost:3000",
	}

	if err := serverInstance.Start(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}

	// serverInstance.Conn.Handler =
}
