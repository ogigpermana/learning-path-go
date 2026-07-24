// File: middle/godoc/main.go
// Level: Middle
// Topik: Dokumentasi dengan Godoc
//
// Go memiliki built-in documentation tool: godoc (atau go doc).
// Dokumentasi ditulis langsung di kode sebagai komentar.
//
// Aturan:
// - Komentar sebelum deklarasi package menjadi "package documentation"
// - Komentar sebelum fungsi/type/variable menjadi dokumentasi item tersebut
// - Mulai dengan nama item yang didokumentasikan

// Package godoc mendemonstrasikan cara membuat dokumentasi di Go.
//
// Dokumentasi package ini akan muncul saat:
//   go doc ./middle/godoc
//   godoc -http=:6060 (buka browser localhost:6060)
package main

import "fmt"

// Person merepresentasikan data seseorang.
// Field Name wajib diisi, Age optional (0 = tidak diketahui).
type Person struct {
	// Name adalah nama lengkap (wajib diisi)
	Name string

	// Age adalah umur dalam tahun (0 jika tidak diketahui)
	Age int
}

// NewPerson membuat Person baru dengan nilai default.
// Mengembalikan pointer ke Person.
//
// Contoh:
//
//	p := NewPerson("Anggi", 20)
//	fmt.Println(p.Greet())
func NewPerson(name string, age int) *Person {
	if age < 0 {
		age = 0
	}
	return &Person{Name: name, Age: age}
}

// Greet mengembalikan string sapaan untuk Person.
// Format: "Hello, my name is {Name}. I am {Age} years old."
func (p *Person) Greet() string {
	return fmt.Sprintf("Hello, my name is %s. I am %d years old.", p.Name, p.Age)
}

// IsAdult mengecek apakah Person sudah dewasa (>= 17 tahun).
func (p *Person) IsAdult() bool {
	return p.Age >= 17
}

func main() {
	fmt.Println("=== Godoc Demo ===")
	fmt.Println()
	fmt.Println("Menulis dokumentasi yang baik:")
	fmt.Println()
	fmt.Println("1. Package comment (sebelum package declaration)")
	fmt.Println("2. Exported item comment (sebelum type/func/variable)")
	fmt.Println("3. Contoh kode (dalam file example_test.go)")
	fmt.Println()
	fmt.Println("Lihat dokumentasi dengan:")
	fmt.Println("  go doc ./middle/godoc")
	fmt.Println("  go doc ./middle/godoc.Person")
	fmt.Println("  go doc ./middle/godoc.NewPerson")
	fmt.Println()
	fmt.Println("Generate dokumentasi HTML:")
	fmt.Println("  godoc -http=:6060")
	fmt.Println("  # buka http://localhost:6060")
	fmt.Println()
	fmt.Println("Akses dokumentasi online:")
	fmt.Println("  pkg.go.dev/std")
	fmt.Println("  pkg.go.dev/github.com/user/repo")

	p := NewPerson("Anggi", 20)
	fmt.Println(p.Greet())
	fmt.Println("Adult:", p.IsAdult())
}

/*
Godoc conventions:
1. Komentar satu baris: // Comment
2. Komentar multi baris: /* Comment * /
3. First word after // harus nama item
4. Gunakan blank line untuk paragraph
5. Indent untuk preformatted text
6. Contoh di file *_test.go dengan fungsi ExampleXxx()

Example file: godoc_test.go
func ExamplePerson_Greet() {
    p := &Person{Name: "Anggi", Age: 20}
    fmt.Println(p.Greet())
    // Output: Hello, my name is Anggi. I am 20 years old.
}

Jalankan: go test -v ./middle/godoc/
*/