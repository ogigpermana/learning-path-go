// File: expert/migration/main.go
// Level: Expert
// Topik: Database Migration
//
// Migration adalah cara versioning schema database.
// Berguna untuk: tracking perubahan skema, rollback, team collaboration.
//
// Alat-alat migration:
// - golang-migrate/migrate: most popular
// - pressly/goose: simple, SQL-based
// - GORM AutoMigrate: automatic (tapi tidak support rollback)
// - bytebase/bytebase: GUI-based (enterprise)
//
// Untuk project sederhana: cukup GORM AutoMigrate atau SQL files.
// Untuk production: gunakan golang-migrate atau goose.

package main

// Contoh SQL migration files
// Simpan di folder migrations/

// 001_create_users.up.sql:
/*
CREATE TABLE users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    age INTEGER DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
*/

// 001_create_users.down.sql:
/*
DROP TABLE IF EXISTS users;
*/

// 002_add_posts.up.sql:
/*
CREATE TABLE posts (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title VARCHAR(200) NOT NULL,
    body TEXT,
    user_id INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id)
);
*/

// 002_add_posts.down.sql:
/*
DROP TABLE IF EXISTS posts;
*/

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// Simple migration runner (demonstration)
// Untuk production, gunakan: github.com/golang-migrate/migrate/v4

type Migration struct {
	Version string
	Name    string
	UpSQL   string
	DownSQL string
}

func main() {
	fmt.Println("=== Migration Example ===")
	fmt.Println()
	fmt.Println("Cara menggunakan golang-migrate:")
	fmt.Println()
	fmt.Println("1. Install CLI:")
	fmt.Println("   go install -tags 'sqlite3' github.com/golang-migrate/migrate/v4/cmd/migrate@latest")
	fmt.Println()
	fmt.Println("2. Buat migration:")
	fmt.Println("   migrate create -ext sql -dir migrations -seq create_users_table")
	fmt.Println()
	fmt.Println("3. Jalankan migration:")
	fmt.Println("   migrate -path migrations -database 'sqlite://app.db' up")
	fmt.Println()
	fmt.Println("4. Rollback:")
	fmt.Println("   migrate -path migrations -database 'sqlite://app.db' down 1")
	fmt.Println()

	// Simulasi manual migration dengan SQL files
	fmt.Println("=== Simulasi Manual Migration ===")
	migrations := loadMigrations("migrations")
	for _, m := range migrations {
		fmt.Printf("Migration %s (%s):\n", m.Version, m.Name)
		fmt.Printf("  UP:   %s\n", truncate(m.UpSQL, 50))
		fmt.Printf("  DOWN: %s\n", truncate(m.DownSQL, 50))
	}

	fmt.Println()
	fmt.Println("=== Best Practices ===")
	fmt.Println("1. Migration harus idempotent (safe dijalankan ulang)")
	fmt.Println("2. Setiap migration punya UP dan DOWN")
	fmt.Println("3. Jangan edit migration yang sudah di-commit")
	fmt.Println("4. Test migration di staging sebelum production")
	fmt.Println("5. Backup database sebelum migration")
	fmt.Println("6. Migration file: YYYYMMDD_description.up.sql")
}

func loadMigrations(dir string) []Migration {
	var migrations []Migration

	entries, err := os.ReadDir(dir)
	if err != nil {
		// Folder belum ada
		return migrations
	}

	files := make(map[string]string)
	for _, entry := range entries {
		files[entry.Name()] = entry.Name()
	}

	// Parse migration files
	seen := make(map[string]bool)
	for name := range files {
		parts := strings.Split(name, ".")
		if len(parts) < 3 {
			continue
		}
		version := parts[0]
		if seen[version] {
			continue
		}
		seen[version] = true

		upPath := filepath.Join(dir, version+".up.sql")
		downPath := filepath.Join(dir, version+".down.sql")

		m := Migration{
			Version: version,
			Name:    parts[1],
		}

		if data, err := os.ReadFile(upPath); err == nil {
			m.UpSQL = string(data)
		}
		if data, err := os.ReadFile(downPath); err == nil {
			m.DownSQL = string(data)
		}

		migrations = append(migrations, m)
	}

	sort.Slice(migrations, func(i, j int) bool {
		return migrations[i].Version < migrations[j].Version
	})

	return migrations
}

func truncate(s string, maxLen int) string {
	if len(s) <= maxLen {
		return strings.ReplaceAll(s, "\n", "\\n")
	}
	return strings.ReplaceAll(s[:maxLen], "\n", "\\n") + "..."
}

/*
Cara install golang-migrate:
go install -tags 'sqlite3' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

Cara install goose:
go install github.com/pressly/goose/v3/cmd/goose@latest

Cara pake goose:
goose sqlite3 app.db create create_users_table sql
goose sqlite3 app.db up
goose sqlite3 app.db down
*/