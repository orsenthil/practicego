package main

import (
	"time"

	"gorm.io/gorm"
)

// MigrationVersion tracks the current database schema version
type MigrationVersion struct {
	ID        uint `gorm:"primaryKey"`
	Version   int  `gorm:"unique;not null"`
	AppliedAt time.Time
}

// Product represents a product in the e-commerce system
type Product struct {
	ID          uint     `gorm:"primaryKey"`
	Name        string   `gorm:"not null"`
	Price       float64  `gorm:"not null"`
	Description string   `gorm:"type:text"`
	CategoryID  uint     `gorm:"not null"`
	Category    Category `gorm:"foreignKey:CategoryID"`
	Stock       int      `gorm:"default:0"`
	SKU         string   `gorm:"unique;not null"`
	IsActive    bool     `gorm:"default:true"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// Category represents a product category
type Category struct {
	ID          uint      `gorm:"primaryKey"`
	Name        string    `gorm:"unique;not null"`
	Description string    `gorm:"type:text"`
	Products    []Product `gorm:"foreignKey:CategoryID"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// ConnectDB establishes a connection to the SQLite database
func ConnectDB() (*gorm.DB, error) {
	// TODO: Implement database connection
	// Hint: Use gorm.Open with sqlite.Open("ecommerce.db")
	// Don't auto-migrate models here - migrations will be handled separately
	return nil, nil
}

// RunMigration runs a specific migration version
func RunMigration(db *gorm.DB, version int) error {
	// TODO: Implement migration execution
	// Hint: Check current version first to avoid re-running migrations
	// Version 1: Create basic products table (without CategoryID, Stock, SKU, IsActive)
	// Version 2: Create categories table and add CategoryID to products
	// Version 3: Add Stock, SKU, IsActive columns to products
	// Record the migration version in MigrationVersion table
	return nil
}

// RollbackMigration rolls back to a specific migration version
func RollbackMigration(db *gorm.DB, version int) error {
	// TODO: Implement migration rollback
	// Hint: Get current version and rollback step by step
	// Rollback from 3 to 2: Drop Stock, SKU, IsActive columns
	// Rollback from 2 to 1: Drop categories table and CategoryID column
	// Rollback from 1 to 0: Drop products table
	// Update MigrationVersion table
	return nil
}

// GetMigrationVersion gets the current migration version
func GetMigrationVersion(db *gorm.DB) (int, error) {
	// TODO: Implement version retrieval
	// Hint: Query MigrationVersion table for the latest version
	// Return 0 if no migrations have been run
	return 0, nil
}

// SeedData populates the database with initial data
func SeedData(db *gorm.DB) error {
	// TODO: Implement data seeding
	// Hint: Create some categories and products for testing
	// Use FirstOrCreate to avoid duplicates
	return nil
}

// CreateProduct creates a new product with validation
func CreateProduct(db *gorm.DB, product *Product) error {
	// TODO: Implement product creation
	// Hint: Validate that price > 0, name is not empty, SKU is unique
	// Use db.Create()
	return nil
}

// GetProductsByCategory retrieves all products in a specific category
func GetProductsByCategory(db *gorm.DB, categoryID uint) ([]Product, error) {
	// TODO: Implement products retrieval by category
	// Hint: Use db.Where("category_id = ?", categoryID).Find()
	// Consider preloading the Category association
	return nil, nil
}

// UpdateProductStock updates the stock quantity of a product
func UpdateProductStock(db *gorm.DB, productID uint, quantity int) error {
	// TODO: Implement stock update
	// Hint: Find the product first, update the Stock field, then save
	// Or use db.Model(&Product{}).Where("id = ?", productID).Update("stock", quantity)
	return nil
}

func main() {
	// TODO: Uncomment and complete this section when you're ready to test
	/*
		// Connect to database
		db, err := ConnectDB()
		if err != nil {
			log.Fatal("Failed to connect to database:", err)
		}

		// Run migrations to version 3
		fmt.Println("Running migrations...")
		for version := 1; version <= 3; version++ {
			if err := RunMigration(db, version); err != nil {
				log.Fatal("Migration failed:", err)
			}
			fmt.Printf("Applied migration version %d\n", version)
		}

		// Check current version
		currentVersion, _ := GetMigrationVersion(db)
		fmt.Printf("Current migration version: %d\n", currentVersion)

		// Seed data
		fmt.Println("\nSeeding data...")
		if err := SeedData(db); err != nil {
			log.Fatal("Seeding failed:", err)
		}

		// Create a product
		product := &Product{
			Name:        "Wireless Mouse",
			Price:       29.99,
			Description: "Ergonomic wireless mouse",
			CategoryID:  1, // Assuming Electronics category exists
			Stock:       50,
			SKU:         "MOUSE-001",
			IsActive:    true,
		}
		if err := CreateProduct(db, product); err != nil {
			log.Fatal("Failed to create product:", err)
		}
		fmt.Printf("Created product: %s (ID: %d)\n", product.Name, product.ID)

		// Get products by category
		products, err := GetProductsByCategory(db, 1)
		if err != nil {
			log.Fatal("Failed to get products:", err)
		}
		fmt.Printf("Found %d products in category 1\n", len(products))

		// Update product stock
		if err := UpdateProductStock(db, product.ID, 45); err != nil {
			log.Fatal("Failed to update stock:", err)
		}
		fmt.Println("Updated product stock")

		// Test rollback
		fmt.Println("\nTesting rollback...")
		if err := RollbackMigration(db, 2); err != nil {
			log.Fatal("Rollback failed:", err)
		}
		fmt.Println("Rolled back to version 2")

		currentVersion, _ = GetMigrationVersion(db)
		fmt.Printf("Current migration version: %d\n", currentVersion)
	*/
}

// Notes:
// - Migrations should be idempotent (safe to run multiple times)
// - Always track migration versions to know current schema state
// - Use transactions for complex migrations to ensure atomicity
// - Test rollbacks thoroughly before using in production
// - Consider using a migration library like golang-migrate for production

