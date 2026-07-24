package main

import "testing"

func Tambah(a, b int) int {
	return a + b
}

func TestTambah(t *testing.T) {
	hasil := Tambah(2, 3)
	if hasil != 5 {
		t.Errorf("Tambah(2, 3) = %d; harus 5", hasil)
	}
}

func TestTambahNegatif(t *testing.T) {
	hasil := Tambah(-1, 1)
	if hasil != 0 {
		t.Errorf("Tambah(-1, 1) = %d; harus 0", hasil)
	}
}