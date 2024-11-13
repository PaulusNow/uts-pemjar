package main

import (
	"bufio"
	"fmt"
	"github.com/gorilla/websocket"
	"net"
	"os"
	"strconv"
	"strings"
)

var wsConn *websocket.Conn

// Struktur User
type User struct {
	ID    string
	Name  string
	Saldo int
}

func connectWebSocket() error {
	var err error
	wsConn, _, err = websocket.DefaultDialer.Dial("ws://localhost:5501/ws", nil)
	if err != nil {
		return fmt.Errorf("Gagal menghubungkan WebSocket: %v", err)
	}
	return nil
}

// Fungsi untuk meminta input dari pengguna
func input(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func main() {
	err := connectWebSocket()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Gagal terhubung ke server:", err)
		return
	}
	defer conn.Close()

	for {
		fmt.Println("\n--- Aplikasi Donasi Real-Time ---")
		fmt.Println("1. Register")
		fmt.Println("2. Login")
		fmt.Println("3. Keluar")
		choice := input("Pilih opsi: ")

		switch choice {
		case "1":
			name := input("Masukkan Nama: ")
			id := input("Masukkan ID: ")

			request := fmt.Sprintf("register|%s|%s", id, name)
			fmt.Fprintln(conn, request)

			response, _ := bufio.NewReader(conn).ReadString('\n')
			fmt.Print(response)
		case "2":
			id := input("Masukkan ID: ")

			request := fmt.Sprintf("login|%s", id)
			fmt.Fprintln(conn, request)

			response, _ := bufio.NewReader(conn).ReadString('\n')
			if strings.Contains(response, "Login berhasil") {
				fmt.Println(response)
				// Tampilkan menu setelah login berhasil
				afterLoginMenu(conn, id)
			} else {
				fmt.Print(response)
			}
		case "3":
			fmt.Println("Keluar dari aplikasi.")
			return
		default:
			fmt.Println("Opsi tidak valid. Silakan coba lagi.")
		}
	}
}

func afterLoginMenu(conn net.Conn, id string) {
	for {
		fmt.Println("\n--- Menu ---")
		fmt.Println("1. Cek Saldo")
		fmt.Println("2. Buat Donasi")
		fmt.Println("3. Kembali ke Menu Utama")
		choice := input("Pilih opsi: ")

		switch choice {
		case "1":
			request := fmt.Sprintf("cekSaldo|%s", id) // sertakan ID saat cek saldo
			fmt.Fprintln(conn, request)

			response, _ := bufio.NewReader(conn).ReadString('\n')
			fmt.Print(response)
		case "2":
			id := input("Masukkan ID Anda: ")
			amountStr := input("Masukkan nominal donasi: ")
			amount, err := strconv.Atoi(amountStr)
			if err != nil || amount <= 0 {
				fmt.Println("Nominal tidak valid.")
				continue
			}
			message := input("Masukkan pesan donasi: ")

			// Formatkan data dalam JSON
			donationData := fmt.Sprintf(`{"id": "%s", "amount": %d, "message": "%s"}`, id, amount, message)
			err = wsConn.WriteMessage(websocket.TextMessage, []byte(donationData))
			if err != nil {
				fmt.Println("Gagal mengirim pesan donasi:", err)
				continue
			}
			fmt.Println("Donasi berhasil dikirim.")
		case "3":
			return
		default:
			fmt.Println("Opsi tidak valid. Silakan coba lagi.")
		}
	}
}
