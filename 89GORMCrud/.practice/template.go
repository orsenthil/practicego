package main

import (
	"time"

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
	// TODO: Implement database connection
	// Hint: Use gorm.Open with sqlite.Open("test.db")
	// Don't forget to auto-migrate the User model
	return nil, nil
}

// CreateUser creates a new user in the database
func CreateUser(db *gorm.DB, user *User) error {
	// TODO: Implement user creation
	// Hint: Use db.Create()
	return nil
}

// GetUserByID retrieves a user by their ID
func GetUserByID(db *gorm.DB, id uint) (*User, error) {
	// TODO: Implement user retrieval by ID
	// Hint: Use db.First() and handle the case where user is not found
	return nil, nil
}

// GetAllUsers retrieves all users from the database
func GetAllUsers(db *gorm.DB) ([]User, error) {
	// TODO: Implement retrieval of all users
	// Hint: Use db.Find()
	return nil, nil
}

// UpdateUser updates an existing user's information
func UpdateUser(db *gorm.DB, user *User) error {
	// TODO: Implement user update
	// Hint: Use db.Save() to update the user
	return nil
}

// DeleteUser removes a user from the database
func DeleteUser(db *gorm.DB, id uint) error {
	// TODO: Implement user deletion
	// Hint: Use db.Delete() with the User ID
	return nil
}

func main() {
	// TODO: Uncomment and complete this section when you're ready to test
	/*
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
	*/
}

// Notes:
// - GORM automatically handles CreatedAt and UpdatedAt timestamps
// - Use gorm.ErrRecordNotFound to check if a record doesn't exist
// - Always check for errors after database operations
// - The database connection should be closed when done (handled by defer)





