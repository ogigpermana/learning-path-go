// File: 18-module/greeter/go.mod
// Level: Beginner
// Topik: Go Module File
//
// go.mod adalah file konfigurasi module Go.
// Berisi:
// - module name: nama module (bisa di-import oleh module lain)
// - go version: versi Go yang digunakan
// - require: dependensi eksternal
// - replace: mengganti path module (untuk local development)

module example.com/greeter

go 1.22

// Untuk menambah dependensi:
// go get github.com/gorilla/mux
//
// Untuk update dependensi:
// go mod tidy