package main

import (
	"errors"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/sqlite"
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
	db, err := gorm.Open(sqlite.Open("ecommerce.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Only migrate the MigrationVersion table
	err = db.AutoMigrate(&MigrationVersion{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

// RunMigration runs a specific migration version
func RunMigration(db *gorm.DB, version int) error {
	// Check if migration already applied
	currentVersion, err := GetMigrationVersion(db)
	if err != nil {
		return err
	}

	if currentVersion >= version {
		return nil // Migration already applied
	}

	// Run migrations in sequence from current to target version
	for v := currentVersion + 1; v <= version; v++ {
		switch v {
		case 1:
			// Version 1: Create basic products table
			err = db.Exec(`
				CREATE TABLE IF NOT EXISTS products (
					id INTEGER PRIMARY KEY AUTOINCREMENT,
					name TEXT NOT NULL,
					price REAL NOT NULL,
					description TEXT,
					created_at DATETIME,
					updated_at DATETIME
				)
			`).Error
			if err != nil {
				return err
			}

		case 2:
			// Version 2: Add categories table and foreign key
			err = db.Exec(`
				CREATE TABLE IF NOT EXISTS categories (
					id INTEGER PRIMARY KEY AUTOINCREMENT,
					name TEXT UNIQUE NOT NULL,
					description TEXT,
					created_at DATETIME,
					updated_at DATETIME
				)
			`).Error
			if err != nil {
				return err
			}

			// Add CategoryID column to products
			err = db.Exec(`
				ALTER TABLE products ADD COLUMN category_id INTEGER
			`).Error
			if err != nil {
				return err
			}

		case 3:
			// Version 3: Add inventory fields to products
			err = db.Exec(`
				ALTER TABLE products ADD COLUMN stock INTEGER DEFAULT 0
			`).Error
			if err != nil {
				return err
			}

			// SQLite doesn't support ADD COLUMN with UNIQUE constraint
			// Add column without constraint first
			err = db.Exec(`
				ALTER TABLE products ADD COLUMN sku TEXT
			`).Error
			if err != nil {
				return err
			}

			// Create unique index separately
			err = db.Exec(`
				CREATE UNIQUE INDEX idx_products_sku ON products(sku)
			`).Error
			if err != nil {
				return err
			}

			err = db.Exec(`
				ALTER TABLE products ADD COLUMN is_active INTEGER DEFAULT 1
			`).Error
			if err != nil {
				return err
			}

		default:
			return errors.New("unknown migration version")
		}

		// Record the migration
		migration := MigrationVersion{
			Version:   v,
			AppliedAt: time.Now(),
		}
		if err := db.Create(&migration).Error; err != nil {
			return err
		}
	}

	return nil
}

// RollbackMigration rolls back to a specific migration version
func RollbackMigration(db *gorm.DB, version int) error {
	currentVersion, err := GetMigrationVersion(db)
	if err != nil {
		return err
	}

	if currentVersion <= version {
		return nil // Already at or before target version
	}

	// Rollback migrations in reverse order
	for v := currentVersion; v > version; v-- {
		switch v {
		case 3:
			// Rollback version 3: Remove inventory fields
			// Drop the unique index first
			db.Exec(`DROP INDEX IF EXISTS idx_products_sku`)
			
			// Note: SQLite doesn't support DROP COLUMN directly
			// In production, you'd need to recreate the table
			// For this example, we'll work around it
			db.Exec(`
				CREATE TABLE products_temp AS 
				SELECT id, name, price, description, category_id, created_at, updated_at 
				FROM products
			`)
			db.Exec(`DROP TABLE products`)
			db.Exec(`ALTER TABLE products_temp RENAME TO products`)

		case 2:
			// Rollback version 2: Remove categories and category_id
			db.Exec(`DROP TABLE IF EXISTS categories`)
			db.Exec(`
				CREATE TABLE products_temp AS 
				SELECT id, name, price, description, created_at, updated_at 
				FROM products
			`)
			db.Exec(`DROP TABLE products`)
			db.Exec(`ALTER TABLE products_temp RENAME TO products`)

		case 1:
			// Rollback version 1: Remove products table
			db.Exec(`DROP TABLE IF EXISTS products`)

		default:
			return errors.New("unknown migration version")
		}

		// Remove the migration record
		db.Where("version = ?", v).Delete(&MigrationVersion{})
	}

	return nil
}

// GetMigrationVersion gets the current migration version
func GetMigrationVersion(db *gorm.DB) (int, error) {
	var migration MigrationVersion
	result := db.Order("version DESC").First(&migration)
	
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return 0, nil // No migrations applied yet
		}
		return 0, result.Error
	}

	return migration.Version, nil
}

// SeedData populates the database with initial data
func SeedData(db *gorm.DB) error {
	// Create categories
	categories := []Category{
		{Name: "Electronics", Description: "Electronic devices and accessories"},
		{Name: "Books", Description: "Physical and digital books"},
		{Name: "Clothing", Description: "Apparel and fashion items"},
	}

	for _, cat := range categories {
		var existing Category
		result := db.Where("name = ?", cat.Name).First(&existing)
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			db.Create(&cat)
		}
	}

	// Create sample products
	var electronicsCategory Category
	db.Where("name = ?", "Electronics").First(&electronicsCategory)

	products := []Product{
		{
			Name:        "Laptop",
			Price:       999.99,
			Description: "High-performance laptop",
			CategoryID:  electronicsCategory.ID,
			Stock:       10,
			SKU:         "LAPTOP-001",
			IsActive:    true,
		},
		{
			Name:        "Keyboard",
			Price:       79.99,
			Description: "Mechanical keyboard",
			CategoryID:  electronicsCategory.ID,
			Stock:       25,
			SKU:         "KEYBOARD-001",
			IsActive:    true,
		},
	}

	for _, prod := range products {
		var existing Product
		result := db.Where("sku = ?", prod.SKU).First(&existing)
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			db.Create(&prod)
		}
	}

	return nil
}

// CreateProduct creates a new product with validation
func CreateProduct(db *gorm.DB, product *Product) error {
	// Validate
	if product.Name == "" {
		return errors.New("product name cannot be empty")
	}
	if product.Price <= 0 {
		return errors.New("product price must be greater than 0")
	}
	if product.SKU == "" {
		return errors.New("product SKU cannot be empty")
	}

	result := db.Create(product)
	return result.Error
}

// GetProductsByCategory retrieves all products in a specific category
func GetProductsByCategory(db *gorm.DB, categoryID uint) ([]Product, error) {
	var products []Product
	result := db.Where("category_id = ?", categoryID).Preload("Category").Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}
	return products, nil
}

// UpdateProductStock updates the stock quantity of a product
func UpdateProductStock(db *gorm.DB, productID uint, quantity int) error {
	result := db.Model(&Product{}).Where("id = ?", productID).Update("stock", quantity)
	return result.Error
}

func main() {
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
}

// Notes:
// - Migrations should be idempotent (safe to run multiple times)
// - Always track migration versions to know current schema state
// - Use transactions for complex migrations to ensure atomicity
// - Test rollbacks thoroughly before using in production
// - Consider using a migration library like golang-migrate for production

