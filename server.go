package main

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/websocket"
)

// Struktur untuk data User
type User struct {
	ID    string
	Name  string
	Saldo int
}

var users = make(map[string]User) // Penyimpanan sementara user

var wsClients = make(map[*websocket.Conn]bool)
var wsUpgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Fungsi untuk menangani koneksi TCP
func handleTCPConnection(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)

	for {
		request, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Koneksi ditutup:", err)
			return
		}
		request = strings.TrimSpace(request)
		parts := strings.Split(request, "|")
		command := parts[0]

		switch command {
		case "register":
			if len(parts) < 3 {
				fmt.Fprintln(conn, "Format salah untuk register.")
				continue
			}
			id, name := parts[1], parts[2]

			if _, exists := users[id]; exists {
				fmt.Fprintln(conn, "ID sudah terdaftar.")
			} else {
				users[id] = User{ID: id, Name: name, Saldo: 0}
				fmt.Fprintln(conn, "Registrasi berhasil.")
			}

		case "login":
			if len(parts) < 2 {
				fmt.Fprintln(conn, "Format salah untuk login.")
				continue
			}
			id := parts[1]

			if user, exists := users[id]; exists {
				fmt.Fprintf(conn, "Login berhasil. Selamat datang, %s!\n", user.Name)
			} else {
				fmt.Fprintln(conn, "ID tidak ditemukan. Silakan daftar terlebih dahulu.")
			}

		case "cekSaldo":
			if len(parts) < 2 {
				fmt.Fprintln(conn, "ID tidak ditemukan. Silakan login terlebih dahulu.")
				continue
			}
			id := parts[1]
			if user, exists := users[id]; exists {
				fmt.Fprintf(conn, "Saldo Anda saat ini: Rp %d\n", user.Saldo)
			} else {
				fmt.Fprintln(conn, "ID tidak ditemukan. Silakan login terlebih dahulu.")
			}

		default:
			fmt.Fprintln(conn, "Perintah tidak dikenal.")
		}
	}
}

// Fungsi untuk menangani koneksi UDP
func handleUDPConnection(conn *net.UDPConn) {
	buffer := make([]byte, 1024)

	for {
		n, addr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Gagal menerima data dari UDP:", err)
			continue
		}

		request := string(buffer[:n])
		parts := strings.Split(request, "|")
		if len(parts) < 3 || parts[0] != "topup" {
			fmt.Println("Format permintaan top up tidak valid.")
			continue
		}

		id := parts[1]
		amount, err := strconv.Atoi(parts[2])
		if err != nil || amount <= 0 {
			fmt.Println("Nominal top up tidak valid.")
			continue
		}

		// Proses top up saldo
		if user, exists := users[id]; exists {
			user.Saldo += amount
			users[id] = user
			response := fmt.Sprintf("Rp %d", user.Saldo)
			_, err = conn.WriteToUDP([]byte(response), addr)
			if err != nil {
				fmt.Println("Gagal mengirim data ke UDP:", err)
			}
		} else {
			_, err = conn.WriteToUDP([]byte("ID tidak ditemukan."), addr)
			if err != nil {
				fmt.Println("Gagal mengirim data ke UDP:", err)
			}
		}
	}
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	wsConn, err := wsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Gagal meng-upgrade koneksi WebSocket:", err)
		return
	}
	defer wsConn.Close()
	wsClients[wsConn] = true

	for {
		_, message, err := wsConn.ReadMessage()
		if err != nil {
			delete(wsClients, wsConn)
			break
		}

		// Kirim pesan donasi ke semua klien WebSocket yang terhubung
		for client := range wsClients {
			err := client.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				delete(wsClients, client)
			}
		}
	}
}

func main() {
	// TCP Listener
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Gagal membuat server TCP:", err)
		return
	}
	defer ln.Close()
	fmt.Println("Server TCP berjalan di port 8080...")

	// UDP Listener
	udpAddr, err := net.ResolveUDPAddr("udp", ":8081")
	if err != nil {
		fmt.Println("Gagal membuat server UDP:", err)
		return
	}
	udpConn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		fmt.Println("Gagal membuat server UDP:", err)
		return
	}
	defer udpConn.Close()
	fmt.Println("Server UDP berjalan di port 8081...")

	// Jalankan goroutine untuk koneksi TCP dan UDP
	go func() {
		for {
			conn, err := ln.Accept()
			if err != nil {
				fmt.Println("Gagal menerima koneksi:", err)
				continue
			}
			go handleTCPConnection(conn)
		}
	}()

	handleUDPConnection(udpConn) // Handle koneksi UDP

	http.HandleFunc("/ws", handleWebSocket)
	go http.ListenAndServe(":5501", nil) // Jalankan server WebSocket pada port 5501
}
