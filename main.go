package main

import (
	"os"

	server "github.com/bmxguy100/battlesnakes_alphabeta/server"
	log "github.com/sirupsen/logrus"
)

func main() {
	command := os.Args[1:]

	if len(command) > 0 && command[0] == "test" {
		log.Info("Testing AI...")
		server.TestSolo()
	} else {
		log.Info("Starting Server...")
		server.StartServer("localhost", 8080)
	}
}
