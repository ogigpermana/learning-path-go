# Belajar Golang dari Nol

Panduan lengkap belajar Go (Golang) dari pemula hingga mahir. **50+ topik** dengan dokumentasi lengkap.

## 📋 Struktur Tutorial

```
learning-path-go/
│
├── exercises/          # Level PEMULA (18 topik)
│   ├── 01-hello.go
│   ├── 02-variabel.go
│   ├── 03-fungsi.go
│   ├── 04-slice.go
│   ├── 05-struct.go
│   ├── 06-struct-method.go
│   ├── 07-interface.go
│   ├── 08-polymorphism.go
│   ├── 09-goroutine.go
│   ├── 10-error.go
│   ├── 11-testing.go
│   ├── 12-api.go
│   ├── 13-pointer.go
│   ├── 14-map.go
│   ├── 15-defer-panic-recover.go
│   ├── 16-loop-switch.go
│   ├── 17-string.go
│   └── 18-module/
│
├── middle/             # Level MENENGAH (24 topik)
│   ├── context/
│   ├── json/
│   ├── file-io/
│   ├── concurrency/
│   ├── testing/
│   ├── database/
│   ├── logging/
│   ├── middleware/
│   ├── env/
│   ├── auth/
│   ├── docker/
│   ├── time/
│   ├── cli/
│   ├── io/
│   ├── template/
│   ├── graceful-shutdown/
│   ├── http-client/
│   ├── embed/
│   ├── regexp/
│   ├── encoding/
│   ├── project-layout/
│   ├── godoc/
│   ├── workspace/
│   └── cross-compile/
│
├── expert/             # Level MAHIR (18 topik)
│   ├── grpc/
│   ├── kubernetes/
│   ├── cicd/
│   ├── profiling/
│   ├── design-patterns/
│   ├── websocket/
│   ├── graphql/
│   ├── reflection/
│   ├── sync-advanced/
│   ├── fuzzing/
│   ├── web-framework/
│   ├── orm/
│   ├── migration/
│   ├── clean-arch/
│   ├── pubsub/
│   ├── event-sourcing/
│   ├── resilience/
│   └── cgo/
│
├── fullstack/          # FULLSTACK PROJECT
│   ├── api/            # Go REST API (CRUD + JSON file storage)
│   └── frontend/       # React frontend
│
└── Belajar-Golang-dari-Nol.pdf  # Buku referensi PDF
```

**Total: 60 topik** (18 beginner + 24 middle + 18 expert)

## 🚀 Cara Belajar

### Level Pemula
```bash
cd exercises
go run 01-hello.go   # Mulai dari sini
go run 02-variabel.go
# ... lanjutkan sampai 18-module
```

### Level Menengah
```bash
cd middle
cd context && go run main.go
cd ../json && go run main.go
# Pilih topik sesuai minat
```

### Level Mahir
```bash
cd expert
cd design-patterns && go run main.go
cd ../pubsub && go run main.go
# Butuh dependensi: go get ...
```

### Fullstack Project
```bash
cd fullstack/api && go run main.go   # Backend :8080
# Buka frontend/index.html di browser
```

## 🔧 Quick Reference

```bash
go run main.go              # Jalankan Go program
go build -o app .           # Build binary
go test ./... -v -cover     # Test dengan coverage
go mod init example.com/app # Init module
go mod tidy                 # Bersihkan dependencies
go fmt ./...                # Format kode
go run -race main.go        # Race detector
go test -bench=.            # Benchmark
go test -fuzz=FuzzXxx       # Fuzzing
GOOS=linux GOARCH=amd64 go build -o app .  # Cross-compile
```

## 🐳 Docker
```bash
cd middle/docker
docker build -t go-app .
docker run -p 8080:8080 go-app
docker compose up -d
```

## 📚 Referensi
- [Go by Example](https://gobyexample.com/)
- [Effective Go](https://go.dev/doc/effective_go)
- [Go Playground](https://go.dev/play/)
- [Standard Go Project Layout](https://github.com/golang-standards/project-layout)

---
**Happy Coding! 🚀**