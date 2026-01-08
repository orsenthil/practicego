package main

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// User represents a user in the social media system
type User struct {
	ID        uint   `gorm:"primaryKey"`
	Username  string `gorm:"unique;not null"`
	Email     string `gorm:"unique;not null"`
	Age       int    `gorm:"not null"`
	Country   string `gorm:"not null"`
	CreatedAt time.Time
	Posts     []Post `gorm:"foreignKey:UserID"`
	Likes     []Like `gorm:"foreignKey:UserID"`
}

// Post represents a social media post
type Post struct {
	ID          uint   `gorm:"primaryKey"`
	Title       string `gorm:"not null"`
	Content     string `gorm:"type:text"`
	UserID      uint   `gorm:"not null"`
	User        User   `gorm:"foreignKey:UserID"`
	Category    string `gorm:"not null"`
	ViewCount   int    `gorm:"default:0"`
	IsPublished bool   `gorm:"default:true"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Likes       []Like `gorm:"foreignKey:PostID"`
}

// Like represents a user's like on a post
type Like struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint `gorm:"not null"`
	PostID    uint `gorm:"not null"`
	User      User `gorm:"foreignKey:UserID"`
	Post      Post `gorm:"foreignKey:PostID"`
	CreatedAt time.Time
}

// UserWithPostCount is a helper struct for aggregation queries
type UserWithPostCount struct {
	User
	PostCount int64
}

// PostWithLikeCount is a helper struct for aggregation queries
type PostWithLikeCount struct {
	Post
	LikeCount int64
}

// ConnectDB establishes a connection to the SQLite database with auto-migration
func ConnectDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("social.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Auto-migrate all models
	err = db.AutoMigrate(&User{}, &Post{}, &Like{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

// GetTopUsersByPostCount retrieves users with the most posts
func GetTopUsersByPostCount(db *gorm.DB, limit int) ([]User, error) {
	var users []User
	
	// Subquery to get post counts
	result := db.
		Preload("Posts").
		Joins("LEFT JOIN posts ON posts.user_id = users.id").
		Group("users.id").
		Order("COUNT(posts.id) DESC").
		Limit(limit).
		Find(&users)
	
	if result.Error != nil {
		return nil, result.Error
	}
	
	return users, nil
}

// GetPostsByCategoryWithUserInfo retrieves posts by category with pagination and user info
func GetPostsByCategoryWithUserInfo(db *gorm.DB, category string, page, pageSize int) ([]Post, int64, error) {
	var posts []Post
	var total int64
	
	// Get total count
	db.Model(&Post{}).Where("category = ?", category).Count(&total)
	
	// Calculate offset
	offset := (page - 1) * pageSize
	
	// Get paginated results
	result := db.
		Where("category = ?", category).
		Preload("User").
		Offset(offset).
		Limit(pageSize).
		Order("created_at DESC").
		Find(&posts)
	
	if result.Error != nil {
		return nil, 0, result.Error
	}
	
	return posts, total, nil
}

// GetUserEngagementStats calculates engagement statistics for a user
func GetUserEngagementStats(db *gorm.DB, userID uint) (map[string]interface{}, error) {
	stats := make(map[string]interface{})
	
	// Total posts
	var totalPosts int64
	db.Model(&Post{}).Where("user_id = ?", userID).Count(&totalPosts)
	stats["total_posts"] = totalPosts
	
	// Total likes received (likes on user's posts)
	var totalLikesReceived int64
	db.Model(&Like{}).
		Joins("JOIN posts ON posts.id = likes.post_id").
		Where("posts.user_id = ?", userID).
		Count(&totalLikesReceived)
	stats["total_likes_received"] = totalLikesReceived
	
	// Total likes given
	var totalLikesGiven int64
	db.Model(&Like{}).Where("user_id = ?", userID).Count(&totalLikesGiven)
	stats["total_likes_given"] = totalLikesGiven
	
	// Average post views
	var avgViews float64
	db.Model(&Post{}).
		Where("user_id = ?", userID).
		Select("AVG(view_count)").
		Scan(&avgViews)
	stats["avg_post_views"] = avgViews
	
	return stats, nil
}

// GetPopularPostsByLikes retrieves popular posts by likes in a time period
func GetPopularPostsByLikes(db *gorm.DB, days int, limit int) ([]Post, error) {
	cutoffDate := time.Now().AddDate(0, 0, -days)
	
	var posts []Post
	result := db.
		Preload("User").
		Preload("Likes").
		Joins("LEFT JOIN likes ON likes.post_id = posts.id").
		Where("posts.created_at >= ?", cutoffDate).
		Group("posts.id").
		Order("COUNT(likes.id) DESC").
		Limit(limit).
		Find(&posts)
	
	if result.Error != nil {
		return nil, result.Error
	}
	
	return posts, nil
}

// GetCountryUserStats retrieves user statistics grouped by country
func GetCountryUserStats(db *gorm.DB) ([]map[string]interface{}, error) {
	var results []map[string]interface{}
	
	rows, err := db.Model(&User{}).
		Select("country, COUNT(*) as user_count, AVG(age) as avg_age").
		Group("country").
		Order("user_count DESC").
		Rows()
	
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	for rows.Next() {
		var country string
		var userCount int64
		var avgAge float64
		
		rows.Scan(&country, &userCount, &avgAge)
		
		results = append(results, map[string]interface{}{
			"country":    country,
			"user_count": userCount,
			"avg_age":    avgAge,
		})
	}
	
	return results, nil
}

// SearchPostsByContent searches posts by content using full-text search
func SearchPostsByContent(db *gorm.DB, query string, limit int) ([]Post, error) {
	var posts []Post
	searchPattern := "%" + query + "%"
	
	result := db.
		Where("title LIKE ? OR content LIKE ?", searchPattern, searchPattern).
		Preload("User").
		Order("created_at DESC").
		Limit(limit).
		Find(&posts)
	
	if result.Error != nil {
		return nil, result.Error
	}
	
	return posts, nil
}

// GetUserRecommendations retrieves user recommendations based on similar interests
func GetUserRecommendations(db *gorm.DB, userID uint, limit int) ([]User, error) {
	var users []User
	
	// Get categories that the user has posted in
	var categories []string
	db.Model(&Post{}).
		Where("user_id = ?", userID).
		Distinct("category").
		Pluck("category", &categories)
	
	if len(categories) == 0 {
		return users, nil
	}
	
	// Find users who have posted in the same categories
	result := db.
		Joins("JOIN posts ON posts.user_id = users.id").
		Where("posts.category IN ?", categories).
		Where("users.id != ?", userID).
		Group("users.id").
		Order("COUNT(DISTINCT posts.category) DESC").
		Limit(limit).
		Find(&users)
	
	if result.Error != nil {
		return nil, result.Error
	}
	
	return users, nil
}

func seedTestData(db *gorm.DB) {
	// Check if data already exists
	var count int64
	db.Model(&User{}).Count(&count)
	if count > 0 {
		fmt.Println("Test data already exists")
		return
	}
	
	// Create users
	users := []User{
		{Username: "alice", Email: "alice@example.com", Age: 28, Country: "USA"},
		{Username: "bob", Email: "bob@example.com", Age: 35, Country: "UK"},
		{Username: "charlie", Email: "charlie@example.com", Age: 22, Country: "Canada"},
		{Username: "diana", Email: "diana@example.com", Age: 30, Country: "USA"},
		{Username: "eve", Email: "eve@example.com", Age: 26, Country: "UK"},
	}
	
	for i := range users {
		db.Create(&users[i])
	}
	
	// Create posts
	posts := []Post{
		{Title: "Introduction to Go", Content: "Go is a great programming language...", UserID: 1, Category: "Technology", ViewCount: 150},
		{Title: "Advanced Go Techniques", Content: "Learn advanced Go programming...", UserID: 1, Category: "Technology", ViewCount: 200},
		{Title: "Web Development with Go", Content: "Building web apps in Go...", UserID: 2, Category: "Technology", ViewCount: 180},
		{Title: "Travel Tips for Europe", Content: "Best places to visit in Europe...", UserID: 2, Category: "Travel", ViewCount: 120},
		{Title: "Cooking 101", Content: "Basic cooking techniques...", UserID: 3, Category: "Lifestyle", ViewCount: 90},
		{Title: "Go Concurrency Patterns", Content: "Mastering goroutines and channels...", UserID: 1, Category: "Technology", ViewCount: 250},
		{Title: "Database Design", Content: "Principles of good database design...", UserID: 4, Category: "Technology", ViewCount: 160},
		{Title: "Fitness Guide", Content: "Stay fit and healthy...", UserID: 5, Category: "Health", ViewCount: 110},
	}
	
	for i := range posts {
		db.Create(&posts[i])
	}
	
	// Create likes
	likes := []Like{
		{UserID: 2, PostID: 1},
		{UserID: 3, PostID: 1},
		{UserID: 4, PostID: 1},
		{UserID: 5, PostID: 1},
		{UserID: 1, PostID: 3},
		{UserID: 3, PostID: 3},
		{UserID: 2, PostID: 6},
		{UserID: 4, PostID: 6},
		{UserID: 5, PostID: 6},
		{UserID: 1, PostID: 4},
		{UserID: 3, PostID: 7},
	}
	
	for i := range likes {
		db.Create(&likes[i])
	}
	
	fmt.Println("Test data seeded successfully")
}

func main() {
	db, err := ConnectDB()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Seed some test data
	fmt.Println("Seeding test data...")
	seedTestData(db)

	// Test GetTopUsersByPostCount
	fmt.Println("\n=== Top Users by Post Count ===")
	topUsers, err := GetTopUsersByPostCount(db, 5)
	if err != nil {
		log.Fatal(err)
	}
	for i, user := range topUsers {
		fmt.Printf("%d. %s - %d posts\n", i+1, user.Username, len(user.Posts))
	}

	// Test GetPostsByCategoryWithUserInfo
	fmt.Println("\n=== Posts in 'Technology' Category (Page 1) ===")
	posts, total, err := GetPostsByCategoryWithUserInfo(db, "Technology", 1, 5)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Total posts in category: %d\n", total)
	for _, post := range posts {
		fmt.Printf("- %s by %s\n", post.Title, post.User.Username)
	}

	// Test GetUserEngagementStats
	fmt.Println("\n=== User Engagement Stats ===")
	stats, err := GetUserEngagementStats(db, 1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Stats: %+v\n", stats)

	// Test GetPopularPostsByLikes
	fmt.Println("\n=== Popular Posts (Last 30 Days) ===")
	popularPosts, err := GetPopularPostsByLikes(db, 30, 5)
	if err != nil {
		log.Fatal(err)
	}
	for i, post := range popularPosts {
		fmt.Printf("%d. %s - %d likes\n", i+1, post.Title, len(post.Likes))
	}

	// Test GetCountryUserStats
	fmt.Println("\n=== User Statistics by Country ===")
	countryStats, err := GetCountryUserStats(db)
	if err != nil {
		log.Fatal(err)
	}
	for _, stat := range countryStats {
		fmt.Printf("%s: %v users, avg age %.1f\n",
			stat["country"], stat["user_count"], stat["avg_age"])
	}

	// Test SearchPostsByContent
	fmt.Println("\n=== Search Posts: 'programming' ===")
	searchResults, err := SearchPostsByContent(db, "programming", 5)
	if err != nil {
		log.Fatal(err)
	}
	for _, post := range searchResults {
		fmt.Printf("- %s\n", post.Title)
	}

	// Test GetUserRecommendations
	fmt.Println("\n=== User Recommendations for User 1 ===")
	recommendations, err := GetUserRecommendations(db, 1, 5)
	if err != nil {
		log.Fatal(err)
	}
	for i, user := range recommendations {
		fmt.Printf("%d. %s\n", i+1, user.Username)
	}
}

// Notes:
// - Use Preload() for loading associations efficiently
// - Use Group(), Having(), Count() for aggregations
// - Use Offset() and Limit() for pagination
// - Use Raw() or Exec() for complex queries when GORM methods aren't sufficient
// - Always consider query performance and indexing
// - Use db.Model() to specify the table without needing a full struct

