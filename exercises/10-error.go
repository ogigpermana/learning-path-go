// File: 10-error.go
// Level: Beginner
// Topik: Error Handling
//
// Go tidak menggunakan try-catch untuk error handling.
// Sebaliknya, fungsi mengembalikan error sebagai return value.
// Pola umum: if err != nil { handle error }
//
// error adalah interface built-in:
// type error interface {
//     Error() string
// }

package main

import (
	"errors"  // package untuk membuat error
	"fmt"
	"strconv" // package untuk konversi string
)

// Fungsi bagi dengan error handling
// Mengembalikan error jika input tidak valid
func bagi(a, b string) (float64, error) {
	// strconv.ParseFloat mengembalikan (float64, error)
	angkaA, err := strconv.ParseFloat(a, 64)
	if err != nil {
		// errors.New() membuat error baru
		return 0, errors.New("angka pertama tidak valid: " + err.Error())
	}

	angkaB, err := strconv.ParseFloat(b, 64)
	if err != nil {
		return 0, errors.New("angka kedua tidak valid: " + err.Error())
	}

	if angkaB == 0 {
		// Error untuk kasus khusus: pembagian dengan nol
		return 0, errors.New("pembagi tidak boleh nol")
	}

	// Jika sukses, kembalikan hasil dan nil (tanpa error)
	return angkaA / angkaB, nil
}

// Custom error type - implementasi interface error
type ValidationError struct {
	Field string
	Value string
	Msg   string
}

// Method Error() membuat ValidationError mengimplementasi interface error
func (e *ValidationError) Error() string {
	return fmt.Sprintf("validasi gagal: %s '%s' - %s",
		e.Field, e.Value, e.Msg)
}

// Fungsi registrasi dengan custom error
func register(nama string, umur int) error {
	if nama == "" {
		return &ValidationError{
			Field: "nama",
			Value: nama,
			Msg:   "nama tidak boleh kosong",
		}
	}
	if umur < 17 {
		return &ValidationError{
			Field: "umur",
			Value: fmt.Sprintf("%d", umur),
			Msg:   "minimal umur 17 tahun",
		}
	}
	return nil
}

func main() {
	// Contoh 1: Basic error handling
	hasil, err := bagi("10", "2")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Hasil:", hasil)
	}

	// Contoh 2: Error karena pembagi nol
	_, err = bagi("10", "0")
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Contoh 3: Custom error dengan type assertion
	err = register("", 20)
	if err != nil {
		// Type assertion untuk mengecek tipe error
		var valErr *ValidationError
		if errors.As(err, &valErr) {
			fmt.Printf("Validation Error: field=%s, msg=%s\n",
				valErr.Field, valErr.Msg)
		} else {
			fmt.Println("Error:", err)
		}
	}

	err = register("Anggi", 15)
	if err != nil {
		var valErr *ValidationError
		if errors.As(err, &valErr) {
			fmt.Printf("Validation Error: field=%s, msg=%s\n",
				valErr.Field, valErr.Msg)
		}
	}

	// Contoh 4: Panic (error fatal) dan Recover
	// defer func() {
	//     if r := recover(); r != nil {
	//         fmt.Println("Recovered:", r)
	//     }
	// }()
	// panic("Terjadi masalah fatal!")
}