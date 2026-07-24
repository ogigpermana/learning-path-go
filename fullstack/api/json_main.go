package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
)

type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

var (
	todos []Todo
	mu    sync.Mutex
	nextID = 1
	dataFile = "todos.json"
)

func loadData() {
	file, err := os.Open(dataFile)
	if err != nil {
		return
	}
	defer file.Close()
	json.NewDecoder(file).Decode(&todos)
	for _, t := range todos {
		if t.ID >= nextID {
			nextID = t.ID + 1
		}
	}
}

func saveData() {
	file, _ := os.Create(dataFile)
	defer file.Close()
	json.NewEncoder(file).Encode(todos)
}

func handleTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	switch r.Method {
	case http.MethodGet:
		mu.Lock()
		defer mu.Unlock()
		json.NewEncoder(w).Encode(todos)

	case http.MethodPost:
		var todo Todo
		json.NewDecoder(r.Body).Decode(&todo)
		mu.Lock()
		todo.ID = nextID
		nextID++
		todos = append(todos, todo)
		mu.Unlock()
		saveData()
		json.NewEncoder(w).Encode(todo)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handleTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	for i, todo := range todos {
		if todo.ID == id {
			switch r.Method {
			case http.MethodPut:
				var updated Todo
				json.NewDecoder(r.Body).Decode(&updated)
				todos[i].Title = updated.Title
				todos[i].Done = updated.Done
				saveData()
				json.NewEncoder(w).Encode(todos[i])

			case http.MethodDelete:
				todos = append(todos[:i], todos[i+1:]...)
				saveData()
				w.WriteHeader(http.StatusNoContent)

			default:
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			}
			return
		}
	}
	http.Error(w, "Todo not found", http.StatusNotFound)
}

func main() {
	loadData()
	http.HandleFunc("/todos", handleTodos)
	http.HandleFunc("/todo", handleTodo)
	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}