// File: middle/testing/main.go
// Level: Middle
// Topik: Testing Lanjutan
//
// Go memiliki package "testing" built-in yang powerful.
// Fitur:
// 1. Table-driven tests - test dengan multiple test cases
// 2. Benchmark - mengukur performa kode
// 3. TestMain - setup/teardown untuk semua test
// 4. Coverage - persentase kode yang di-test
//
// NOTE: File ini berisi fungsi yang akan di-test DAN fungsi test.
// Untuk menjalankan, rename file jadi *_test.go atau gunakan go test -run.

package main

import "testing"

// ===== FUNGSI YANG AKAN DI-TEST =====

// Tambah menjumlahkan dua bilangan
func Tambah(a, b int) int {
	return a + b
}

// IsEven mengecek apakah bilangan genap
func IsEven(n int) bool {
	return n%2 == 0
}

// ===== UNIT TEST =====

// TestTambah adalah contoh table-driven test
// Table-driven test: mendefinisikan test cases dalam slice struct
func TestTambah(t *testing.T) {
	// Test cases dalam bentuk slice struct
	tests := []struct {
		name     string // nama test case
		a, b     int    // input
		expected int    // expected output
	}{
		{name: "positive numbers", a: 2, b: 3, expected: 5},
		{name: "negative numbers", a: -1, b: 1, expected: 0},
		{name: "zero", a: 0, b: 0, expected: 0},
		{name: "large numbers", a: 1000, b: 2000, expected: 3000},
	}

	// Iterasi test cases
	for _, tt := range tests {
		// t.Run() menjalankan sub-test dengan nama
		t.Run(tt.name, func(t *testing.T) {
			result := Tambah(tt.a, tt.b)
			if result != tt.expected {
				// t.Errorf melaporkan error tapi melanjutkan test
				t.Errorf("Tambah(%d, %d) = %d; want %d",
					tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

// TestIsEven adalah test sederhana
func TestIsEven(t *testing.T) {
	if !IsEven(2) {
		t.Error("2 harus even")
	}
	if IsEven(3) {
		t.Error("3 harus odd")
	}
}

// ===== BENCHMARK =====

// BenchmarkTambah mengukur performa fungsi Tambah
// Benchmark dijalankan dengan: go test -bench=.
// b.N diatur oleh Go untuk mendapatkan hasil yang stabil
func BenchmarkTambah(b *testing.B) {
	// Loop b.N kali - jumlah loop diatur oleh Go
	for i := 0; i < b.N; i++ {
		Tambah(100, 200)
	}
}

// ===== TEST MAIN =====

// TestMain adalah entry point untuk semua test dalam package
// Berguna untuk setup/teardown global
func TestMain(m *testing.M) {
	// Setup: dijalankan sebelum semua test
	println("=== SETUP: Membuat database test ===")

	// m.Run() menjalankan semua test
	code := m.Run()

	// Teardown: dijalankan setelah semua test selesai
	println("=== TEARDOWN: Membersihkan database test ===")

	// os.Exit() dengan kode dari test
	// NOTE: untuk actual code, gunakan os.Exit(code)
	println("Exit code:", code)
}

// ===== RUNNING TESTS =====
// go test -v              -> verbose, lihat detail setiap test
// go test -run TestTambah -> jalankan test spesifik
// go test -bench=.        -> jalankan benchmark
// go test -bench=Bench    -> jalankan semua benchmark
// go test -cover          -> lihat code coverage
// go test -coverprofile=coverage.out -> output coverage ke file