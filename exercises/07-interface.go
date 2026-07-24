// File: 07-interface.go
// Level: Beginner
// Topik: Interface
//
// Interface adalah kumpulan method signatures (kontrak).
// Tipe data yang memiliki semua method dalam interface
// secara otomatis dianggap mengimplementasi interface tersebut.
// Ini disebut "structural typing" atau "duck typing".

package main

import "fmt"

// Mendefinisikan interface dengan satu method GetInfo()
// Semua tipe yang memiliki method GetInfo() string otomatis
// mengimplementasi interface ini
type Info interface {
	GetInfo() string
}

// Struct Benda dengan field Nama dan Harga
type Benda struct {
	Nama  string
	Harga int
}

// Struct Benda mengimplementasi interface Info
// karena memiliki method GetInfo() string
func (b Benda) GetInfo() string {
	return fmt.Sprintf("Benda: %s, Harga: Rp%d", b.Nama, b.Harga)
}

// Struct Hewan dengan field Nama dan Umur
type Hewan struct {
	Nama string
	Umur int
}

// Struct Hewan juga mengimplementasi interface Info
func (h Hewan) GetInfo() string {
	return fmt.Sprintf("Hewan: %s, Umur: %d tahun", h.Nama, h.Umur)
}

// Fungsi yang menerima interface Info
// Bisa menerima Benda, Hewan, atau tipe lain yang
// mengimplementasi interface Info
func TampilInfo(item Info) {
	// Memanggil method yang didefinisikan di interface
	fmt.Println(item.GetInfo())
}

// Type assertion: memeriksa tipe asli dari interface
func Jenis(item Info) {
	switch v := item.(type) {
	case Benda:
		fmt.Printf("Ini benda dengan harga Rp%d\n", v.Harga)
	case Hewan:
		fmt.Printf("Ini hewan umur %d tahun\n", v.Umur)
	default:
		fmt.Println("Tipe tidak dikenal")
	}
}

func main() {
	// Membuat instance Benda dan Hewan
	sepatu := Benda{Nama: "Sepatu", Harga: 500000}
	kucing := Hewan{Nama: "Kucing", Umur: 3}

	// Memanggil fungsi dengan parameter interface
	TampilInfo(sepatu)
	TampilInfo(kucing)

	// Type assertion
	Jenis(sepatu)
	Jenis(kucing)

	// Interface kosong (any) - bisa menampung tipe apapun
	var apapun any = "Hello"
	fmt.Println("any:", apapun)
	apapun = 42
	fmt.Println("any:", apapun)
	apapun = Benda{"Meja", 100000}
	fmt.Println("any:", apapun.(Benda).Nama)
}