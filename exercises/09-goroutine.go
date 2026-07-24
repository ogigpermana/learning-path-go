// File: 09-goroutine.go
// Level: Beginner
// Topik: Goroutine dan Channel
//
// Goroutine: thread ringan di Go, dijalankan dengan keyword "go"
// Channel: jembatan komunikasi antar goroutine
// - ch <- value  : kirim data ke channel
// - <-ch         : terima data dari channel
// Paradigma: "Don't communicate by sharing memory, share memory by communicating"

package main

import (
	"fmt"
	"time"
)

// Fungsi yang akan dijalankan sebagai goroutine
// nama: identifier goroutine, ch: channel untuk mengirim pesan
func cetakPesan(nama string, ch chan string) {
	for i := 0; i < 3; i++ {
		time.Sleep(100 * time.Millisecond)      // simulasi kerja
		ch <- fmt.Sprintf("%s: pesan %d", nama, i+1) // kirim ke channel
	}
}

func main() {
	// Membuat channel dengan make()
	// Channel bisa mengirim data string
	ch := make(chan string)

	// Menjalankan goroutine dengan keyword "go"
	// Dua goroutine berjalan secara concurrent
	go cetakPesan("Goroutine 1", ch)
	go cetakPesan("Goroutine 2", ch)

	// Menerima data dari channel sebanyak 6 kali
	// (3 dari goroutine 1 + 3 dari goroutine 2)
	for i := 0; i < 6; i++ {
		pesan := <-ch // blocking: menunggu sampai ada data
		fmt.Println(pesan)
	}

	// Contoh channel buffered
	buffered := make(chan int, 3) // buffer capacity 3
	buffered <- 1                  // tidak blocking karena buffer masih kosong
	buffered <- 2
	buffered <- 3
	// buffered <- 4 // ini akan blocking karena buffer penuh

	fmt.Println("Buffered channel:")
	fmt.Println(<-buffered)
	fmt.Println(<-buffered)
	fmt.Println(<-buffered)

	// Contoh: unbuffered channel akan blocking sampai ada receiver
	// go func() {
	//     ch2 := make(chan string)
	//     ch2 <- "test" // blocking jika tidak ada yang menerima
	// }()
}