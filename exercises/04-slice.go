// File: 04-slice.go
// Level: Beginner
// Topik: Array dan Slice
//
// Array: kumpulan data dengan ukuran tetap
// Slice: array dinamis (ukuran bisa berubah)

package main

import "fmt"

func main() {
	// ARRAY - ukuran tetap [5] berarti 5 elemen
	var angka [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array:", angka)
	fmt.Println("Elemen ke-2:", angka[1]) // Index dimulai dari 0

	// SLICE - dinamis, tidak perlu ukuran
	var fruits []string           // Deklarasi slice kosong
	fruits = []string{"apel", "jeruk", "mangga"}
	fruits = append(fruits, "pisang") // append untuk menambah elemen
	fmt.Println("Buah:", fruits)
	fmt.Println("Panjang:", len(fruits))
	fmt.Println("Kapasitas:", cap(fruits))

	// Slice dari array
	sliceDariArray := angka[1:4] // index 1 sampai 3 (4 tidak termasuk)
	fmt.Println("Slice dari array:", sliceDariArray)

	// Make - membuat slice dengan kapasitas awal
	numbers := make([]int, 3, 5) // len=3, cap=5
	numbers[0] = 10
	numbers[1] = 20
	numbers[2] = 30
	numbers = append(numbers, 40, 50)
	fmt.Println("Numbers:", numbers)

	// Loop over slice dengan range
	fmt.Println("\nDaftar Buah:")
	for i, fruit := range fruits {
		fmt.Printf("%d: %s\n", i, fruit)
	}

	// Loop tanpa index (gunakan _)
	fmt.Println("\nNama Buah:")
	for _, fruit := range fruits {
		fmt.Println("-", fruit)
	}

	// Slice 2 dimensi
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println("Matrix:", matrix)

	// Copy slice
	src := []int{1, 2, 3}
	dst := make([]int, len(src))
	copy(dst, src)
	fmt.Println("Copy:", dst)
}