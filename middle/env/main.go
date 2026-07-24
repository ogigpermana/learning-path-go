// File: middle/env/main.go
// Level: Middle
// Topik: Environment Variables & Konfigurasi
//
// Environment variables (env vars) adalah cara standard untuk:
// 1. Konfigurasi environment (development/production)
// 2. Menyimpan secrets (API keys, database password)
// 3. Konfigurasi port, database path, dll
//
// Go: os.Getenv() untuk membaca, os.Setenv() untuk menulis
// Untuk production: go get github.com/joho/godotenv

package main

import (
	"fmt"
	"os"
	"strconv"
)

// Config menyimpan semua konfigurasi aplikasi
type Config struct {
	Port    int    // Port HTTP server
	DBPath  string // Path file database
	Env     string // Environment (dev/staging/prod)
	Debug   bool   // Mode debug
}

// getEnv membaca env var dengan fallback value
// Jika variabel tidak ada, gunakan fallback
func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}

// loadConfig membaca semua env var dan mengembalikan Config
func loadConfig() Config {
	port, _ := strconv.Atoi(getEnv("PORT", "8080"))
	debug, _ := strconv.ParseBool(getEnv("DEBUG", "false"))

	return Config{
		Port:   port,
		DBPath: getEnv("DB_PATH", "./data.db"),
		Env:    getEnv("ENV", "development"),
		Debug:  debug,
	}
}

func main() {
	// 1. SET ENV VARS (simulasi)
	// Di production, env var di-set di Docker/docker-compose/CI
	fmt.Println("=== Set Environment Variables ===")
	os.Setenv("PORT", "3000")
	os.Setenv("ENV", "production")
	os.Setenv("DEBUG", "true")
	os.Setenv("DB_PATH", "/app/data/production.db")

	// 2. BACA CONFIG
	fmt.Println("\n=== Load Config ===")
	config := loadConfig()
	fmt.Printf("Config: %+v\n", config)
	fmt.Printf("Server akan mulai di port %d (%s mode)\n",
		config.Port, config.Env)

	if config.Debug {
		fmt.Println("Debug mode: ON - log akan lebih detail")
	}

	// 3. BACA SINGLE ENV
	fmt.Println("\n=== Read Single Env ===")
	home := os.Getenv("HOME")
	user := os.Getenv("USER")
	fmt.Printf("Home: %s\n", home)
	fmt.Printf("User: %s\n", user)

	// 4. LIST ALL ENV VARS
	fmt.Println("\n=== All Env Vars (filtered) ===")
	for _, env := range os.Environ() {
		// Tampilkan hanya env vars yang relevan
		if len(env) > 20 {
			fmt.Println(env[:20], "...")
		}
	}

	// 5. UNSET - membersihkan (untuk production: jangan set lewat kode)
	os.Unsetenv("PORT")
	os.Unsetenv("ENV")
	os.Unsetenv("DEBUG")
	os.Unsetenv("DB_PATH")
	fmt.Println("\nEnv vars cleaned up")
}