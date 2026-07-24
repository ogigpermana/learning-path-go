// File: middle/time/main.go
// Level: Middle
// Topik: Time - Manipulasi Waktu
//
// Package "time" adalah salah satu package paling penting di Go.
// Digunakan untuk: format waktu, parsing, timing, delay, ticker, timer.
//
// Konsep penting:
// - time.Time : representasi waktu
// - time.Duration : durasi (nanosecond)
// - time.Location : zona waktu
// - Layout: Go menggunakan reference time: Mon Jan 2 15:04:05 MST 2006

package main

import (
	"fmt"
	"time"
)

func main() {
	// 1. SEKARANG
	fmt.Println("=== time.Now() ===")
	now := time.Now()
	fmt.Println("Waktu sekarang:", now)
	fmt.Println("Unix timestamp:", now.Unix())
	fmt.Println("UnixNano:", now.UnixNano())

	// 2. FORMAT WAKTU (format menggunakan reference time)
	fmt.Println("\n=== Format Waktu ===")
	// Reference: Mon Jan 2 15:04:05 MST 2006
	// Layout: 01-02-2006 (bulan-tanggal-tahun)
	fmt.Println("Format 1:", now.Format("2006-01-02 15:04:05"))
	fmt.Println("Format 2:", now.Format("02 Jan 2006"))
	fmt.Println("Format 3:", now.Format("Monday, 02 January 2006"))
	fmt.Println("Format 4:", now.Format(time.RFC3339))
	fmt.Println("Format 5:", now.Format(time.Kitchen))

	// 3. PARSE STRING KE TIME
	fmt.Println("\n=== Parse ===")
	parsed, _ := time.Parse("2006-01-02", "2024-12-25")
	fmt.Println("Parsed:", parsed)
	fmt.Println("Hari:", parsed.Weekday())

	// Parse dengan zona waktu
	parsed2, _ := time.Parse(time.RFC3339, "2024-12-25T10:00:00+07:00")
	fmt.Println("Parsed RFC3339:", parsed2)

	// 4. DURATION
	fmt.Println("\n=== Duration ===")
	dur := 5 * time.Second
	fmt.Println("Duration:", dur)
	fmt.Println("In minutes:", dur.Minutes())
	fmt.Println("In milliseconds:", dur.Milliseconds())

	// Sleep
	fmt.Println("Sleep 500ms...")
	start := time.Now()
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Slept for:", time.Since(start))

	// 5. PERBANDINGAN WAKTU
	fmt.Println("\n=== Perbandingan ===")
	t1 := time.Date(2024, 12, 25, 0, 0, 0, 0, time.UTC)
	t2 := time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC)

	fmt.Println("t1:", t1)
	fmt.Println("t2:", t2)
	fmt.Println("t1 == t2:", t1.Equal(t2))
	fmt.Println("t1.Before(t2):", t1.Before(t2))
	fmt.Println("t1.After(t2):", t1.After(t2))
	fmt.Println("Selisih:", t2.Sub(t1))
	fmt.Println("Selisih hari:", t2.Sub(t1).Hours()/24)

	// 6. TIMER & TICKER
	fmt.Println("\n=== Timer ===")
	timer := time.NewTimer(1 * time.Second)
	go func() {
		<-timer.C
		fmt.Println("Timer fired!")
	}()
	time.Sleep(1100 * time.Millisecond)

	fmt.Println("\n=== Ticker ===")
	ticker := time.NewTicker(300 * time.Millisecond)
	done := make(chan bool)
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println("Tick:", <-ticker.C)
		}
		ticker.Stop()
		done <- true
	}()
	<-done

	// 7. TIME ZONE
	fmt.Println("\n=== Time Zone ===")
	loc, _ := time.LoadLocation("Asia/Jakarta")
	nowJakarta := now.In(loc)
	fmt.Println("Waktu Jakarta:", nowJakarta)
	fmt.Println("UTC:", now.UTC())

	// 8. STOPWATCH (hitung durasi)
	fmt.Println("\n=== Stopwatch ===")
	start = time.Now()
	time.Sleep(123 * time.Millisecond)
	elapsed := time.Since(start)
	fmt.Printf("Elapsed: %v (%d ms)\n", elapsed, elapsed.Milliseconds())
}