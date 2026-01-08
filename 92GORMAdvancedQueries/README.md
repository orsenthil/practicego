# 92GORMAdvancedQueries - Advanced Queries & Analytics

## Overview

This practice module teaches advanced querying techniques, aggregations, and complex data analysis using GORM. You'll build a Social Media Analytics System that demonstrates how to write efficient, complex queries for real-world applications.

## Challenge: Social Media Analytics System

Build a Social Media Analytics System using GORM that demonstrates advanced querying techniques, aggregations, and complex data analysis.

## Concepts Covered

- **Complex Queries**: Advanced filtering, sorting, and joins
- **Aggregations**: GROUP BY, COUNT, SUM, AVG operations
- **Pagination**: Efficient page-based data retrieval
- **Subqueries**: Nested and correlated subqueries
- **Full-Text Search**: Content-based search queries
- **Query Optimization**: Efficient data retrieval patterns
- **Analytics**: User engagement and recommendation algorithms
- **Performance**: Using appropriate indexes and query patterns

## Data Models

### User Model

```go
type User struct {
    ID        uint
    Username  string
    Email     string
    Age       int
    Country   string
    CreatedAt time.Time
    Posts     []Post  // One-to-many
    Likes     []Like  // One-to-many
}
```

### Post Model

```go
type Post struct {
    ID          uint
    Title       string
    Content     string
    UserID      uint
    User        User    // Belongs to
    Category    string
    ViewCount   int
    IsPublished bool
    CreatedAt   time.Time
    UpdatedAt   time.Time
    Likes       []Like  // One-to-many
}
```

### Like Model

```go
type Like struct {
    ID        uint
    UserID    uint
    PostID    uint
    User      User  // Belongs to
    Post      Post  // Belongs to
    CreatedAt time.Time
}
```

## Required Functions

Implement these eight advanced query functions:

### 1. GetTopUsersByPostCount
**Purpose**: Find users with the most posts (leaderboard)

**Query Pattern**: Aggregation with JOIN and GROUP BY
```sql
SELECT users.*, COUNT(posts.id) as post_count
FROM users
LEFT JOIN posts ON users.id = posts.user_id
GROUP BY users.id
ORDER BY post_count DESC
LIMIT ?
```

### 2. GetPostsByCategoryWithUserInfo
**Purpose**: Paginated posts with category filter and user details

**Query Pattern**: Filtering, pagination, and preloading
```sql
-- Count query
SELECT COUNT(*) FROM posts WHERE category = ?

-- Data query
SELECT * FROM posts 
WHERE category = ? 
ORDER BY created_at DESC 
LIMIT ? OFFSET ?
```

### 3. GetUserEngagementStats
**Purpose**: Calculate engagement metrics for a user

**Metrics**:
- Total posts created
- Total likes received (on their posts)
- Total likes given (by them)
- Average post view count

**Query Pattern**: Multiple aggregations and subqueries

### 4. GetPopularPostsByLikes
**Purpose**: Find trending posts based on like count

**Query Pattern**: Time-based filtering with aggregation
```sql
SELECT posts.*, COUNT(likes.id) as like_count
FROM posts
LEFT JOIN likes ON posts.id = likes.post_id
WHERE posts.created_at >= ?
GROUP BY posts.id
ORDER BY like_count DESC
LIMIT ?
```

### 5. GetCountryUserStats
**Purpose**: User demographics by country

**Query Pattern**: Aggregation with GROUP BY
```sql
SELECT country, COUNT(*) as user_count, AVG(age) as avg_age
FROM users
GROUP BY country
ORDER BY user_count DESC
```

### 6. SearchPostsByContent
**Purpose**: Full-text search in posts

**Query Pattern**: LIKE queries with wildcards
```sql
SELECT * FROM posts
WHERE title LIKE ? OR content LIKE ?
ORDER BY created_at DESC
LIMIT ?
```

### 7. GetUserRecommendations
**Purpose**: Recommend users with similar interests

**Algorithm**:
1. Find categories where target user has posted
2. Find other users who post in same categories
3. Rank by number of shared categories

**Query Pattern**: Subquery with IN clause and GROUP BY

## Key Learning Points

1. **Aggregation Functions**: COUNT(), AVG(), SUM(), MIN(), MAX()
2. **GROUP BY**: Grouping data for analysis
3. **JOIN Operations**: Combining data from multiple tables
4. **Pagination**: OFFSET and LIMIT for large datasets
5. **Preloading**: Avoiding N+1 query problems
6. **Subqueries**: Nested SELECT statements
7. **Time-Based Filtering**: Working with date ranges
8. **Query Optimization**: Minimizing database roundtrips

## How to Practice

1. Navigate to the `.practice` directory
2. Open `template.go` and complete the TODOs
3. Uncomment the main function code to test
4. Run the code: `go run template.go`
5. Compare with `solution.go` if you get stuck

## Expected Output

```
Seeding test data...
Test data seeded successfully

=== Top Users by Post Count ===
1. alice - 3 posts
2. bob - 2 posts
3. diana - 1 posts

=== Posts in 'Technology' Category (Page 1) ===
Total posts in category: 5
- Go Concurrency Patterns by alice
- Database Design by diana
- Advanced Go Techniques by alice
- Web Development with Go by bob
- Introduction to Go by alice

=== User Engagement Stats ===
Stats: map[avg_post_views:200 total_likes_given:2 total_likes_received:7 total_posts:3]

=== Popular Posts (Last 30 Days) ===
1. Introduction to Go - 4 likes
2. Go Concurrency Patterns - 3 likes
3. Web Development with Go - 2 likes

=== User Statistics by Country ===
USA: 2 users, avg age 29.0
UK: 2 users, avg age 30.5
Canada: 1 users, avg age 22.0

=== Search Posts: 'programming' ===
- Advanced Go Techniques
- Introduction to Go

=== User Recommendations for User 1 ===
1. bob
2. diana
```

## Testing Requirements

Your solution should:
- ✅ Retrieve top users by post count with aggregation
- ✅ Paginate posts with proper offset calculation
- ✅ Calculate user engagement statistics accurately
- ✅ Filter popular posts by time period
- ✅ Group user statistics by country
- ✅ Search posts by content (title or body)
- ✅ Recommend users based on shared interests
- ✅ Preload associations to avoid N+1 queries
- ✅ Handle edge cases (no data, empty results)

## Advanced Query Techniques

### Aggregation with GROUP BY

```go
db.Model(&User{}).
    Select("country, COUNT(*) as user_count, AVG(age) as avg_age").
    Group("country").
    Order("user_count DESC").
    Scan(&results)
```

### Pagination Pattern

```go
offset := (page - 1) * pageSize
db.Where("category = ?", category).
    Offset(offset).
    Limit(pageSize).
    Find(&posts)
```

### JOIN with Aggregation

```go
db.Joins("LEFT JOIN posts ON posts.user_id = users.id").
    Group("users.id").
    Order("COUNT(posts.id) DESC").
    Find(&users)
```

### Time-Based Filtering

```go
cutoff := time.Now().AddDate(0, 0, -days)
db.Where("created_at >= ?", cutoff).Find(&posts)
```

### Subquery Pattern

```go
// Get categories user has posted in
var categories []string
db.Model(&Post{}).
    Where("user_id = ?", userID).
    Distinct("category").
    Pluck("category", &categories)

// Find users who post in those categories
db.Joins("JOIN posts ON posts.user_id = users.id").
    Where("posts.category IN ?", categories).
    Where("users.id != ?", userID).
    Group("users.id").
    Find(&users)
```

## Performance Optimization Tips

### 1. Use Preload for Associations
```go
// ❌ Bad: N+1 query problem
db.Find(&posts)
for _, post := range posts {
    db.First(&post.User, post.UserID) // Extra query per post!
}

// ✅ Good: Single additional query
db.Preload("User").Find(&posts)
```

### 2. Use Select for Specific Fields
```go
// Only load needed fields
db.Select("id, username, email").Find(&users)
```

### 3. Use Count Separately
```go
// Get count and data in separate queries
var total int64
db.Model(&Post{}).Where("category = ?", cat).Count(&total)
db.Where("category = ?", cat).Limit(10).Find(&posts)
```

### 4. Use Indexes
```sql
CREATE INDEX idx_posts_category ON posts(category);
CREATE INDEX idx_posts_created_at ON posts(created_at);
CREATE INDEX idx_likes_post_id ON likes(post_id);
```

## Common Pitfalls

1. **N+1 Queries**: Always use Preload() for associations
2. **Missing Pagination**: Can cause memory issues with large datasets
3. **Inefficient Counting**: Count and data queries should be separate
4. **Ignoring NULL Values**: Use LEFT JOIN for optional relationships
5. **String Matching Performance**: LIKE queries can be slow without indexes
6. **Time Zone Issues**: Be consistent with time zone handling

## Query Performance Comparison

```go
// Slow: Multiple queries
for _, user := range users {
    db.Where("user_id = ?", user.ID).Find(&posts)  // N queries
}

// Fast: Single query with JOIN
db.Preload("Posts").Find(&users)  // 2 queries total
```

## Real-World Applications

### Analytics Dashboard
- User growth by country
- Content engagement metrics
- Popular content identification

### Recommendation Engine
- Similar user suggestions
- Content-based recommendations
- Collaborative filtering

### Search Functionality
- Full-text search
- Faceted search (filters + search)
- Auto-complete suggestions

## Learning Resources

- [GORM Queries](https://gorm.io/docs/query.html)
- [GORM Advanced Query](https://gorm.io/docs/advanced_query.html)
- [GORM Raw SQL](https://gorm.io/docs/sql_builder.html)
- [SQL Aggregation Functions](https://www.postgresql.org/docs/current/functions-aggregate.html)
- [Database Indexing](https://use-the-index-luke.com/)

## Extensions (Optional Challenges)

After completing the basic implementation:

1. **Advanced Analytics**: Calculate engagement rate (likes/views ratio)
2. **Trending Algorithm**: Combine recency and popularity
3. **Full-Text Search**: Implement ranking/relevance scoring
4. **Caching**: Add Redis caching for popular queries
5. **Aggregation Pipeline**: Calculate multiple metrics efficiently
6. **Time Series Analysis**: Track metrics over time
7. **Recommendation Improvements**: Use collaborative filtering
8. **Query Builder**: Create reusable query components
9. **Batch Operations**: Process large datasets efficiently
10. **Export Functionality**: Generate reports from queries

## Advanced Topics

- **Window Functions**: Running totals, rankings
- **CTE (Common Table Expressions)**: Complex recursive queries
- **EXPLAIN ANALYZE**: Query performance analysis
- **Query Caching**: Result caching strategies
- **Materialized Views**: Pre-computed aggregations
- **Denormalization**: Trading space for speed

## Next Steps

After completing this module:
- **Performance Tuning**: Learn query optimization
- **NoSQL Alternatives**: Explore document databases
- **GraphQL**: Alternative query approaches
- **Elasticsearch**: Advanced search capabilities

---

**Note**: This module creates a `social.db` SQLite database file. Delete it to start fresh.

