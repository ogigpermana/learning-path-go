// File: 06-struct-method.go
// Level: Beginner
// Topik: Struct dan Method (lanjutan)
//
// Method adalah fungsi yang receiver-nya adalah struct.
// Dua jenis receiver:
// 1. Value receiver  (m Mobil) - tidak mengubah nilai asli
// 2. Pointer receiver (m *Mobil) - bisa mengubah nilai asli

package main

import "fmt"

// Mendefinisikan struct Mobil
type Mobil struct {
	Merek string
	Warna string
	Tahun int
}

// Value receiver - hanya membaca data
// Method ini tidak mengubah struct aslinya
func (m Mobil) TampilInfo() string {
	return fmt.Sprintf("Mobil %s berwarna %s tahun %d",
		m.Merek, m.Warna, m.Tahun)
}

// Pointer receiver - bisa mengubah data struct asli
// Mengubah tahun mobil menjadi 2025
func (m *Mobil) Perbaiki() {
	m.Tahun = 2025
}

// Value receiver - method dengan parameter tambahan
func (m Mobil) IsTua() bool {
	return m.Tahun < 2020
}

func main() {
	// Inisialisasi struct Mobil
	gojek := Mobil{Merek: "Gojek", Warna: "Hitam", Tahun: 2020}

	// Memanggil method value receiver
	fmt.Println(gojek.TampilInfo())

	// Memanggil method value receiver (cek kondisi)
	fmt.Println("Apakah mobil tua?", gojek.IsTua())

	// Memanggil method pointer receiver - mengubah data
	gojek.Perbaiki()
	fmt.Println("Setelah diperbaiki:", gojek.TampilInfo())

	// Method bisa dipanggil dari pointer atau value
	// Go otomatis mengkonversi
	mobilPtr := &gojek
	mobilPtr.Perbaiki()

	// Chain method tidak default di Go, tapi bisa dibuat
	// dengan mengembalikan pointer
	type Builder struct {
		parts []string
	}
	(builder)
}