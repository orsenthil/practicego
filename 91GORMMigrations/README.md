# 91GORMMigrations - Database Migrations & Schema Evolution

## Overview

This practice module teaches you how to manage database schema changes over time using migrations. You'll build an E-commerce System that evolves through multiple schema versions, demonstrating version-controlled database migrations, rollbacks, and data seeding with GORM.

## Challenge: E-commerce System with Migrations

Build an E-commerce System using GORM that demonstrates database migrations, schema evolution, and version control for database changes.

## Concepts Covered

- **Database Migrations**: Version-controlled schema changes
- **Schema Evolution**: Adding, modifying, and removing database structures
- **Migration Rollbacks**: Reverting schema changes safely
- **Migration Tracking**: Recording which migrations have been applied
- **Data Seeding**: Populating database with initial data
- **Idempotent Migrations**: Safe to run multiple times
- **Forward and Backward Migrations**: Up and down migrations

## Schema Evolution Journey

### Version 1: Basic Product System

```go
type Product struct {
    ID          uint
    Name        string
    Price       float64
    Description string
    CreatedAt   time.Time
    UpdatedAt   time.Time
}
```

**What it creates**: Simple products table with basic fields

### Version 2: Add Categories

```go
type Category struct {
    ID          uint
    Name        string
    Description string
    Products    []Product  // One-to-many relationship
}

type Product struct {
    // ... existing fields ...
    CategoryID  uint       // Foreign key added
}
```

**What it adds**:
- Categories table
- Foreign key relationship between products and categories

### Version 3: Enhanced Product with Inventory

```go
type Product struct {
    // ... existing fields ...
    Stock       int        // Inventory tracking
    SKU         string     // Stock Keeping Unit
    IsActive    bool       // Product availability flag
}
```

**What it adds**:
- Stock management field
- Unique SKU identifier
- Active/inactive flag

## Data Models

### MigrationVersion (Tracking Table)

```go
type MigrationVersion struct {
    ID        uint
    Version   int       // Current schema version
    AppliedAt time.Time // When migration was applied
}
```

This table tracks which migrations have been applied to the database.

### Product (Final Schema - Version 3)

```go
type Product struct {
    ID          uint
    Name        string
    Price       float64
    Description string
    CategoryID  uint
    Category    Category
    Stock       int
    SKU         string
    IsActive    bool
    CreatedAt   time.Time
    UpdatedAt   time.Time
}
```

### Category

```go
type Category struct {
    ID          uint
    Name        string
    Description string
    Products    []Product
    CreatedAt   time.Time
    UpdatedAt   time.Time
}
```

## Required Functions

Implement these functions in the template:

1. **ConnectDB() (*gorm.DB, error)**
   - Establish database connection
   - Auto-migrate only the MigrationVersion table

2. **RunMigration(db *gorm.DB, version int) error**
   - Run migrations up to specified version
   - Check current version to avoid re-running
   - Execute SQL for schema changes
   - Record migration in MigrationVersion table

3. **RollbackMigration(db *gorm.DB, version int) error**
   - Rollback to specified version
   - Remove schema changes in reverse order
   - Update MigrationVersion table

4. **GetMigrationVersion(db *gorm.DB) (int, error)**
   - Query current migration version
   - Return 0 if no migrations applied

5. **SeedData(db *gorm.DB) error**
   - Populate database with initial data
   - Create sample categories and products
   - Use idempotent approach (check before creating)

6. **CreateProduct(db *gorm.DB, product *Product) error**
   - Create product with validation
   - Validate name, price, SKU

7. **GetProductsByCategory(db *gorm.DB, categoryID uint) ([]Product, error)**
   - Retrieve products in specific category
   - Preload category association

8. **UpdateProductStock(db *gorm.DB, productID uint, quantity int) error**
   - Update product stock quantity

## Migration Flow Diagram

```
Version 0 (Empty)
    ↓ [Migrate Up to V1]
Version 1 (Products table)
    ↓ [Migrate Up to V2]
Version 2 (+ Categories table, + CategoryID)
    ↓ [Migrate Up to V3]
Version 3 (+ Stock, SKU, IsActive)
    ↓ [Rollback to V2]
Version 2 (- Stock, SKU, IsActive)
    ↓ [Rollback to V1]
Version 1 (- Categories, - CategoryID)
    ↓ [Rollback to V0]
Version 0 (- Products)
```

## Key Learning Points

1. **Version Tracking**: Always track which migrations have been applied
2. **Idempotency**: Migrations should be safe to run multiple times
3. **Atomic Changes**: Each migration version should be a complete, atomic change
4. **Rollback Strategy**: Always plan how to undo migrations
5. **Data Migration**: Consider existing data when changing schemas
6. **Testing**: Test both forward and backward migrations
7. **Production Safety**: Never modify existing migrations after they're deployed

## How to Practice

1. Navigate to the `.practice` directory
2. Open `template.go` and complete the TODOs
3. Uncomment the main function code to test your implementation
4. Run the code: `go run template.go`
5. Compare with `solution.go` if you get stuck

## Expected Output

```
Running migrations...
Applied migration version 1
Applied migration version 2
Applied migration version 3
Current migration version: 3

Seeding data...
Created product: Wireless Mouse (ID: 3)
Found 3 products in category 1
Updated product stock

Testing rollback...
Rolled back to version 2
Current migration version: 2
```

## Testing Requirements

Your solution should:
- ✅ Run migrations sequentially (1 → 2 → 3)
- ✅ Track current migration version accurately
- ✅ Support idempotent migrations (safe to re-run)
- ✅ Rollback migrations in reverse order
- ✅ Seed data without creating duplicates
- ✅ Create products with validation
- ✅ Query products by category
- ✅ Update product inventory
- ✅ Handle migration version 0 (no migrations)
- ✅ Complete end-to-end workflow

## SQLite-Specific Considerations

SQLite has limitations compared to other databases:

1. **No DROP COLUMN**: SQLite doesn't support dropping columns directly
   - Workaround: Create new table without the column, copy data, rename table
2. **Limited ALTER TABLE**: Cannot modify existing columns
3. **Foreign Keys**: Must be enabled explicitly (not required for this exercise)

## Migration Best Practices

### DO:
- ✅ Version all schema changes
- ✅ Test migrations on a copy of production data
- ✅ Make migrations reversible when possible
- ✅ Keep migrations small and focused
- ✅ Document migration purpose and impact
- ✅ Use transactions for complex migrations
- ✅ Back up data before running migrations in production

### DON'T:
- ❌ Modify existing migrations after deployment
- ❌ Delete old migrations
- ❌ Skip migration versions
- ❌ Mix schema and data changes excessively
- ❌ Ignore migration failures
- ❌ Run untested migrations in production

## Common Pitfalls

1. **Not Checking Current Version**: Always check before running migrations
2. **Non-Idempotent Migrations**: Migrations should handle re-runs gracefully
3. **Losing Data**: Be careful when dropping tables or columns
4. **Broken Rollbacks**: Test rollback path as thoroughly as forward path
5. **Hardcoded IDs**: Don't rely on specific IDs in seed data

## Advanced Migration Patterns

### Data Migration Example

```go
// Migration that transforms existing data
err = db.Exec(`
    UPDATE products 
    SET sku = 'PROD-' || id 
    WHERE sku IS NULL
`).Error
```

### Conditional Migration

```go
// Check if column exists before adding
var count int
db.Raw("SELECT COUNT(*) FROM pragma_table_info('products') WHERE name='stock'").Scan(&count)
if count == 0 {
    db.Exec("ALTER TABLE products ADD COLUMN stock INTEGER DEFAULT 0")
}
```

### Transaction-Based Migration

```go
err = db.Transaction(func(tx *gorm.DB) error {
    // Multiple operations in transaction
    if err := tx.Exec("CREATE TABLE ...").Error; err != nil {
        return err
    }
    if err := tx.Exec("ALTER TABLE ...").Error; err != nil {
        return err
    }
    return nil
})
```

## Production Migration Tools

For production systems, consider using dedicated migration tools:

- **[golang-migrate](https://github.com/golang-migrate/migrate)**: Full-featured migration tool
- **[goose](https://github.com/pressly/goose)**: Database migration tool with Go support
- **[sql-migrate](https://github.com/rubenv/sql-migrate)**: SQL schema migration tool
- **[atlas](https://atlasgo.io/)**: Modern schema migration tool

These tools provide:
- CLI for managing migrations
- Migration file generation
- Automatic rollback on failure
- Migration status tracking
- Support for multiple databases

## Learning Resources

- [GORM Migration Guide](https://gorm.io/docs/migration.html)
- [Database Schema Migrations](https://martinfowler.com/articles/evodb.html)
- [Evolutionary Database Design](https://www.martinfowler.com/articles/evodb.html)
- [golang-migrate Documentation](https://github.com/golang-migrate/migrate)
- [SQLite ALTER TABLE](https://www.sqlite.org/lang_altertable.html)

## Extensions (Optional Challenges)

After completing the basic implementation, try these extensions:

1. **Transaction Safety**: Wrap each migration in a transaction
2. **Migration Timestamps**: Add timestamp to migration file names
3. **Dry Run Mode**: Preview migrations without applying them
4. **Migration Status**: Show which migrations are pending
5. **Data Validation**: Validate data consistency after migrations
6. **Migration Hooks**: Add before/after migration callbacks
7. **Multiple Databases**: Support PostgreSQL and MySQL
8. **Parallel Migrations**: Handle concurrent migration attempts
9. **Partial Rollback**: Rollback only specific migration
10. **Migration Dependencies**: Handle dependent migrations

## Real-World Scenarios

### Adding a New Field

```go
// Step 1: Add nullable column
ALTER TABLE products ADD COLUMN rating REAL

// Step 2: Set default values for existing rows
UPDATE products SET rating = 5.0 WHERE rating IS NULL

// Step 3: Add NOT NULL constraint (next migration)
```

### Renaming a Column

```go
// Step 1: Add new column
ALTER TABLE products ADD COLUMN product_name TEXT

// Step 2: Copy data
UPDATE products SET product_name = name

// Step 3: Drop old column (next migration)
ALTER TABLE products DROP COLUMN name
```

### Splitting a Table

```go
// Create new table
CREATE TABLE product_details ...

// Copy data
INSERT INTO product_details SELECT ... FROM products

// Remove columns from products
```

## Next Steps

After completing this module, explore:
- **Advanced GORM Features**: Hooks, scopes, and plugins
- **Database Performance**: Indexing, query optimization
- **Microservices**: Managing schemas across services
- **Testing Strategies**: Integration testing with databases
- **CI/CD Integration**: Automated migration testing

---

**Note**: This module creates an `ecommerce.db` SQLite database file. Delete it to start fresh.

