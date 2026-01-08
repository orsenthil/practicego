package main

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/sqlite"
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
	db, err := gorm.Open(sqlite.Open("blog.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Auto-migrate all models
	err = db.AutoMigrate(&User{}, &Post{}, &Tag{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

// CreateUserWithPosts creates a new user with associated posts
func CreateUserWithPosts(db *gorm.DB, user *User) error {
	// GORM automatically creates associated Posts when creating a User
	result := db.Create(user)
	return result.Error
}

// GetUserWithPosts retrieves a user with all their posts preloaded
func GetUserWithPosts(db *gorm.DB, userID uint) (*User, error) {
	var user User
	result := db.Preload("Posts").First(&user, userID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// CreatePostWithTags creates a new post with specified tags
func CreatePostWithTags(db *gorm.DB, post *Post, tagNames []string) error {
	// Find or create tags
	var tags []Tag
	for _, name := range tagNames {
		var tag Tag
		// FirstOrCreate will find existing tag or create new one
		result := db.Where(Tag{Name: name}).FirstOrCreate(&tag)
		if result.Error != nil {
			return result.Error
		}
		tags = append(tags, tag)
	}

	// Assign tags to post
	post.Tags = tags

	// Create the post with associated tags
	result := db.Create(post)
	return result.Error
}

// GetPostsByTag retrieves all posts that have a specific tag
func GetPostsByTag(db *gorm.DB, tagName string) ([]Post, error) {
	var tag Tag
	// Find the tag by name
	result := db.Where("name = ?", tagName).First(&tag)
	if result.Error != nil {
		return nil, result.Error
	}

	// Get all posts associated with this tag
	var posts []Post
	err := db.Model(&tag).Association("Posts").Find(&posts)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

// AddTagsToPost adds tags to an existing post
func AddTagsToPost(db *gorm.DB, postID uint, tagNames []string) error {
	// Find the post
	var post Post
	result := db.First(&post, postID)
	if result.Error != nil {
		return result.Error
	}

	// Find or create tags
	var tags []Tag
	for _, name := range tagNames {
		var tag Tag
		result := db.Where(Tag{Name: name}).FirstOrCreate(&tag)
		if result.Error != nil {
			return result.Error
		}
		tags = append(tags, tag)
	}

	// Add tags to the post using Association API
	err := db.Model(&post).Association("Tags").Append(tags)
	return err
}

// GetPostWithUserAndTags retrieves a post with user and tags preloaded
func GetPostWithUserAndTags(db *gorm.DB, postID uint) (*Post, error) {
	var post Post
	result := db.Preload("User").Preload("Tags").First(&post, postID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &post, nil
}

func main() {
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
}

// Notes:
// - Use Preload() to eagerly load associations
// - GORM automatically handles foreign key constraints
// - Many-to-many relationships use a join table (post_tags)
// - Association() API is useful for managing relationships
// - Always handle errors when working with associations

