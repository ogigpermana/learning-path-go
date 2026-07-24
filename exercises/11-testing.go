// File: 11-testing.go
// Level: Beginner
// Topik: Unit Testing
//
// Go memiliki built-in testing package.
// Aturan:
// 1. File test harus berakhiran _test.go
// 2. Fungsi test harus dimulai dengan Test
// 3. Parameter fungsi test: t *testing.T
// 4. Jalankan dengan: go test

package main

import "testing"

// Fungsi yang akan di-test
func Tambah(a, b int) int {
	return a + b
}

// TestTambah adalah fungsi unit test
// Nama fungsi harus dimulai dengan "Test"
func TestTambah(t *testing.T) {
	hasil := Tambah(2, 3)
	expected := 5

	// Jika hasil tidak sesuai, laporkan error
	if hasil != expected {
		t.Errorf("Tambah(2, 3) = %d; harus %d", hasil, expected)
	}
}

// Test dengan multiple cases
func TestTambahMulti(t *testing.T) {
	type testCase struct {
		a, b     int
		expected int
	}

	cases := []testCase{
		{2, 3, 5},
		{-1, 1, 0},
		{0, 5, 5},
		{100, -50, 50},
	}

	for _, tc := range cases {
		hasil := Tambah(tc.a, tc.b)
		if hasil != tc.expected {
			t.Errorf("Tambah(%d, %d) = %d; harus %d",
				tc.a, tc.b, hasil, tc.expected)
		}
	}
}

// Test helper - untuk setup dan teardown
func setupTest(t *testing.T) func() {
	t.Log("Setup sebelum test")
	return func() {
		t.Log("Teardown setelah test")
	}
}

func TestDenganSetup(t *testing.T) {
	teardown := setupTest(t)
	defer teardown() // dipanggil setelah test selesai

	hasil := Tambah(1, 1)
	if hasil != 2 {
		t.Fail()
	}
}

// Cara menjalankan:
// go test -v           -> verbose output
// go test -run TestTambah -> jalankan test spesifik
// go test -cover       -> lihat code coverage
//
// NOTE: file ini harus di-rename menjadi *_test.go
// agar bisa dijalankan dengan `go test`
// Contoh: go test 11-testing.go 11-testing_test.go -v