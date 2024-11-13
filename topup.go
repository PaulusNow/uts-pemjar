package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

// Fungsi untuk meminta input dari pengguna
func input(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func main() {
	// Atur alamat server UDP
	serverAddr, err := net.ResolveUDPAddr("udp", "localhost:8081")
	if err != nil {
		fmt.Println("Gagal mengatur alamat server UDP:", err)
		return
	}

	// Hubungkan ke server UDP
	conn, err := net.DialUDP("udp", nil, serverAddr)
	if err != nil {
		fmt.Println("Gagal terhubung ke server UDP:", err)
		return
	}
	defer conn.Close()

	for {
		fmt.Println("\n--- Fitur Top Up ---")
		id := input("Masukkan ID untuk top up: ")
		amountStr := input("Masukkan nominal top up: ")

		amount, err := strconv.Atoi(amountStr)
		if err != nil || amount <= 0 {
			fmt.Println("Nominal tidak valid. Silakan coba lagi.")
			continue
		}

		// Kirim permintaan top up ke server
		request := fmt.Sprintf("topup|%s|%d", id, amount)
		_, err = conn.Write([]byte(request))
		if err != nil {
			fmt.Println("Gagal mengirim data ke server:", err)
			continue
		}

		// Terima respons saldo terbaru dari server
		buffer := make([]byte, 1024)
		n, _, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Gagal menerima data dari server:", err)
			continue
		}

		response := string(buffer[:n])
		fmt.Println("Saldo terbaru:", response)
	}
}
