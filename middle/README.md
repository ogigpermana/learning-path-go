# Tutorial Go Level Menengah

Setelah menguasai dasar-dasar Go, lanjut ke topik menengah untuk membangun aplikasi production-ready.

## Topik

| # | Topik | Deskripsi | File |
|---|-------|-----------|------|
| 1 | **Context** | Timeout, cancellation, passing value antar goroutine | [context/main.go](context/main.go) |
| 2 | **JSON** | Encoding/decoding JSON, struct tags, file I/O | [json/main.go](json/main.go) |
| 3 | **File I/O** | Read/write file, scanner, directory operations | [file-io/main.go](file-io/main.go) |
| 4 | **Concurrency** | WaitGroup, Mutex, Select (multiplexing channel) | [concurrency/main.go](concurrency/main.go) |
| 5 | **Testing** | Table-driven test, benchmark, TestMain, coverage | [testing/main.go](testing/main.go) |
| 6 | **Database** | SQLite CRUD, transaction, prepared statement | [database/main.go](database/main.go) |
| 7 | **Logging** | Standard log, multi-writer, log ke file | [logging/main.go](logging/main.go) |
| 8 | **Middleware** | HTTP middleware: logging, auth, recovery | [middleware/main.go](middleware/main.go) |
| 9 | **Env Config** | Environment variables, config struct, fallback | [env/main.go](env/main.go) |
| 10 | **JWT Auth** | JWT token create/validate (manual implementasi) | [auth/main.go](auth/main.go) |
| 11 | **Docker** | Multi-stage build, Docker Compose, health check | [docker/](docker/) |
| 12 | **Time** | Manipulasi waktu, format, parse, timer, ticker | [time/main.go](time/main.go) |
| 13 | **CLI/Flag** | Command-line flag parsing, subcommands | [cli/main.go](cli/main.go) |
| 14 | **io.Reader/Writer** | Interface I/O fundamental Go | [io/main.go](io/main.go) |
| 15 | **Templates** | html/template & text/template | [template/main.go](template/main.go) |
| 16 | **Graceful Shutdown** | Signal handling, server shutdown | [graceful-shutdown/main.go](graceful-shutdown/main.go) |
| 17 | **HTTP Client** | HTTP client best practices, timeout, context | [http-client/main.go](http-client/main.go) |
| 18 | **Embed** | //go:embed directive | [embed/main.go](embed/main.go) |
| 19 | **Regexp** | Regular expression patterns | [regexp/main.go](regexp/main.go) |
| 20 | **Encoding** | CSV, XML, Base64, Hex | [encoding/main.go](encoding/main.go) |
| 21 | **Project Layout** | Standard Go project structure | [project-layout/main.go](project-layout/main.go) |
| 22 | **Godoc** | Dokumentasi kode dengan godoc | [godoc/main.go](godoc/main.go) |
| 23 | **Workspace** | Go workspace (go.work) & vendor | [workspace/main.go](workspace/main.go) |
| 24 | **Cross Compile** | Build untuk berbagai platform | [cross-compile/main.go](cross-compile/main.go) |

## Cara Belajar

1. Baca kode di setiap file (ada komentar lengkap)
2. Jalankan dengan `go run` atau `go test`
3. Modifikasi kode untuk eksperimen
4. Pahami konsep, bukan hanya syntax

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