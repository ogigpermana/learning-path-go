// File: middle/encoding/main.go
// Level: Middle
// Topik: Encoding (CSV, XML, Base64, Hex)
//
// Selain JSON, Go memiliki package untuk format data lain:
// - encoding/csv: membaca/menulis file CSV
// - encoding/xml: XML parsing & serialization
// - encoding/base64: Base64 encoding
// - encoding/hex: Hexadecimal encoding

package main

import (
	"bytes"
	"encoding/base64"
	"encoding/csv"
	"encoding/hex"
	"encoding/xml"
	"fmt"
	"strings"
)

type PersonXML struct {
	XMLName xml.Name `xml:"person"`
	Name    string   `xml:"name"`
	Age     int      `xml:"age"`
	Email   string   `xml:"email,omitempty"`
}

func main() {
	// 1. CSV - Write
	fmt.Println("=== CSV Write ===")
	var csvBuf bytes.Buffer
	writer := csv.NewWriter(&csvBuf)
	writer.Write([]string{"Nama", "Umur", "Kota"})
	writer.Write([]string{"Anggi", "20", "Jakarta"})
	writer.Write([]string{"Budi", "25", "Bandung"})
	writer.Write([]string{"Citra", "22", "Surabaya"})
	writer.Flush()

	csvStr := csvBuf.String()
	fmt.Println("CSV Output:")
	fmt.Println(csvStr)

	// 2. CSV - Read
	fmt.Println("=== CSV Read ===")
	reader := csv.NewReader(strings.NewReader(csvStr))
	records, _ := reader.ReadAll()
	for i, record := range records {
		if i == 0 {
			fmt.Println("Header:", record)
		} else {
			fmt.Printf("Row %d: %+v\n", i, record)
		}
	}
	fmt.Println()

	// 3. XML - Marshal
	fmt.Println("=== XML Marshal ===")
	p := PersonXML{Name: "Anggi", Age: 20, Email: "anggi@mail.com"}
	xmlData, _ := xml.MarshalIndent(p, "", "  ")
	fmt.Println("XML:")
	fmt.Println(string(xmlData))
	fmt.Println()

	// 4. XML - Unmarshal
	fmt.Println("=== XML Unmarshal ===")
	xmlStr := `<person><name>Budi</name><age>25</age></person>`
	var p2 PersonXML
	xml.Unmarshal([]byte(xmlStr), &p2)
	fmt.Printf("Decoded: %+v\n", p2)
	fmt.Println()

	// 5. BASE64
	fmt.Println("=== Base64 ===")
	original := "Hello, Golang!"
	encoded := base64.StdEncoding.EncodeToString([]byte(original))
	decoded, _ := base64.StdEncoding.DecodeString(encoded)

	fmt.Printf("Original: %s\n", original)
	fmt.Printf("Encoded: %s\n", encoded)
	fmt.Printf("Decoded: %s\n", decoded)
	fmt.Println()

	// 6. BASE64 URL (untuk URL & filename)
	fmt.Println("=== Base64 URL ===")
	urlEncoded := base64.URLEncoding.EncodeToString([]byte(original))
	fmt.Printf("URL Encoded: %s\n", urlEncoded)
	fmt.Println()

	// 7. HEX
	fmt.Println("=== Hex ===")
	hexEncoded := hex.EncodeToString([]byte(original))
	hexDecoded, _ := hex.DecodeString(hexEncoded)

	fmt.Printf("Original: %s\n", original)
	fmt.Printf("Hex: %s\n", hexEncoded)
	fmt.Printf("Decoded: %s\n", hexDecoded)
	fmt.Println()

	// 8. HEX dengan format (uppercase)
	fmt.Printf("Hex Upper: %X\n", []byte(original))
}

/*
Pilihan encoding:
- JSON: untuk REST API, web services
- XML: untuk SOAP, legacy systems, config files
- CSV: untuk spreadsheet, data export
- Base64: untuk binary data di text format (email, JWT)
- Hex: untuk debugging, cryptography, low-level protocols
- YAML: untuk config files (gunakan go get gopkg.in/yaml.v3)
*/