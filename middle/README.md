# Tutorial Go Level Menengah

Setelah menguasai dasar-dasar Go, lanjut ke topik menengah untuk membangun aplikasi production-ready.

## Topik

| # | Topik | Deskripsi | File | Cara Run |
|---|-------|-----------|------|----------|
| 1 | **Context** | Timeout, cancellation, passing value antar goroutine | [context/main.go](context/main.go) | `go run context/main.go` |
| 2 | **JSON** | Encoding/decoding JSON, struct tags, file I/O | [json/main.go](json/main.go) | `go run json/main.go` |
| 3 | **File I/O** | Read/write file, scanner, directory operations | [file-io/main.go](file-io/main.go) | `go run file-io/main.go` |
| 4 | **Concurrency** | WaitGroup, Mutex, Select (multiplexing channel) | [concurrency/main.go](concurrency/main.go) | `go run concurrency/main.go` |
| 5 | **Testing** | Table-driven test, benchmark, TestMain, coverage | [testing/main.go](testing/main.go) | `go test -v -bench=. -cover` |
| 6 | **Database** | SQLite CRUD, transaction, prepared statement | [database/main.go](database/main.go) | `go run database/main.go` |
| 7 | **Logging** | Standard log, multi-writer, log ke file | [logging/main.go](logging/main.go) | `go run logging/main.go` |
| 8 | **Middleware** | HTTP middleware: logging, auth, recovery | [middleware/main.go](middleware/main.go) | `go run middleware/main.go` |
| 9 | **Env Config** | Environment variables, config struct, fallback | [env/main.go](env/main.go) | `go run env/main.go` |
| 10 | **JWT Auth** | JWT token create/validate (manual implementasi) | [auth/main.go](auth/main.go) | `go run auth/main.go` |
| 11 | **Docker** | Multi-stage build, Docker Compose, health check | [docker/](docker/) | `docker build -t go-app .` |

## Cara Belajar

1. Baca kode di setiap file (ada komentar lengkap)
2. Jalankan dengan `go run` atau `go test`
3. Modifikasi kode untuk eksperimen
4. Pahami konsep, bukan hanya syntax

## Prasyarat

Sebelum masuk level menengah, pastikan sudah paham:
- ✅ Variabel, fungsi, struct, interface
- ✅ Pointer, slice, map
- ✅ Goroutine & channel dasar
- ✅ Error handling & testing dasar

## Instalasi Dependensi

Beberapa topik butuh dependensi eksternal:

```bash
# Database (SQLite)
cd database && go get modernc.org/sqlite

# WebSocket (expert level)
cd ../websocket && go get github.com/gorilla/websocket

# GraphQL (expert level)
cd ../graphql && go get github.com/graphql-go/graphql
```