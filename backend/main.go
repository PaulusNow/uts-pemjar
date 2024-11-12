package main

import (
    "fmt"
    "net/http"
    "github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool {
        return true // Allow all origins for development; adjust in production
    },
}

var clients = make(map[*websocket.Conn]bool)

// Struct untuk format data pesan JSON
type Message struct {
    Sender  string `json:"sender"`
    Amount  string `json:"amount"`
    Message string `json:"message"`
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
    // Tambahkan header CORS
    w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
    w.Header().Set("Access-Control-Allow-Credentials", "true") // Izinkan kredensial jika diperlukan

    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        fmt.Println("Error upgrading connection:", err)
        return
    }
    defer conn.Close()

    clients[conn] = true
    fmt.Println("Client connected")

    for {
        _, msg, err := conn.ReadMessage()
        if err != nil {
            fmt.Println("Error reading message:", err)
            delete(clients, conn)
            break
        }

        fmt.Printf("Message received: %s\n", msg)

        // Broadcast message to all clients
        for client := range clients {
            if err := client.WriteMessage(websocket.TextMessage, msg); err != nil {
                fmt.Println("Error sending message:", err)
                client.Close()
                delete(clients, client)
            }
        }
    }
}

func main() {
    http.HandleFunc("/ws", handleWebSocket)
    fmt.Println("Server started at :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Println("Error starting server:", err)
    }
}
