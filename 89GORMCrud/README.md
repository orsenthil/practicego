# 89GORMCrud - GORM CRUD Operations

## Overview

This practice module introduces GORM, the fantastic ORM library for Golang. GORM provides a complete solution for working with databases using an object-relational mapping approach. In this module, you'll learn to implement fundamental CRUD (Create, Read, Update, Delete) operations for a User Management System.

## Challenge: User Management System

Build a User Management System using GORM that demonstrates fundamental database operations with SQLite.

## Concepts Covered

- **Database Connection**: Setting up and configuring a GORM connection to SQLite
- **Auto Migration**: Automatically creating database tables from Go structs
- **Create Operations**: Inserting new records into the database
- **Read Operations**: Querying single records and retrieving all records
- **Update Operations**: Modifying existing records in the database
- **Delete Operations**: Removing records from the database
- **Error Handling**: Properly handling database errors

## Data Model

```go
type User struct {
    ID        uint      `gorm:"primaryKey"`
    Name      string    `gorm:"not null"`
    Email     string    `gorm:"unique;not null"`
    Age       int       `gorm:"check:age > 0"`
    CreatedAt time.Time
    UpdatedAt time.Time
}
```

### Field Tags Explained

- `primaryKey` - Marks ID as the primary key
- `not null` - Ensures the field cannot be null
- `unique` - Ensures email addresses are unique
- `check:age > 0` - Adds a check constraint for positive ages
- `CreatedAt/UpdatedAt` - Automatically managed by GORM

## Required Functions

Implement these functions in the template:

1. **ConnectDB() (*gorm.DB, error)** - Establish database connection and run migrations
2. **CreateUser(db *gorm.DB, user *User) error** - Create a new user
3. **GetUserByID(db *gorm.DB, id uint) (*User, error)** - Retrieve user by ID
4. **GetAllUsers(db *gorm.DB) ([]User, error)** - Retrieve all users
5. **UpdateUser(db *gorm.DB, user *User) error** - Update existing user
6. **DeleteUser(db *gorm.DB, id uint) error** - Delete user by ID

## Key Learning Points

1. **GORM Configuration**: Learn to configure GORM with different database drivers
2. **Auto Migration**: Understand how GORM creates/updates tables automatically
3. **Struct Tags**: Use GORM tags to define database constraints and behaviors
4. **CRUD Patterns**: Implement standard database operation patterns
5. **Error Handling**: Handle common database errors like record not found
6. **Timestamp Management**: Leverage GORM's automatic timestamp handling

## How to Practice

1. Navigate to the `.practice` directory
2. Open `template.go` and complete the TODOs
3. Uncomment the main function code to test your implementation
4. Run the code: `go run template.go`
5. Compare with `solution.go` if you get stuck

## Expected Output

```
Created user with ID: 1
Fetched user: &{ID:1 Name:John Doe Email:john@example.com Age:30 CreatedAt:... UpdatedAt:...}
User updated successfully
Total users: 1
User deleted successfully
```

## Testing Requirements

Your solution should:
- ✅ Successfully connect to SQLite database
- ✅ Create database table with proper constraints
- ✅ Insert new users with validation
- ✅ Query users by ID and handle not found cases
- ✅ Retrieve all users from database
- ✅ Update user information correctly
- ✅ Delete users by ID
- ✅ Handle errors appropriately

## Common Pitfalls

1. **Forgetting Auto Migration**: Always call `AutoMigrate()` after connecting
2. **Ignoring Errors**: Always check and handle error returns
3. **Pointer vs Value**: GORM works with pointers for struct operations
4. **Record Not Found**: Use `gorm.ErrRecordNotFound` to check if a record exists

## Learning Resources

- [GORM Official Documentation](https://gorm.io/docs/)
- [GORM CRUD Interface](https://gorm.io/docs/create.html)
- [GORM SQLite Driver](https://github.com/gorm-io/sqlite)
- [Database Migration Guide](https://gorm.io/docs/migration.html)

## Extensions (Optional Challenges)

After completing the basic implementation, try these extensions:

1. **Query Filters**: Add a function to find users by name or email
2. **Batch Operations**: Create multiple users at once
3. **Soft Deletes**: Implement soft delete functionality
4. **Relationships**: Add a Profile struct with one-to-one relationship
5. **Transactions**: Wrap operations in database transactions
6. **Pagination**: Implement pagination for GetAllUsers

## Next Steps

After completing this module, explore:
- **Advanced GORM Features**: Associations, Hooks, and Scopes
- **Database Migrations**: Version-controlled schema changes
- **Performance Optimization**: Preloading, Joins, and Indexes
- **Testing**: Writing unit tests for GORM operations

---

**Note**: This module creates a `test.db` SQLite database file in the current directory. You can delete it to start fresh.





