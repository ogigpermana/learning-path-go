// File: middle/io/main.go
// Level: Middle
// Topik: io.Reader & io.Writer
//
// io.Reader dan io.Writer adalah interface paling fundamental di Go.
// Hampir semua I/O di Go menggunakan interface ini.
//
// type Reader interface {
//     Read(p []byte) (n int, err error)
// }
//
// type Writer interface {
//     Write(p []byte) (n int, err error)
// }
//
// Keuntungan: kode jadi reusable untuk berbagai sumber data
// (file, network, memory buffer, compression, encryption, dll)

package main

import (
	"bytes"     // bytes.Buffer implements Reader & Writer
	"compress/gzip" // gzip.Reader & gzip.Writer
	"crypto/md5"
	"fmt"
	"io"        // Reader, Writer interface
	"os"
	"strings"   // strings.Reader implements Reader
)

func main() {
	// 1. STRINGS.READER (io.Reader dari string)
	fmt.Println("=== strings.Reader ===")
	reader := strings.NewReader("Hello, io.Reader!")
	
	// Membaca byte by byte
	buf := make([]byte, 4)
	for {
		n, err := reader.Read(buf)
		if err == io.EOF {
			break // selesai membaca
		}
		fmt.Printf("Read %d bytes: %s\n", n, buf[:n])
	}
	fmt.Println()

	// 2. BYTES.BUFFER (io.Reader + io.Writer + in-memory buffer)
	fmt.Println("=== bytes.Buffer ===")
	var buf2 bytes.Buffer
	
	// Write ke buffer
	buf2.WriteString("Data pertama\n")
	buf2.Write([]byte("Data kedua\n"))
	fmt.Printf("Buffer size: %d bytes\n", buf2.Len())
	
	// Read dari buffer
	line, _ := buf2.ReadString('\n')
	fmt.Printf("Read line: %s", line)
	fmt.Println()

	// 3. io.Copy - menyalin data dari Reader ke Writer
	fmt.Println("=== io.Copy ===")
	src := strings.NewReader("Data yang akan di-copy")
	var dst bytes.Buffer
	
	n, _ := io.Copy(&dst, src)
	fmt.Printf("Copied %d bytes: %s\n", n, dst.String())
	fmt.Println()

	// 4. io.MultiReader - menggabungkan multiple Reader
	fmt.Println("=== io.MultiReader ===")
	r1 := strings.NewReader("Bagian 1, ")
	r2 := strings.NewReader("Bagian 2, ")
	r3 := strings.NewReader("Bagian 3")
	mr := io.MultiReader(r1, r2, r3)
	
	io.Copy(os.Stdout, mr)
	fmt.Println("\n")

	// 5. io.MultiWriter - menulis ke multiple Writer
	fmt.Println("=== io.MultiWriter ===")
	var buf3, buf4 bytes.Buffer
	mw := io.MultiWriter(&buf3, &buf4)
	
	mw.Write([]byte("Data ke semua writer"))
	fmt.Println("buf3:", buf3.String())
	fmt.Println("buf4:", buf4.String())
	fmt.Println()

	// 6. io.Pipe - Reader/Writer terhubung (seperti pipeline)
	fmt.Println("=== io.Pipe ===")
	pr, pw := io.Pipe()
	
	go func() {
		pw.Write([]byte("Data dari pipe"))
		pw.Close()
	}()
	
	pipeData, _ := io.ReadAll(pr)
	fmt.Println("Pipe data:", string(pipeData))
	fmt.Println()

	// 7. COMPRESSION (gzip) - contoh chain Reader/Writer
	fmt.Println("=== gzip Compression ===")
	var compressed bytes.Buffer
	gw := gzip.NewWriter(&compressed)
	gw.Write([]byte("Data yang akan di-compress"))
	gw.Close()
	
	fmt.Printf("Compressed: %d bytes\n", compressed.Len())
	
	gr, _ := gzip.NewReader(&compressed)
	decompressed, _ := io.ReadAll(gr)
	fmt.Printf("Decompressed: %s\n", decompressed)
	fmt.Println()

	// 8. io.TeeReader - baca data sambil tulis ke Writer lain
	fmt.Println("=== io.TeeReader ===")
	var logBuffer bytes.Buffer
	originalData := strings.NewReader("Data penting")
	teeReader := io.TeeReader(originalData, &logBuffer)
	
	// Baca data, otomatis log ke logBuffer
	finalData, _ := io.ReadAll(teeReader)
	fmt.Printf("Final: %s\n", finalData)
	fmt.Printf("Log: %s\n", logBuffer.String())
	fmt.Println()

	// 9. REAL WORLD: hash file
	fmt.Println("=== Hash String with MD5 ===")
	h := md5.New()
	io.WriteString(h, "password123")
	fmt.Printf("MD5: %x\n", h.Sum(nil))
}

/*
io.Reader dan io.Writer ada di mana-mana:
- os.File (implements Reader & Writer)
- http.Response.Body (implements Reader)
- TCP/UDP connections (implements Reader & Writer)
- bytes.Buffer (implements Reader & Writer)
- gzip.Reader/Writer (wraps another Reader/Writer)
- crypto cipher (wraps Reader/Writer for encryption)
*/