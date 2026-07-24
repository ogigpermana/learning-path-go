package main

import "fmt"

func tambah(a, b int) int {
	return a + b
}

func sapa(nama string) string {
	return "Halo, " + nama + "!"
}

func info() (string, int) {
	return "Golang", 2024
}

func main() {
	hasil := tambah(5, 3)
	fmt.Println("5 + 3 =", hasil)
	
	fmt.Println(sapa("Teman"))
	
	bahasa, tahun := info()
	fmt.Println(bahasa, "dirilis", tahun)
}