// File: middle/logging/main.go
// Level: Middle
// Topik: Logging
//
// Package "log" adalah logging bawaan Go.
// Level: Print (info), Fatal (exit), Panic (panic)
// Format: log.SetFlags() untuk konfigurasi format
// Logger bisa di-custom: prefix, output destination, format
//
// Untuk production, gunakan library seperti logrus atau zap:
// go get github.com/sirupsen/logrus
// go get go.uber.org/zap

package main

import (
	"io"       // io.MultiWriter
	"log"      // standard logger
	"os"       // file operations
	"time"     // time formatting
)

func main() {
	// 1. BASIC LOG
	fmt := log.New(os.Stdout, "", 0) // hanya untuk fmt.Println-style
	fmt.Println("=== Basic Log ===")
	log.Println("Ini log biasa")
	log.Printf("User %s login pada %s\n", "Anggi", time.Now().Format(time.RFC3339))

	// 2. LOG DENGAN FORMAT
	fmt.Println("\n=== Log Format ===")
	// log.Ldate: tanggal, Ltime: waktu, Lshortfile: file & line number
	warnLog := log.New(os.Stdout,
		"WARN: ",
		log.Ldate|log.Ltime|log.Lshortfile,
	)
	warnLog.Println("Ini warning dengan lokasi file")
	warnLog.Println("Format: tanggal + waktu + file:line")

	// 3. LOG KE FILE
	fmt.Println("\n=== Log to File ===")
	file, err := os.Create("app.log")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	defer os.Remove("app.log")

	fileLog := log.New(file, "INFO: ", log.LstdFlags)
	fileLog.Println("Aplikasi dimulai")
	fileLog.Println("Aplikasi berjalan")
	fileLog.Println("Aplikasi selesai")

	// 4. MULTI WRITER - log ke stdout dan file sekaligus
	fmt.Println("\n=== Multi Writer ===")
	multiWriter := io.MultiWriter(os.Stdout, file)
	multiLog := log.New(multiWriter, "MULTI: ", log.LstdFlags)
	multiLog.Println("Log ini muncul di stdout DAN file")
	multiLog.Println("Berguna untuk development sekaligus logging ke file")

	// 5. FATAL & PANIC (HATI-HATI!)
	// log.Fatal: mencetak log lalu exit(1)
	// log.Panic: mencetak log lalu panic
	// Uncomment untuk melihat efek:
	// log.Fatal("Fatal error - program akan exit!")
	// log.Panic("Panic error - program akan panic!")
}