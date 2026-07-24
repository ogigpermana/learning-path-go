// File: 18-module/main/go.mod
// Level: Beginner
// Topik: Go Module dengan Local Dependensi
//
// go.mod ini memiliki local dependensi ke module greeter.
// "replace" digunakan untuk menunjuk ke path lokal.
// Di production, "replace" diganti dengan version number.

module example.com/main

go 1.22

// Module main bergantung pada module greeter
require example.com/greeter v0.0.0

// replace mengarahkan Go ke folder local
// Ini berguna saat development module paralel
replace example.com/greeter => ../greeter

// Untuk dependensi eksternal (contoh):
// require github.com/gorilla/mux v1.8.1