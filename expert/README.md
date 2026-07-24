# Tutorial Go untuk Pengguna Mahir

## Daftar Isi
1. Microservices dengan gRPC
2. Docker & Kubernetes
3. CI/CD Pipeline
4. Performance Optimization
5. Security Best Practices

## 1. Microservices dengan gRPC

```
protoc --go_out=. --go-grpc_out=. proto/todo.proto
```

## 2. Dockerize

Dockerfile:
```dockerfile
FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o main .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/main .
CMD ["./main"]
```

Build: `docker build -t go-todo .`

## 3. Kubernetes Deployment

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: go-todo
spec:
  replicas: 3
  selector:
    matchLabels:
      app: go-todo
  template:
    metadata:
      labels:
        app: go-todo
    spec:
      containers:
      - name: go-todo
        image: go-todo:latest
        ports:
        - containerPort: 8080
```

## 4. Performance

- Profiling: `go tool pprof`
- Memory: `go test -memprofile=mem.out`

## 5. Security

- Validation library
- JWT authentication
- Rate limiting