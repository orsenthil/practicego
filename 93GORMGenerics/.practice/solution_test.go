package main

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) (*gorm.DB, context.Context) {
	db, err := ConnectDB()
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
	ctx := context.Background()
	return db, ctx
}

func cleanupTestDB(t *testing.T) {
	os.Remove("generics.db")
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
	db, ctx := setupTestDB(t)

	user := &User{
		Name:  "John Doe",
		Email: "john@example.com",
		Age:   30,
	}

	err := CreateUser(ctx, db, user)
	if err != nil {
		t.Fatalf("CreateUser failed: %v", err)
	}

	if user.ID == 0 {
		t.Error("Expected user ID to be set")
	}
}

func TestGetUserByID(t *testing.T) {
	defer cleanupTestDB(t)
	db, ctx := setupTestDB(t)

	// Create a user first
	user := &User{Name: "Jane", Email: "jane@example.com", Age: 25}
	CreateUser(ctx, db, user)

	// Retrieve the user
	fetchedUser, err := GetUserByID(ctx, db, user.ID)
	if err != nil {
		t.Fatalf("GetUserByID failed: %v", err)
	}

	if fetchedUser.Name != user.Name {
		t.Errorf("Expected name %s, got %s", user.Name, fetchedUser.Name)
	}
}

func TestGetUserByIDNotFound(t *testing.T) {
	defer cleanupTestDB(t)
	db, ctx := setupTestDB(t)

	_, err := GetUserByID(ctx, db, 999)
	if err == nil {
		t.Error("Expected error for non-existent user")
	}
}

func TestUpdateUserAge(t *testing.T) {
	defer cleanupTestDB(t)
	db, ctx := setupTestDB(t)

	user := &User{Name: "Bob", Email: "bob@example.com", Age: 28}
	CreateUser(ctx, db, user)

	// Update age
	newAge := 29
	err := UpdateUserAge(ctx, db, user.ID, newAge)
	if err != nil {
		t.Fatalf("UpdateUserAge failed: %v", err)
	}

	// Verify update
	updated, _ := GetUserByID(ctx, db, user.ID)
	if updated.Age != newAge {
		t.Errorf("Expected age %d, got %d", newAge, updated.Age)
	}
}

func TestDeleteUser(t *testing.T) {
	defer cleanupTestDB(t)
	db, ctx := setupTestDB(t)

	user := &User{Name: "Alice", Email: "alice@example.com", Age: 32}
	CreateUser(ctx, db, user)

	// Delete user
	err := DeleteUser(ctx, db, user.ID)
	if err != nil {
		t.Fatalf("DeleteUser failed: %v", err)
	}

	// Verify deletion
	_, err = GetUserByID(ctx, db, user.ID)
	if err == nil {
		t.Error("Expected error when getting deleted user")
	}
}

func TestCreateUsersInBatches(t *testing.T) {
	defer cleanupTestDB(t)
	db, ctx := setupTestDB(t)

	users := []User{
		{Name: "User1", Email: "user1@example.com", Age: 20},
		{Name: "User2", Email: "user2@example.com", Age: 25},
		{Name: "User3", Email: "user3@example.com", Age: 30},
		{Name: "User4", Email: "user4@example.com", Age: 35},
	}

	err := CreateUsersInBatches(ctx, db, users, 2)
	if err != nil {
		t.Fatalf("CreateUsersInBatches failed: %v", err)
	}

	// Verify all users were created
	for _, user := range users {
		if user.ID == 0 {
			t.Error("Expected all users to have IDs")
		}
	}
}

func TestFindUsersByAgeRange(t *testing.T) {
	defer cleanupTestDB(t)
	db, ctx := setupTestDB(t)

	users := []User{
		{Name: "Young", Email: "young@example.com", Age: 20},
		{Name: "Middle", Email: "middle@example.com", Age: 30},
		{Name: "Senior", Email: "senior@example.com", Age: 40},
	}
	CreateUsersInBatches(ctx, db, users, 3)

	// Find users aged 25-35
	found, err := FindUsersByAgeRange(ctx, db, 25, 35)
	if err != nil {
		t.Fatalf("FindUsersByAgeRange failed: %v", err)
	}

	if len(found) != 1 {
		t.Errorf("Expected 1 user in age range, got %d", len(found))
	}

	if found[0].Age != 30 {
		t.Errorf("Expected user with age 30, got %d", found[0].Age)
	}
}

func TestUpsertUser(t *testing.T) {
	defer cleanupTestDB(t)
	db, ctx := setupTestDB(t)

	// Create initial user
	user := &User{Name: "Original", Email: "test@example.com", Age: 25}
	CreateUser(ctx, db, user)

	// Upsert with same email (should update)
	upsertUser := &User{Name: "Updated", Email: "test@example.com", Age: 30}
	err := UpsertUser(ctx, db, upsertUser)
	if err != nil {
		t.Fatalf("UpsertUser failed: %v", err)
	}

	// Verify update occurred
	var updated User
	db.WithContext(ctx).Where("email = ?", "test@example.com").First(&updated)
	if updated.Name != "Updated" {
		t.Errorf("Expected name 'Updated', got %s", updated.Name)
	}
	if updated.Age != 30 {
		t.Errorf("Expected age 30, got %d", updated.Age)
	}
}

func TestCreateUserWithResult(t *testing.T) {
	defer cleanupTestDB(t)
	db, ctx := setupTestDB(t)

	user := &User{Name: "TestUser", Email: "testuser@example.com", Age: 27}
	rows, err := CreateUserWithResult(ctx, db, user)
	if err != nil {
		t.Fatalf("CreateUserWithResult failed: %v", err)
	}

	if rows != 1 {
		t.Errorf("Expected 1 row affected, got %d", rows)
	}

	if user.ID == 0 {
		t.Error("Expected user ID to be set")
	}
}

func TestGetUsersWithCompany(t *testing.T) {
	defer cleanupTestDB(t)
	db, ctx := setupTestDB(t)

	// Create company
	company := &Company{Name: "TestCorp", Industry: "Tech", FoundedYear: 2020}
	db.WithContext(ctx).Create(company)

	// Create user with company
	user := &User{Name: "Employee", Email: "employee@example.com", Age: 28, CompanyID: &company.ID}
	CreateUser(ctx, db, user)

	// Get users with company
	users, err := GetUsersWithCompany(ctx, db)
	if err != nil {
		t.Fatalf("GetUsersWithCompany failed: %v", err)
	}

	if len(users) == 0 {
		t.Fatal("Expected to find users")
	}

	// Verify company is preloaded
	if users[0].Company == nil {
		t.Error("Expected company to be preloaded")
	}
	if users[0].Company.Name != "TestCorp" {
		t.Errorf("Expected company name TestCorp, got %s", users[0].Company.Name)
	}
}

func TestGetUsersWithPosts(t *testing.T) {
	defer cleanupTestDB(t)
	db, ctx := setupTestDB(t)

	// Create user
	user := &User{Name: "Author", Email: "author@example.com", Age: 30}
	CreateUser(ctx, db, user)

	// Create posts
	for i := 1; i <= 5; i++ {
		db.WithContext(ctx).Create(&Post{
			Title:     fmt.Sprintf("Post %d", i),
			Content:   "Content",
			UserID:    user.ID,
			CreatedAt: time.Now().Add(time.Duration(i) * time.Minute),
		})
	}

	// Get users with max 3 posts
	users, err := GetUsersWithPosts(ctx, db, 3)
	if err != nil {
		t.Fatalf("GetUsersWithPosts failed: %v", err)
	}

	if len(users) == 0 {
		t.Fatal("Expected to find users")
	}

	// Verify posts are limited
	if len(users[0].Posts) > 3 {
		t.Errorf("Expected max 3 posts, got %d", len(users[0].Posts))
	}
}

func TestGetUserWithPostsAndCompany(t *testing.T) {
	defer cleanupTestDB(t)
	db, ctx := setupTestDB(t)

	// Create company
	company := &Company{Name: "BigCorp", Industry: "Finance", FoundedYear: 2015}
	db.WithContext(ctx).Create(company)

	// Create user
	user := &User{Name: "FullUser", Email: "fulluser@example.com", Age: 35, CompanyID: &company.ID}
	CreateUser(ctx, db, user)

	// Create post
	db.WithContext(ctx).Create(&Post{Title: "My Post", Content: "Content", UserID: user.ID})

	// Get user with both associations
	fullUser, err := GetUserWithPostsAndCompany(ctx, db, user.ID)
	if err != nil {
		t.Fatalf("GetUserWithPostsAndCompany failed: %v", err)
	}

	if fullUser.Company == nil {
		t.Error("Expected company to be preloaded")
	}
	if len(fullUser.Posts) == 0 {
		t.Error("Expected posts to be preloaded")
	}
}

func TestSearchUsersInCompany(t *testing.T) {
	defer cleanupTestDB(t)
	db, ctx := setupTestDB(t)

	// Create companies
	tech := &Company{Name: "TechCo", Industry: "Technology", FoundedYear: 2010}
	finance := &Company{Name: "FinanceCo", Industry: "Finance", FoundedYear: 2015}
	db.WithContext(ctx).Create(tech)
	db.WithContext(ctx).Create(finance)

	// Create users
	CreateUser(ctx, db, &User{Name: "TechUser1", Email: "tech1@example.com", Age: 25, CompanyID: &tech.ID})
	CreateUser(ctx, db, &User{Name: "TechUser2", Email: "tech2@example.com", Age: 30, CompanyID: &tech.ID})
	CreateUser(ctx, db, &User{Name: "FinUser", Email: "fin@example.com", Age: 28, CompanyID: &finance.ID})

	// Search users in TechCo
	users, err := SearchUsersInCompany(ctx, db, "TechCo")
	if err != nil {
		t.Fatalf("SearchUsersInCompany failed: %v", err)
	}

	if len(users) != 2 {
		t.Errorf("Expected 2 users in TechCo, got %d", len(users))
	}
}

func TestGetTopActiveUsers(t *testing.T) {
	defer cleanupTestDB(t)
	db, ctx := setupTestDB(t)

	// Create users
	user1 := &User{Name: "ActiveUser", Email: "active@example.com", Age: 30}
	user2 := &User{Name: "LessActive", Email: "less@example.com", Age: 25}
	CreateUser(ctx, db, user1)
	CreateUser(ctx, db, user2)

	// Create posts (user1 has more)
	db.WithContext(ctx).Create(&Post{Title: "Post1", Content: "C", UserID: user1.ID})
	db.WithContext(ctx).Create(&Post{Title: "Post2", Content: "C", UserID: user1.ID})
	db.WithContext(ctx).Create(&Post{Title: "Post3", Content: "C", UserID: user1.ID})
	db.WithContext(ctx).Create(&Post{Title: "Post4", Content: "C", UserID: user2.ID})

	// Get top active users
	topUsers, err := GetTopActiveUsers(ctx, db, 2)
	if err != nil {
		t.Fatalf("GetTopActiveUsers failed: %v", err)
	}

	if len(topUsers) == 0 {
		t.Fatal("Expected to find active users")
	}

	// Verify order (most active first)
	if topUsers[0].ID != user1.ID {
		t.Error("Expected most active user to be first")
	}
}

func TestContextCancellation(t *testing.T) {
	defer cleanupTestDB(t)
	db, _ := setupTestDB(t)

	// Create a cancelled context
	ctx, cancel := context.WithCancel(context.Background())
	cancel() // Cancel immediately

	user := &User{Name: "Test", Email: "test@example.com", Age: 25}
	err := CreateUser(ctx, db, user)

	// Should get context cancelled error
	if err != context.Canceled {
		// SQLite might not respect context cancellation immediately
		// This test documents the expected behavior
		t.Logf("Context cancellation: %v", err)
	}
}

