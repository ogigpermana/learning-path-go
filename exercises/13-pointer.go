// File: 13-pointer.go
// Level: Beginner
// Topik: Pointer
//
// Pointer adalah variabel yang menyimpan alamat memori variabel lain.
// Alih-alih menyimpan nilai, pointer menyimpan lokasi di memori.
//
// Operator:
// &  -> mengambil alamat memori (address-of)
// *  -> mengambil nilai dari alamat (dereference)

package main

import "fmt"

func main() {
	// Variabel biasa
	var x int = 10
	fmt.Println("Nilai x:", x)
	fmt.Println("Alamat x:", &x) // &x = alamat memori x

	// Pointer p menunjuk ke alamat x
	// *int berarti "pointer to int"
	var p *int = &x
	fmt.Println("Nilai p (alamat x):", p)
	fmt.Println("Dereference p:", *p) // *p = nilai di alamat tersebut

	// Mengubah nilai melalui pointer
	*p = 20 // set nilai di alamat yang ditunjuk p menjadi 20
	fmt.Println("Setelah *p = 20, x menjadi:", x) // x ikut berubah

	// Pointer ke struct
	type User struct {
		Name string
		Age  int
	}

	// Membuat pointer langsung ke struct dengan &
	u := &User{Name: "Anggi", Age: 20}
	fmt.Println("User:", u.Name, u.Age) // Go otomatis dereference

	// Pointer sebagai parameter fungsi - pass by reference
	// Tanpa pointer, fungsi hanya menerima copy (pass by value)
	angka := 5
	double(angka)          // pass by value - tidak mengubah asli
	fmt.Println("Setelah double:", angka) // masih 5
	doublePtr(&angka)      // pass by reference - mengubah asli
	fmt.Println("Setelah doublePtr:", angka) // jadi 10

	// Nil pointer - pointer yang belum punya alamat
	var ptr *int
	if ptr == nil {
		fmt.Println("ptr adalah nil (belum指向 alamat manapun)")
	}

	// Zero value dari pointer adalah nil
	// Mengakses *ptr saat nil akan menyebabkan panic
}

// Pass by value - parameter adalah copy
func double(n int) {
	n = n * 2 // hanya mengubah copy lokal
}

// Pass by reference - parameter adalah pointer
func doublePtr(n *int) {
	*n = *n * 2 // mengubah nilai asli melalui pointer
}