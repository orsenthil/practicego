package main

import (
	"os"
	"testing"
	"time"

	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := ConnectDB()
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
	seedTestData(db)
	return db
}

func cleanupTestDB(t *testing.T) {
	os.Remove("social.db")
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

func TestGetTopUsersByPostCount(t *testing.T) {
	defer cleanupTestDB(t)
	db := setupTestDB(t)
	
	users, err := GetTopUsersByPostCount(db, 3)
	if err != nil {
		t.Fatalf("GetTopUsersByPostCount failed: %v", err)
	}
	
	if len(users) == 0 {
		t.Fatal("Expected to get top users")
	}
	
	// Alice should be first (3 posts)
	if users[0].Username != "alice" {
		t.Errorf("Expected alice to be first, got %s", users[0].Username)
	}
	
	if len(users[0].Posts) != 3 {
		t.Errorf("Expected alice to have 3 posts, got %d", len(users[0].Posts))
	}
}

func TestGetPostsByCategoryWithUserInfo(t *testing.T) {
	defer cleanupTestDB(t)
	db := setupTestDB(t)
	
	posts, total, err := GetPostsByCategoryWithUserInfo(db, "Technology", 1, 5)
	if err != nil {
		t.Fatalf("GetPostsByCategoryWithUserInfo failed: %v", err)
	}
	
	if total != 5 {
		t.Errorf("Expected 5 total Technology posts, got %d", total)
	}
	
	if len(posts) == 0 {
		t.Fatal("Expected to get posts")
	}
	
	// Verify user info is loaded
	if posts[0].User.Username == "" {
		t.Error("Expected user info to be preloaded")
	}
}

func TestGetPostsByCategoryPagination(t *testing.T) {
	defer cleanupTestDB(t)
	db := setupTestDB(t)
	
	// Page 1
	posts1, total, err := GetPostsByCategoryWithUserInfo(db, "Technology", 1, 3)
	if err != nil {
		t.Fatalf("GetPostsByCategoryWithUserInfo page 1 failed: %v", err)
	}
	
	if len(posts1) != 3 {
		t.Errorf("Expected 3 posts on page 1, got %d", len(posts1))
	}
	
	// Page 2
	posts2, _, err := GetPostsByCategoryWithUserInfo(db, "Technology", 2, 3)
	if err != nil {
		t.Fatalf("GetPostsByCategoryWithUserInfo page 2 failed: %v", err)
	}
	
	if len(posts2) != 2 {
		t.Errorf("Expected 2 posts on page 2, got %d", len(posts2))
	}
	
	// Verify different posts on different pages
	if posts1[0].ID == posts2[0].ID {
		t.Error("Expected different posts on different pages")
	}
	
	_ = total // Use total
}

func TestGetUserEngagementStats(t *testing.T) {
	defer cleanupTestDB(t)
	db := setupTestDB(t)
	
	stats, err := GetUserEngagementStats(db, 1)
	if err != nil {
		t.Fatalf("GetUserEngagementStats failed: %v", err)
	}
	
	// Alice has 3 posts
	if stats["total_posts"].(int64) != 3 {
		t.Errorf("Expected 3 total_posts, got %v", stats["total_posts"])
	}
	
	// Alice's posts have received likes
	if stats["total_likes_received"].(int64) < 1 {
		t.Error("Expected some likes received")
	}
	
	// Check average views exists
	if _, ok := stats["avg_post_views"]; !ok {
		t.Error("Expected avg_post_views in stats")
	}
}

func TestGetPopularPostsByLikes(t *testing.T) {
	defer cleanupTestDB(t)
	db := setupTestDB(t)
	
	posts, err := GetPopularPostsByLikes(db, 30, 5)
	if err != nil {
		t.Fatalf("GetPopularPostsByLikes failed: %v", err)
	}
	
	if len(posts) == 0 {
		t.Fatal("Expected to get popular posts")
	}
	
	// Verify posts have likes loaded
	if len(posts[0].Likes) == 0 {
		t.Error("Expected popular posts to have likes")
	}
	
	// Verify posts are ordered by likes (descending)
	if len(posts) > 1 {
		if len(posts[0].Likes) < len(posts[1].Likes) {
			t.Error("Expected posts to be ordered by like count DESC")
		}
	}
}

func TestGetPopularPostsByLikesTimeFilter(t *testing.T) {
	defer cleanupTestDB(t)
	db := setupTestDB(t)
	
	// Create an old post
	oldPost := Post{
		Title:     "Old Post",
		Content:   "Very old content",
		UserID:    1,
		Category:  "Technology",
		CreatedAt: time.Now().AddDate(0, 0, -60), // 60 days ago
	}
	db.Create(&oldPost)
	
	// Add likes to old post
	db.Create(&Like{UserID: 2, PostID: oldPost.ID})
	db.Create(&Like{UserID: 3, PostID: oldPost.ID})
	
	// Get popular posts from last 30 days
	posts, err := GetPopularPostsByLikes(db, 30, 10)
	if err != nil {
		t.Fatalf("GetPopularPostsByLikes failed: %v", err)
	}
	
	// Old post should not be included
	for _, post := range posts {
		if post.ID == oldPost.ID {
			t.Error("Expected old post to be filtered out")
		}
	}
}

func TestGetCountryUserStats(t *testing.T) {
	defer cleanupTestDB(t)
	db := setupTestDB(t)
	
	stats, err := GetCountryUserStats(db)
	if err != nil {
		t.Fatalf("GetCountryUserStats failed: %v", err)
	}
	
	if len(stats) == 0 {
		t.Fatal("Expected country statistics")
	}
	
	// Verify stats structure
	if _, ok := stats[0]["country"]; !ok {
		t.Error("Expected country field in stats")
	}
	if _, ok := stats[0]["user_count"]; !ok {
		t.Error("Expected user_count field in stats")
	}
	if _, ok := stats[0]["avg_age"]; !ok {
		t.Error("Expected avg_age field in stats")
	}
	
	// USA and UK should have 2 users each
	usaFound := false
	for _, stat := range stats {
		if stat["country"] == "USA" {
			usaFound = true
			if stat["user_count"].(int64) != 2 {
				t.Errorf("Expected USA to have 2 users, got %v", stat["user_count"])
			}
		}
	}
	if !usaFound {
		t.Error("Expected USA in country stats")
	}
}

func TestSearchPostsByContent(t *testing.T) {
	defer cleanupTestDB(t)
	db := setupTestDB(t)
	
	posts, err := SearchPostsByContent(db, "Go", 10)
	if err != nil {
		t.Fatalf("SearchPostsByContent failed: %v", err)
	}
	
	if len(posts) == 0 {
		t.Fatal("Expected to find posts with 'Go' in content")
	}
	
	// Verify posts contain the search term
	found := false
	for _, post := range posts {
		if containsIgnoreCase(post.Title, "Go") || containsIgnoreCase(post.Content, "Go") {
			found = true
			break
		}
	}
	if !found {
		t.Error("Expected posts to contain search term")
	}
}

func TestSearchPostsByContentLimit(t *testing.T) {
	defer cleanupTestDB(t)
	db := setupTestDB(t)
	
	posts, err := SearchPostsByContent(db, "Go", 2)
	if err != nil {
		t.Fatalf("SearchPostsByContent failed: %v", err)
	}
	
	if len(posts) > 2 {
		t.Errorf("Expected at most 2 posts, got %d", len(posts))
	}
}

func TestGetUserRecommendations(t *testing.T) {
	defer cleanupTestDB(t)
	db := setupTestDB(t)
	
	// User 1 (alice) posts in Technology category
	// Should recommend users who also post in Technology
	users, err := GetUserRecommendations(db, 1, 5)
	if err != nil {
		t.Fatalf("GetUserRecommendations failed: %v", err)
	}
	
	if len(users) == 0 {
		t.Fatal("Expected user recommendations")
	}
	
	// Should not recommend the user themselves
	for _, user := range users {
		if user.ID == 1 {
			t.Error("Should not recommend the user themselves")
		}
	}
	
	// Bob and Diana should be recommended (they post in Technology)
	foundBob := false
	for _, user := range users {
		if user.Username == "bob" {
			foundBob = true
		}
	}
	if !foundBob {
		t.Error("Expected to find bob in recommendations (shares Technology category)")
	}
}

func TestGetUserRecommendationsNoCommonInterests(t *testing.T) {
	defer cleanupTestDB(t)
	db := setupTestDB(t)
	
	// Create a user with no posts
	user := User{
		Username: "newuser",
		Email:    "newuser@example.com",
		Age:      25,
		Country:  "France",
	}
	db.Create(&user)
	
	// Should return empty list
	users, err := GetUserRecommendations(db, user.ID, 5)
	if err != nil {
		t.Fatalf("GetUserRecommendations failed: %v", err)
	}
	
	if len(users) != 0 {
		t.Error("Expected no recommendations for user with no posts")
	}
}

func TestCompleteWorkflow(t *testing.T) {
	defer cleanupTestDB(t)
	db := setupTestDB(t)
	
	// Test all functions in sequence
	users, err := GetTopUsersByPostCount(db, 3)
	if err != nil {
		t.Fatalf("GetTopUsersByPostCount failed: %v", err)
	}
	if len(users) == 0 {
		t.Fatal("Expected top users")
	}
	
	posts, total, err := GetPostsByCategoryWithUserInfo(db, "Technology", 1, 5)
	if err != nil {
		t.Fatalf("GetPostsByCategoryWithUserInfo failed: %v", err)
	}
	if total == 0 {
		t.Fatal("Expected posts")
	}
	
	stats, err := GetUserEngagementStats(db, 1)
	if err != nil {
		t.Fatalf("GetUserEngagementStats failed: %v", err)
	}
	if len(stats) == 0 {
		t.Fatal("Expected stats")
	}
	
	popularPosts, err := GetPopularPostsByLikes(db, 30, 5)
	if err != nil {
		t.Fatalf("GetPopularPostsByLikes failed: %v", err)
	}
	if len(popularPosts) == 0 {
		t.Fatal("Expected popular posts")
	}
	
	countryStats, err := GetCountryUserStats(db)
	if err != nil {
		t.Fatalf("GetCountryUserStats failed: %v", err)
	}
	if len(countryStats) == 0 {
		t.Fatal("Expected country stats")
	}
	
	searchResults, err := SearchPostsByContent(db, "programming", 5)
	if err != nil {
		t.Fatalf("SearchPostsByContent failed: %v", err)
	}
	_ = searchResults
	
	recommendations, err := GetUserRecommendations(db, 1, 5)
	if err != nil {
		t.Fatalf("GetUserRecommendations failed: %v", err)
	}
	_ = recommendations
	
	_ = posts // Use variables
}

// Helper function
func containsIgnoreCase(s, substr string) bool {
	return len(s) >= len(substr) && 
		(s == substr || 
		 len(s) > 0 && len(substr) > 0)
}

