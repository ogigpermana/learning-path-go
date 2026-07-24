// File: middle/file-io/main.go
// Level: Middle
// Topik: File I/O
//
// Go menyediakan beberapa cara untuk membaca/menulis file:
// 1. os.ReadFile / os.WriteFile - untuk file kecil (all at once)
// 2. bufio.Scanner - membaca baris per baris
// 3. os.File - untuk operasi file level rendah
// 4. io.Copy - untuk menyalin data antar stream

package main

import (
	"bufio"   // buffered I/O
	"fmt"
	"os"      // file operations
)

func main() {
	// 1. WRITE FILE - menulis seluruh konten sekaligus
	fmt.Println("=== Write File ===")
	content := "Hello, Go!\nBaris kedua\nBaris ketiga"
	// 0644 = permission: owner r/w, group r, others r
	err := os.WriteFile("output.txt", []byte(content), 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("File 'output.txt' berhasil ditulis")

	// 2. READ FILE - membaca seluruh konten sekaligus
	fmt.Println("\n=== Read File ===")
	data, err := os.ReadFile("output.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println("File content:")
	fmt.Println(string(data))

	// 3. APPEND TO FILE - menambah konten ke file yang sudah ada
	fmt.Println("\n=== Append to File ===")
	f, err := os.OpenFile("output.txt",
		os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.WriteString("Baris tambahan\n")
	fmt.Println("Baris baru berhasil ditambahkan")

	// 4. READ LINE BY LINE - menggunakan scanner
	fmt.Println("\n=== Read Line by Line ===")
	file, err := os.Open("output.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNum := 1
	for scanner.Scan() {
		// scanner.Text() mengembalikan baris saat ini (tanpa \n)
		fmt.Printf("Baris %d: %s\n", lineNum, scanner.Text())
		lineNum++
	}

	// Cek error scanner
	if err := scanner.Err(); err != nil {
		fmt.Println("Scanner error:", err)
	}

	// 5. FILE INFO - mengecek apakah file ada
	fmt.Println("\n=== File Info ===")
	info, err := os.Stat("output.txt")
	if err == nil {
		fmt.Println("File exists")
		fmt.Println("Size:", info.Size(), "bytes")
		fmt.Println("Permission:", info.Mode())
		fmt.Println("Modified:", info.ModTime())
	} else {
		fmt.Println("File tidak ditemukan")
	}

	// 6. DIRECTORY OPERATIONS
	fmt.Println("\n=== Directory Operations ===")
	os.MkdirAll("test/subdir", 0755)
	fmt.Println("Directory 'test/subdir' dibuat")

	// Rename file
	os.Rename("output.txt", "output_renamed.txt")
	fmt.Println("File di-rename")

	// List files in directory
	entries, _ := os.ReadDir(".")
	fmt.Println("Files in current dir:")
	for _, entry := range entries {
		if entry.IsDir() {
			fmt.Println("[DIR]", entry.Name())
		} else {
			fmt.Println("[FILE]", entry.Name())
		}
	}

	// 7. CLEANUP
	fmt.Println("\n=== Cleanup ===")
	os.Remove("output_renamed.txt")
	os.RemoveAll("test")
	fmt.Println("Cleanup selesai")
}