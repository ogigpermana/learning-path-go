// File: middle/workspace/main.go
// Level: Middle
// Topik: Go Workspace (go.work)
//
// Go workspace (1.18+) mengelola multiple modules dalam satu repositori.
// Berguna untuk: monorepo, development multiple module paralel.
//
// Sebelumnya: menggunakan replace directive di go.mod
// Workspace: file go.work di root

// go.work file:
/*
go 1.22

use (
    ./api
    ./shared
    ./cmd
)

replace (
    example.com/legacy => ../legacy
)
*/

package main

import (
	"fmt"
	"os/exec"
)

func main() {
	fmt.Println("=== Go Workspace (go.work) ===")
	fmt.Println()
	fmt.Println("Cara menggunakan workspace:")
	fmt.Println()
	fmt.Println("1. Inisialisasi workspace di root proyek:")
	fmt.Println("   go work init ./module1 ./module2")
	fmt.Println()
	fmt.Println("2. Tambah module baru:")
	fmt.Println("   go work use ./module3")
	fmt.Println()
	fmt.Println("3. Hapus module:")
	fmt.Println("   go work use -drop ./module3")
	fmt.Println()
	fmt.Println("4. Lihat struktur:")
	cmd := exec.Command("go", "work", "edit", "--json")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("   (belum ada go.work file)")
	} else {
		fmt.Println(string(output))
	}
	fmt.Println()
	fmt.Println("=== Vendor Mode ===")
	fmt.Println()
	fmt.Println("Vendor: copy dependencies ke folder vendor/")
	fmt.Println("   go mod vendor")
	fmt.Println("   go build -mod=vendor")
	fmt.Println()
	fmt.Println("Kapan pakai vendor:")
	fmt.Println("- Air-gapped environment (no internet)")
	fmt.Println("- Reproducible build (pinned versions)")
	fmt.Println("- Review dependency code")
	fmt.Println()
	fmt.Println("=== Best Practices ===")
	fmt.Println("1. go.work hanya untuk DEVELOPMENT (jangan commit)")
	fmt.Println("2. Di CI: gunakan replace directive di go.mod")
	fmt.Println("3. go.work di .gitignore")
	fmt.Println("4. go mod tidy untuk bersihkan dependensi")
	fmt.Println("5. go mod verify untuk cek integritas")

	fmt.Println()
	fmt.Println("Contoh struktur monorepo:")
	fmt.Println(`
myproject/
├── go.work             # workspace (development only)
├── .gitignore          # go.work di sini
├── api/
│   ├── go.mod
│   └── main.go
├── shared/
│   ├── go.mod
│   └── pkg.go
└── cmd/
    ├── go.mod
    └── main.go
`)
}

/*
Commands Reference:
go work init [modules...]   # Init workspace
go work use [module]        # Add module
go work use -drop [module]  # Remove module
go work sync                # Sync dependencies
go mod vendor               # Create vendor dir
go mod tidy                 # Clean dependencies
go mod verify               # Verify dependencies
*/