package main

import (
	"time"

	"gorm.io/gorm"
)

// User represents a user in the blog system
type User struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"not null"`
	Email     string `gorm:"unique;not null"`
	Posts     []Post `gorm:"foreignKey:UserID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Post represents a blog post
type Post struct {
	ID        uint   `gorm:"primaryKey"`
	Title     string `gorm:"not null"`
	Content   string `gorm:"type:text"`
	UserID    uint   `gorm:"not null"`
	User      User   `gorm:"foreignKey:UserID"`
	Tags      []Tag  `gorm:"many2many:post_tags;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Tag represents a tag for categorizing posts
type Tag struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"unique;not null"`
	Posts []Post `gorm:"many2many:post_tags;"`
}

// ConnectDB establishes a connection to the SQLite database and auto-migrates the models
func ConnectDB() (*gorm.DB, error) {
	// TODO: Implement database connection with auto-migration
	// Hint: Use gorm.Open with sqlite.Open("blog.db")
	// Don't forget to auto-migrate all three models: User, Post, Tag
	return nil, nil
}

// CreateUserWithPosts creates a new user with associated posts
func CreateUserWithPosts(db *gorm.DB, user *User) error {
	// TODO: Implement user creation with posts
	// Hint: GORM automatically handles associated Posts when you create a User
	// Just use db.Create() - GORM will create the user and all associated posts
	return nil
}

// GetUserWithPosts retrieves a user with all their posts preloaded
func GetUserWithPosts(db *gorm.DB, userID uint) (*User, error) {
	// TODO: Implement user retrieval with posts
	// Hint: Use db.Preload("Posts") to load associated posts
	// Then use db.First() to get the user
	return nil, nil
}

// CreatePostWithTags creates a new post with specified tags
func CreatePostWithTags(db *gorm.DB, post *Post, tagNames []string) error {
	// TODO: Implement post creation with tags
	// Hint: First, find or create tags by name
	// Then assign them to post.Tags
	// Finally, create the post with db.Create()
	return nil
}

// GetPostsByTag retrieves all posts that have a specific tag
func GetPostsByTag(db *gorm.DB, tagName string) ([]Post, error) {
	// TODO: Implement posts retrieval by tag
	// Hint: First find the tag by name
	// Then use db.Model(&tag).Association("Posts").Find(&posts)
	// Or use db.Joins with the post_tags table
	return nil, nil
}

// AddTagsToPost adds tags to an existing post
func AddTagsToPost(db *gorm.DB, postID uint, tagNames []string) error {
	// TODO: Implement adding tags to existing post
	// Hint: First find the post
	// Then find or create the tags
	// Finally use db.Model(&post).Association("Tags").Append(tags)
	return nil
}

// GetPostWithUserAndTags retrieves a post with user and tags preloaded
func GetPostWithUserAndTags(db *gorm.DB, postID uint) (*Post, error) {
	// TODO: Implement post retrieval with user and tags
	// Hint: Use db.Preload("User").Preload("Tags") to load both associations
	// Then use db.First() to get the post
	return nil, nil
}

func main() {
	// TODO: Uncomment and complete this section when you're ready to test
	/*
		// Connect to database
		db, err := ConnectDB()
		if err != nil {
			log.Fatal("Failed to connect to database:", err)
		}

		// Create a user with posts
		user := &User{
			Name:  "Alice Johnson",
			Email: "alice@example.com",
			Posts: []Post{
				{Title: "First Post", Content: "This is my first post"},
				{Title: "Second Post", Content: "This is my second post"},
			},
		}
		if err := CreateUserWithPosts(db, user); err != nil {
			log.Fatal("Failed to create user with posts:", err)
		}
		fmt.Printf("Created user %s with %d posts\n", user.Name, len(user.Posts))

		// Get user with posts
		fetchedUser, err := GetUserWithPosts(db, user.ID)
		if err != nil {
			log.Fatal("Failed to get user with posts:", err)
		}
		fmt.Printf("Fetched user %s with %d posts\n", fetchedUser.Name, len(fetchedUser.Posts))

		// Create a post with tags
		post := &Post{
			Title:   "Go Programming Tips",
			Content: "Here are some tips for Go programming",
			UserID:  user.ID,
		}
		tagNames := []string{"golang", "programming", "tutorial"}
		if err := CreatePostWithTags(db, post, tagNames); err != nil {
			log.Fatal("Failed to create post with tags:", err)
		}
		fmt.Printf("Created post '%s' with %d tags\n", post.Title, len(tagNames))

		// Get posts by tag
		posts, err := GetPostsByTag(db, "golang")
		if err != nil {
			log.Fatal("Failed to get posts by tag:", err)
		}
		fmt.Printf("Found %d posts with tag 'golang'\n", len(posts))

		// Add tags to existing post
		if err := AddTagsToPost(db, user.Posts[0].ID, []string{"beginner", "guide"}); err != nil {
			log.Fatal("Failed to add tags to post:", err)
		}
		fmt.Println("Added tags to existing post")

		// Get post with user and tags
		fullPost, err := GetPostWithUserAndTags(db, post.ID)
		if err != nil {
			log.Fatal("Failed to get post with associations:", err)
		}
		fmt.Printf("Post '%s' by %s with %d tags\n", fullPost.Title, fullPost.User.Name, len(fullPost.Tags))
	*/
}

// Notes:
// - Use Preload() to eagerly load associations
// - GORM automatically handles foreign key constraints
// - Many-to-many relationships use a join table (post_tags)
// - Association() API is useful for managing relationships
// - Always handle errors when working with associations

