// File: middle/middleware/main.go
// Level: Middle
// Topik: Middleware Pattern pada HTTP Server
//
// Middleware adalah fungsi yang membungkus HTTP handler.
// Berguna untuk: logging, auth, recovery, rate limiting, CORS, dll.
// Pola: middleware(http.Handler) http.Handler
//
// Middleware bisa di-chain: handler = m1(m2(m3(handler)))
// Atau dibalik: handler = chain(handler, m1, m2, m3)

package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// loggingMiddleware mencatat setiap request (method, path, durasi)
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Before handler
		start := time.Now()

		// Panggil handler asli
		next.ServeHTTP(w, r)

		// After handler
		log.Printf("[%s] %s %s %v",
			r.Method, r.RequestURI, r.RemoteAddr, time.Since(start))
	})
}

// authMiddleware memvalidasi token dari header Authorization
func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Ambil token dari header
		token := r.Header.Get("Authorization")

		// Validasi (sederhana)
		if token != "valid-token" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return // stop: jangan panggil handler berikutnya
		}

		// Token valid, lanjut ke handler
		next.ServeHTTP(w, r)
	})
}

// recoveryMiddleware menangkap panic agar server tidak crash
func recoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				// Log error tanpa crash
				log.Printf("Recovered from panic: %v", err)
				http.Error(w, "Internal Server Error",
					http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

// chainMiddlewares menerapkan middleware secara berurutan
// Urutan penting: recovery harus paling luar
func chainMiddlewares(h http.Handler, middlewares ...func(http.Handler) http.Handler) http.Handler {
	for _, m := range middlewares {
		h = m(h)
	}
	return h
}

// Handler: mengembalikan JSON response
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, `{"message": "Hello, World!"}`)
}

// Handler: sengaja panic untuk demo recovery
func panicHandler(w http.ResponseWriter, r *http.Request) {
	panic("Terjadi error tak terduga!")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/panic", panicHandler)

	// Chain middlewares
	handler := chainMiddlewares(mux,
		recoveryMiddleware, // paling luar
		loggingMiddleware,
		authMiddleware, // paling dalam (dijalankan pertama sebelum handler)
	)

	log.Println("Server running on :8080")
	log.Println("Test: curl -H 'Authorization: valid-token' localhost:8080/hello")
	log.Println("Test: curl -H 'Authorization: valid-token' localhost:8080/panic")
	http.ListenAndServe(":8080", handler)
}