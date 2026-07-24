// File: expert/web-framework/main.go
// Level: Expert
// Topik: Web Framework (Gin)
//
// Go memiliki banyak web framework populer:
// - Gin: fast, lightweight, most popular
// - Echo: similar to Gin, more features
// - Fiber: Express.js-like, fastest
// - Chi: lightweight, idiomatic
//
// Framework mempermudah: routing, middleware, parameter binding, validation.
//
// Install: go get github.com/gin-gonic/gin

package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Model
type Todo struct {
	ID        int       `json:"id"`
	Title     string    `json:"title" binding:"required"`
	Done      bool      `json:"done"`
	CreatedAt time.Time `json:"created_at"`
}

// Database in-memory
var todos = []Todo{
	{ID: 1, Title: "Belajar Go", Done: true, CreatedAt: time.Now()},
	{ID: 2, Title: "Buat REST API", Done: false, CreatedAt: time.Now()},
}

var nextID = 3

func main() {
	r := gin.Default() // dengan Logger & Recovery middleware

	// ===== ROUTES =====

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok", "time": time.Now()})
	})

	// GET /todos - list all
	r.GET("/todos", func(c *gin.Context) {
		c.JSON(200, todos)
	})

	// GET /todos/:id - get by ID
	r.GET("/todos/:id", func(c *gin.Context) {
		id := c.Param("id")
		for _, todo := range todos {
			if fmt.Sprintf("%d", todo.ID) == id {
				c.JSON(200, todo)
				return
			}
		}
		c.JSON(404, gin.H{"error": "Todo not found"})
	})

	// POST /todos - create
	r.POST("/todos", func(c *gin.Context) {
		var todo Todo
		if err := c.ShouldBindJSON(&todo); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		todo.ID = nextID
		nextID++
		todo.CreatedAt = time.Now()
		todos = append(todos, todo)
		c.JSON(201, todo)
	})

	// PUT /todos/:id - update
	r.PUT("/todos/:id", func(c *gin.Context) {
		id := c.Param("id")
		var updated Todo
		if err := c.ShouldBindJSON(&updated); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		for i, todo := range todos {
			if fmt.Sprintf("%d", todo.ID) == id {
				todos[i].Title = updated.Title
				todos[i].Done = updated.Done
				c.JSON(200, todos[i])
				return
			}
		}
		c.JSON(404, gin.H{"error": "Todo not found"})
	})

	// DELETE /todos/:id - delete
	r.DELETE("/todos/:id", func(c *gin.Context) {
		id := c.Param("id")
		for i, todo := range todos {
			if fmt.Sprintf("%d", todo.ID) == id {
				todos = append(todos[:i], todos[i+1:]...)
				c.JSON(204, nil)
				return
			}
		}
		c.JSON(404, gin.H{"error": "Todo not found"})
	})

	// ===== MIDDLEWARE DEMO =====
	// Custom middleware
	r.Use(func(c *gin.Context) {
		start := time.Now()
		c.Next() // lanjut ke handler
		fmt.Printf("[Gin] %s %s %v\n",
			c.Request.Method, c.Request.URL.Path, time.Since(start))
	})

	// Group routes dengan prefix /api
	api := r.Group("/api")
	{
		api.GET("/stats", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"total_todos": len(todos),
				"done_count":  countDone(),
			})
		})
	}

	// Run server
	r.Run(":8080")
}

func countDone() int {
	count := 0
	for _, t := range todos {
		if t.Done {
			count++
		}
	}
	return count
}

/*
Coba:
go run main.go

# List
curl localhost:8080/todos

# Create
curl -X POST -H "Content-Type: application/json" \
  -d '{"title":"Belajar Gin"}' \
  localhost:8080/todos

# Get by ID
curl localhost:8080/todos/1

# Update
curl -X PUT -H "Content-Type: application/json" \
  -d '{"title":"Belajar Gin","done":true}' \
  localhost:8080/todos/3

# Delete
curl -X DELETE localhost:8080/todos/3

Alternatif framework:
- Echo: go get github.com/labstack/echo/v4
- Fiber: go get github.com/gofiber/fiber/v2
- Chi: go get github.com/go-chi/chi/v5
*/