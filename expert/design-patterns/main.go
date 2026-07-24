// File: expert/design-patterns/main.go
// Level: Expert
// Topik: Design Patterns di Go
//
// Design patterns adalah solusi umum untuk masalah umum dalam programming.
// Go tidak mewarisi pola OOP klasik (inheritance), tapi punya cara sendiri.
//
// Pola-pola yang umum di Go:
// 1. Repository Pattern - abstraksi database
// 2. Factory Pattern - membuat objek berdasarkan kondisi
// 3. Builder Pattern - konstruksi objek kompleks bertahap
// 4. Singleton Pattern - satu instance global
// 5. Strategy Pattern - algoritma interchangeable

package main

import "fmt"

// ==========================================
// 1. REPOSITORY PATTERN
// Abstraksi layer database, memudahkan switching DB
// ==========================================

// User adalah entity/model
type User struct {
	ID   int
	Name string
}

// UserRepository adalah kontrak/interface untuk operasi user
// Dengan interface, kita bisa ganti implementasi database kapan saja
type UserRepository interface {
	GetUser(id int) (*User, error)
	CreateUser(name string) (*User, error)
	DeleteUser(id int) error
}

// InMemoryUserRepo adalah implementasi in-memory (untuk development)
type InMemoryUserRepo struct {
	users map[int]*User
	nextID int
}

func NewInMemoryUserRepo() *InMemoryUserRepo {
	return &InMemoryUserRepo{
		users:  make(map[int]*User),
		nextID: 1,
	}
}

func (r *InMemoryUserRepo) GetUser(id int) (*User, error) {
	if user, ok := r.users[id]; ok {
		return user, nil
	}
	return nil, fmt.Errorf("user %d tidak ditemukan", id)
}

func (r *InMemoryUserRepo) CreateUser(name string) (*User, error) {
	user := &User{ID: r.nextID, Name: name}
	r.users[r.nextID] = user
	r.nextID++
	return user, nil
}

func (r *InMemoryUserRepo) DeleteUser(id int) error {
	if _, ok := r.users[id]; !ok {
		return fmt.Errorf("user %d tidak ditemukan", id)
	}
	delete(r.users, id)
	return nil
}

// UserRepositoryPostgres adalah implementasi PostgreSQL (simulasi)
type UserRepositoryPostgres struct {
	// db *sql.DB  // koneksi database
}

func (r *UserRepositoryPostgres) GetUser(id int) (*User, error) {
	// return db.QueryRow("SELECT ...", id)
	panic("implementasi postgres")
}
func (r *UserRepositoryPostgres) CreateUser(name string) (*User, error) {
	panic("implementasi postgres")
}
func (r *UserRepositoryPostgres) DeleteUser(id int) error {
	panic("implementasi postgres")
}

// UserService menggunakan UserRepository (dependency injection)
type UserService struct {
	repo UserRepository // tergantung interface, bukan implementasi
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Register(name string) (*User, error) {
	return s.repo.CreateUser(name)
}

// ==========================================
// 2. FACTORY PATTERN
// Membuat objek berdasarkan tipe yang diminta
// ==========================================

// Logger interface
type Logger interface {
	Log(msg string)
}

// ConsoleLogger mencetak ke console
type ConsoleLogger struct{}
func (c ConsoleLogger) Log(msg string) {
	fmt.Println("CONSOLE:", msg)
}

// FileLogger menulis ke file (simulasi)
type FileLogger struct{}
func (f FileLogger) Log(msg string) {
	// f.write("[FILE] " + msg)
	fmt.Println("FILE:", msg)
}

// NewLogger adalah factory function
func NewLogger(loggerType string) Logger {
	switch loggerType {
	case "file":
		return FileLogger{}
	default:
		return ConsoleLogger{}
	}
}

// ==========================================
// 3. BUILDER PATTERN
// Membangun objek kompleks step-by-step
// ==========================================

type Pizza struct {
	Size      string
	Crust     string
	Cheese    bool
	Pepperoni bool
	Mushrooms bool
	Olives    bool
}

type PizzaBuilder struct {
	pizza Pizza
}

func (b *PizzaBuilder) SetSize(size string) *PizzaBuilder {
	b.pizza.Size = size
	return b // return builder untuk method chaining
}

func (b *PizzaBuilder) SetCrust(crust string) *PizzaBuilder {
	b.pizza.Crust = crust
	return b
}

func (b *PizzaBuilder) AddCheese() *PizzaBuilder {
	b.pizza.Cheese = true
	return b
}

func (b *PizzaBuilder) AddPepperoni() *PizzaBuilder {
	b.pizza.Pepperoni = true
	return b
}

func (b *PizzaBuilder) Build() Pizza {
	return b.pizza
}

// ==========================================
// 4. SINGLETON PATTERN
// Satu instance global untuk suatu objek
// ==========================================

// singleton menggunakan sync.Once untuk thread safety
type AppConfig struct {
	Version string
	Env     string
}

var instance *AppConfig

func GetAppConfig() *AppConfig {
	// Initialize only once (sederhana, tanpa sync.Once)
	if instance == nil {
		instance = &AppConfig{
			Version: "1.0.0",
			Env:     "production",
		}
	}
	return instance
}

// ==========================================
// MAIN
// ==========================================

func main() {
	fmt.Println("=== Repository Pattern ===")
	repo := NewInMemoryUserRepo()
	service := NewUserService(repo)

	user, _ := service.Register("Anggi")
	fmt.Printf("Created: %+v\n", user)
	user, _ = service.Register("Budi")
	fmt.Printf("Created: %+v\n", user)
	fetched, _ := repo.GetUser(1)
	fmt.Printf("Fetched: %+v\n", fetched)

	fmt.Println("\n=== Factory Pattern ===")
	consoleLogger := NewLogger("console")
	consoleLogger.Log("test message")
	fileLogger := NewLogger("file")
	fileLogger.Log("test message")

	fmt.Println("\n=== Builder Pattern ===")
	pizza := (&PizzaBuilder{}).
		SetSize("Large").
		SetCrust("Thin").
		AddCheese().
		AddPepperoni().
		Build()
	fmt.Printf("Pizza: size=%s, crust=%s, cheese=%t, pepperoni=%t\n",
		pizza.Size, pizza.Crust, pizza.Cheese, pizza.Pepperoni)

	fmt.Println("\n=== Singleton Pattern ===")
	cfg1 := GetAppConfig()
	cfg2 := GetAppConfig()
	fmt.Printf("Config: %+v\n", cfg1)
	fmt.Println("Same instance:", cfg1 == cfg2)
}