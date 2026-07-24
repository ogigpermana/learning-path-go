package main

import "fmt"

type BangunDatar interface {
	Luas() float64
	Nama() string
}

type Persegi struct {
	Sisi float64
}

func (p Persegi) Luas() float64 {
	return p.Sisi * p.Sisi
}

func (p Persegi) Nama() string {
	return "Persegi"
}

type Lingkaran struct {
	JariJari float64
}

func (l Lingkaran) Luas() float64 {
	return 3.14159 * l.JariJari * l.JariJari
}

func (l Lingkaran) Nama() string {
	return "Lingkaran"
}

func CetakLuas(bd BangunDatar) {
	fmt.Printf("%s memiliki luas %.2f\n", bd.Nama(), bd.Luas())
}

func main() {
	benda := []BangunDatar{
		Persegi{Sisi: 5},
		Lingkaran{JariJari: 7},
		Persegi{Sisi: 10},
		Lingkaran{JariJari: 3.5},
	}

	for _, item := range benda {
		CetakLuas(item)
	}
}