// File: expert/cgo/main.go
// Level: Expert
// Topik: CGo - Memanggil C dari Go
//
// CGo memungkinkan Go memanggil kode C.
// Digunakan untuk: sistem programming, legacy library, performa kritis.
//
// Syntax: import "C" (sebelumnya ada komentar dengan kode C)
//
// Kekurangan: build lebih lambat, cross-compile lebih susah,
// no goroutine safety, memory management manual.
//
// Tips: hindari CGo jika ada alternatif Go murni.

package main

/*
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

// Fungsi C sederhana
void say_hello(const char* name) {
    printf("Hello from C: %s!\n", name);
}

// Fungsi C yang mengembalikan nilai
int add(int a, int b) {
    return a + b;
}

// Fungsi C dengan struct
typedef struct {
    char name[50];
    int age;
} Person;

void print_person(Person p) {
    printf("Person: %s, Age: %d\n", p.name, p.age);
}
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println("=== CGo Demo ===")

	// 1. Panggil fungsi C
	fmt.Println("\n1. Call C function (void):")
	C.say_hello(C.CString("Anggi"))

	// 2. Panggil fungsi C dengan return value
	fmt.Println("\n2. Call C function (return int):")
	result := C.add(C.int(10), C.int(20))
	fmt.Printf("C.add(10, 20) = %d\n", int(result))

	// 3. C struct
	fmt.Println("\n3. C struct:")
	p := C.Person{}
	// Isi field struct
	name := C.CString("Anggi")
	defer C.free(unsafe.Pointer(name))
	C.strcpy(&p.name[0], name)
	p.age = C.int(20)

	C.print_person(p)

	// 4. String conversion
	fmt.Println("\n4. String conversion:")
	goStr := "Hello from Go!"
	cStr := C.CString(goStr)
	defer C.free(unsafe.Pointer(cStr)) // FREE memory!
	fmt.Printf("Go string -> C string: %s\n", C.GoString(cStr))

	// 5. Memory management
	fmt.Println("\n5. Memory management:")
	cMemory := C.malloc(C.size_t(100))
	defer C.free(cMemory)
	if cMemory == nil {
		fmt.Println("Failed to allocate memory")
		return
	}
	fmt.Println("C memory allocated successfully")

	// 6. C array
	fmt.Println("\n6. C array:")
	cArray := C.malloc(C.size_t(5 * C.sizeof_int))
	defer C.free(cArray)

	// Isi array
	intPtr := (*[5]C.int)(cArray)
	for i := range intPtr {
		intPtr[i] = C.int(i * 10)
	}

	fmt.Printf("C array: ")
	for i := range intPtr {
		fmt.Printf("%d ", int(intPtr[i]))
	}
	fmt.Println()
}

/*
CGo build:
go build -o app .                    # build biasa (butuh GCC)
CGO_ENABLED=0 go build -o app .     # build tanpa CGo

Cross-compile dengan CGo:
CC=aarch64-linux-gnu-gcc GOOS=linux GOARCH=arm64 CGO_ENABLED=1 go build

Alternatif tanpa CGo:
- package syscall (untuk OS system calls)
- package os/exec (untuk menjalankan program external)
- c-for-go (generate Go bindings dari C headers)
- ebpf (untuk kernel programming)
*/