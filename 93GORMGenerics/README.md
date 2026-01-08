# 93GORMGenerics - Context-Aware Operations & Modern GORM

## Overview

This final practice module in the GORM series teaches context-aware database operations and modern GORM patterns. You'll build a User & Post Management System that demonstrates production-ready patterns including context support, batch operations, and advanced preloading.

## Challenge: Modern User & Post Management System

Build a comprehensive system using GORM's modern API with full context support, demonstrating enterprise-grade database operations.

## Concepts Covered

- **Context-Aware Operations**: All operations support context for cancellation and timeouts
- **Batch Operations**: Efficient bulk inserts with `CreateInBatches`
- **Advanced Preloading**: Custom conditions and limits per record
- **Upsert Operations**: Handle conflicts with `OnConflict` clauses
- **Result Metadata**: Capture rows affected and operation details
- **Complex Joins**: Multi-table queries with custom conditions
- **Performance Optimization**: Minimize database roundtrips

## Data Models

### User Model (with Company Association)

```go
type User struct {
    ID        uint
    Name      string
    Email     string    // unique
    Age       int       // > 0
    CompanyID *uint     // optional foreign key
    Company   *Company  // belongs to
    Posts     []Post    // has many
    CreatedAt time.Time
    UpdatedAt time.Time
}
```

### Company Model

```go
type Company struct {
    ID          uint
    Name        string  // unique
    Industry    string
    FoundedYear int
    Users       []User  // has many
    CreatedAt   time.Time
}
```

### Post Model

```go
type Post struct {
    ID        uint
    Title     string
    Content   string
    UserID    uint      // foreign key
    User      User      // belongs to
    ViewCount int
    CreatedAt time.Time
    UpdatedAt time.Time
}
```

## Required Functions

Implement these 14 context-aware functions:

### Basic Operations (Context-Aware)

1. **ConnectDB() (*gorm.DB, error)**
   - Connect to SQLite and auto-migrate models

2. **CreateUser(ctx, db, *User) error**
   - Create user with context support

3. **GetUserByID(ctx, db, uint) (*User, error)**
   - Retrieve user with context

4. **UpdateUserAge(ctx, db, uint, int) error**
   - Update specific field with context

5. **DeleteUser(ctx, db, uint) error**
   - Delete user with context

### Batch Operations

6. **CreateUsersInBatches(ctx, db, []User, int) error**
   - Bulk insert with configurable batch size

7. **FindUsersByAgeRange(ctx, db, int, int) ([]User, error)**
   - Range queries with context

### Advanced Features

8. **UpsertUser(ctx, db, *User) error**
   - Insert or update on conflict

9. **CreateUserWithResult(ctx, db, *User) (int64, error)**
   - Return operation metadata

### Enhanced Associations

10. **GetUsersWithCompany(ctx, db) ([]User, error)**
    - Preload company association

11. **GetUsersWithPosts(ctx, db, int) ([]User, error)**
    - Preload posts with limit per user

12. **GetUserWithPostsAndCompany(ctx, db, uint) (*User, error)**
    - Multiple preloads simultaneously

### Complex Queries

13. **SearchUsersInCompany(ctx, db, string) ([]User, error)**
    - Join with filter conditions

14. **GetTopActiveUsers(ctx, db, int) ([]User, error)**
    - Aggregation with joins

## Key Modern GORM Patterns

### 1. Context Support

```go
ctx := context.Background()
db.WithContext(ctx).Create(&user)
```

**Benefits:**
- Request cancellation
- Timeout handling
- Trace propagation
- Request-scoped values

### 2. Batch Operations

```go
db.WithContext(ctx).CreateInBatches(users, 100)
```

**Benefits:**
- Reduced database roundtrips
- Better performance for bulk data
- Configurable batch size

### 3. Advanced Preloading

```go
db.WithContext(ctx).Preload("Posts", func(db *gorm.DB) *gorm.DB {
    return db.Order("created_at DESC").Limit(5)
}).Find(&users)
```

**Features:**
- Custom ordering
- Limit per record
- Conditional loading

### 4. OnConflict Handling

```go
db.WithContext(ctx).Clauses(clause.OnConflict{
    Columns:   []clause.Column{{Name: "email"}},
    DoUpdates: clause.AssignmentColumns([]string{"name", "age"}),
}).Create(&user)
```

**Use Cases:**
- Upsert operations
- Handling unique constraints
- Idempotent inserts

### 5. Result Metadata

```go
result := db.WithContext(ctx).Create(&user)
rowsAffected := result.RowsAffected
err := result.Error
```

## How to Practice

1. Navigate to the `.practice` directory
2. Open `template.go` and complete the TODOs
3. Uncomment the main function code to test
4. Run: `go run template.go`
5. Compare with `solution.go` if needed

## Expected Output

```
Created companies
Created users in batches
Created posts
Found user: Alice (Age: 30)
Updated user age to 31
Found 3 users in age range 28-35
Upserted user (handled email conflict)
Created user, rows affected: 1
Found 4 users with companies:
  - Alice Updated works at TechCorp
  - Bob works at TechCorp
  - Charlie works at FinanceInc
  - David works at TechCorp

Users with posts (max 2 per user):
  - Alice Updated has 1 post(s)
  - Bob has 1 post(s)

Full user info for Alice Updated:
  Company: TechCorp (Technology)
  Posts: 2

3 users work at TechCorp

Top 3 active users:
  1. Alice Updated (2 posts)
  2. Bob (1 posts)

Deleted user with ID 5
```

## Testing Requirements

Your solution should pass all 16 tests:
- ✅ Database connection
- ✅ Create user with context
- ✅ Get user by ID
- ✅ Handle not found errors
- ✅ Update specific fields
- ✅ Delete operations
- ✅ Batch insertions
- ✅ Range queries
- ✅ Upsert with conflict handling
- ✅ Result metadata capture
- ✅ Preload associations
- ✅ Custom preload conditions
- ✅ Multiple preloads
- ✅ Join with filters
- ✅ Aggregation queries
- ✅ Context cancellation

## Performance Benefits

### Context Support
- **Cancellation**: Stop long-running queries
- **Timeouts**: Prevent resource exhaustion
- **Tracing**: Track requests across services

### Batch Operations
- **Reduced Latency**: Fewer database roundtrips
- **Better Throughput**: Process thousands of records efficiently
- **Resource Efficiency**: Lower connection overhead

### Smart Preloading
- **N+1 Prevention**: Load associations in one query
- **Selective Loading**: Only load what you need
- **Custom Conditions**: Filter associations efficiently

## Production Patterns

### 1. Request Context Pattern

```go
func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
    ctx := r.Context() // Request context with timeout
    
    user := &User{...}
    if err := CreateUser(ctx, db, user); err != nil {
        if err == context.DeadlineExceeded {
            http.Error(w, "Request timeout", 504)
            return
        }
        http.Error(w, err.Error(), 500)
        return
    }
}
```

### 2. Transaction Pattern

```go
err := db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
    if err := CreateUser(ctx, tx, user); err != nil {
        return err // Rollback
    }
    if err := CreatePost(ctx, tx, post); err != nil {
        return err // Rollback
    }
    return nil // Commit
})
```

### 3. Timeout Pattern

```go
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

users, err := GetUsersWithCompany(ctx, db)
if err == context.DeadlineExceeded {
    log.Println("Query timed out")
}
```

## Common Pitfalls & Solutions

### 1. Forgetting Context
```go
// ❌ Bad: No context support
db.Create(&user)

// ✅ Good: Context-aware
db.WithContext(ctx).Create(&user)
```

### 2. Not Checking Errors
```go
// ❌ Bad: Ignoring errors
db.WithContext(ctx).Create(&user)

// ✅ Good: Proper error handling
if err := db.WithContext(ctx).Create(&user).Error; err != nil {
    return err
}
```

### 3. Inefficient Batch Operations
```go
// ❌ Bad: Individual inserts
for _, user := range users {
    db.WithContext(ctx).Create(&user)
}

// ✅ Good: Batch insert
db.WithContext(ctx).CreateInBatches(users, 100)
```

## Learning Resources

- [GORM Documentation](https://gorm.io/docs/)
- [GORM Context Support](https://gorm.io/docs/context.html)
- [GORM Advanced Query](https://gorm.io/docs/advanced_query.html)
- [Go Context Package](https://pkg.go.dev/context)
- [Database Best Practices](https://www.alexedwards.net/blog/organising-database-access)

## Extensions (Optional Challenges)

1. **Connection Pooling**: Configure max connections and idle timeout
2. **Prepared Statements**: Use prepared statement mode
3. **Read Replicas**: Route reads to replicas
4. **Middleware**: Add logging and metrics middleware
5. **Soft Deletes**: Implement soft delete with DeletedAt
6. **Optimistic Locking**: Use version field for concurrency
7. **Custom Types**: Create custom scanner/valuer
8. **Hooks**: Implement Before/After callbacks
9. **Scopes**: Create reusable query scopes
10. **Database Sharding**: Partition data across databases

## Migration from Previous Modules

### From 89GORMCrud (Basic CRUD)
- Add context to all operations
- Use batch operations for bulk data
- Implement proper error handling

### From 90GORMAssociations (Associations)
- Add custom preload conditions
- Use joins for complex queries
- Optimize association loading

### From 91GORMMigrations (Migrations)
- Use context in migration scripts
- Batch process data transformations
- Handle timeouts in long migrations

### From 92GORMAdvancedQueries (Advanced Queries)
- Add context to all aggregations
- Optimize with proper preloading
- Use result metadata for monitoring

## Complete GORM Series 🎓

Congratulations! You've completed all GORM modules:

1. **89GORMCrud** - Basic CRUD operations ✅
2. **90GORMAssociations** - Relationships & associations ✅
3. **91GORMMigrations** - Schema evolution & migrations ✅
4. **92GORMAdvancedQueries** - Analytics & complex queries ✅
5. **93GORMGenerics** - Context-aware & modern patterns ✅

You now have comprehensive knowledge of GORM from basics to advanced production patterns!

---

**Note**: This module creates a `generics.db` SQLite database file. Delete it to start fresh.

