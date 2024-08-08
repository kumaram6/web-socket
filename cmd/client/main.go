package main

import (
	"log"
	"web-socket/pkg/websocket"
)

func main() {
	url := "ws://localhost:8085/ws"
	log.Printf("Connecting to %s", url)
	websocket.StartClient(url)
}
