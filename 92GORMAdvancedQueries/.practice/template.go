package main

import (
	"time"

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

// ConnectDB establishes a connection to the SQLite database with auto-migration
func ConnectDB() (*gorm.DB, error) {
	// TODO: Implement database connection with auto-migration
	// Hint: Use gorm.Open with sqlite.Open("social.db")
	// Auto-migrate all three models: User, Post, Like
	return nil, nil
}

// GetTopUsersByPostCount retrieves users with the most posts
func GetTopUsersByPostCount(db *gorm.DB, limit int) ([]User, error) {
	// TODO: Implement top users by post count aggregation
	// Hint: Use db.Model(&User{}).Select("users.*, COUNT(posts.id) as post_count")
	// Join with posts, group by user, order by count, preload Posts
	return nil, nil
}

// GetPostsByCategoryWithUserInfo retrieves posts by category with pagination and user info
func GetPostsByCategoryWithUserInfo(db *gorm.DB, category string, page, pageSize int) ([]Post, int64, error) {
	// TODO: Implement paginated posts retrieval with user info
	// Hint: Calculate offset = (page - 1) * pageSize
	// Use db.Where("category = ?").Preload("User").Offset().Limit()
	// Also get total count with db.Model(&Post{}).Where("category = ?").Count()
	return nil, 0, nil
}

// GetUserEngagementStats calculates engagement statistics for a user
func GetUserEngagementStats(db *gorm.DB, userID uint) (map[string]interface{}, error) {
	// TODO: Implement user engagement statistics
	// Hint: Calculate:
	// - total_posts: Count of posts by user
	// - total_likes_received: Count of likes on user's posts
	// - total_likes_given: Count of likes by user
	// - avg_post_views: Average view count of user's posts
	// Return as map[string]interface{}
	return nil, nil
}

// GetPopularPostsByLikes retrieves popular posts by likes in a time period
func GetPopularPostsByLikes(db *gorm.DB, days int, limit int) ([]Post, error) {
	// TODO: Implement popular posts by likes
	// Hint: Use time.Now().AddDate(0, 0, -days) to get cutoff date
	// Join with likes, filter by created_at >= cutoff, group by post
	// Order by like count DESC, preload User and Likes
	return nil, nil
}

// GetCountryUserStats retrieves user statistics grouped by country
func GetCountryUserStats(db *gorm.DB) ([]map[string]interface{}, error) {
	// TODO: Implement country-based user statistics
	// Hint: Use db.Model(&User{}).Select("country, COUNT(*) as user_count, AVG(age) as avg_age")
	// Group by country, order by user_count DESC
	// Scan into []map[string]interface{}
	return nil, nil
}

// SearchPostsByContent searches posts by content using full-text search
func SearchPostsByContent(db *gorm.DB, query string, limit int) ([]Post, error) {
	// TODO: Implement full-text search
	// Hint: Use db.Where("title LIKE ? OR content LIKE ?", "%"+query+"%", "%"+query+"%")
	// Preload User, limit results, order by created_at DESC
	return nil, nil
}

// GetUserRecommendations retrieves user recommendations based on similar interests
func GetUserRecommendations(db *gorm.DB, userID uint, limit int) ([]User, error) {
	// TODO: Implement user recommendations algorithm
	// Hint: Find users who have posted in the same categories as the given user
	// Exclude the given user, group by user, order by shared categories count
	// Use subquery to get categories of user's posts
	return nil, nil
}

func main() {
	// TODO: Uncomment and complete this section when you're ready to test
	/*
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
	*/
}

// Notes:
// - Use Preload() for loading associations efficiently
// - Use Group(), Having(), Count() for aggregations
// - Use Offset() and Limit() for pagination
// - Use Raw() or Exec() for complex queries when GORM methods aren't sufficient
// - Always consider query performance and indexing
// - Use db.Model() to specify the table without needing a full struct

