package main

import (
	"fmt"
)

const NMAX = 1000000
var data [NMAX]int

func main() {
	var n, k int

	// Membaca baris pertama (n dan k)
	fmt.Scan(&n, &k)

	// Mengisi array data
	isiArray(n)

	// Mencari posisi k
	pos := posisi(n, k)

	// Menampilkan hasil pencarian
	if pos == -1 {
		fmt.Println("TIDAK ADA")
	} else {
		fmt.Println(pos)
	}
}

func isiArray(n int) {
	// Mengisi array data dari input
	for i := 0; i < n; i++ {
		fmt.Scan(&data[i])
	}
}

func posisi(n, k int) int {
	// Binary search untuk mencari k
	low, high := 0, n-1
	for low <= high {
		mid := (low + high) / 2
		if data[mid] == k {
			return mid
		} else if data[mid] < k {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1 // Jika tidak ditemukan
}

