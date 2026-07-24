// File: 18-module/main/main.go
// Level: Beginner
// Topik: Go Modules
//
// Go Modules adalah sistem manajemen dependensi resmi Go.
// Setiap project Go dimulai dengan `go mod init nama-module`.
//
// Langkah-langkah:
// 1. Buat folder project
// 2. go mod init nama-module
// 3. Tulis kode
// 4. go mod tidy (jika ada dependensi eksternal)
//
// Struktur:
// module-example/
// ├── greeter/        # package sendiri
// │   ├── go.mod      # module greeter
// │   └── greeter.go  # kode package
// └── main/           # program utama
//     ├── go.mod      # module main (depend ke greeter)
//     └── main.go     # entry point

package main

import (
	"fmt"

	// Mengimport package lokal (local module)
	// Path mengikuti module name di go.mod
	"example.com/greeter"
)

func main() {
	// Memanggil function dari package greeter
	// Function yang bisa diakses: diawali huruf BESAR (exported)
	message := greeter.Hello("Anggi")
	fmt.Println(message)

	// Contoh lain
	names := []string{"Anggi", "Budi", "Citra"}
	greeter.Greet(names)
}

// Cara menjalankan:
// cd 18-module/main
// go run main.go
//
// Cara membuat module sendiri:
// cd 18-module/greeter
// go mod init example.com/greeter
//
// Cara menghubungkan main ke greeter:
// cd 18-module/main
// go mod init example.com/main
// go mod edit -replace example.com/greeter=../greeter
// go mod tidy
//
// Untuk dependensi eksternal:
// go get github.com/gorilla/mux