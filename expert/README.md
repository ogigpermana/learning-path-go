# Tutorial Go Level Mahir

Level mahir mencakup topik-topik production-grade: arsitektur microservices, deployment, CI/CD, dan optimisasi.

## Topik

| # | Topik | Deskripsi | File | Cara Run |
|---|-------|-----------|------|----------|
| 1 | **gRPC** | High-performance RPC dengan Protocol Buffers | [grpc/proto/todo.proto](grpc/proto/todo.proto) | `protoc --go_out=. --go-grpc_out=. proto/todo.proto` |
| 2 | **Kubernetes** | Deployment, Service, HPA (auto-scaling) | [kubernetes/deployment.yaml](kubernetes/deployment.yaml) | `kubectl apply -f deployment.yaml` |
| 3 | **CI/CD** | GitHub Actions: test → build → docker → deploy | [cicd/.github/workflows/go.yml](cicd/.github/workflows/go.yml) | Push ke GitHub → otomatis jalan |
| 4 | **Profiling** | Pprof, race detector, benchmark, fan-out pattern | [profiling/main.go](profiling/main.go) | `go run profiling/main.go` |
| 5 | **Design Patterns** | Repository, Factory, Builder, Singleton di Go | [design-patterns/main.go](design-patterns/main.go) | `go run design-patterns/main.go` |
| 6 | **WebSocket** | Real-time bidirectional communication | [websocket/main.go](websocket/main.go) | `go run websocket/main.go` |
| 7 | **GraphQL** | Flexible API query language | [graphql/main.go](graphql/main.go) | `go run graphql/main.go` |

## Instalasi Dependensi

```bash
# gRPC tools
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# WebSocket
cd websocket && go get github.com/gorilla/websocket

# GraphQL
cd graphql && go get github.com/graphql-go/graphql
```

## Commands Penting

```bash
# Profiling
go test -bench=. -cpuprofile=cpu.out -memprofile=mem.out
go tool pprof -http=:8081 cpu.out

# Race detection
go run -race main.go
go test -race ./...

# Build untuk production
GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o app .
# ldflags -s -w: strip debug info (smaller binary)

# WebAssembly (eksekusi Go di browser)
GOOS=js GOARCH=wasm go build -o main.wasm
```

## Prasyarat

Sebelum masuk level mahir, pastikan sudah paham:
- ✅ Semua topik beginner (exercises/)
- ✅ Semua topik middle (context, testing, middleware, database)
- ✅ Docker dasar
- ✅ HTTP & REST API