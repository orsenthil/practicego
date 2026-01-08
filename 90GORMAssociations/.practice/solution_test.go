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
	os.Remove("blog.db")
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

func TestCreateUserWithPosts(t *testing.T) {
	defer cleanupTestDB(t)
	db := setupTestDB(t)

	user := &User{
		Name:  "John Doe",
		Email: "john@example.com",
		Posts: []Post{
			{Title: "Post 1", Content: "Content 1"},
			{Title: "Post 2", Content: "Content 2"},
		},
	}

	err := CreateUserWithPosts(db, user)
	if err != nil {
		t.Fatalf("CreateUserWithPosts failed: %v", err)
	}

	if user.ID == 0 {
		t.Error("Expected user ID to be set")
	}

	// Verify posts were created
	var postCount int64
	db.Model(&Post{}).Where("user_id = ?", user.ID).Count(&postCount)
	if postCount != 2 {
		t.Errorf("Expected 2 posts, got %d", postCount)
	}
}

func TestGetUserWithPosts(t *testing.T) {
	defer cleanupTestDB(t)
	db := setupTestDB(t)

	// Create user with posts
	user := &User{
		Name:  "Jane Smith",
		Email: "jane@example.com",
		Posts: []Post{
			{Title: "Jane's Post 1", Content: "Content 1"},
			{Title: "Jane's Post 2", Content: "Content 2"},
			{Title: "Jane's Post 3", Content: "Content 3"},
		},
	}
	CreateUserWithPosts(db, user)

	// Get user with posts
	fetchedUser, err := GetUserWithPosts(db, user.ID)
	if err != nil {
		t.Fatalf("GetUserWithPosts failed: %v", err)
	}

	if fetchedUser.Name != user.Name {
		t.Errorf("Expected name %s, got %s", user.Name, fetchedUser.Name)
	}

	if len(fetchedUser.Posts) != 3 {
		t.Errorf("Expected 3 posts, got %d", len(fetchedUser.Posts))
	}
}

func TestCreatePostWithTags(t *testing.T) {
	defer cleanupTestDB(t)
	db := setupTestDB(t)

	// Create a user first
	user := &User{Name: "Bob", Email: "bob@example.com"}
	db.Create(user)

	// Create post with tags
	post := &Post{
		Title:   "Go Tutorial",
		Content: "Learning Go",
		UserID:  user.ID,
	}
	tagNames := []string{"golang", "programming", "tutorial"}

	err := CreatePostWithTags(db, post, tagNames)
	if err != nil {
		t.Fatalf("CreatePostWithTags failed: %v", err)
	}

	if post.ID == 0 {
		t.Error("Expected post ID to be set")
	}

	// Verify tags were created
	var tagCount int64
	db.Model(&Tag{}).Count(&tagCount)
	if tagCount != 3 {
		t.Errorf("Expected 3 tags, got %d", tagCount)
	}

	// Verify associations
	var postWithTags Post
	db.Preload("Tags").First(&postWithTags, post.ID)
	if len(postWithTags.Tags) != 3 {
		t.Errorf("Expected 3 tags associated with post, got %d", len(postWithTags.Tags))
	}
}

func TestGetPostsByTag(t *testing.T) {
	defer cleanupTestDB(t)
	db := setupTestDB(t)

	// Create user
	user := &User{Name: "Alice", Email: "alice@example.com"}
	db.Create(user)

	// Create posts with tags
	post1 := &Post{Title: "Go Basics", Content: "Learn Go", UserID: user.ID}
	CreatePostWithTags(db, post1, []string{"golang", "beginner"})

	post2 := &Post{Title: "Advanced Go", Content: "Advanced concepts", UserID: user.ID}
	CreatePostWithTags(db, post2, []string{"golang", "advanced"})

	post3 := &Post{Title: "Python Tutorial", Content: "Learn Python", UserID: user.ID}
	CreatePostWithTags(db, post3, []string{"python", "beginner"})

	// Get posts by tag
	posts, err := GetPostsByTag(db, "golang")
	if err != nil {
		t.Fatalf("GetPostsByTag failed: %v", err)
	}

	if len(posts) != 2 {
		t.Errorf("Expected 2 posts with tag 'golang', got %d", len(posts))
	}
}

func TestAddTagsToPost(t *testing.T) {
	defer cleanupTestDB(t)
	db := setupTestDB(t)

	// Create user and post
	user := &User{Name: "Charlie", Email: "charlie@example.com"}
	db.Create(user)

	post := &Post{Title: "My Post", Content: "Content", UserID: user.ID}
	CreatePostWithTags(db, post, []string{"initial"})

	// Add more tags
	err := AddTagsToPost(db, post.ID, []string{"added1", "added2"})
	if err != nil {
		t.Fatalf("AddTagsToPost failed: %v", err)
	}

	// Verify tags were added
	var postWithTags Post
	db.Preload("Tags").First(&postWithTags, post.ID)
	if len(postWithTags.Tags) != 3 {
		t.Errorf("Expected 3 tags, got %d", len(postWithTags.Tags))
	}
}

func TestGetPostWithUserAndTags(t *testing.T) {
	defer cleanupTestDB(t)
	db := setupTestDB(t)

	// Create user
	user := &User{Name: "David", Email: "david@example.com"}
	db.Create(user)

	// Create post with tags
	post := &Post{
		Title:   "Complete Post",
		Content: "Full content",
		UserID:  user.ID,
	}
	CreatePostWithTags(db, post, []string{"tag1", "tag2", "tag3"})

	// Get post with associations
	fullPost, err := GetPostWithUserAndTags(db, post.ID)
	if err != nil {
		t.Fatalf("GetPostWithUserAndTags failed: %v", err)
	}

	if fullPost.User.Name != user.Name {
		t.Errorf("Expected user name %s, got %s", user.Name, fullPost.User.Name)
	}

	if len(fullPost.Tags) != 3 {
		t.Errorf("Expected 3 tags, got %d", len(fullPost.Tags))
	}
}

func TestManyToManyRelationship(t *testing.T) {
	defer cleanupTestDB(t)
	db := setupTestDB(t)

	// Create user
	user := &User{Name: "Emma", Email: "emma@example.com"}
	db.Create(user)

	// Create multiple posts with shared tags
	post1 := &Post{Title: "Post 1", Content: "Content 1", UserID: user.ID}
	CreatePostWithTags(db, post1, []string{"shared", "tag1"})

	post2 := &Post{Title: "Post 2", Content: "Content 2", UserID: user.ID}
	CreatePostWithTags(db, post2, []string{"shared", "tag2"})

	// Verify the shared tag exists only once
	var tagCount int64
	db.Model(&Tag{}).Where("name = ?", "shared").Count(&tagCount)
	if tagCount != 1 {
		t.Errorf("Expected 1 'shared' tag, got %d", tagCount)
	}

	// Verify both posts have the shared tag
	posts, err := GetPostsByTag(db, "shared")
	if err != nil {
		t.Fatalf("Failed to get posts by tag: %v", err)
	}
	if len(posts) != 2 {
		t.Errorf("Expected 2 posts with 'shared' tag, got %d", len(posts))
	}
}

func TestForeignKeyConstraint(t *testing.T) {
	defer cleanupTestDB(t)
	db := setupTestDB(t)

	// Create a user
	user := &User{Name: "Test User", Email: "test@example.com"}
	db.Create(user)

	// Create a post for that user
	post := &Post{
		Title:   "Valid Post",
		Content: "Has user",
		UserID:  user.ID,
	}

	err := db.Create(post).Error
	if err != nil {
		t.Errorf("Expected no error when creating post with valid UserID, got: %v", err)
	}

	// Verify the post was created with correct foreign key
	var fetchedPost Post
	db.First(&fetchedPost, post.ID)
	if fetchedPost.UserID != user.ID {
		t.Errorf("Expected UserID %d, got %d", user.ID, fetchedPost.UserID)
	}
}

