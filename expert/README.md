# Tutorial Go Level Mahir

Level mahir mencakup topik-topik production-grade: arsitektur microservices, deployment, CI/CD, dan optimisasi.

## Topik Utama

| # | Topik | Deskripsi | File |
|---|-------|-----------|------|
| 1 | **gRPC** | High-performance RPC dengan Protocol Buffers | [grpc/proto/todo.proto](grpc/proto/todo.proto) |
| 2 | **Kubernetes** | Deployment, Service, HPA (auto-scaling) | [kubernetes/deployment.yaml](kubernetes/deployment.yaml) |
| 3 | **CI/CD** | GitHub Actions: test → build → docker → deploy | [cicd/.github/workflows/go.yml](cicd/.github/workflows/go.yml) |
| 4 | **Profiling** | Pprof, race detector, benchmark, fan-out pattern | [profiling/main.go](profiling/main.go) |
| 5 | **Design Patterns** | Repository, Factory, Builder, Singleton di Go | [design-patterns/main.go](design-patterns/main.go) |
| 6 | **WebSocket** | Real-time bidirectional communication | [websocket/main.go](websocket/main.go) |
| 7 | **GraphQL** | Flexible API query language | [graphql/main.go](graphql/main.go) |
| 8 | **Reflection** | reflect package: type introspection, struct tags | [reflection/main.go](reflection/main.go) |
| 9 | **Advanced Sync** | sync.Map, sync.Pool, sync.Once, atomic | [sync-advanced/main.go](sync-advanced/main.go) |
| 10 | **Fuzzing** | Automated testing dengan random input | [fuzzing/main.go](fuzzing/main.go) |
| 11 | **Web Framework** | Gin framework: routing, middleware, binding | [web-framework/main.go](web-framework/main.go) |
| 12 | **ORM** | GORM: auto migrate, CRUD, preloading, transactions | [orm/main.go](orm/main.go) |
| 13 | **Migration** | Database migration tools & patterns | [migration/main.go](migration/main.go) |
| 14 | **Clean Architecture** | Hexagonal architecture, DI, layers | [clean-arch/main.go](clean-arch/main.go) |
| 15 | **Pub/Sub** | Publish-subscribe pattern, event-driven | [pubsub/main.go](pubsub/main.go) |
| 16 | **Event Sourcing** | Event store, audit trail, CQRS | [event-sourcing/main.go](event-sourcing/main.go) |
| 17 | **Resilience** | Circuit breaker, rate limiter, retry patterns | [resilience/main.go](resilience/main.go) |
| 18 | **CGo** | Calling C code from Go | [cgo/main.go](cgo/main.go) |

## Instalasi Dependensi

```bash
# gRPC tools
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# WebSocket
cd websocket && go get github.com/gorilla/websocket

# GraphQL
cd graphql && go get github.com/graphql-go/graphql

# Web Framework (Gin)
cd web-framework && go get github.com/gin-gonic/gin

# ORM (GORM)
cd orm && go get -u gorm.io/gorm gorm.io/driver/sqlite

# Migration
go install -tags 'sqlite3' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

## Commands Penting

```bash
# Profiling
go test -bench=. -cpuprofile=cpu.out -memprofile=mem.out
go tool pprof -http=:8081 cpu.out

# Race detection
go run -race main.go
go test -race ./...

# Fuzzing
go test -fuzz=FuzzReverse -fuzztime=10s

# Build untuk production
GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o app .
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o app .

# WebAssembly (eksekusi Go di browser)
GOOS=js GOARCH=wasm go build -o main.wasm
```

## Prasyarat

Sebelum masuk level mahir, pastikan sudah paham:
- ✅ Semua topik beginner (exercises/)
- ✅ Semua topik middle (context, testing, middleware, database)
- ✅ Docker dasar
- ✅ HTTP & REST API