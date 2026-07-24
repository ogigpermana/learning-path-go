package main

import "fmt"

func main() {
	// Array
	angka := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array:", angka)
	
	// Slice
	fruits := []string{"apel", "jeruk", "mangga"}
	fruits = append(fruits, "pisang")
	fmt.Println("Buah:", fruits)
	
	// Loop
	for i, buah := range fruits {
		fmt.Printf("%d: %s\n", i, buah)
	}
}