package main

import (
	"fmt"
	"log"
	"net/http"

	"web-socket/pkg/websocket"
)

func main() {
	http.HandleFunc("/ws", websocket.HandleConnections)
	fmt.Println("WebSocket server started at :8080")
	log.Fatal(http.ListenAndServe(":8085", nil))
}
