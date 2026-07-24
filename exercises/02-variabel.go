// File: 02-variabel.go
// Level: Beginner
// Topik: Variabel dan Tipe Data
//
// Go memiliki beberapa cara untuk mendeklarasikan variabel:
// 1. var nama tipe = nilai  -> deklarasi eksplisit
// 2. nama := nilai          -> deklarasi singkat (type inference)
// 3. const nama = nilai     -> konstanta (tidak bisa diubah)

package main

import "fmt"

func main() {
	// Deklarasi eksplisit dengan var
	var nama string = "Golang" // string: kumpulan karakter
	var umur int = 2           // int: bilangan bulat

	// Deklarasi singkat dengan :=
	// Go akan otomatis mendeteksi tipe data
	aktif := true // bool: true/false
	versi := 1.21 // float64: bilangan desimal

	fmt.Println("Bahasa:", nama)
	fmt.Println("Versi:", umur)
	fmt.Println("Sedang belajar:", aktif)

	// Konstanta - nilainya tidak bisa diubah setelah dideklarasi
	const appName = "Go Learning"
	fmt.Println("Aplikasi:", appName)

	// Multiple variabel dalam satu baris
	var a, b int = 1, 2
	fmt.Println("a =", a, "b =", b)

	// Zero value: variabel tanpa nilai awal
	var defaultInt int     // 0
	var defaultString string // "" (string kosong)
	var defaultBool bool   // false
	fmt.Printf("Default: int=%d, string='%s', bool=%v\n",
		defaultInt, defaultString, defaultBool)
}