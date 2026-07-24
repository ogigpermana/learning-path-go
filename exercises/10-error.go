package main

import (
	"errors"
	"fmt"
	"strconv"
)

func bagi(a, b string) (float64, error) {
	angkaA, err := strconv.ParseFloat(a, 64)
	if err != nil {
		return 0, errors.New(" angka pertama tidak valid")
	}
	
	angkaB, err := strconv.ParseFloat(b, 64)
	if err != nil {
		return 0, errors.New(" angka kedua tidak valid")
	}
	
	if angkaB == 0 {
		return 0, errors.New(" pembagi tidak boleh nol")
	}
	
	return angkaA / angkaB, nil
}

func main() {
	hasil, err := bagi("10", "2")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Hasil:", hasil)
	}
	
	_, err = bagi("10", "0")
	if err != nil {
		fmt.Println("Error:", err)
	}
}