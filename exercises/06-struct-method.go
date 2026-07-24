package main

import "fmt"

type Mobil struct {
	Warna  string
	Merek  string
	Tahun  int
}

func (m Mobil) TampilInfo() string {
	return fmt.Sprintf("Mobil %s berwarna %s tahun %d", m.Merek, m.Warna, m.Tahun)
}

func (m *Mobil) Perbaiki() {
	m.Tahun = 2025
}

func main() {
	gojek := Mobil{Warna: "Hitam", Merek: "Gojek", Tahun: 2020}
	
	fmt.Println(gojek.TampilInfo())
	
	gojek.Perbaiki()
	fmt.Println("Setelah diperbaiki:", gojek.TampilInfo())
}