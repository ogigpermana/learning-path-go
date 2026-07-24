// File: 18-module/greeter/greeter.go
// Level: Beginner
// Topik: Membuat Package
//
// Package greeter adalah contoh pembuatan package sendiri.
// Function yang diawali huruf BESAR bisa diakses dari luar package (exported).
// Function yang diawali huruf kecil hanya bisa diakses dalam package (unexported).

package greeter

import "fmt"

// Hello adalah exported function (huruf besar)
// Karena exported, bisa dipanggil dari package lain
// Menerima nama dan mengembalikan sapaan
func Hello(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}

// Greet adalah exported function
// Menerima slice nama dan menyapa semua
func Greet(names []string) {
	for _, name := range names {
		fmt.Println(Hello(name))
	}
}

// helloInternal adalah unexported function (huruf kecil)
// Hanya bisa dipanggil dari dalam package greeter
func helloInternal() string {
	return "Internal greeting"
}