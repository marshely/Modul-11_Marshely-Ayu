package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	// Menentukan ketua dan wakil
	type Candidate struct {
		Number int
		Votes  int
	}
	var candidates []Candidate
	for i, suara := range count {
		if suara > 0 {
			candidates = append(candidates, Candidate{Number: i + 1, Votes: suara})
		}
	}

	// Sortir berdasarkan suara terbanyak, lalu nomor terkecil
	sort.Slice(candidates, func(i, j int) bool {
		if candidates[i].Votes == candidates[j].Votes {
			return candidates[i].Number < candidates[j].Number
		}
		return candidates[i].Votes > candidates[j].Votes
	})

	// Menentukan ketua dan wakil
	ketua := Candidate{}
	wakil := Candidate{}
	if len(candidates) > 0 {
		ketua = candidates[0]
	}
	if len(candidates) > 1 {
		wakil = candidates[1]
	}

	// Menampilkan hasil
	fmt.Printf("Total suara terbaca: %d\n", totalSuara)
	fmt.Printf("Total suara valid: %d\n", validSuara)
	fmt.Printf("Ketua RT: Calon nomor %d dengan %d suara\n", ketua.Number, ketua.Votes)
	if wakil.Number > 0 {
		fmt.Printf("Wakil Ketua RT: Calon nomor %d dengan %d suara\n", wakil.Number, wakil.Votes)
	} else {
		fmt.Println("Wakil Ketua RT: Tidak ada")
	}
}
