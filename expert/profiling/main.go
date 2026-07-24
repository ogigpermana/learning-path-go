// File: expert/profiling/main.go
// Level: Expert
// Topik: Profiling & Optimization Patterns
//
// Profiling adalah proses mengukur performa aplikasi:
// - CPU usage: fungsi mana yang paling boros CPU
// - Memory usage: berapa banyak memory yang dipakai
// - Goroutine: jumlah goroutine yang berjalan
//
// Go profiling tools:
// 1. pprof: CPU & Memory profiler
// 2. trace: Goroutine tracing
// 3. race detector: Data race detection
//
// Commands:
// go test -bench=. -cpuprofile=cpu.out -memprofile=mem.out
// go tool pprof cpu.out
// go run -race main.go

package main

import (
	"fmt"
	"sync"
	"time"
)

// ===== CONCURRENCY PATTERNS =====

// 1. SEQUENTIAL VS CONCURRENT
// Demonstrasi perbedaan waktu eksekusi

// Post struct
type Post struct {
	ID   int
	Body string
}

// Comment struct
type Comment struct {
	ID      int
	Content string
}

// fetchPost simulasi API call (100ms)
func fetchPost(id int) Post {
	time.Sleep(100 * time.Millisecond)
	return Post{ID: id, Body: "Post content"}
}

// fetchComments simulasi API call (100ms)
func fetchComments(postID int) []Comment {
	time.Sleep(100 * time.Millisecond)
	return []Comment{
		{ID: 1, Content: "Comment 1"},
		{ID: 2, Content: "Comment 2"},
	}
}

// Sequential: post diambil, baru comment (total: ~200ms)
func renderPageSequential(postID int) time.Duration {
	start := time.Now()
	post := fetchPost(postID)
	comments := fetchComments(post.ID)
	elapsed := time.Since(start)
	fmt.Printf("Sequential - %d comments, waktu: %v\n",
		len(comments), elapsed)
	return elapsed
}

// Concurrent: post dan comment diambil bersamaan (total: ~100ms)
func renderPageConcurrent(postID int) time.Duration {
	start := time.Now()

	// Struct untuk menampung hasil goroutine
	type result struct {
		post     Post
		comments []Comment
	}

	// Channel untuk hasil
	ch := make(chan result)

	// Jalankan goroutine
	go func() {
		ch <- result{
			fetchPost(postID),
			fetchComments(postID),
		}
	}()

	res := <-ch
	elapsed := time.Since(start)
	fmt.Printf("Concurrent - %d comments, waktu: %v\n",
		len(res.comments), elapsed)
	return elapsed
}

// 2. FAN-OUT PATTERN
// Membagi pekerjaan ke beberapa worker goroutine
func fanOutExample() {
	fmt.Println("\n=== Fan-Out Pattern ===")
	items := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	start := time.Now()
	results := processFanOut(items, 3) // 3 workers
	fmt.Printf("Hasil: %v\n", results)
	fmt.Printf("Waktu: %v (sequential akan ~%dms)\n",
		time.Since(start), len(items)*10)
}

// processFanOut memproses items dengan multiple workers
func processFanOut(items []int, numWorkers int) []int {
	// Buat channel untuk jobs dan results
	jobs := make(chan int, len(items))
	results := make(chan int, len(items))

	// Start workers
	var wg sync.WaitGroup
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go workerFn(i, jobs, results, &wg)
	}

	// Kirim jobs ke channel
	for _, item := range items {
		jobs <- item
	}
	close(jobs) // tidak ada job lagi

	// Tunggu semua worker selesai
	wg.Wait()
	close(results) // tidak ada result lagi

	// Kumpulkan hasil
	var output []int
	for result := range results {
		output = append(output, result)
	}
	return output
}

// workerFn adalah worker yang memproses job
func workerFn(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		// Simulasi kerja
		time.Sleep(10 * time.Millisecond)
		fmt.Printf("Worker %d memproses: %d\n", id, job)
		results <- job * 2
	}
}

// ===== MAIN =====

func main() {
	// 1. Sequential vs Concurrent
	fmt.Println("=== Sequential vs Concurrent ===")
	seqTime := renderPageSequential(1)
	concTime := renderPageConcurrent(1)
	fmt.Printf("Percepatan: %.fx\n", float64(seqTime)/float64(concTime))

	// 2. Fan-Out Pattern
	fanOutExample()

	fmt.Println("\n=== Tips Optimization ===")
	fmt.Println("1. Gunakan goroutine untuk I/O bound tasks")
	fmt.Println("2. Gunakan channel untuk komunikasi")
	fmt.Println("3. Hindari data race: gunakan mutex atau channel")
	fmt.Println("4. Profiling: go test -bench=.")
	fmt.Println("5. Race detector: go run -race main.go")
	fmt.Println("6. Pprof: go tool pprof cpu.out")
}