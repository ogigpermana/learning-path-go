// File: middle/cross-compile/main.go
// Level: Middle
// Topik: Cross Compilation
//
// Cross-compile = build binary untuk OS/architecture lain.
// Go memiliki dukungan cross-compile terbaik (built-in, tanpa toolchain tambahan).
//
// Env vars:
// GOOS: target operating system
// GOARCH: target architecture
// CGO_ENABLED=0: disable CGo untuk static binary

package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("=== Cross Compilation Guide ===")
	fmt.Println()
	fmt.Printf("Current platform: %s/%s\n", runtime.GOOS, runtime.GOARCH)
	fmt.Println()

	fmt.Println("Build untuk berbagai platform:")
	fmt.Println()

	targets := []struct {
		OS      string
		Arch    string
		Output  string
		Desc    string
	}{
		{"linux", "amd64", "app-linux-amd64", "Linux 64-bit (most servers)"},
		{"linux", "arm64", "app-linux-arm64", "Linux ARM 64-bit (Raspberry Pi 4, AWS Graviton)"},
		{"linux", "arm", "app-linux-arm", "Linux ARM 32-bit (Raspberry Pi, IoT)"},
		{"darwin", "amd64", "app-darwin-amd64", "macOS Intel"},
		{"darwin", "arm64", "app-darwin-arm64", "macOS Apple Silicon (M1/M2)"},
		{"windows", "amd64", "app-windows-amd64.exe", "Windows 64-bit"},
		{"windows", "arm64", "app-windows-arm64.exe", "Windows ARM"},
		{"freebsd", "amd64", "app-freebsd-amd64", "FreeBSD 64-bit"},
		{"android", "arm64", "app-android-arm64", "Android ARM 64-bit"},
		{"ios", "arm64", "app-ios-arm64", "iOS ARM 64-bit"},
		{"js", "wasm", "main.wasm", "WebAssembly (browser)"},
	}

	for _, t := range targets {
		fmt.Printf("  GOOS=%s GOARCH=%s  => %s  (%s)\n",
			t.OS, t.Arch, t.Output, t.Desc)
	}

	fmt.Println()
	fmt.Println("=== Build Commands ===")
	fmt.Println()
	fmt.Printf("# Build untuk Linux AMD64\n")
	fmt.Printf("GOOS=linux GOARCH=amd64 go build -o app-linux-amd64 .\n")
	fmt.Println()
	fmt.Printf("# Build untuk macOS ARM64\n")
	fmt.Printf("GOOS=darwin GOARCH=arm64 go build -o app-darwin-arm64 .\n")
	fmt.Println()
	fmt.Printf("# Build untuk Windows\n")
	fmt.Printf("GOOS=windows GOARCH=amd64 go build -o app.exe .\n")
	fmt.Println()
	fmt.Printf("# Build untuk WebAssembly\n")
	fmt.Printf("GOOS=js GOARCH=wasm go build -o main.wasm .\n")
	fmt.Println()
	fmt.Printf("# Build tanpa CGo (static binary)\n")
	fmt.Printf("CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app .\n")
	fmt.Println()
	fmt.Printf("# Build dengan optimasi ukuran\n")
	fmt.Printf("GOOS=linux GOARCH=amd64 go build -ldflags=\"-s -w\" -o app .\n")
	fmt.Println()
	fmt.Printf("# Build dengan UPX compression\n")
	fmt.Printf("go build -o app . && upx --best app\n")
	fmt.Println()
	fmt.Println("=== Build Tags ===")
	fmt.Println("Build tags: conditional compilation")
	fmt.Println()
	fmt.Printf("//go:build debug\n")
	fmt.Printf("package main\n")
	fmt.Println()
	fmt.Printf("Build: go build -tags debug .\n")
	fmt.Println()
	fmt.Println("=== Tips ===")
	fmt.Println("1. CGO_ENABLED=0 untuk static binary (deploy mudah)")
	fmt.Println("2. -ldflags=\"-s -w\" untuk binary lebih kecil")
	fmt.Println("3. Go binary sudah static (tidak perlu runtime)")
	fmt.Println("4. Cross-compile Go TIDAK butuh toolchain target")
	fmt.Println("5. Untuk CGo cross-compile, butuh cross-compiler C")

	fmt.Println("\nCurrent build info:")
	fmt.Printf("  OS: %s\n", runtime.GOOS)
	fmt.Printf("  Arch: %s\n", runtime.GOARCH)
	fmt.Printf("  CPUs: %d\n", runtime.NumCPU())
	fmt.Printf("  Go Version: %s\n", runtime.Version())
}

/*
Quick build script (build.sh):
#!/bin/bash
platforms=("linux/amd64" "linux/arm64" "darwin/amd64" "darwin/arm64" "windows/amd64")
for platform in "${platforms[@]}"; do
    split=(${platform//\// })
    GOOS=${split[0]}
    GOARCH=${split[1]}
    output="app-${GOOS}-${GOARCH}"
    if [ $GOOS = "windows" ]; then
        output+='.exe'
    fi
    GOOS=$GOOS GOARCH=$GOARCH go build -ldflags="-s -w" -o $output .
    echo "Built: $output"
done
*/