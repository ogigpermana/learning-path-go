// File: middle/project-layout/main.go
// Level: Middle
// Topik: Standard Go Project Layout
//
// Go tidak memaksa struktur proyek tertentu, tapi ada convention.
// Standard Go Project Layout: https://github.com/golang-standards/project-layout
//
// Struktur umum:
/*
myproject/
├── cmd/               # Entry points (main.go per aplikasi)
│   └── app/
│       └── main.go
├── internal/          # Private code (tidak bisa di-import dari luar)
│   └── service/
│       └── user.go
├── pkg/               # Public code (bisa di-import)
│   └── api/
│       └── handler.go
├── api/               # API definitions (OpenAPI, protobuf)
├── web/               # Static web files
├── configs/           # Config files
├── scripts/           # Build scripts
├── migrations/        # Database migrations
├── test/              # Integration/E2E tests
├── docs/              # Documentation
├── Makefile           # Build automation
├── Dockerfile
├── go.mod
└── README.md
*/
//
// Untuk proyek kecil (seperti tutorial ini):
// - exercises/   -> satu file per topik (simple)
// - middle/      -> satu folder per topik
// - expert/      -> kompleks, multiple files

package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	fmt.Println("=== Standard Go Project Layout ===")
	fmt.Println()
	fmt.Println("Struktur untuk production Go project:")
	fmt.Println()

	// Simulasi struktur proyek
	// root adalah proyek yang direpresentasikan
	root := "my-go-app/"
	structure := map[string]string{
		"cmd/api/main.go":             "Entry point API server",
		"cmd/worker/main.go":          "Entry point background worker",
		"internal/handler/user.go":    "HTTP handler (private)",
		"internal/middleware/auth.go": "Auth middleware (private)",
		"internal/repository/user.go": "Database access (private)",
		"internal/service/user.go":    "Business logic (private)",
		"pkg/api/response.go":         "Shared API utilities",
		"pkg/validator/email.go":      "Validation helpers",
		"api/openapi.yaml":            "OpenAPI specification",
		"configs/config.yaml":         "Config file",
		"migrations/001_init.sql":     "Database migration",
		"web/templates/index.html":    "HTML templates",
		"scripts/build.sh":            "Build script",
		"test/integration/user_test.go": "Integration test",
		".github/workflows/ci.yml":    "CI pipeline",
		"Dockerfile":                  "Docker build",
		"Makefile":                    "Build automation",
		"README.md":                   "Documentation",
	}

	for path, desc := range structure {
		// Buat path lengkap
		fullPath := filepath.Join(root, path)
		dir := filepath.Dir(fullPath)

		// Tampilkan dengan indentasi
		indent := ""
		for i := 0; i < len(dir)-len(root); i++ {
			if dir[i] == '/' || dir[i] == '\\' {
				indent += "  "
			}
		}

		fmt.Printf("%s%s/  %s\n", indent, filepath.Base(dir), desc)
	}

	fmt.Println()
	fmt.Println("=== Rules ===")
	fmt.Println("1. cmd/: satu folder per binary/executable")
	fmt.Println("2. internal/: private code (tidak bisa di-import pihak luar)")
	fmt.Println("3. pkg/: public library code")
	fmt.Println("4. jangan gunakan src/ atau lib/ (bukan Go style)")
	fmt.Println("5. hindari circular dependencies")
	fmt.Println()
	fmt.Println("=== Project size guide ===")
	fmt.Println("- Small (<500 lines): flat structure (seperti exercises/)")
	fmt.Println("- Medium (500-5000): cmd/ + internal/ + pkg/")
	fmt.Println("- Large (>5000): tambah domain/subdomain folders")

	fmt.Println("\nCurrent working directory:", getCWD())
}

func getCWD() string {
	cwd, _ := os.Getwd()
	return cwd
}

/*
Tips:
1. Mulai flat dulu, pisah kalau perlu
2. internal/ mencegah import cycle
3. cmd/ multiple binaries dari satu repo
4. Jangan over-engineer untuk proyek kecil
*/