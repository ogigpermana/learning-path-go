// File: 03-fungsi.go
// Level: Beginner
// Topik: Fungsi
//
// Fungsi adalah blok kode yang bisa dipanggil berulang kali.
// Syntax: func namaFungsi(parameter) (returnType) { body }

package main

import "fmt"

// Fungsi dengan parameter dan return value
// a dan b bertipe int, return value bertipe int
func tambah(a int, b int) int {
	return a + b
}

// Fungsi dengan satu parameter dan return string
func sapa(nama string) string {
	return "Halo, " + nama + "!"
}

// Fungsi dengan multiple return values
// Mengembalikan dua nilai: string dan int
func info() (string, int) {
	return "Golang", 2024
}

// Fungsi dengan named return values
// Variabel sudah dideklarasi di signature fungsi
func bagi(a, b int) (hasil int, err bool) {
	if b == 0 {
		return 0, true // error: pembagi nol
	}
	hasil = a / b // tidak perlu := karena sudah dideklarasi
	err = false
	return // naked return - mengembalikan hasil dan err
}

// Variadic function - menerima jumlah parameter tidak tetap
func jumlahkan(angka ...int) int {
	total := 0
	for _, n := range angka {
		total += n
	}
	return total
}

func main() {
	// Memanggil fungsi tambah
	hasil := tambah(5, 3)
	fmt.Println("5 + 3 =", hasil)

	// Memanggil fungsi sapa
	fmt.Println(sapa("Teman"))

	// Menangkap multiple return values
	bahasa, tahun := info()
	fmt.Println(bahasa, "dirilis", tahun)

	// Named return
	h, isErr := bagi(10, 2)
	fmt.Println("10/2 =", h, "Error:", isErr)

	// Variadic function
	total := jumlahkan(1, 2, 3, 4, 5)
	fmt.Println("Total:", total)

	// Anonymous function - fungsi tanpa nama
	sayHello := func(name string) string {
		return "Hello " + name
	}
	fmt.Println(sayHello("World"))
}