// File: middle/concurrency/main.go
// Level: Middle
// Topik: Concurrency Lanjutan - WaitGroup, Mutex, Select
//
// Lanjutan dari goroutine & channel di level Beginner.
// Pola-pola penting:
// 1. WaitGroup - menunggu goroutine selesai
// 2. Mutex - mengamankan akses ke shared data
// 3. Select - multiplexing multiple channels

package main

import (
	"fmt"
	"sync" // WaitGroup, Mutex
	"time"
)

func main() {
	// 1. WAITGROUP - menunggu kumpulan goroutine selesai
	// WaitGroup seperti counter: Add() menambah, Done() mengurangi, Wait() menunggu
	fmt.Println("=== WaitGroup ===")
	var wg sync.WaitGroup

	// Menjalankan 5 worker goroutine
	for i := 0; i < 5; i++ {
		wg.Add(1) // increment counter
		go worker(i, &wg)
	}

	wg.Wait() // blocking sampai counter = 0
	fmt.Println("Semua worker selesai")
	fmt.Println()

	// 2. MUTEX - mutual exclusion
	// Mencegah race condition ketika multiple goroutine mengakses data yang sama
	fmt.Println("=== Mutex ===")
	var counter int
	var mu sync.Mutex

	// 10 goroutine mencoba increment counter secara concurrent
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go increment(&counter, &mu, &wg)
	}

	wg.Wait()
	fmt.Println("Counter akhir:", counter) // harus 10
	fmt.Println()

	// 3. SELECT - multiplexing channel
	// Menunggu multiple channel, merespon yang pertama siap
	fmt.Println("=== Select ===")
	c1 := make(chan string)
	c2 := make(chan string)

	// Goroutine 1: kirim setelah 100ms
	go func() {
		time.Sleep(100 * time.Millisecond)
		c1 <- "dari channel 1"
	}()

	// Goroutine 2: kirim setelah 200ms
	go func() {
		time.Sleep(200 * time.Millisecond)
		c2 <- "dari channel 2"
	}()

	// Select: merespon channel yang pertama siap
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		case <-time.After(500 * time.Millisecond):
			// Timeout jika tidak ada channel yang siap dalam 500ms
			fmt.Println("Timeout!")
		}
	}
}

// worker mensimulasikan pekerjaan goroutine
func worker(id int, wg *sync.WaitGroup) {
	// Done() harus dipanggil, gunakan defer untuk keamanan
	defer wg.Done()

	time.Sleep(100 * time.Millisecond) // simulasi kerja
	fmt.Printf("Worker %d selesai\n", id)
}

// increment menambah counter dengan aman menggunakan Mutex
func increment(counter *int, mu *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()

	// Lock: goroutine lain harus menunggu sampai Unlock
	mu.Lock()
	*counter++ // critical section - hanya satu goroutine boleh akses
	mu.Unlock()
	// Unlock: goroutine lain bisa melanjutkan
}