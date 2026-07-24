// File: middle/cli/main.go
// Level: Middle
// Topik: CLI dengan Flag
//
// Package "flag" digunakan untuk membuat command-line interface.
// Fitur: string flag, int flag, bool flag, custom usage, subcommands.
//
// Alternatif untuk production: cobra (github.com/spf13/cobra)

package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// 1. DEFINISI FLAG
	// flag.Tipe(nama, default, deskripsi)
	// Mengembalikan pointer ke value

	// String flag: --name atau -n
	name := flag.String("name", "World", "Nama untuk greeting")

	// Int flag: --age atau -a
	age := flag.Int("age", 0, "Umur (0 = tidak ditampilkan)")

	// Bool flag: --verbose atau -v
	verbose := flag.Bool("verbose", false, "Tampilkan log detail")

	// Float64 flag
	rate := flag.Float64("rate", 1.0, "Rate multiplier")

	// Alternatif: binding ke existing variable
	var message string
	flag.StringVar(&message, "message", "Hello", "Pesan greeting")

	// Custom usage
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Penggunaan: %s [options]\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
	}

	// Parse flag (WAJIB dipanggil setelah definisi flag)
	flag.Parse()

	// 2. AKSES NILAI FLAG
	if *verbose {
		fmt.Println("=== Mode Verbose ===")
		fmt.Printf("Args: %v\n", os.Args)
		fmt.Printf("Parsed flags: name=%s, age=%d, verbose=%t, rate=%.2f\n",
			*name, *age, *verbose, *rate)
		fmt.Println()
	}

	// 3. ARGUMEN POSITIONAL
	// flag.Args() mengembalikan argumen non-flag
	fmt.Printf("%s, %s!\n", message, *name)

	if *age > 0 {
		fmt.Printf("Umur: %d tahun\n", *age)
	}

	if *rate != 1.0 {
		fmt.Printf("Rate: %.2fx\n", *rate)
	}

	// Tampilkan argumen positional
	if args := flag.Args(); len(args) > 0 {
		fmt.Println("\nArgumen lain:", args)
	}

	// 4. SUBCOMMAND PATTERN (sederhana)
	fmt.Println("\n--- Subcommand Demo ---")
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "greet":
			greetCmd := flag.NewFlagSet("greet", flag.ExitOnError)
			greetName := greetCmd.String("name", "User", "Nama")
			greetCmd.Parse(os.Args[2:])
			fmt.Printf("Halo %s!\n", *greetName)

		case "math":
			mathCmd := flag.NewFlagSet("math", flag.ExitOnError)
			x := mathCmd.Int("x", 0, "Angka pertama")
			y := mathCmd.Int("y", 0, "Angka kedua")
			op := mathCmd.String("op", "add", "Operasi (add/sub/mul/div)")
			mathCmd.Parse(os.Args[2:])

			switch *op {
			case "add":
				fmt.Printf("%d + %d = %d\n", *x, *y, *x+*y)
			case "sub":
				fmt.Printf("%d - %d = %d\n", *x, *y, *x-*y)
			case "mul":
				fmt.Printf("%d * %d = %d\n", *x, *y, *x**y)
			case "div":
				if *y != 0 {
					fmt.Printf("%d / %d = %d\n", *x, *y, *x / *y)
				}
			}
		}
	}
}

/*
Coba:
go run main.go --name Anggi --age 20 --verbose
go run main.go --name Budi --rate 2.5
go run main.go --help
go run main.go greet --name Citra
go run main.go math --x 10 --y 5 --op add
*/