// File: 15-defer-panic-recover.go
// Level: Beginner
// Topik: Defer, Panic, Recover
//
// Defer  : menjadwalkan eksekusi fungsi setelah fungsi induk selesai
//          (biasanya untuk cleanup: close file, close connection)
// Panic  : menghentikan program secara paksa (seperti exception)
// Recover: menangkap panic agar program tidak crash
//
// Eksekusi defer menggunakan LIFO (Last In First Out)

package main

import "fmt"

func main() {
	// Defer: fungsi dijalankan saat fungsi main() selesai
	// Defer menggunakan stack (LIFO): yang terakhir didefer dijalankan pertama
	defer fmt.Println("1. Ini dijalankan terakhir (defer 1)")
	defer fmt.Println("2. Ini dijalankan kedua terakhir (defer 2)")
	fmt.Println("3. Ini dijalankan pertama")
	fmt.Println()

	// Defer untuk cleanup - pattern umum di Go
	readFile()

	// Panic & Recover - mencegah program crash
	fmt.Println("\n--- Demo safeDivision ---")
	safeDivision(10, 2) // sukses
	safeDivision(10, 0) // panic, tapi di-recover

	fmt.Println("Program tetap berjalan tanpa crash!")
	fmt.Println()

	// Contoh penggunaan: HTTP handler dengan recover
	handleRequest("valid")
	handleRequest("panic")
	fmt.Println("Server tetap berjalan setelah panic handler")
}

// readFile mensimulasikan pembacaan file
func readFile() {
	fmt.Println("Membuka file...")
	// defer akan selalu dieksekusi, bahkan jika terjadi panic
	defer fmt.Println("Menutup file...") // cleanup
	fmt.Println("Membaca konten file...")
	// File otomatis "tertutup" karena defer di atas
}

// safeDivision melakukan pembagian dengan proteksi panic
func safeDivision(a, b int) {
	// Recover hanya berguna di dalam deferred function
	defer func() {
		if r := recover(); r != nil {
			// recover() mengembalikan nilai panic
			// Program tidak jadi crash
			fmt.Println("Recover dari panic:", r)
		}
	}()

	if b == 0 {
		// panic menghentikan fungsi saat ini
		panic("Pembagian dengan nol!")
	}
	result := a / b
	fmt.Printf("%d / %d = %d\n", a, b, result)
}

// handleRequest mensimulasikan handler HTTP dengan recover
func handleRequest(req string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Request '%s' menyebabkan panic: %v\n", req, r)
			fmt.Println("Request ditangani tanpa crash server")
		}
	}()

	fmt.Printf("Memproses request: %s\n", req)
	if req == "panic" {
		panic("sesuatu error!")
	}
	fmt.Println("Request sukses")
}