// File: expert/fuzzing/main.go
// Level: Expert
// Topik: Fuzzing
//
// Fuzzing adalah teknik testing dengan input random/otomatis
// untuk menemukan bug, crash, atau security vulnerability.
//
// Go 1.18+ memiliki built-in fuzzing.
// Cara jalan: go test -fuzz=FuzzFuncName ./...
//
// Fuzz function signature:
// func FuzzXxx(f *testing.F) {
//     f.Add("seed_corpus") // sample inputs
//     f.Fuzz(func(t *testing.T, data string) {
//         // test dengan input random
//     })
// }

package main

import "testing"

// Fungsi yang akan di-fuzz
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// Fungsi yang RENTAN terhadap input tertentu (contoh crash)
func ParseAndProcess(data string) (string, error) {
	if len(data) == 0 {
		return "", nil
	}

	// Bug: panic jika data[0] adalah karakter tertentu
	if data[0] == 0xFF {
		panic("crash: invalid byte sequence")
	}

	// Bug lain: integer overflow
	if len(data) > 10000 {
		panic("crash: too long")
	}

	return data, nil
}

// Fuzz test untuk Reverse
// Jalankan: go test -fuzz=FuzzReverse -fuzztime=10s
func FuzzReverse(f *testing.F) {
	// Seed corpus (input awal yang diketahui benar)
	f.Add("Hello, World!")
	f.Add("12345")
	f.Add("Go is awesome")
	f.Add("") // empty string
	f.Add("!@#$%")

	f.Fuzz(func(t *testing.T, input string) {
		// Property-based testing:
		// Property: Reverse(Reverse(s)) == s
		reversed := Reverse(input)
		doubleReversed := Reverse(reversed)

		if input != doubleReversed {
			t.Errorf("Double reverse failed: input=%q, reversed=%q, double=%q",
				input, reversed, doubleReversed)
		}

		// Property: panjang string tetap sama
		if len(input) != len(reversed) {
			t.Errorf("Length mismatch: input len=%d, reversed len=%d",
				len(input), len(reversed))
		}
	})
}

// Fuzz test untuk ParseAndProcess
// Cari crash/panic: go test -fuzz=FuzzParse
func FuzzParse(f *testing.F) {
	f.Add("valid input")
	f.Add("another test")

	f.Fuzz(func(t *testing.T, data string) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("Panic dengan input %q: %v", data, r)
			}
		}()

		result, err := ParseAndProcess(data)
		if err == nil && result != data {
			t.Errorf("Hasil tidak sesuai: got %q, want %q", result, data)
		}
	})
}

/*
Cara menjalankan fuzzing:

# Basic fuzzing
go test -fuzz=FuzzReverse -fuzztime=10s

# Dengan verbose
go test -fuzz=FuzzReverse -fuzztime=30s -v

# Stop setelah menemukan bug
go test -fuzz=FuzzParse -fuzztime=60s

# Test corpus tanpa fuzzing
go test -run=FuzzReverse

Tips:
- Semakin lama fuzztime, semakin banyak input random yang dicoba
- Cache korpus ada di folder testdata/fuzz/
- Fix bug lalu jalankan ulang untuk verifikasi
- Fuzzing sangat efektif untuk menemukan crash dan security issues
*/