// File: middle/embed/main.go
// Level: Middle
// Topik: Embed Files (//go:embed)
//
// //go:embed adalah directive Go (1.16+) untuk embed file ke binary.
// Berguna untuk: embed HTML templates, static files, config, SQL migrations.
//
// Keuntungan:
// - Satu binary deployment (no need to ship files separately)
// - Versioning otomatis (file selalu sesuai binary)
// - Aman (file tidak bisa diubah setelah build)

package main

import (
	"embed"  // import embed package
	"fmt"
	"html/template"
	"net/http"
)

//go:embed templates/*
var templateFS embed.FS // embed seluruh folder templates/

//go:embed static/*
var staticFS embed.FS // embed static files

//go:embed config.yaml
var configYAML string // embed single file as string

//go:embed logo.png
var logoData []byte // embed binary file as []byte

//go:embed version.txt
var version string // embed single file as string

func main() {
	// 1. STRING EMBED
	fmt.Println("=== String Embed ===")
	fmt.Printf("Version: %s", version)
	fmt.Println("Config length:", len(configYAML), "bytes")

	// 2. BYTES EMBED (binary files)
	fmt.Println("\n=== Bytes Embed ===")
	fmt.Printf("Logo size: %d bytes\n", len(logoData))

	// 3. FILESYSTEM EMBED
	fmt.Println("\n=== Filesystem Embed ===")
	entries, _ := templateFS.ReadDir("templates")
	fmt.Println("Templates:")
	for _, entry := range entries {
		fmt.Println(" -", entry.Name())
	}

	// 4. WEB SERVER DENGAN EMBED
	fmt.Println("\n=== Web Server with Embed ===")
	mux := http.NewServeMux()

	// Serve embedded static files
	mux.Handle("/static/", http.FileServer(http.FS(staticFS)))

	// Render embedded template
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmplContent, _ := templateFS.ReadFile("templates/index.html")
		tmpl := template.Must(template.New("index").Parse(string(tmplContent)))
		tmpl.Execute(w, map[string]string{
			"Title": "Go Embed Example",
			"Version": version,
		})
	})

	fmt.Println("Server: http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

/*
Cara build & run:
go run main.go        # development
go build -o app .     # production (semua file ter-embed di binary)
./app                 # run tanpa perlu folder templates/ atau static/

Note: Buat folder templates/ dan static/ dulu dengan file-file yang dibutuhkan.
*/