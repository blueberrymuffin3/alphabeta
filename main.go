package main

import (
	server "github.com/bmxguy100/battlesnakes_alphabeta/server"
	// log "github.com/sirupsen/logrus"
)

func main() {
	// server.TestSolo(nil)
	server.StartServer("", 8080)
}
