package main

import "fmt"

func main() {
	var nama string = "Golang"
	umur := 2
	aktif := true
	
	fmt.Println("Bahasa:", nama)
	fmt.Println("Versi:", umur)
	fmt.Println("Sedang belajar:", aktif)
	
	// Konstanta
	const versi = "1.21"
	fmt.Println("Versi Go:", versi)
}