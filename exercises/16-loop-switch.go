// File: 16-loop-switch.go
// Level: Beginner
// Topik: Perulangan dan Percabangan
//
// Go hanya punya satu keyword untuk loop: "for"
// - for init; condition; post {}
// - for condition {} (seperti while)
// - for {} (infinite loop)
//
// Switch: percabangan multi-kondisi
// - switch value {}
// - switch {} (dengan kondisi)

package main

import "fmt"

func main() {
	// 1. FOR BASIC
	fmt.Println("=== For basic ===")
	// init: i := 0, condition: i < 5, post: i++
	for i := 0; i < 5; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()

	// 2. FOR SEBAGAI WHILE
	fmt.Println("\n=== For as while ===")
	sum := 0
	for sum < 10 { // hanya condition, tidak ada init/post
		sum += 3
		fmt.Print(sum, " ")
	}
	fmt.Println()

	// 3. INFINITE LOOP + BREAK
	fmt.Println("\n=== Infinite loop dengan break ===")
	count := 0
	for {
		count++
		if count > 3 {
			break // keluar dari loop
		}
		fmt.Print(count, " ")
	}
	fmt.Println()

	// 4. CONTINUE - skip iterasi
	fmt.Println("\n=== Continue ===")
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue // skip angka genap
		}
		fmt.Print(i, " ")
	}
	fmt.Println()

	// 5. RANGE OVER SLICE
	fmt.Println("\n=== Range slice ===")
	fruits := []string{"apel", "jeruk", "mangga"}
	// i: index, fruit: value
	for i, fruit := range fruits {
		fmt.Printf("%d: %s\n", i, fruit)
	}

	// Range tanpa index
	fmt.Println("\nRange tanpa index:")
	for _, fruit := range fruits {
		fmt.Println("-", fruit)
	}

	// Range over map
	fmt.Println("\nRange over map:")
	scores := map[string]int{"Anggi": 90, "Budi": 85}
	for name, score := range scores {
		fmt.Printf("%s: %d\n", name, score)
	}

	// 6. SWITCH - percabangan nilai
	fmt.Println("\n=== Switch value ===")
	day := "senin"
	switch day {
	case "senin":
		fmt.Println("Hari kerja pertama")
	case "sabtu", "minggu": // multiple case
		fmt.Println("Akhir pekan")
	default:
		fmt.Println("Hari biasa")
	}

	// 7. SWITCH DENGAN KONDISI
	fmt.Println("\n=== Switch condition ===")
	score := 85
	switch {
	case score >= 90:
		fmt.Println("A")
	case score >= 80:
		fmt.Println("B")
	case score >= 70:
		fmt.Println("C")
	default:
		fmt.Println("D")
	}

	// 8. SWITCH TYPE - untuk interface
	fmt.Println("\n=== Switch type ===")
	var value any = 42
	switch v := value.(type) {
	case int:
		fmt.Println("Integer:", v)
	case string:
		fmt.Println("String:", v)
	default:
		fmt.Println("Tipe lain:", v)
	}
}