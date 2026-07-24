// File: expert/clean-arch/main.go
// Level: Expert
// Topik: Clean Architecture / Hexagonal Architecture
//
// Clean Architecture (Robert C. Martin) = memisahkan kode menjadi layers:
// 1. Entity/Business: aturan bisnis inti (tidak tergantung apapun)
// 2. Use Case: alur use case aplikasi
// 3. Repository: interface untuk database
// 4. Handler/Delivery: HTTP handler, gRPC, CLI
//
// Dependency Inversion: layer dalam tidak tahu layer luar (hanya interface)

package main

import (
	"fmt"
	"time"
)

// ==========================================
// DOMAIN / ENTITY (Layer paling dalam)
// ==========================================

// User adalah entity bisnis, tidak tergantung framework/database
type User struct {
	ID        int
	Name      string
	Email     string
	CreatedAt time.Time
}

// ==========================================
// REPOSITORY INTERFACE (Port)
// ==========================================

// UserRepository adalah port (interface) yang didefinisikan dalam domain
type UserRepository interface {
	FindByID(id int) (*User, error)
	FindAll() ([]User, error)
	Save(user *User) error
	Delete(id int) error
}

// ==========================================
// USE CASE (Aplikasi Layer)
// ==========================================

// UserService adalah use case / service layer
type UserService struct {
	repo UserRepository // tergantung interface, bukan implementasi
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Register(name, email string) (*User, error) {
	// Business logic: validasi
	if name == "" {
		return nil, fmt.Errorf("name cannot be empty")
	}
	if email == "" {
		return nil, fmt.Errorf("email cannot be empty")
	}

	user := &User{
		Name:      name,
		Email:     email,
		CreatedAt: time.Now(),
	}

	// Simpan via repository (abstraksi database)
	if err := s.repo.Save(user); err != nil {
		return nil, fmt.Errorf("failed to save user: %w", err)
	}

	return user, nil
}

func (s *UserService) GetUser(id int) (*User, error) {
	return s.repo.FindByID(id)
}

func (s *UserService) ListUsers() ([]User, error) {
	return s.repo.FindAll()
}

// ==========================================
// INFRASTRUCTURE (Layer Luar)
// ==========================================

// InMemoryUserRepo adalah implementasi repository (infrastructure)
type InMemoryUserRepo struct {
	users  map[int]User
	nextID int
}

func NewInMemoryUserRepo() *InMemoryUserRepo {
	return &InMemoryUserRepo{
		users:  make(map[int]User),
		nextID: 1,
	}
}

func (r *InMemoryUserRepo) FindByID(id int) (*User, error) {
	if user, ok := r.users[id]; ok {
		return &user, nil
	}
	return nil, fmt.Errorf("user %d not found", id)
}

func (r *InMemoryUserRepo) FindAll() ([]User, error) {
	var result []User
	for _, user := range r.users {
		result = append(result, user)
	}
	return result, nil
}

func (r *InMemoryUserRepo) Save(user *User) error {
	if user.ID == 0 {
		user.ID = r.nextID
		r.nextID++
	}
	r.users[user.ID] = *user
	return nil
}

func (r *InMemoryUserRepo) Delete(id int) error {
	delete(r.users, id)
	return nil
}

// ==========================================
// DELIVERY / HTTP HANDLER (Layer Luar)
// ==========================================

// UserHandler adalah HTTP adapter
type UserHandler struct {
	service *UserService
}

func NewUserHandler(service *UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) HandleRegister(name, email string) {
	user, err := h.service.Register(name, email)
	if err != nil {
		fmt.Printf("[ERROR] Register gagal: %v\n", err)
		return
	}
	fmt.Printf("[SUCCESS] User registered: %+v\n", user)
}

func (h *UserHandler) HandleList() {
	users, err := h.service.ListUsers()
	if err != nil {
		fmt.Printf("[ERROR] List gagal: %v\n", err)
		return
	}
	fmt.Printf("[SUCCESS] Users: %d\n", len(users))
	for _, u := range users {
		fmt.Printf("  %d: %s (%s)\n", u.ID, u.Name, u.Email)
	}
}

// ==========================================
// MAIN (Composition Root)
// ==========================================

func main() {
	fmt.Println("=== Clean Architecture Demo ===")
	fmt.Println()

	// Dependency Injection: wire komponen di main()
	repo := NewInMemoryUserRepo()
	service := NewUserService(repo)
	handler := NewUserHandler(service)

	// Gunakan handler
	handler.HandleRegister("Anggi", "anggi@mail.com")
	handler.HandleRegister("Budi", "budi@mail.com")
	handler.HandleList()

	fmt.Println()
	fmt.Println("=== Keuntungan Clean Architecture ===")
	fmt.Println("1. Domain tidak tergantung framework/database")
	fmt.Println("2. Mudah ganti database (cukup buat implementasi baru)")
	fmt.Println("3. Mudah di-test (mock repository)")
	fmt.Println("4. Kode terorganisir dengan jelas")
	fmt.Println("5. Business logic terisolasi")
}

/*
Struktur folder untuk production:
clean-arch/
├── domain/
│   └── user.go           # Entity, Repository interface
├── usecase/
│   └── user_service.go   # Business logic
├── repository/
│   └── postgres.go       # Implementasi database
├── handler/
│   └── http.go           # HTTP handler
├── middleware/
│   └── auth.go           # Middleware
└── main.go               # Dependency injection
*/