// File: middle/graceful-shutdown/main.go
// Level: Middle
// Topik: Graceful Shutdown
//
// Graceful shutdown = menutup server dengan rapi saat menerima sinyal berhenti.
// Penting untuk production: menyelesaikan request yang sedang berlangsung,
// menutup koneksi database, menyimpan state, dll.
//
// Sinyal:
// - SIGINT (Ctrl+C)
// - SIGTERM (docker stop, systemd stop)
// - SIGHUP (reload config)

package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// 1. BUAT HTTP SERVER
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Simulasi request yang butuh waktu proses
		fmt.Fprintln(w, "Memproses request...")
		time.Sleep(2 * time.Second) // kerja berat
		fmt.Fprintln(w, "Selesai!")
	})

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{"status":"ok"}`)
	})

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	// 2. JALANKAN SERVER DI GOROUTINE
	go func() {
		log.Println("Server running on :8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	// 3. GRACEFUL SHUTDOWN
	// Buat channel untuk menangkap sinyal OS
	quit := make(chan os.Signal, 1)

	// Daftarkan sinyal yang ingin ditangkap
	signal.Notify(quit,
		syscall.SIGINT,  // Ctrl+C
		syscall.SIGTERM, // Docker stop / systemd stop
		// syscall.SIGHUP, // reload (uncomment jika perlu)
	)

	// Blocking sampai menerima sinyal
	sig := <-quit
	log.Printf("Menerima sinyal: %v. Memulai shutdown...\n", sig)

	// 4. SHUTDOWN DENGAN TIMEOUT
	// Beri waktu 30 detik untuk request selesai
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Shutdown: stop menerima request baru, tunggu request yang ada selesai
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	// 5. CLEANUP RESOURCE LAIN
	// Tutup database connection
	// db.Close()
	// Tutup message queue connection
	// mq.Close()
	// Save state ke file
	// saveState()

	log.Println("Server stopped gracefully!")
}

/*
Cara test:
1. go run main.go
2. curl localhost:8080 (di terminal lain)
3. Tekan Ctrl+C
4. Lihat server menunggu request selesai sebelum shutdown

Pattern yang sama untuk:
- Database connection pool
- Kafka/RabbitMQ consumers
- gRPC servers
- Background workers
*/