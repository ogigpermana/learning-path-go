# Tutorial Go untuk Pengguna Menengah

## Daftar Isi
1. Testing Lanjutan
2. REST API dengan JSON
3. Middleware & Logging
4. Environment Variables
5. Dockerize Go App

## 1. Testing Lanjutan

### Table-Driven Tests

```go
func TestBagi(t *testing.T) {
    tests := []struct {
        name     string
        a, b     int
        expected float64
        hasError bool
    }{
        {"normal division", 10, 2, 5, false},
        {"division by zero", 10, 0, 0, true},
        {"negative numbers", -10, 2, -5, false},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // test implementation
        })
    }
}
```

### Mocking Database

Gunakan `testify/mock`:
```go
go get github.com/stretchr/testify/mock
```

### Benchmarking

```go
func BenchmarkTambah(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Tambah(100, 200)
    }
}
```

Run: `go test -bench=.`

## 2. REST API dengan JSON

Lihat contoh di `api/todo-api.go`