# Belajar Golang dari Nol

Panduan langkah demi langkah untuk belajar bahasa Go (Golang).

## Daftar Isi
1. [Setup Lingkungan](#setup-lingkungan)
2. [Dasar-Dasar](#dasar-dasar)
3. [Variabel & Tipe Data](#variabel--tipe-data)
4. [Fungsi](#fungsi)
5. [Struct & Method](#struct--method)
6. [Interface](#interface)
7. [Goroutine & Channel](#goroutine--channel)
8. [Error Handling](#error-handling)
9. [Testing](#testing)
10. [Project Praktis](#project-praktis)

## Setup Lingkungan

Pastikan Go terinstall:
```bash
go version
```

Buat file pertama: `hello.go`
```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

Jalankan:
```bash
go run hello.go
```

## Cara Belajar

Ikuti urutan materi di atas, latih dengan mengisi kode di file `.go` yang sudah disiapkan di folder `exercises/`.