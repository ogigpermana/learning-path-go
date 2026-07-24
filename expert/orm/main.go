// File: expert/orm/main.go
// Level: Expert
// Topik: ORM dengan GORM
//
// ORM (Object-Relational Mapping) memetakan struct Go ke tabel database.
// GORM adalah ORM paling populer di Go.
//
// Fitur: Auto Migrate, CRUD, Preloading, Hooks, Transactions, Scopes.
//
// Install: go get -u gorm.io/gorm gorm.io/driver/sqlite

package main

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// ===== MODEL DEFINITIONS =====

// User dengan GORM tags
type User struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"size:100;not null"`
	Email     string    `gorm:"uniqueIndex;size:255"`
	Age       int       `gorm:"default:0"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"` // soft delete
	Posts     []Post         // has many relationship
}

// Post dengan foreign key ke User
type Post struct {
	ID        uint   `gorm:"primaryKey"`
	Title     string `gorm:"not null"`
	Body      string `gorm:"type:text"`
	UserID    uint   // foreign key
	CreatedAt time.Time
}

func main() {
	// 1. KONEKSI DATABASE
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Warn),
	})
	if err != nil {
		log.Fatal("Gagal koneksi:", err)
	}

	// 2. AUTO MIGRATE (buat tabel otomatis)
	db.AutoMigrate(&User{}, &Post{})
	fmt.Println("Migration selesai")
	fmt.Println()

	// Cleanup
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()
	defer db.Migrator().DropTable(&User{}, &Post{})

	// 3. CREATE
	fmt.Println("=== CREATE ===")
	user := User{
		Name:  "Anggi",
		Email: "anggi@mail.com",
		Age:   20,
		Posts: []Post{
			{Title: "Post 1", Body: "Content 1"},
			{Title: "Post 2", Body: "Content 2"},
		},
	}
	result := db.Create(&user)
	fmt.Printf("Created user ID %d (rows: %d)\n", user.ID, result.RowsAffected)
	fmt.Println()

	// 4. QUERY SINGLE
	fmt.Println("=== QUERY ===")
	var found User
	db.First(&found, user.ID) // by primary key
	fmt.Printf("First: %+v\n", found)

	db.First(&found, "email = ?", "anggi@mail.com")
	fmt.Printf("By email: %+v\n", found)
	fmt.Println()

	// 5. QUERY MULTIPLE
	fmt.Println("=== QUERY ALL ===")
	var users []User
	db.Find(&users)
	fmt.Printf("Total users: %d\n", len(users))
	for _, u := range users {
		fmt.Printf("  %d: %s (%s)\n", u.ID, u.Name, u.Email)
	}
	fmt.Println()

	// 6. WHERE & CONDITIONS
	fmt.Println("=== WHERE ===")
	var adults []User
	db.Where("age >= ?", 17).Find(&adults)
	fmt.Printf("Adult users: %d\n", len(adults))
	fmt.Println()

	// 7. PRELOAD (EAGER LOADING)
	fmt.Println("=== PRELOAD ===")
	var userWithPosts User
	db.Preload("Posts").First(&userWithPosts, user.ID)
	fmt.Printf("User: %s\n", userWithPosts.Name)
	for _, post := range userWithPosts.Posts {
		fmt.Printf("  Post: %s - %s\n", post.Title, post.Body)
	}
	fmt.Println()

	// 8. UPDATE
	fmt.Println("=== UPDATE ===")
	db.Model(&found).Update("age", 21)
	db.First(&found, user.ID)
	fmt.Printf("Updated age: %d\n", found.Age)

	// Update multiple fields
	db.Model(&found).Updates(User{Name: "Anggi Ganteng", Age: 22})
	db.First(&found, user.ID)
	fmt.Printf("Updated: %+v\n", found)
	fmt.Println()

	// 9. DELETE (soft delete)
	fmt.Println("=== DELETE ===")
	db.Delete(&found) // soft delete (set DeletedAt)
	var deletedCount int64
	db.Model(&User{}).Where("deleted_at IS NOT NULL").Count(&deletedCount)
	fmt.Printf("Soft deleted: %d\n", deletedCount)

	// Unscoped: true delete
	// db.Unscoped().Delete(&found)
	fmt.Println()

	// 10. TRANSACTION
	fmt.Println("=== TRANSACTION ===")
	db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&User{Name: "Budi", Email: "budi@mail.com"}).Error; err != nil {
			return err // rollback
		}
		if err := tx.Create(&User{Name: "Citra", Email: "citra@mail.com"}).Error; err != nil {
			return err // rollback
		}
		return nil // commit
	})
	fmt.Println("Transaction selesai")
}
/*
GORM Fitur Lain:
- Hooks: BeforeCreate, AfterCreate, BeforeUpdate, dll
- Scopes: reusable query conditions
- Raw SQL: db.Raw("SELECT * FROM users WHERE age > ?", 17)
- Joins: db.Joins("JOIN posts ON posts.user_id = users.id")
- Pagination: db.Limit(10).Offset(20).Find(&users)
- Count: db.Model(&User{}).Where("age > ?", 17).Count(&count)
- Pluck: db.Model(&User{}).Pluck("name", &names)

Alternatif ORM: Ent (facebook), sqlx (near SQL), bun
*/