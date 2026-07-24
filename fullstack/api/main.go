// File: fullstack/api/main.go
// Level: Middle (Fullstack)
// Topik: CRUD REST API dengan JSON File Storage
//
// REST API untuk Todo CRUD dengan data disimpan di JSON file.
// Fitur: GET, POST, PUT, DELETE dengan persistensi ke file.
// CORS enabled untuk frontend React.
//
// Endpoints:
// GET    /todos      -> List semua todos
// POST   /todos      -> Tambah todo baru (body: {"title":"...","done":false})
// PUT    /todo?id=1  -> Update todo (body: {"title":"...","done":true})
// DELETE /todo?id=1  -> Hapus todo

package main

import (
	"encoding/json" // untuk encode/decode JSON
	"log"           // untuk logging
	"net/http"      // HTTP server
	"os"            // file I/O
	"strconv"       // konversi string ke int
	"sync"          // mutual exclusion untuk concurrent access
)

// Todo merepresentasikan item todo
type Todo struct {
	ID    int    `json:"id"`    // ID unik
	Title string `json:"title"` // Judul todo
	Done  bool   `json:"done"`  // Status selesai/belum
}

// Global state
var (
	todos    []Todo    // slice untuk menyimpan todos di memori
	mu       sync.Mutex // mutex untuk thread safety
	nextID   = 1       // auto-increment ID
	dataFile = "todos.json" // file untuk persistensi data
)

// loadData membaca data dari JSON file ke memory
func loadData() {
	file, err := os.Open(dataFile)
	if err != nil {
		// File belum ada (pertama kali jalan)
		return
	}
	defer file.Close()

	json.NewDecoder(file).Decode(&todos)

	// Set nextID berdasarkan ID terakhir di file
	for _, t := range todos {
		if t.ID >= nextID {
			nextID = t.ID + 1
		}
	}
}

// saveData menulis data dari memory ke JSON file
func saveData() {
	file, err := os.Create(dataFile)
	if err != nil {
		log.Printf("Error saving data: %v", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // pretty print JSON
	encoder.Encode(todos)
}

// handleTodos menangani GET /todos dan POST /todos
func handleTodos(w http.ResponseWriter, r *http.Request) {
	// Set response headers
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*") // CORS
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// Handle CORS preflight
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	// Routing berdasarkan HTTP method
	switch r.Method {
	case http.MethodGet:
		// GET /todos - Kembalikan semua todos
		mu.Lock()
		defer mu.Unlock()
		json.NewEncoder(w).Encode(todos)

	case http.MethodPost:
		// POST /todos - Tambah todo baru
		var todo Todo
		err := json.NewDecoder(r.Body).Decode(&todo)
		if err != nil {
			http.Error(w, "Invalid JSON body", http.StatusBadRequest)
			return
		}

		mu.Lock()
		todo.ID = nextID // generate ID baru
		nextID++
		todos = append(todos, todo)
		mu.Unlock()

		saveData() // persist ke file

		w.WriteHeader(http.StatusCreated) // HTTP 201
		json.NewEncoder(w).Encode(todo)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// handleTodo menangani PUT /todo?id=1 dan DELETE /todo?id=1
func handleTodo(w http.ResponseWriter, r *http.Request) {
	// Set response headers
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	// Parse ID dari query parameter
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, `{"error":"Invalid ID"}`, http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	// Cari todo dengan ID yang sesuai
	for i, todo := range todos {
		if todo.ID == id {
			switch r.Method {
			case http.MethodPut:
				// PUT /todo?id=1 - Update todo
				var updated Todo
				err := json.NewDecoder(r.Body).Decode(&updated)
				if err != nil {
					http.Error(w, "Invalid JSON body", http.StatusBadRequest)
					return
				}
				todos[i].Title = updated.Title
				todos[i].Done = updated.Done
				saveData()
				json.NewEncoder(w).Encode(todos[i])

			case http.MethodDelete:
				// DELETE /todo?id=1 - Hapus todo
				todos = append(todos[:i], todos[i+1:]...)
				saveData()
				w.WriteHeader(http.StatusNoContent) // HTTP 204

			default:
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			}
			return // todo ditemukan, selesai
		}
	}

	// Todo tidak ditemukan
	http.Error(w, `{"error":"Todo not found"}`, http.StatusNotFound)
}

func main() {
	// Load data dari file saat startup
	loadData()
	log.Printf("Loaded %d todos from %s", len(todos), dataFile)

	// Register routes
	http.HandleFunc("/todos", handleTodos)  // collection route
	http.HandleFunc("/todo", handleTodo)    // single item route

	// Start server
	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

/*
Cara menjalankan:
cd fullstack/api
go run main.go

Coba endpoints:
# List semua todos
curl localhost:8080/todos

# Tambah todo
curl -X POST -H "Content-Type: application/json" \
  -d '{"title":"Belajar Go","done":false}' \
  localhost:8080/todos

# Update todo
curl -X PUT -H "Content-Type: application/json" \
  -d '{"title":"Belajar Go","done":true}' \
  "localhost:8080/todo?id=1"

# Hapus todo
curl -X DELETE "localhost:8080/todo?id=1"
*/