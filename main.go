package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	fmt.Println("Client connected")
	initialMessage := "Welcome to the WebSocket server!"
	if err := conn.WriteMessage(websocket.TextMessage, []byte(initialMessage)); err != nil {
		fmt.Println("Error sending initial message:", err)
		return
	}

	for {
		// Read message from the client
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Client disconnected")
			return
		}

		if string(p) == "force-error" {
			errMessage := "Simulated error: Forced from the client side"
			fmt.Println(errMessage)
			conn.WriteMessage(messageType, []byte(errMessage))
			return
		}

		// Print the received message
		fmt.Printf("Received message: %s\n", p)

		// Echo the message back to the client
		if err := conn.WriteMessage(messageType, p); err != nil {
			fmt.Println("Error writing message:", err)
			return
		}

	}

}

func main() {
	http.HandleFunc("/ws", handleWebSocket)
	http.ListenAndServe(":8080", nil)
}
