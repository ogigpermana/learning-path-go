// File: middle/context/main.go
// Level: Middle
// Topik: Context
//
// Context digunakan untuk:
// 1. Timeout & Deadline - membatasi waktu eksekusi
// 2. Cancellation - membatalkan operasi yang berjalan
// 3. Passing values - mengirim data antar goroutine (request-scoped)
//
// Context penting untuk HTTP server, database query, dan API calls.
// Selalu gunakan context.Background() atau context.TODO() sebagai parent.

package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// 1. CONTEXT WITH TIMEOUT
	// Membuat context yang otomatis selesai setelah 2 detik
	// Berguna untuk: timeout HTTP request, database query
	fmt.Println("=== Context with Timeout ===")
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel() // selalu panggil cancel untuk membebaskan resource

	// Menjalankan goroutine dengan context
	go task(ctx)

	// Menunggu context selesai (timeout) atau timer 3 detik
	select {
	case <-ctx.Done():
		// ctx.Done() mengembalikan channel yang close saat context selesai
		// ctx.Err() memberikan alasan: DeadlineExceeded atau Canceled
		fmt.Println("Main: context selesai:", ctx.Err())
	case <-time.After(3 * time.Second):
		fmt.Println("Main: timeout tidak terjadi")
	}

	// Beri waktu goroutine mencetak output
	time.Sleep(500 * time.Millisecond)
	fmt.Println()

	// 2. CONTEXT WITH VALUE
	// Untuk passing request-scoped data (seperti userID, requestID)
	// CATATAN: hanya untuk data yang penting untuk request flow
	fmt.Println("=== Context with Value ===")
	ctx2 := context.WithValue(context.Background(), "userID", 12345)
	processRequest(ctx2)
	fmt.Println()

	// 3. CONTEXT WITH CANCEL
	// Manual cancel - menghentikan goroutine dari luar
	fmt.Println("=== Context with Cancel ===")
	ctx3, cancel3 := context.WithCancel(context.Background())
	go cancellableTask(ctx3)

	// Biarkan task berjalan 500ms lalu cancel
	time.Sleep(500 * time.Millisecond)
	cancel3() // mengirim sinyal cancel ke goroutine

	// Beri waktu goroutine merespon cancel
	time.Sleep(200 * time.Millisecond)
}

// task melakukan pekerjaan dengan context timeout
func task(ctx context.Context) {
	select {
	case <-time.After(3 * time.Second):
		// Simulasi pekerjaan yang butuh 3 detik
		fmt.Println("Task selesai (3 detik)")
	case <-ctx.Done():
		// Context selesai (timeout) sebelum task selesai
		fmt.Println("Task dibatalkan karena:", ctx.Err())
	}
}

// processRequest mengambil data dari context
func processRequest(ctx context.Context) {
	// ctx.Value() mengembalikan interface{}, perlu type assertion
	userID := ctx.Value("userID")
	fmt.Println("Processing request untuk user:", userID)
}

// cancellableTask melakukan loop sampai di-cancel
func cancellableTask(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			// Menerima sinyal cancel
			fmt.Println("Cancellable task dihentikan")
			return
		default:
			// Melakukan pekerjaan
			fmt.Print(".")
			time.Sleep(100 * time.Millisecond)
		}
	}
}