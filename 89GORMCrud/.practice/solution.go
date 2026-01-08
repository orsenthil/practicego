package main

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// User represents a user in the system
type User struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"not null"`
	Email     string `gorm:"unique;not null"`
	Age       int    `gorm:"check:age > 0"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// ConnectDB establishes a connection to the SQLite database
func ConnectDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Auto-migrate the User model
	err = db.AutoMigrate(&User{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

// CreateUser creates a new user in the database
func CreateUser(db *gorm.DB, user *User) error {
	result := db.Create(user)
	return result.Error
}

// GetUserByID retrieves a user by their ID
func GetUserByID(db *gorm.DB, id uint) (*User, error) {
	var user User
	result := db.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// GetAllUsers retrieves all users from the database
func GetAllUsers(db *gorm.DB) ([]User, error) {
	var users []User
	result := db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

// UpdateUser updates an existing user's information
func UpdateUser(db *gorm.DB, user *User) error {
	result := db.Save(user)
	return result.Error
}

// DeleteUser removes a user from the database
func DeleteUser(db *gorm.DB, id uint) error {
	result := db.Delete(&User{}, id)
	return result.Error
}

func main() {
	// Connect to database
	db, err := ConnectDB()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Create a new user
	user := &User{
		Name:  "John Doe",
		Email: "john@example.com",
		Age:   30,
	}
	if err := CreateUser(db, user); err != nil {
		log.Fatal("Failed to create user:", err)
	}
	fmt.Printf("Created user with ID: %d\n", user.ID)

	// Get user by ID
	fetchedUser, err := GetUserByID(db, user.ID)
	if err != nil {
		log.Fatal("Failed to get user:", err)
	}
	fmt.Printf("Fetched user: %+v\n", fetchedUser)

	// Update user
	fetchedUser.Age = 31
	if err := UpdateUser(db, fetchedUser); err != nil {
		log.Fatal("Failed to update user:", err)
	}
	fmt.Println("User updated successfully")

	// Get all users
	users, err := GetAllUsers(db)
	if err != nil {
		log.Fatal("Failed to get all users:", err)
	}
	fmt.Printf("Total users: %d\n", len(users))

	// Delete user
	if err := DeleteUser(db, user.ID); err != nil {
		log.Fatal("Failed to delete user:", err)
	}
	fmt.Println("User deleted successfully")
}

// Notes:
// - GORM automatically handles CreatedAt and UpdatedAt timestamps
// - Use gorm.ErrRecordNotFound to check if a record doesn't exist
// - Always check for errors after database operations
// - The database connection should be closed when done (handled by defer)





