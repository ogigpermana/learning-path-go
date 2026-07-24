// File: 08-polymorphism.go
// Level: Beginner
// Topik: Polymorphism (Polimorfisme) via Interface
//
// Polimorfisme = kemampuan objek memiliki banyak bentuk.
// Di Go, polimorfisme dicapai melalui interface.
// Satu interface bisa menampung berbagai tipe data selama
// tipe tersebut mengimplementasi semua method di interface.

package main

import (
	"fmt"
	"math"
)

// Interface BangunDatar - kontrak untuk semua bangun datar
// Setiap bangun datar harus memiliki method Luas() dan Nama()
type BangunDatar interface {
	Luas() float64 // method untuk menghitung luas
	Nama() string  // method untuk mengembalikan nama bangun
}

// Struct Persegi
type Persegi struct {
	Sisi float64
}

// Persegi mengimplementasi interface BangunDatar
func (p Persegi) Luas() float64 {
	return p.Sisi * p.Sisi // rumus luas persegi: sisi x sisi
}
func (p Persegi) Nama() string {
	return "Persegi"
}

// Struct Lingkaran
type Lingkaran struct {
	JariJari float64
}

// Lingkaran mengimplementasi interface BangunDatar
func (l Lingkaran) Luas() float64 {
	return math.Pi * l.JariJari * l.JariJari // rumus luas: π x r²
}
func (l Lingkaran) Nama() string {
	return "Lingkaran"
}

// Struct Segitiga
type Segitiga struct {
	Alas   float64
	Tinggi float64
}

// Segitiga mengimplementasi interface BangunDatar
func (s Segitiga) Luas() float64 {
	return 0.5 * s.Alas * s.Tinggi // rumus luas: ½ x a x t
}
func (s Segitiga) Nama() string {
	return "Segitiga"
}

// Fungsi yang menerima interface BangunDatar
// Bisa menerima Persegi, Lingkaran, Segitiga, dll
// Inilah polimorfisme: satu fungsi bisa menangani banyak tipe
func CetakLuas(bd BangunDatar) {
	fmt.Printf("%s memiliki luas %.2f\n", bd.Nama(), bd.Luas())
}

func main() {
	// Slice berisi berbagai tipe yang implementasi BangunDatar
	benda := []BangunDatar{
		Persegi{Sisi: 5},
		Lingkaran{JariJari: 7},
		Segitiga{Alas: 5, Tinggi: 10},
		Persegi{Sisi: 10},
		Lingkaran{JariJari: 3.5},
	}

	// Loop dan panggil CetakLuas() untuk setiap elemen
	// Meskipun tipenya berbeda, mereka bisa diperlakukan sama
	for _, item := range benda {
		CetakLuas(item)
	}
}