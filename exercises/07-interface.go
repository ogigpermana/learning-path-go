package main

import "fmt"

type Benda struct {
	Nama  string
	Harga int
}

type Hewan struct {
	Nama string
	Umur int
}

type Info interface {
	GetInfo() string
}

func (b Benda) GetInfo() string {
	return fmt.Sprintf("Benda: %s, Harga: Rp%d", b.Nama, b.Harga)
}

func (h Hewan) GetInfo() string {
	return fmt.Sprintf("Hewan: %s, Umur: %d tahun", h.Nama, h.Umur)
}

func TampilInfo(item Info) {
	fmt.Println(item.GetInfo())
}

func main() {
	sepatu := Benda{Nama: "Sepatu", Harga: 500000}
	kucing := Hewan{Nama: "Kucing", Umur: 3}
	
	TampilInfo(sepatu)
	TampilInfo(kucing)
}