// File: middle/database/main.go
// Level: Middle
// Topik: Database SQL dengan SQLite
//
// Go memiliki package "database/sql" - interface standar untuk database SQL.
// Dibutuhkan driver spesifik untuk tiap database:
// - SQLite: modernc.org/sqlite (pure Go, no C compiler needed)
// - PostgreSQL: github.com/lib/pq
// - MySQL: github.com/go-sql-driver/mysql
//
// Step install:
// cd middle/database
// go get modernc.org/sqlite

package main

import (
	"database/sql"
	"fmt"
	"log"

	// Import driver SQLite (side effect: menggunakan _)
	_ "modernc.org/sqlite"
)

// User merepresentasikan tabel users
type User struct {
	ID    int
	Name  string
	Email string
}

func main() {
	// 1. KONEKSI DATABASE
	// sql.Open() tidak langsung connect, hanya menyiapkan koneksi
	db, err := sql.Open("sqlite", "test.db")
	if err != nil {
		log.Fatal("Gagal buka database:", err)
	}
	defer db.Close() // tutup koneksi saat fungsi selesai

	// Test koneksi
	err = db.Ping()
	if err != nil {
		log.Fatal("Gagal koneksi:", err)
	}
	fmt.Println("Koneksi database berhasil")

	// Cleanup tabel jika ada
	defer db.Exec("DROP TABLE IF EXISTS users")

	// 2. CREATE TABLE
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		email TEXT UNIQUE NOT NULL
	)`)
	if err != nil {
		log.Fatal("Gagal create table:", err)
	}
	fmt.Println("Tabel users siap")
	fmt.Println()

	// 3. INSERT
	fmt.Println("=== INSERT ===")
	result, err := db.Exec(
		"INSERT INTO users (name, email) VALUES (?, ?)",
		"Anggi", "anggi@mail.com",
	)
	if err != nil {
		log.Fatal("Gagal insert:", err)
	}
	id, _ := result.LastInsertId() // ambil ID yang baru di-generate
	fmt.Println("Inserted user ID:", id)

	// Insert multiple dalam 1 transaction
	fmt.Println("Insert multiple dengan transaction:")
	tx, _ := db.Begin() // mulai transaction
	tx.Exec("INSERT INTO users (name, email) VALUES (?, ?)", "Budi", "budi@mail.com")
	tx.Exec("INSERT INTO users (name, email) VALUES (?, ?)", "Citra", "citra@mail.com")
	tx.Commit() // commit transaction
	fmt.Println("Transaction selesai")
	fmt.Println()

	// 4. QUERY SINGLE ROW
	fmt.Println("=== Query Single Row ===")
	var u User
	// QueryRow mengambil baris pertama
	row := db.QueryRow("SELECT id, name, email FROM users WHERE id = ?", 1)
	// Scan membaca hasil query ke variabel (urutan harus sesuai SELECT)
	err = row.Scan(&u.ID, &u.Name, &u.Email)
	if err != nil {
		log.Fatal("Gagal query:", err)
	}
	fmt.Printf("User: id=%d, name=%s, email=%s\n", u.ID, u.Name, u.Email)
	fmt.Println()

	// 5. QUERY MULTIPLE ROWS
	fmt.Println("=== Query Multiple Rows ===")
	rows, err := db.Query("SELECT id, name, email FROM users ORDER BY id")
	if err != nil {
		log.Fatal("Gagal query:", err)
	}
	defer rows.Close() // tutup rows saat selesai

	// Iterasi hasil query
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			log.Fatal("Gagal scan:", err)
		}
		fmt.Printf("  id=%d, name=%s, email=%s\n", user.ID, user.Name, user.Email)
	}
	// Cek error setelah iterasi
	if err = rows.Err(); err != nil {
		log.Fatal("Error rows:", err)
	}
	fmt.Println()

	// 6. UPDATE
	fmt.Println("=== UPDATE ===")
	db.Exec("UPDATE users SET name = ? WHERE id = ?", "Anggi Updated", 1)
	fmt.Println("User 1 updated")

	// 7. DELETE
	fmt.Println("\n=== DELETE ===")
	db.Exec("DELETE FROM users WHERE id = ?", 3)
	fmt.Println("User 3 deleted")

	// 8. COUNT
	fmt.Println("\n=== COUNT ===")
	var count int
	db.QueryRow("SELECT COUNT(*) FROM users").Scan(&count)
	fmt.Println("Total users tersisa:", count)
}