// File: 17-string.go
// Level: Beginner
// Topik: Manipulasi String
//
// String di Go adalah immutable (tidak bisa diubah).
// Manipulasi string menggunakan package "strings" dan "strconv".
// Untuk operasi string yang banyak, gunakan strings.Builder.

package main

import (
	"fmt"
	"strconv"  // konversi string <-> number
	"strings"  // manipulasi string
)

func main() {
	// 1. OPERASI DASAR STRING
	s := "Hello, Golang!"
	fmt.Println("String asli:", s)
	fmt.Println("Panjang:", len(s))           // jumlah byte
	fmt.Println("Char pertama:", string(s[0])) // H
	fmt.Println("Char ke-7:", string(s[7]))   // G
	fmt.Println("Substring [0:5]:", s[0:5])   // Hello
	fmt.Println("Substring [7:]:", s[7:])     // Golang!

	// 2. PACKAGE STRINGS
	fmt.Println("\n=== Package strings ===")
	fmt.Println("ToUpper:", strings.ToUpper(s))
	fmt.Println("ToLower:", strings.ToLower(s))
	fmt.Println("Contains 'Go':", strings.Contains(s, "Go"))
	fmt.Println("HasPrefix 'Hello':", strings.HasPrefix(s, "Hello"))
	fmt.Println("HasSuffix '!':", strings.HasSuffix(s, "!"))

	// Replace
	fmt.Println("Replace Go->Rust:", strings.Replace(s, "Go", "Rust", 1))
	fmt.Println("ReplaceAll a->x:", strings.ReplaceAll(s, "a", "x"))

	// Split dan Join
	parts := strings.Split(s, ", ")
	fmt.Println("Split by ', ':", parts)
	fmt.Println("Join:", strings.Join([]string{"a", "b", "c"}, "-"))

	// Trim
	text := "  hello world  "
	fmt.Println("TrimSpace:", "'"+strings.TrimSpace(text)+"'")
	fmt.Println("Trim ' ': ", "'"+strings.Trim(text, " ")+"'")

	// Count dan Repeat
	fmt.Println("Count 'l':", strings.Count(s, "l"))
	fmt.Println("Repeat 'Go' x3:", strings.Repeat("Go", 3))

	// 3. KONVERSI STRING <-> NUMBER
	fmt.Println("\n=== Konversi ===")
	num := 42
	strNum := strconv.Itoa(num) // Int to string
	fmt.Printf("%d -> '%s'\n", num, strNum)

	parsed, err := strconv.Atoi("100") // String to int
	if err != nil {
		fmt.Println("Error parsing:", err)
	} else {
		fmt.Println("Parsed + 1:", parsed+1)
	}

	// ParseFloat
	floatVal, _ := strconv.ParseFloat("3.14", 64)
	fmt.Println("ParseFloat:", floatVal)

	// 4. STRINGS.BUILDER (performance untuk concatenation)
	fmt.Println("\n=== strings.Builder ===")
	var sb strings.Builder
	sb.WriteString("Ini ")
	sb.WriteString("adalah ")
	sb.WriteString("kalimat ")
	sb.WriteString("panjang")
	fmt.Println("Builder result:", sb.String())
	fmt.Println("Builder len:", sb.Len())
}