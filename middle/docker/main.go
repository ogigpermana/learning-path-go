// File: middle/docker/main.go
// Level: Middle
// Topik: Aplikasi Go untuk Docker
//
// Aplikasi ini di-build menjadi Docker image.
// Dockerfile menggunakan multi-stage build:
// Stage 1: build binary (golang:alpine)
// Stage 2: run binary (alpine minimal)
//
// Keuntungan multi-stage: image size kecil (~15MB vs ~800MB)

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// Baca port dari environment variable
	// Convention: PORT env var (standard di Heroku, Railway, dll)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // default port
	}

	// Health check endpoint - penting untuk Docker/Kubernetes
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"status":"healthy","service":"go-app"}`)
	})

	// Root endpoint
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		hostname, _ := os.Hostname()
		fmt.Fprintf(w,
			"Hello from Dockerized Go App!\n\n"+
				"Hostname: %s\n"+
				"Port: %s\n"+
				"API Base: http://localhost:%s\n",
			hostname, port, port,
		)
	})

	log.Printf("Server running on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

/*
Cara build & run:

1. Build Docker image:
   cd fullstack/middle/docker
   docker build -t go-app .

2. Run container:
   docker run -p 8080:8080 go-app

3. Test:
   curl localhost:8080
   curl localhost:8080/health

4. Docker Compose (multi-service):
   docker compose up -d
*/