// File: 12-api.go
// Level: Beginner
// Topik: REST API sederhana dengan net/http
//
// Go memiliki package net/http untuk membuat HTTP server.
// Tidak perlu framework tambahan untuk API sederhana.
// Package encoding/json untuk serialisasi JSON.

package main

import (
	"encoding/json" // encode/decode JSON
	"fmt"
	"log"   // logging
	"net/http" // HTTP server & client
)

// Struct Pengguna dengan JSON tags
// Tags memberi tahu encoder/decoder JSON tentang nama field
type Pengguna struct {
	ID   int    `json:"id"`   // di JSON jadi "id"
	Nama string `json:"nama"` // di JSON jadi "nama"
}

// Data dummy - slice of Pengguna
var pengguna = []Pengguna{
	{ID: 1, Nama: "Anggi"},
	{ID: 2, Nama: "Budi"},
}

// Handler untuk endpoint /pengguna
// http.ResponseWriter: untuk menulis response
// http.Request: data request dari client
func handlerPengguna(w http.ResponseWriter, r *http.Request) {
	// Set header Content-Type ke JSON
	w.Header().Set("Content-Type", "application/json")

	// Encode slice pengguna ke JSON dan kirim ke response
	err := json.NewEncoder(w).Encode(pengguna)
	if err != nil {
		// Jika error, kirim HTTP 500 Internal Server Error
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	// Mendaftarkan handler ke path /pengguna
	http.HandleFunc("/pengguna", handlerPengguna)

	// Handler untuk root path
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Selamat datang di API Go!")
		fmt.Fprintln(w, "Endpoint: /pengguna")
	})

	// Menjalankan server di port 8080
	// log.Fatal akan mencatat error dan exit jika server gagal start
	fmt.Println("Server berjalan di http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Untuk mengakses:
// curl http://localhost:8080/pengguna
// curl http://localhost:8080/