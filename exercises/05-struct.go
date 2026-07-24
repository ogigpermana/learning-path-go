// File: 05-struct.go
// Level: Beginner
// Topik: Struct
//
// Struct adalah kumpulan field (seperti class di OOP lain)
// Digunakan untuk membuat tipe data kustom

package main

import "fmt"

// Mendefinisikan struct Mahasiswa
// Field: Nama (string), Usia (int), NIM (string)
type Mahasiswa struct {
	Nama string
	Usia int
	NIM  string
}

// Method: fungsi yang menempel pada struct
// (m Mahasiswa) disebut receiver - mirip method di OOP
func (m Mahasiswa) Tampil() string {
	return fmt.Sprintf("%s (%s, %d th)", m.Nama, m.NIM, m.Usia)
}

// Method dengan pointer receiver
// Bisa mengubah nilai asli struct
func (m *Mahasiswa) Ultah() {
	m.Usia++ // menambah usia, perubahan terasa di luar fungsi
}

func main() {
	// Cara 1: inisialisasi dengan field names
	anggi := Mahasiswa{
		Nama: "Anggi",
		Usia: 20,
		NIM:  "12345",
	}

	// Cara 2: inisialisasi berdasarkan urutan field
	budi := Mahasiswa{"Budi", 22, "67890"}

	// Cara 3: deklarasi lalu isi
	var citra Mahasiswa
	citra.Nama = "Citra"
	citra.Usia = 19
	citra.NIM = "11111"

	fmt.Println("Anggi:", anggi.Tampil())
	fmt.Println("Budi:", budi.Tampil())
	fmt.Println("Citra:", citra.Tampil())

	// Method pointer receiver
	fmt.Println("\nSebelum ultah:", anggi.Usia)
	anggi.Ultah()
	fmt.Println("Setelah ultah:", anggi.Usia)

	// Struct embedding (seperti inheritance)
	type MahasiswaAktif struct {
		Mahasiswa          // embedded struct
		Semester   int
		IPK        float64
	}

	aktif := MahasiswaAktif{
		Mahasiswa: Mahasiswa{"Dewi", 21, "22222"},
		Semester:  4,
		IPK:       3.5,
	}
	fmt.Printf("\nMahasiswa Aktif: %s, Semester %d, IPK %.2f\n",
		aktif.Nama, aktif.Semester, aktif.IPK)
}