// File: expert/websocket/main.go
// Level: Expert
// Topik: WebSocket Server
//
// WebSocket memungkinkan komunikasi real-time bidirectional
// antara client (browser) dan server.
//
// Cocok untuk:
// - Chat application
// - Live notifications
// - Real-time data (sports, stocks)
// - Collaborative editing
//
// Library: github.com/gorilla/websocket

package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

// upgrader mengubah HTTP connection menjadi WebSocket connection
var upgrader = websocket.Upgrader{
	// CheckOrigin: izinkan koneksi dari origin manapun (development)
	CheckOrigin: func(r *http.Request) bool { return true },
	// Di production, batasi origin:
	// CheckOrigin: func(r *http.Request) bool {
	//     return r.Header.Get("Origin") == "https://app.example.com"
	// },
}

// handleWS adalah handler untuk WebSocket connections
func handleWS(w http.ResponseWriter, r *http.Request) {
	// Upgrade HTTP ke WebSocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}
	defer conn.Close() // tutup koneksi saat selesai

	log.Printf("Client connected: %s", r.RemoteAddr)

	// Loop membaca pesan dari client
	for {
		// ReadMessage: blocking sampai ada pesan
		messageType, msg, err := conn.ReadMessage()
		if err != nil {
			// Client disconnect atau error
			if websocket.IsUnexpectedCloseError(err,
				websocket.CloseGoingAway, websocket.CloseNormalClosure) {
				log.Printf("Error: %v", err)
			}
			break
		}

		// Log pesan yang diterima
		log.Printf("Received: %s", msg)

		// Kirim response ke client
		response := fmt.Sprintf("Server menerima pada %s: %s",
			time.Now().Format(time.RFC3339), msg)

		err = conn.WriteMessage(messageType, []byte(response))
		if err != nil {
			log.Println("Write error:", err)
			break
		}
	}
}

func main() {
	// WebSocket endpoint
	http.HandleFunc("/ws", handleWS)

	// Serve HTML client
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	log.Println("WebSocket server running on :8080")
	log.Println("Buka http://localhost:8080 di browser")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Cara install:
// cd expert/websocket
// go get github.com/gorilla/websocket
// go run main.go