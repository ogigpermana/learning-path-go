// File: expert/sync-advanced/main.go
// Level: Expert
// Topik: Advanced Sync (sync.Map, sync.Pool, sync.Once, atomic)
//
// Package sync lanjutan untuk concurrent programming yang lebih kompleks.
//
// sync.Map: map yang aman untuk concurrent access (tanpa Mutex manual)
// sync.Pool: object pooling untuk mengurangi GC pressure
// sync.Once: menjalankan fungsi hanya sekali (e.g., singleton initialization)
// sync/atomic: operasi atomik tanpa locking

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	// 1. SYNC.ONCE - jalan sekali
	fmt.Println("=== sync.Once ===")
	var once sync.Once
	var config *AppConfig

	initConfig := func() {
		fmt.Println("Initializing config...")
		config = &AppConfig{Version: "1.0.0"}
	}

	// Panggil 5 kali, tapi initConfig hanya jalan sekali
	for i := 0; i < 5; i++ {
		go func(id int) {
			once.Do(initConfig) // hanya jalan sekali, meski dari banyak goroutine
			fmt.Printf("Goroutine %d: config=%+v\n", id, config)
		}(i)
	}

	time.Sleep(100 * time.Millisecond)
	fmt.Println()

	// 2. SYNC.MAP - concurrent-safe map
	fmt.Println("=== sync.Map ===")
	var sm sync.Map

	// Store
	sm.Store("key1", "value1")
	sm.Store("key2", "value2")

	// Load
	if val, ok := sm.Load("key1"); ok {
		fmt.Println("key1:", val)
	}

	// LoadOrStore
	actual, loaded := sm.LoadOrStore("key1", "newvalue")
	fmt.Printf("LoadOrStore key1: actual=%v, loaded=%v\n", actual, loaded)

	actual, loaded = sm.LoadOrStore("key3", "value3")
	fmt.Printf("LoadOrStore key3: actual=%v, loaded=%v\n", actual, loaded)

	// Delete
	sm.Delete("key2")

	// Range
	sm.Range(func(key, value any) bool {
		fmt.Printf("  %v: %v\n", key, value)
		return true
	})
	fmt.Println()

	// 3. SYNC.POOL - object pooling
	fmt.Println("=== sync.Pool ===")
	var pool = sync.Pool{
		New: func() any {
			fmt.Println("  Membuat buffer baru")
			return make([]byte, 1024)
		},
	}

	// Ambil dari pool
	buf1 := pool.Get().([]byte)
	fmt.Println("  buf1 length:", len(buf1))

	// Kembalikan ke pool
	pool.Put(buf1)

	// Ambil lagi (akan reuse buf1, bukan create baru)
	buf2 := pool.Get().([]byte)
	fmt.Println("  buf2 length:", len(buf2)) // reuse, bukan create baru
	fmt.Println()

	// 4. SYNC/ATOMIC - atomic operations
	fmt.Println("=== sync/atomic ===")
	var counter atomic.Int64
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Add(1) // atomic increment, tanpa Mutex
		}()
	}

	wg.Wait()
	fmt.Printf("Atomic counter: %d (harus 100)\n", counter.Load())

	// CompareAndSwap
	old := counter.Load()
	if counter.CompareAndSwap(old, 999) {
		fmt.Printf("CAS success: %d -> %d\n", old, counter.Load())
	}

	// Store
	counter.Store(0)
	fmt.Println("After reset:", counter.Load())
	fmt.Println()

	// 5. PERBANDINGAN: Mutex vs Atomic
	fmt.Println("=== Mutex vs Atomic ===")
	var mu sync.Mutex
	var mutexCounter int64
	var atomicCounter atomic.Int64

	var wg2 sync.WaitGroup
	start := time.Now()
	for i := 0; i < 10000; i++ {
		wg2.Add(1)
		go func() {
			defer wg2.Done()
			mu.Lock()
			mutexCounter++
			mu.Unlock()
		}()
	}
	wg2.Wait()
	fmt.Printf("Mutex: %d, time: %v\n", mutexCounter, time.Since(start))

	start = time.Now()
	for i := 0; i < 10000; i++ {
		wg2.Add(1)
		go func() {
			defer wg2.Done()
			atomicCounter.Add(1)
		}()
	}
	wg2.Wait()
	fmt.Printf("Atomic: %d, time: %v\n", atomicCounter.Load(), time.Since(start))
}

type AppConfig struct {
	Version string
}

/*
Pilih yang mana?
- sync.Mutex: untuk critical section yang kompleks (multiple operations)
- sync/atomic: untuk counter, flag, single variable operations
- sync.Map: untuk map dengan concurrent baca/tulis yang frequent
- sync.Pool: untuk object yang sering create/destroy (e.g., buffer)
- sync.Once: untuk singleton, lazy initialization
- sync.WaitGroup: untuk menunggu goroutine selesai
*/