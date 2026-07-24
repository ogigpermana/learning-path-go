// File: 14-map.go
// Level: Beginner
// Topik: Map
//
// Map adalah struktur data key-value (seperti dictionary/hash map).
// Syntax: map[KeyType]ValueType
// - Key: harus comparable (bisa dibandingkan dengan ==)
// - Value: tipe data apapun

package main

import "fmt"

func main() {
	// Deklarasi map literal
	// map[string]int berarti key string, value int
	scores := map[string]int{
		"Anggi": 90, // key: "Anggi", value: 90
		"Budi":  85,
	}
	fmt.Println("Scores:", scores)

	// Menambah dan mengupdate data
	scores["Citra"] = 95  // menambah key baru
	scores["Anggi"] = 100 // mengupdate key yang sudah ada
	fmt.Println("Updated:", scores)

	// Mengakses value dengan dua return value
	// val  : nilai jika key ada
	// exists: true jika key ada, false jika tidak
	val, exists := scores["Deni"]
	if exists {
		fmt.Println("Deni:", val)
	} else {
		fmt.Println("Deni tidak ditemukan")
	}

	// Menghapus key dengan delete()
	delete(scores, "Budi")
	fmt.Println("Setelah delete Budi:", scores)

	// Iterasi map dengan range (urutan tidak dijamin)
	fmt.Println("\nDaftar scores:")
	for name, score := range scores {
		fmt.Printf("%s: %d\n", name, score)
	}

	// Map sebagai set (hanya butuh key, tanpa value berarti)
	visited := map[string]bool{}
	visited["/home"] = true
	visited["/about"] = true

	if visited["/home"] {
		fmt.Println("Home sudah dikunjungi")
	}

	// Map dengan nilai kompleks
	students := map[string]map[string]int{
		"Anggi": {"matematika": 90, "fisika": 85},
		"Budi":  {"matematika": 80, "fisika": 95},
	}
	fmt.Println("\nNilai matematika Anggi:", students["Anggi"]["matematika"])

	// Nil map - harus diinisialisasi dengan make() sebelum digunakan
	var m map[string]int
	if m == nil {
		fmt.Println("m adalah nil map")
		m = make(map[string]int) // inisialisasi
		m["test"] = 1
		fmt.Println("Nilai m[\"test\"]:", m["test"])
	}

	// Panjang map
	fmt.Println("Jumlah students:", len(scores))
}