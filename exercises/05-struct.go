package main

import "fmt"

type Mahasiswa struct {
	Nama  string
	Usia  int
	NIM   string
}

func (m Mahasiswa) Tampil() string {
	return fmt.Sprintf("%s (%s, %d th)", m.Nama, m.NIM, m.Usia)
}

func main() {
	// Membuat struct
	anggi := Mahasiswa{
		Nama: "Anggi",
		Usia: 20,
		NIM:  "12345",
	}
	
	fmt.Println(anggi.Tampil())
}