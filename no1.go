package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Membaca input dari pengguna
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Masukkan suara (pisahkan dengan spasi, akhiri dengan 0):")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	// Memisahkan input menjadi slice string
	data := strings.Split(input, " ")

	// Inisialisasi variabel untuk menghitung suara
	totalSuara := 0
	validSuara := 0
	count := make([]int, 20)

	// Memproses data suara
	for _, value := range data {
		num, err := strconv.Atoi(value)
		if err != nil {
			continue // Abaikan jika tidak valid
		}
		if num == 0 {
			break // Akhiri jika ditemukan angka 0
		}
		totalSuara++
		if num >= 1 && num <= 20 {
			validSuara++
			count[num-1]++ // Hitung suara untuk calon tertentu
		}
	}

	// Menampilkan hasil
	fmt.Printf("Total suara terbaca: %d\n", totalSuara)
	fmt.Printf("Total suara valid: %d\n", validSuara)
	for i, suara := range count {
		if suara > 0 {
			fmt.Printf("Calon nomor %d: %d suara\n", i+1, suara)
		}
	}
}
