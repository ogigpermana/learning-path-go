// File: middle/http-client/main.go
// Level: Middle
// Topik: HTTP Client
//
// Package net/http juga menyediakan HTTP client.
// http.Client dengan timeout, transport, dan best practices.
//
// Penting: selalu set Timeout untuk mencegah goroutine leak!

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

func main() {
	// 1. HTTP CLIENT DENGAN TIMEOUT
	fmt.Println("=== HTTP Client dengan Timeout ===")
	client := &http.Client{
		Timeout: 10 * time.Second, // timeout total (termasuk koneksi, redirect, baca body)
	}

	// GET request
	resp, err := client.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		log.Fatal("GET error:", err)
	}
	defer resp.Body.Close()

	// Baca response body
	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("Status: %s\n", resp.Status)
	fmt.Printf("Body: %s\n", body)
	fmt.Println()

	// 2. CUSTOM TRANSPORT (lebih detail)
	fmt.Println("=== Custom Transport ===")
	transport := &http.Transport{
		MaxIdleConns:        10,              // max koneksi idle
		IdleConnTimeout:     30 * time.Second, // timeout koneksi idle
		DisableCompression:  false,            // izinkan compression
	}

	customClient := &http.Client{
		Transport: transport,
		Timeout:   5 * time.Second,
		// CheckRedirect: policy redirect (nil = follow sampai 10 redirect)
	}

	resp2, _ := customClient.Get("https://api.github.com")
	defer resp2.Body.Close()
	body2, _ := io.ReadAll(resp2.Body)
	fmt.Printf("GitHub API: %s\n", body2[:100])
	fmt.Println()

	// 3. POST REQUEST (JSON body)
	fmt.Println("=== POST Request ===")
	todo := Todo{Title: "Belajar HTTP Client", Done: false}
	jsonData, _ := json.Marshal(todo)

	resp3, err := client.Post(
		"https://jsonplaceholder.typicode.com/todos",
		"application/json",
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		log.Fatal("POST error:", err)
	}
	defer resp3.Body.Close()

	var createdTodo Todo
	json.NewDecoder(resp3.Body).Decode(&createdTodo)
	fmt.Printf("Created: %+v\n", createdTodo)
	fmt.Println()

	// 4. REQUEST DENGAN HEADERS
	fmt.Println("=== Request dengan Headers ===")
	req, _ := http.NewRequest("GET", "https://api.github.com/users/ogigpermana", nil)
	req.Header.Set("User-Agent", "Go-Learning/1.0")
	req.Header.Set("Accept", "application/json")

	resp4, _ := client.Do(req)
	defer resp4.Body.Close()
	body4, _ := io.ReadAll(resp4.Body)
	fmt.Printf("Response: %s\n", body4[:200])
	fmt.Println()

	// 5. REQUEST TIMEOUT
	fmt.Println("=== Timeout Demo ===")
	fastClient := &http.Client{Timeout: 1 * time.Millisecond}
	_, err = fastClient.Get("https://httpbin.org/delay/3") // butuh 3 detik
	if err != nil {
		fmt.Printf("Timeout terjadi seperti yang diharapkan: %v\n", err)
	}
	fmt.Println()

	// 6. CONTEXT DENGAN HTTP REQUEST
	fmt.Println("=== Context dengan HTTP ===")
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	req2, _ := http.NewRequestWithContext(ctx, "GET",
		"https://httpbin.org/delay/2", nil)
	_, err = client.Do(req2)
	if err != nil {
		fmt.Printf("Context timeout: %v\n", err)
	}
}

/*
Best Practices HTTP Client:
1. SELALU set Timeout (default: no timeout = goroutine leak risk)
2. Reuse http.Client untuk multiple requests
3. SELALU close resp.Body (even if you don't read it)
4. Gunakan http.NewRequestWithContext untuk cancellable requests
5. Set User-Agent header untuk API requests
6. Gunakan transport pooling untuk performance
*/