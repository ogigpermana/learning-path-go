// File: middle/auth/main.go
// Level: Middle
// Topik: JWT Authentication (Implementasi Manual)
//
// JWT (JSON Web Token) format: header.payload.signature
// 1. Header: algoritma & tipe token
// 2. Payload: data (userID, username, exp)
// 3. Signature: verifikasi token tidak dimodifikasi
//
// Implementasi ini untuk PEMBELAJARAN.
// Untuk production: go get github.com/golang-jwt/jwt/v5

package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

// jwtSecret adalah kunci rahasia untuk sign & verify token
// Dalam production: simpan di environment variable
var jwtSecret = []byte("rahasia-kunci-jwt-2024")

// createToken membuat JWT token baru
// Format: base64URL(header).base64URL(payload).base64URL(signature)
func createToken(userID int, username string) string {
	// Header: algoritma HS256 & tipe JWT
	header := `{"alg":"HS256","typ":"JWT"}`
	// Payload: data user + expired time (24 jam dari sekarang)
	payload := fmt.Sprintf(
		`{"user_id":%d,"username":"%s","exp":%d}`,
		userID, username, time.Now().Add(24*time.Hour).Unix(),
	)

	// Encode header & payload ke base64 URL-safe
	encodedHeader := base64URLEncode([]byte(header))
	encodedPayload := base64URLEncode([]byte(payload))

	// Buat signing input: header.payload
	signingInput := encodedHeader + "." + encodedPayload
	// Sign menggunakan HMAC-SHA256
	signature := sign(signingInput)

	// Token final: header.payload.signature
	return signingInput + "." + signature
}

// sign membuat signature menggunakan HMAC-SHA256
func sign(input string) string {
	mac := hmac.New(sha256.New, jwtSecret)
	mac.Write([]byte(input))
	return base64URLEncode(mac.Sum(nil))
}

// base64URLEncode encode ke base64 tanpa padding (=)
func base64URLEncode(data []byte) string {
	return strings.TrimRight(base64.URLEncoding.EncodeToString(data), "=")
}

// validateToken memverifikasi token dan mengembalikan payload
func validateToken(token string) (map[string]interface{}, error) {
	// Split token menjadi 3 bagian
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return nil, fmt.Errorf("format token tidak valid")
	}

	// Verifikasi signature
	signingInput := parts[0] + "." + parts[1]
	expectedSig := sign(signingInput)
	if parts[2] != expectedSig {
		return nil, fmt.Errorf("signature tidak valid")
	}

	// Decode payload
	payloadBytes, err := base64.URLEncoding.WithPadding(base64.NoPadding).DecodeString(parts[1])
	if err != nil {
		return nil, fmt.Errorf("gagal decode payload")
	}

	// Parse payload ke map
	var payload map[string]interface{}
	json.Unmarshal(payloadBytes, &payload)

	// Verifikasi expired
	exp := payload["exp"].(float64)
	if time.Now().Unix() > int64(exp) {
		return nil, fmt.Errorf("token sudah expired")
	}

	return payload, nil
}

// authHandler - endpoint yang membutuhkan auth
func authHandler(w http.ResponseWriter, r *http.Request) {
	// Ambil token dari header Authorization: Bearer <token>
	token := r.Header.Get("Authorization")
	token = strings.TrimPrefix(token, "Bearer ")

	payload, err := validateToken(token)
	if err != nil {
		http.Error(w, "Unauthorized: "+err.Error(), http.StatusUnauthorized)
		return
	}

	fmt.Fprintf(w, "Welcome %s! (User ID: %.0f)\n",
		payload["username"], payload["user_id"])
}

// loginHandler - endpoint untuk mendapatkan token
func loginHandler(w http.ResponseWriter, r *http.Request) {
	// Di production: validasi username & password dari database
	token := createToken(1, "Anggi")
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"token":"%s"}`, token)
}

func main() {
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/protected", authHandler)

	fmt.Println("Auth server running on :8080")
	fmt.Println("\nCoba:")
	fmt.Println("1. curl localhost:8080/login")
	fmt.Println("2. TOKEN=$(curl -s localhost:8080/login | grep -o '\"token\":\"[^\"]*\"' | cut -d'\"' -f4)")
	fmt.Println("3. curl -H 'Authorization: Bearer $TOKEN' localhost:8080/protected")
	http.ListenAndServe(":8080", nil)
}