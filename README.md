# Belajar Golang dari Nol

Panduan lengkap belajar Go (Golang) dari pemula hingga mahir.

## 📋 Struktur Tutorial

```
learning-path-go/
│
├── exercises/          # Level PEMULA (18 topik)
│   ├── 01-hello.go     - Hello World
│   ├── 02-variabel.go  - Variabel & Tipe Data
│   ├── 03-fungsi.go    - Fungsi
│   ├── 04-slice.go     - Array & Slice
│   ├── 05-struct.go    - Struct
│   ├── 06-struct-method.go - Struct & Method
│   ├── 07-interface.go - Interface
│   ├── 08-polymorphism.go  - Polymorphism
│   ├── 09-goroutine.go - Goroutine & Channel
│   ├── 10-error.go     - Error Handling
│   ├── 11-testing.go   - Testing
│   ├── 12-api.go       - REST API sederhana
│   ├── 13-pointer.go   - Pointer
│   ├── 14-map.go       - Map
│   ├── 15-defer-panic-recover.go - Defer/Panic/Recover
│   ├── 16-loop-switch.go  - Loop & Switch
│   ├── 17-string.go    - Manipulasi String
│   └── 18-module/      - Go Modules
│
├── middle/             # Level MENENGAH (11 topik)
│   ├── context/        - Context (timeout, cancel, value)
│   ├── json/           - JSON encoding/decoding
│   ├── file-io/        - File & directory operations
│   ├── concurrency/    - WaitGroup, Mutex, Select
│   ├── testing/        - Table-driven test, benchmark
│   ├── database/       - SQLite CRUD (go get modernc.org/sqlite)
│   ├── logging/        - Standard log, multi-writer
│   ├── middleware/     - HTTP middleware pattern
│   ├── env/            - Environment variables config
│   ├── auth/           - JWT authentication
│   └── docker/         - Docker multi-stage build
│
├── expert/             # Level MAHIR (7 topik)
│   ├── grpc/           - gRPC & Protocol Buffers
│   ├── kubernetes/     - K8s deployment & HPA
│   ├── cicd/           - GitHub Actions pipeline
│   ├── profiling/      - Pprof, race detector, patterns
│   ├── design-patterns/ - Repository, Factory, Builder, Singleton
│   ├── websocket/      - Real-time WebSocket (go get gorilla/websocket)
│   └── graphql/        - GraphQL API (go get graphql-go/graphql)
│
└── fullstack/          # FULLSTACK PROJECT
    ├── api/            - Go REST API (CRUD + JSON file storage)
    └── frontend/       - React frontend (CDN atau Vite)
```

## 🚀 Cara Belajar

### 1. Level Pemula
```bash
# Jalankan setiap file secara berurutan
cd exercises
go run 01-hello.go
go run 02-variabel.go
# ... lanjutkan sampai 18-module
```

### 2. Level Menengah
```bash
cd middle
# Topik 1: Context
cd context && go run main.go
# Topik 2: JSON
cd ../json && go run main.go
# ... dan seterusnya
```

### 3. Level Mahir
```bash
cd expert
# Desain Patterns
cd design-patterns && go run main.go
# WebSocket (install dulu: go get github.com/gorilla/websocket)
cd ../websocket && go run main.go
# GraphQL
cd ../graphql && go run main.go
```

### 4. Fullstack Project
```bash
cd fullstack

# Terminal 1: Backend
cd api && go run main.go

# Terminal 2: Frontend (buka di browser)
cd frontend
# Buka index.html langsung di browser
# Atau jalankan: python3 -m http.server 3000
```

## 🔧 Quick Reference

### Commands Dasar
```bash
go run main.go              # Jalankan Go program
go build -o app .           # Build binary
go test ./... -v -cover     # Test dengan coverage
go mod init example.com/app # Init module
go get github.com/package  # Install dependency
go mod tidy                 # Bersihkan dependencies
go fmt ./...                # Format kode
```

### Debug & Profiling
```bash
go run -race main.go        # Race detector
go test -bench=. -benchmem # Benchmark
go tool pprof cpu.out       # Profiling
```

### Build untuk Production
```bash
GOOS=linux GOARCH=amd64 go build -o app .
# Output: binary Linux 64-bit (static)
```

## 💡 Tips Belajar
1. **Praktek, bukan hanya baca**: Jalankan setiap contoh
2. **Modifikasi kode**: Ubah nilai, tambah fungsi baru
3. **Cari tahu error**: Error adalah guru terbaik
4. **Baca dokumentasi**: [go.dev/doc](https://go.dev/doc/)
5. **Konsisten**: Sedikit setiap hari lebih baik dari banyak sekali

## 📚 Referensi
- [Go by Example](https://gobyexample.com/) - Belajar dari contoh
- [Effective Go](https://go.dev/doc/effective_go) - Best practices
- [Go Wiki](https://github.com/golang/go/wiki) - Komunitas
- [Go Playground](https://go.dev/play/) - Coba Go di browser

## 🐳 Docker
```bash
cd middle/docker
docker build -t go-app .
docker run -p 8080:8080 go-app
docker compose up -d        # Dengan database
```

## 📄 License
MIT - Belajar, gunakan, dan bagikan dengan bebas.

---
**Happy Coding! 🚀**