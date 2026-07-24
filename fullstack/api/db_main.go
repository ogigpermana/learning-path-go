package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

var db *sql.DB

func initDB() {
	var err error
	db, err = sql.Open("sqlite3", "./todos.db")
	if err != nil {
		log.Fatal(err)
	}

	db.Exec(`CREATE TABLE IF NOT EXISTS todos (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		done INTEGER DEFAULT 0
	)`)
}

func handleTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	switch r.Method {
	case http.MethodGet:
		rows, _ := db.Query("SELECT id, title, done FROM todos")
		defer rows.Close()
		var todos []Todo
		for rows.Next() {
			var t Todo
			var doneInt int
			rows.Scan(&t.ID, &t.Title, &doneInt)
			t.Done = doneInt == 1
			todos = append(todos, t)
		}
		json.NewEncoder(w).Encode(todos)

	case http.MethodPost:
		var todo Todo
		json.NewDecoder(r.Body).Decode(&todo)
		result, _ := db.Exec("INSERT INTO todos (title, done) VALUES (?, ?)", todo.Title, todo.Done)
		id, _ := result.LastInsertId()
		todo.ID = int(id)
		json.NewEncoder(w).Encode(todo)
	}
}

func handleTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	id, _ := strconv.Atoi(r.URL.Query().Get("id"))

	switch r.Method {
	case http.MethodPut:
		var todo Todo
		json.NewDecoder(r.Body).Decode(&todo)
		db.Exec("UPDATE todos SET title = ?, done = ? WHERE id = ?", todo.Title, todo.Done, id)
		todo.ID = id
		json.NewEncoder(w).Encode(todo)

	case http.MethodDelete:
		db.Exec("DELETE FROM todos WHERE id = ?", id)
		w.WriteHeader(http.StatusNoContent)
	}
}

func main() {
	initDB()
	http.HandleFunc("/todos", handleTodos)
	http.HandleFunc("/todo", handleTodo)
	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}