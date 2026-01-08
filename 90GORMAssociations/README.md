# 90GORMAssociations - GORM Associations & Relationships

## Overview

This practice module teaches you how to work with database relationships and associations using GORM. You'll build a Blog System that demonstrates one-to-many and many-to-many relationships, along with advanced querying techniques like preloading and association management.

## Challenge: Blog System with Relationships

Build a Blog System using GORM that demonstrates database relationships and associations between models.

## Concepts Covered

- **One-to-Many Relationships**: Users can have multiple posts
- **Many-to-Many Relationships**: Posts can have multiple tags, tags can belong to multiple posts
- **Foreign Keys**: Setting up and managing foreign key constraints
- **Preloading**: Efficiently loading related data to avoid N+1 queries
- **Association API**: Creating, querying, and managing related data
- **Join Tables**: Understanding how GORM handles many-to-many relationships
- **FirstOrCreate**: Finding or creating records to avoid duplicates

## Data Models

### User Model (One-to-Many with Posts)

```go
type User struct {
    ID        uint      `gorm:"primaryKey"`
    Name      string    `gorm:"not null"`
    Email     string    `gorm:"unique;not null"`
    Posts     []Post    `gorm:"foreignKey:UserID"`
    CreatedAt time.Time
    UpdatedAt time.Time
}
```

### Post Model (Belongs to User, Many-to-Many with Tags)

```go
type Post struct {
    ID        uint      `gorm:"primaryKey"`
    Title     string    `gorm:"not null"`
    Content   string    `gorm:"type:text"`
    UserID    uint      `gorm:"not null"`
    User      User      `gorm:"foreignKey:UserID"`
    Tags      []Tag     `gorm:"many2many:post_tags;"`
    CreatedAt time.Time
    UpdatedAt time.Time
}
```

### Tag Model (Many-to-Many with Posts)

```go
type Tag struct {
    ID    uint   `gorm:"primaryKey"`
    Name  string `gorm:"unique;not null"`
    Posts []Post `gorm:"many2many:post_tags;"`
}
```

## Relationship Explanations

### One-to-Many (User → Posts)
- A User can have multiple Posts
- Each Post belongs to exactly one User
- Foreign key `UserID` in Post table references User's ID
- Use `foreignKey` tag to specify the relationship

### Many-to-Many (Post ↔ Tags)
- A Post can have multiple Tags
- A Tag can belong to multiple Posts
- GORM automatically creates a join table `post_tags`
- Use `many2many` tag to specify the join table name

## Required Functions

Implement these functions in the template:

1. **ConnectDB() (*gorm.DB, error)**
   - Establish database connection
   - Auto-migrate all three models

2. **CreateUserWithPosts(db *gorm.DB, user *User) error**
   - Create user with associated posts in one operation

3. **GetUserWithPosts(db *gorm.DB, userID uint) (*User, error)**
   - Retrieve user with all their posts preloaded

4. **CreatePostWithTags(db *gorm.DB, post *Post, tagNames []string) error**
   - Create post and associate it with tags (create tags if needed)

5. **GetPostsByTag(db *gorm.DB, tagName string) ([]Post, error)**
   - Find all posts that have a specific tag

6. **AddTagsToPost(db *gorm.DB, postID uint, tagNames []string) error**
   - Add tags to an existing post

7. **GetPostWithUserAndTags(db *gorm.DB, postID uint) (*Post, error)**
   - Retrieve post with both user and tags preloaded

## Key Learning Points

1. **Association Auto-Creation**: When you create a User with Posts, GORM automatically creates all associated Posts
2. **Preloading**: Use `Preload()` to load associations and avoid N+1 query problems
3. **Association API**: Use `Association()` methods to manage relationships dynamically
4. **FirstOrCreate**: Prevent duplicate tags by finding existing ones first
5. **Join Tables**: GORM handles the complexity of many-to-many join tables automatically
6. **Foreign Key Constraints**: Database enforces referential integrity

## How to Practice

1. Navigate to the `.practice` directory
2. Open `template.go` and complete the TODOs
3. Uncomment the main function code to test your implementation
4. Run the code: `go run template.go`
5. Compare with `solution.go` if you get stuck

## Expected Output

```
Created user Alice Johnson with 2 posts
Fetched user Alice Johnson with 2 posts
Created post 'Go Programming Tips' with 3 tags
Found 1 posts with tag 'golang'
Added tags to existing post
Post 'Go Programming Tips' by Alice Johnson with 3 tags
```

## Testing Requirements

Your solution should:
- ✅ Create users with multiple associated posts
- ✅ Create posts with multiple tags
- ✅ Query users with their posts preloaded
- ✅ Query posts by tag name
- ✅ Add tags to existing posts
- ✅ Load posts with both user and tags
- ✅ Handle foreign key constraints properly
- ✅ Prevent duplicate tags
- ✅ Support many-to-many relationships

## Common Pitfalls

1. **N+1 Queries**: Always use `Preload()` when you need related data
2. **Forgetting Auto-Migration**: Migrate all models including join tables
3. **Duplicate Tags**: Use `FirstOrCreate` to avoid creating duplicate tags
4. **Missing Foreign Keys**: Ensure foreign key fields are properly set
5. **Association Not Loaded**: Check if you've preloaded associations before accessing them

## GORM Preloading Techniques

```go
// Single association
db.Preload("User").Find(&posts)

// Multiple associations
db.Preload("User").Preload("Tags").Find(&posts)

// Nested preloading
db.Preload("Posts.Tags").Find(&users)

// Conditional preloading
db.Preload("Posts", "created_at > ?", someDate).Find(&users)
```

## GORM Association API Examples

```go
// Append associations
db.Model(&post).Association("Tags").Append([]Tag{tag1, tag2})

// Replace associations
db.Model(&post).Association("Tags").Replace([]Tag{tag3, tag4})

// Delete associations
db.Model(&post).Association("Tags").Delete(tag1)

// Clear all associations
db.Model(&post).Association("Tags").Clear()

// Count associations
count := db.Model(&post).Association("Tags").Count()

// Find associations
var tags []Tag
db.Model(&post).Association("Tags").Find(&tags)
```

## Database Schema

GORM will create these tables:

1. **users** - Stores user information
2. **posts** - Stores post information with `user_id` foreign key
3. **tags** - Stores unique tag names
4. **post_tags** - Join table for many-to-many relationship

## Learning Resources

- [GORM Associations](https://gorm.io/docs/associations.html)
- [GORM Belongs To](https://gorm.io/docs/belongs_to.html)
- [GORM Has Many](https://gorm.io/docs/has_many.html)
- [GORM Many To Many](https://gorm.io/docs/many_to_many.html)
- [GORM Preloading](https://gorm.io/docs/preload.html)
- [GORM Association Mode](https://gorm.io/docs/associations.html#Association-Mode)

## Extensions (Optional Challenges)

After completing the basic implementation, try these extensions:

1. **Nested Preloading**: Get users with posts AND tags in one query
2. **Custom Join Table**: Define a custom struct for the join table with additional fields
3. **Has One Relationship**: Add a Profile model with one-to-one relationship to User
4. **Polymorphic Associations**: Make comments that can belong to either posts or users
5. **Batch Loading**: Optimize loading multiple users with their posts
6. **Soft Delete**: Implement soft deletes and see how it affects associations
7. **Aggregation**: Count posts per user or tags per post
8. **Preload Conditions**: Preload only posts created in the last week

## Advanced Topics to Explore

- **Eager Loading vs Lazy Loading**: Understanding performance implications
- **Select Specific Fields**: Load only required fields from associations
- **Join Queries**: Use joins instead of preload for specific cases
- **Association Callbacks**: Hooks that run when associations change
- **Constraint Settings**: Configure foreign key constraints (CASCADE, SET NULL, etc.)

## Next Steps

After completing this module, explore:
- **Advanced Querying**: Complex joins, subqueries, and raw SQL
- **Performance Optimization**: Indexing, query analysis, connection pooling
- **Database Migrations**: Managing schema changes in production
- **Testing Strategies**: Mocking databases, test fixtures, and factories

---

**Note**: This module creates a `blog.db` SQLite database file in the current directory. You can delete it to start fresh.

