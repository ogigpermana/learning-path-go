package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Pengguna struct {
	ID   int    `json:"id"`
	Nama string `json:"nama"`
}

var pengguna = []Pengguna{
	{ID: 1, Nama: "Anggi"},
	{ID: 2, Nama: "Budi"},
}

func handlerPengguna(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pengguna)
}

func main() {
	http.HandleFunc("/pengguna", handlerPengguna)
	
	fmt.Println("Server running di http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}