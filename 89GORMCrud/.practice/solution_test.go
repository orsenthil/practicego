package main

import (
	"os"
	"testing"

	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := ConnectDB()
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
	return db
}

func cleanupTestDB(t *testing.T) {
	os.Remove("test.db")
}

func TestConnectDB(t *testing.T) {
	defer cleanupTestDB(t)
	
	db, err := ConnectDB()
	if err != nil {
		t.Fatalf("ConnectDB failed: %v", err)
	}
	if db == nil {
		t.Fatal("Expected non-nil database connection")
	}
}

func TestCreateUser(t *testing.T) {
	defer cleanupTestDB(t)
	db := setupTestDB(t)
	
	user := &User{
		Name:  "Alice Smith",
		Email: "alice@example.com",
		Age:   25,
	}
	
	err := CreateUser(db, user)
	if err != nil {
		t.Fatalf("CreateUser failed: %v", err)
	}
	
	if user.ID == 0 {
		t.Error("Expected user ID to be set after creation")
	}
}

func TestGetUserByID(t *testing.T) {
	defer cleanupTestDB(t)
	db := setupTestDB(t)
	
	// Create a user first
	user := &User{
		Name:  "Bob Johnson",
		Email: "bob@example.com",
		Age:   30,
	}
	CreateUser(db, user)
	
	// Retrieve the user
	fetchedUser, err := GetUserByID(db, user.ID)
	if err != nil {
		t.Fatalf("GetUserByID failed: %v", err)
	}
	
	if fetchedUser.Name != user.Name {
		t.Errorf("Expected name %s, got %s", user.Name, fetchedUser.Name)
	}
	if fetchedUser.Email != user.Email {
		t.Errorf("Expected email %s, got %s", user.Email, fetchedUser.Email)
	}
}

func TestGetUserByIDNotFound(t *testing.T) {
	defer cleanupTestDB(t)
	db := setupTestDB(t)
	
	_, err := GetUserByID(db, 999)
	if err == nil {
		t.Error("Expected error for non-existent user")
	}
}

func TestGetAllUsers(t *testing.T) {
	defer cleanupTestDB(t)
	db := setupTestDB(t)
	
	// Create multiple users
	users := []*User{
		{Name: "User1", Email: "user1@example.com", Age: 20},
		{Name: "User2", Email: "user2@example.com", Age: 30},
		{Name: "User3", Email: "user3@example.com", Age: 40},
	}
	
	for _, user := range users {
		CreateUser(db, user)
	}
	
	allUsers, err := GetAllUsers(db)
	if err != nil {
		t.Fatalf("GetAllUsers failed: %v", err)
	}
	
	if len(allUsers) != 3 {
		t.Errorf("Expected 3 users, got %d", len(allUsers))
	}
}

func TestUpdateUser(t *testing.T) {
	defer cleanupTestDB(t)
	db := setupTestDB(t)
	
	// Create a user
	user := &User{
		Name:  "Charlie Brown",
		Email: "charlie@example.com",
		Age:   28,
	}
	CreateUser(db, user)
	
	// Update the user
	user.Age = 29
	user.Name = "Charlie B."
	err := UpdateUser(db, user)
	if err != nil {
		t.Fatalf("UpdateUser failed: %v", err)
	}
	
	// Fetch and verify the update
	fetchedUser, _ := GetUserByID(db, user.ID)
	if fetchedUser.Age != 29 {
		t.Errorf("Expected age 29, got %d", fetchedUser.Age)
	}
	if fetchedUser.Name != "Charlie B." {
		t.Errorf("Expected name 'Charlie B.', got %s", fetchedUser.Name)
	}
}

func TestDeleteUser(t *testing.T) {
	defer cleanupTestDB(t)
	db := setupTestDB(t)
	
	// Create a user
	user := &User{
		Name:  "David Wilson",
		Email: "david@example.com",
		Age:   35,
	}
	CreateUser(db, user)
	
	// Delete the user
	err := DeleteUser(db, user.ID)
	if err != nil {
		t.Fatalf("DeleteUser failed: %v", err)
	}
	
	// Verify the user is deleted
	_, err = GetUserByID(db, user.ID)
	if err == nil {
		t.Error("Expected error when getting deleted user")
	}
}

func TestCreateUserValidation(t *testing.T) {
	defer cleanupTestDB(t)
	db := setupTestDB(t)
	
	// Test duplicate email
	user1 := &User{
		Name:  "User One",
		Email: "duplicate@example.com",
		Age:   25,
	}
	CreateUser(db, user1)
	
	user2 := &User{
		Name:  "User Two",
		Email: "duplicate@example.com", // Same email
		Age:   30,
	}
	err := CreateUser(db, user2)
	if err == nil {
		t.Error("Expected error for duplicate email")
	}
}





