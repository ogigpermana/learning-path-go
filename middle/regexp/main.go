// File: middle/regexp/main.go
// Level: Middle
// Topik: Regular Expression (regexp)
//
// Package regexp untuk pencocokan pola teks.
// Syntax: RE2 (sama dengan RE2 di C++, Python, dll)
// - Compile: regexp.MustCompile (panic) / regexp.Compile (return error)
// - Methods: MatchString, FindString, FindAllString, ReplaceAllString, Split

package main

import (
	"fmt"
	"regexp"
)

func main() {
	// 1. MATCH - cek apakah string cocok dengan pattern
	fmt.Println("=== MatchString ===")
	emailPattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailPattern)

	emails := []string{
		"anggi@mail.com",
		"budi@company.co.id",
		"invalid-email",
		"test@.com",
	}

	for _, email := range emails {
		match := re.MatchString(email)
		fmt.Printf("%-25s -> %v\n", email, match)
	}
	fmt.Println()

	// 2. FIND - mencari string yang cocok
	fmt.Println("=== Find ===")
	text := "Hubungi kami di support@mail.com atau admin@company.com"
	re2 := regexp.MustCompile(`[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)

	fmt.Println("First match:", re2.FindString(text))
	fmt.Println("All matches:", re2.FindAllString(text, -1))
	fmt.Println()

	// 3. FIND WITH INDEX
	fmt.Println("=== Find with Index ===")
	loc := re2.FindStringIndex(text)
	fmt.Printf("Match at position %d-%d: %s\n", loc[0], loc[1], text[loc[0]:loc[1]])
	fmt.Println()

	// 4. SUBMATCH (capture groups)
	fmt.Println("=== Submatch (Capture Groups) ===")
	logLine := `2024-01-15 10:30:45 ERROR User login failed: invalid password`
	logPattern := `(\d{4}-\d{2}-\d{2}) (\d{2}:\d{2}:\d{2}) (\w+) (.*)`
	re3 := regexp.MustCompile(logPattern)

	matches := re3.FindStringSubmatch(logLine)
	fmt.Println("Full match:", matches[0])
	fmt.Println("Date:", matches[1])
	fmt.Println("Time:", matches[2])
	fmt.Println("Level:", matches[3])
	fmt.Println("Message:", matches[4])
	fmt.Println()

	// 5. REPLACE
	fmt.Println("=== Replace ===")
	tweet := "RT @user: Cek website kami di https://example.com! #golang"
	re4 := regexp.MustCompile(`https?://[^\s]+`)

	replaced := re4.ReplaceAllString(tweet, "[LINK]")
	fmt.Println("Original:", tweet)
	fmt.Println("Replaced:", replaced)
	fmt.Println()

	// 6. REPLACE WITH FUNCTION
	fmt.Println("=== Replace with Function ===")
	re5 := regexp.MustCompile(`\d+`)
	result := re5.ReplaceAllStringFunc("Harga: Rp10000 dan Rp25000", func(s string) string {
		return fmt.Sprintf("Rp%s", s)
	})
	fmt.Println("Result:", result)
	fmt.Println()

	// 7. SPLIT
	fmt.Println("=== Split ===")
	re6 := regexp.MustCompile(`[,;|\s]+`)
	parts := re6.Split("apel,jeruk;mangga|pisang anggur", -1)
	fmt.Println("Split result:", parts)
	fmt.Println()

	// 8. COMPILED REUSABLE
	fmt.Println("=== Compiled ===")
	phonePattern := regexp.MustCompile(`^(\+62|0)8[1-9][0-9]{7,11}$`)

	phones := []string{"081234567890", "+6281234567890", "12345", "081234"}
	for _, phone := range phones {
		fmt.Printf("%-15s -> valid: %v\n", phone, phonePattern.MatchString(phone))
	}
}

/*
Regex patterns:
\d = digit, \w = word char, \s = whitespace
.  = any char, * = 0 or more, + = 1 or more, ? = 0 or 1
{n} = exactly n, {n,} = n or more, {n,m} = n to m
[abc] = one of, [^abc] = not one of
^ = start, $ = end
| = OR, () = group
*/