package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// User represents a user in the system with company association
type User struct {
	ID        uint     `gorm:"primaryKey"`
	Name      string   `gorm:"not null"`
	Email     string   `gorm:"unique;not null"`
	Age       int      `gorm:"check:age > 0"`
	CompanyID *uint    `gorm:"index"`
	Company   *Company `gorm:"foreignKey:CompanyID"`
	Posts     []Post   `gorm:"foreignKey:UserID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Company represents a company that users can belong to
type Company struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"not null;unique"`
	Industry    string `gorm:"not null"`
	FoundedYear int    `gorm:"not null"`
	Users       []User `gorm:"foreignKey:CompanyID"`
	CreatedAt   time.Time
}

// Post represents a blog post by a user
type Post struct {
	ID        uint   `gorm:"primaryKey"`
	Title     string `gorm:"not null"`
	Content   string `gorm:"type:text"`
	UserID    uint   `gorm:"not null;index"`
	User      User   `gorm:"foreignKey:UserID"`
	ViewCount int    `gorm:"default:0"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// ConnectDB establishes a connection to the SQLite database and auto-migrates models
func ConnectDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("generics.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Auto-migrate all models
	err = db.AutoMigrate(&User{}, &Company{}, &Post{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

// CreateUser creates a new user with context support
func CreateUser(ctx context.Context, db *gorm.DB, user *User) error {
	result := db.WithContext(ctx).Create(user)
	return result.Error
}

// GetUserByID retrieves a user by ID with context
func GetUserByID(ctx context.Context, db *gorm.DB, id uint) (*User, error) {
	var user User
	result := db.WithContext(ctx).First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// UpdateUserAge updates a user's age with context
func UpdateUserAge(ctx context.Context, db *gorm.DB, userID uint, age int) error {
	result := db.WithContext(ctx).Model(&User{}).Where("id = ?", userID).Update("age", age)
	return result.Error
}

// DeleteUser deletes a user by ID with context
func DeleteUser(ctx context.Context, db *gorm.DB, userID uint) error {
	result := db.WithContext(ctx).Delete(&User{}, userID)
	return result.Error
}

// CreateUsersInBatches creates multiple users in batches for better performance
func CreateUsersInBatches(ctx context.Context, db *gorm.DB, users []User, batchSize int) error {
	result := db.WithContext(ctx).CreateInBatches(users, batchSize)
	return result.Error
}

// FindUsersByAgeRange finds users within an age range
func FindUsersByAgeRange(ctx context.Context, db *gorm.DB, minAge, maxAge int) ([]User, error) {
	var users []User
	result := db.WithContext(ctx).Where("age BETWEEN ? AND ?", minAge, maxAge).Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

// UpsertUser creates or updates a user handling conflicts using Clauses
func UpsertUser(ctx context.Context, db *gorm.DB, user *User) error {
	result := db.WithContext(ctx).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "email"}},
		DoUpdates: clause.AssignmentColumns([]string{"name", "age"}),
	}).Create(user)
	return result.Error
}

// CreateUserWithResult creates a user and returns result metadata
func CreateUserWithResult(ctx context.Context, db *gorm.DB, user *User) (int64, error) {
	result := db.WithContext(ctx).Create(user)
	return result.RowsAffected, result.Error
}

// GetUsersWithCompany retrieves users with their company information using Preload
func GetUsersWithCompany(ctx context.Context, db *gorm.DB) ([]User, error) {
	var users []User
	result := db.WithContext(ctx).Preload("Company").Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

// GetUsersWithPosts retrieves users with their posts (limited per user)
func GetUsersWithPosts(ctx context.Context, db *gorm.DB, limit int) ([]User, error) {
	var users []User
	result := db.WithContext(ctx).Preload("Posts", func(db *gorm.DB) *gorm.DB {
		return db.Order("created_at DESC").Limit(limit)
	}).Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

// GetUserWithPostsAndCompany retrieves a user with both posts and company preloaded
func GetUserWithPostsAndCompany(ctx context.Context, db *gorm.DB, userID uint) (*User, error) {
	var user User
	result := db.WithContext(ctx).
		Preload("Company").
		Preload("Posts").
		First(&user, userID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// SearchUsersInCompany finds users working in a specific company
func SearchUsersInCompany(ctx context.Context, db *gorm.DB, companyName string) ([]User, error) {
	var users []User
	result := db.WithContext(ctx).
		Joins("Company").
		Where("Company.name = ?", companyName).
		Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

// GetTopActiveUsers retrieves users with the most posts
func GetTopActiveUsers(ctx context.Context, db *gorm.DB, limit int) ([]User, error) {
	var users []User
	result := db.WithContext(ctx).
		Joins("LEFT JOIN posts ON posts.user_id = users.id").
		Group("users.id").
		Order("COUNT(posts.id) DESC").
		Limit(limit).
		Preload("Posts").
		Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func main() {
	ctx := context.Background()

	db, err := ConnectDB()
	if err != nil {
		log.Fatal("Failed to connect:", err)
	}

	// Create companies
	tech := &Company{Name: "TechCorp", Industry: "Technology", FoundedYear: 2010}
	db.WithContext(ctx).Create(tech)

	finance := &Company{Name: "FinanceInc", Industry: "Finance", FoundedYear: 2015}
	db.WithContext(ctx).Create(finance)

	fmt.Println("Created companies")

	// Create users in batches
	users := []User{
		{Name: "Alice", Email: "alice@example.com", Age: 30, CompanyID: &tech.ID},
		{Name: "Bob", Email: "bob@example.com", Age: 35, CompanyID: &tech.ID},
		{Name: "Charlie", Email: "charlie@example.com", Age: 28, CompanyID: &finance.ID},
	}

	if err := CreateUsersInBatches(ctx, db, users, 2); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Created users in batches")

	// Create posts
	db.WithContext(ctx).Create(&Post{Title: "Go Basics", Content: "Learn Go programming", UserID: users[0].ID, ViewCount: 100})
	db.WithContext(ctx).Create(&Post{Title: "Advanced Go", Content: "Master Go concurrency", UserID: users[0].ID, ViewCount: 150})
	db.WithContext(ctx).Create(&Post{Title: "Go Testing", Content: "Testing in Go", UserID: users[1].ID, ViewCount: 80})
	fmt.Println("Created posts")

	// Get user by ID
	user, err := GetUserByID(ctx, db, users[0].ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Found user: %s (Age: %d)\n", user.Name, user.Age)

	// Update age
	if err := UpdateUserAge(ctx, db, users[0].ID, 31); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Updated user age to 31")

	// Find by age range
	ageRangeUsers, err := FindUsersByAgeRange(ctx, db, 28, 35)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Found %d users in age range 28-35\n", len(ageRangeUsers))

	// Upsert user (update on conflict)
	existingUser := &User{Name: "Alice Updated", Email: "alice@example.com", Age: 32, CompanyID: &tech.ID}
	if err := UpsertUser(ctx, db, existingUser); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Upserted user (handled email conflict)")

	// Create with result metadata
	newUser := &User{Name: "David", Email: "david@example.com", Age: 29, CompanyID: &tech.ID}
	rows, err := CreateUserWithResult(ctx, db, newUser)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Created user, rows affected: %d\n", rows)

	// Get users with company
	usersWithCompany, err := GetUsersWithCompany(ctx, db)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Found %d users with companies:\n", len(usersWithCompany))
	for _, u := range usersWithCompany {
		if u.Company != nil {
			fmt.Printf("  - %s works at %s\n", u.Name, u.Company.Name)
		}
	}

	// Get users with posts (limited)
	usersWithPosts, err := GetUsersWithPosts(ctx, db, 2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nUsers with posts (max 2 per user):\n")
	for _, u := range usersWithPosts {
		if len(u.Posts) > 0 {
			fmt.Printf("  - %s has %d post(s)\n", u.Name, len(u.Posts))
		}
	}

	// Get user with posts and company
	fullUser, err := GetUserWithPostsAndCompany(ctx, db, users[0].ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nFull user info for %s:\n", fullUser.Name)
	if fullUser.Company != nil {
		fmt.Printf("  Company: %s (%s)\n", fullUser.Company.Name, fullUser.Company.Industry)
	}
	fmt.Printf("  Posts: %d\n", len(fullUser.Posts))

	// Search users in company
	techUsers, err := SearchUsersInCompany(ctx, db, "TechCorp")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\n%d users work at TechCorp\n", len(techUsers))

	// Get top active users
	topUsers, err := GetTopActiveUsers(ctx, db, 3)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nTop %d active users:\n", len(topUsers))
	for i, u := range topUsers {
		fmt.Printf("  %d. %s (%d posts)\n", i+1, u.Name, len(u.Posts))
	}

	// Delete user
	if err := DeleteUser(ctx, db, newUser.ID); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nDeleted user with ID %d\n", newUser.ID)
}

// Notes:
// - Context support enables cancellation, timeouts, and request-scoped values
// - Use db.WithContext(ctx) for all database operations
// - Batch operations improve performance for bulk inserts
// - Preload with custom conditions allows efficient association loading
// - Clauses like OnConflict handle edge cases elegantly
// - Always handle errors and gorm.ErrRecordNotFound

