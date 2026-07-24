// File: expert/graphql/main.go
// Level: Expert
// Topik: GraphQL API Server
//
// GraphQL adalah query language untuk API yang dikembangkan oleh Facebook.
// Kelebihan dibanding REST: client bisa meminta field spesifik yang dibutuhkan.
//
// Library: github.com/graphql-go/graphql
//
// Cara install:
// cd expert/graphql
// go get github.com/graphql-go/graphql
// go run main.go

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/graphql-go/graphql"
)

// ===== TYPE DEFINITIONS =====

// Todo adalah model data
type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

// Data dummy (biasanya dari database)
var todos = []Todo{
	{ID: 1, Title: "Belajar Go", Done: true},
	{ID: 2, Title: "Buat REST API", Done: false},
	{ID: 3, Title: "Belajar GraphQL", Done: false},
}

// ===== GRAPHQL SCHEMA DEFINITION =====

// todoType adalah representasi GraphQL dari struct Todo
var todoType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Todo",
	Description: "Todo item dengan ID, title, dan status done",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type:        graphql.Int,
			Description: "ID unik todo",
		},
		"title": &graphql.Field{
			Type:        graphql.String,
			Description: "Judul todo",
		},
		"done": &graphql.Field{
			Type:        graphql.Boolean,
			Description: "Status selesai atau belum",
		},
	},
})

// queryType adalah entry point untuk query
var queryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		// Query: todo(id: 1) { id title done }
		"todo": &graphql.Field{
			Type:        todoType,
			Description: "Mendapatkan todo berdasarkan ID",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type:        graphql.Int,
					Description: "ID dari todo",
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				// Ambil argumen id
				id, _ := p.Args["id"].(int)
				// Cari di data
				for _, todo := range todos {
					if todo.ID == id {
						return todo, nil
					}
				}
				return nil, nil
			},
		},
		// Query: todos { id title done }
		"todos": &graphql.Field{
			Type:        graphql.NewList(todoType),
			Description: "Mendapatkan semua todos",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return todos, nil
			},
		},
	},
})

// Schema adalah GraphQL schema yang sudah dikompilasi
var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query: queryType,
})

// ===== HTTP HANDLER =====

// graphqlHandler menangani semua request GraphQL
func graphqlHandler(w http.ResponseWriter, r *http.Request) {
	// Parse request body
	var params struct {
		Query string `json:"query"`
	}
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Eksekusi query GraphQL
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: params.Query,
	})

	// Set response headers
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Kirim response
	json.NewEncoder(w).Encode(result)
}

func main() {
	http.HandleFunc("/graphql", graphqlHandler)

	log.Println("GraphQL server running on :8080")
	log.Println("Test:")
	log.Println(`curl -X POST -d '{"query":"{ todos { id title done } }"}' localhost:8080/graphql`)
	log.Println(`curl -X POST -d '{"query":"{ todo(id: 1) { id title done } }"}' localhost:8080/graphql`)
	log.Fatal(http.ListenAndServe(":8080", nil))
}