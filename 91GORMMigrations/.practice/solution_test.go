package main

import (
	"os"
	"testing"

	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := ConnectDB()
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
	return db
}

func cleanupTestDB(t *testing.T) {
	os.Remove("ecommerce.db")
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

func TestRunMigrationVersion1(t *testing.T) {
	defer cleanupTestDB(t)
	db := setupTestDB(t)

	err := RunMigration(db, 1)
	if err != nil {
		t.Fatalf("RunMigration(1) failed: %v", err)
	}

	// Verify migration version
	version, err := GetMigrationVersion(db)
	if err != nil {
		t.Fatalf("GetMigrationVersion failed: %v", err)
	}
	if version != 1 {
		t.Errorf("Expected version 1, got %d", version)
	}

	// Verify products table exists
	var count int64
	db.Raw("SELECT COUNT(*) FROM sqlite_master WHERE type='table' AND name='products'").Scan(&count)
	if count == 0 {
		t.Error("Expected products table to exist")
	}
}

func TestRunMigrationVersion2(t *testing.T) {
	defer cleanupTestDB(t)
	db := setupTestDB(t)

	// Run migration to version 2
	err := RunMigration(db, 2)
	if err != nil {
		t.Fatalf("RunMigration(2) failed: %v", err)
	}

	// Verify migration version
	version, err := GetMigrationVersion(db)
	if err != nil {
		t.Fatalf("GetMigrationVersion failed: %v", err)
	}
	if version != 2 {
		t.Errorf("Expected version 2, got %d", version)
	}

	// Verify categories table exists
	var count int64
	db.Raw("SELECT COUNT(*) FROM sqlite_master WHERE type='table' AND name='categories'").Scan(&count)
	if count == 0 {
		t.Error("Expected categories table to exist")
	}
}

func TestRunMigrationVersion3(t *testing.T) {
	defer cleanupTestDB(t)
	db := setupTestDB(t)

	// Run migration to version 3
	err := RunMigration(db, 3)
	if err != nil {
		t.Fatalf("RunMigration(3) failed: %v", err)
	}

	// Verify migration version
	version, err := GetMigrationVersion(db)
	if err != nil {
		t.Fatalf("GetMigrationVersion failed: %v", err)
	}
	if version != 3 {
		t.Errorf("Expected version 3, got %d", version)
	}

	// Verify stock column exists in products
	var columnExists int
	db.Raw("SELECT COUNT(*) FROM pragma_table_info('products') WHERE name='stock'").Scan(&columnExists)
	if columnExists == 0 {
		t.Error("Expected stock column to exist in products table")
	}
}

func TestMigrationIdempotent(t *testing.T) {
	defer cleanupTestDB(t)
	db := setupTestDB(t)

	// Run migration twice
	err := RunMigration(db, 2)
	if err != nil {
		t.Fatalf("First RunMigration(2) failed: %v", err)
	}

	err = RunMigration(db, 2)
	if err != nil {
		t.Fatalf("Second RunMigration(2) failed: %v", err)
	}

	// Verify version is still 2
	version, err := GetMigrationVersion(db)
	if err != nil {
		t.Fatalf("GetMigrationVersion failed: %v", err)
	}
	if version != 2 {
		t.Errorf("Expected version 2, got %d", version)
	}
}

func TestGetMigrationVersionNoMigrations(t *testing.T) {
	defer cleanupTestDB(t)
	db := setupTestDB(t)

	version, err := GetMigrationVersion(db)
	if err != nil {
		t.Fatalf("GetMigrationVersion failed: %v", err)
	}
	if version != 0 {
		t.Errorf("Expected version 0 with no migrations, got %d", version)
	}
}

func TestRollbackMigration(t *testing.T) {
	defer cleanupTestDB(t)
	db := setupTestDB(t)

	// Run to version 3
	RunMigration(db, 3)

	// Rollback to version 2
	err := RollbackMigration(db, 2)
	if err != nil {
		t.Fatalf("RollbackMigration(2) failed: %v", err)
	}

	// Verify version
	version, err := GetMigrationVersion(db)
	if err != nil {
		t.Fatalf("GetMigrationVersion failed: %v", err)
	}
	if version != 2 {
		t.Errorf("Expected version 2 after rollback, got %d", version)
	}
}

func TestRollbackMigrationToZero(t *testing.T) {
	defer cleanupTestDB(t)
	db := setupTestDB(t)

	// Run to version 2
	RunMigration(db, 2)

	// Rollback to version 0
	err := RollbackMigration(db, 0)
	if err != nil {
		t.Fatalf("RollbackMigration(0) failed: %v", err)
	}

	// Verify version
	version, err := GetMigrationVersion(db)
	if err != nil {
		t.Fatalf("GetMigrationVersion failed: %v", err)
	}
	if version != 0 {
		t.Errorf("Expected version 0 after rollback, got %d", version)
	}

	// Verify products table doesn't exist
	var count int64
	db.Raw("SELECT COUNT(*) FROM sqlite_master WHERE type='table' AND name='products'").Scan(&count)
	if count != 0 {
		t.Error("Expected products table to not exist after rollback to 0")
	}
}

func TestSeedData(t *testing.T) {
	defer cleanupTestDB(t)
	db := setupTestDB(t)

	// Run migrations first
	RunMigration(db, 3)

	// Seed data
	err := SeedData(db)
	if err != nil {
		t.Fatalf("SeedData failed: %v", err)
	}

	// Verify categories were created
	var categoryCount int64
	db.Model(&Category{}).Count(&categoryCount)
	if categoryCount == 0 {
		t.Error("Expected categories to be created")
	}

	// Verify products were created
	var productCount int64
	db.Model(&Product{}).Count(&productCount)
	if productCount == 0 {
		t.Error("Expected products to be created")
	}
}

func TestCreateProduct(t *testing.T) {
	defer cleanupTestDB(t)
	db := setupTestDB(t)

	RunMigration(db, 3)
	SeedData(db)

	product := &Product{
		Name:        "Test Product",
		Price:       19.99,
		Description: "Test description",
		CategoryID:  1,
		Stock:       100,
		SKU:         "TEST-001",
		IsActive:    true,
	}

	err := CreateProduct(db, product)
	if err != nil {
		t.Fatalf("CreateProduct failed: %v", err)
	}

	if product.ID == 0 {
		t.Error("Expected product ID to be set")
	}
}

func TestCreateProductValidation(t *testing.T) {
	defer cleanupTestDB(t)
	db := setupTestDB(t)

	RunMigration(db, 3)
	SeedData(db)

	// Test empty name
	product := &Product{
		Name:  "",
		Price: 19.99,
		SKU:   "TEST-002",
	}
	err := CreateProduct(db, product)
	if err == nil {
		t.Error("Expected error for empty product name")
	}

	// Test invalid price
	product = &Product{
		Name:  "Test",
		Price: -10,
		SKU:   "TEST-003",
	}
	err = CreateProduct(db, product)
	if err == nil {
		t.Error("Expected error for negative price")
	}

	// Test empty SKU
	product = &Product{
		Name:  "Test",
		Price: 19.99,
		SKU:   "",
	}
	err = CreateProduct(db, product)
	if err == nil {
		t.Error("Expected error for empty SKU")
	}
}

func TestGetProductsByCategory(t *testing.T) {
	defer cleanupTestDB(t)
	db := setupTestDB(t)

	RunMigration(db, 3)
	SeedData(db)

	// Create additional product
	product := &Product{
		Name:        "Extra Product",
		Price:       29.99,
		Description: "Extra product",
		CategoryID:  1,
		Stock:       50,
		SKU:         "EXTRA-001",
		IsActive:    true,
	}
	CreateProduct(db, product)

	// Get products by category
	products, err := GetProductsByCategory(db, 1)
	if err != nil {
		t.Fatalf("GetProductsByCategory failed: %v", err)
	}

	if len(products) == 0 {
		t.Error("Expected to find products in category 1")
	}
}

func TestUpdateProductStock(t *testing.T) {
	defer cleanupTestDB(t)
	db := setupTestDB(t)

	RunMigration(db, 3)
	SeedData(db)

	// Create a product
	product := &Product{
		Name:        "Stock Test Product",
		Price:       39.99,
		Description: "For stock testing",
		CategoryID:  1,
		Stock:       100,
		SKU:         "STOCK-001",
		IsActive:    true,
	}
	CreateProduct(db, product)

	// Update stock
	newStock := 75
	err := UpdateProductStock(db, product.ID, newStock)
	if err != nil {
		t.Fatalf("UpdateProductStock failed: %v", err)
	}

	// Verify stock was updated
	var updatedProduct Product
	db.First(&updatedProduct, product.ID)
	if updatedProduct.Stock != newStock {
		t.Errorf("Expected stock %d, got %d", newStock, updatedProduct.Stock)
	}
}

func TestMigrationSequence(t *testing.T) {
	defer cleanupTestDB(t)
	db := setupTestDB(t)

	// Test running migrations in sequence
	for version := 1; version <= 3; version++ {
		err := RunMigration(db, version)
		if err != nil {
			t.Fatalf("RunMigration(%d) failed: %v", version, err)
		}

		currentVersion, err := GetMigrationVersion(db)
		if err != nil {
			t.Fatalf("GetMigrationVersion failed: %v", err)
		}

		if currentVersion != version {
			t.Errorf("After migration %d, expected version %d, got %d", version, version, currentVersion)
		}
	}
}

func TestCompleteWorkflow(t *testing.T) {
	defer cleanupTestDB(t)
	db := setupTestDB(t)

	// Run full migration
	err := RunMigration(db, 3)
	if err != nil {
		t.Fatalf("Migration failed: %v", err)
	}

	// Seed data
	err = SeedData(db)
	if err != nil {
		t.Fatalf("Seeding failed: %v", err)
	}

	// Create product
	product := &Product{
		Name:        "Complete Test Product",
		Price:       49.99,
		Description: "Complete workflow test",
		CategoryID:  1,
		Stock:       200,
		SKU:         "COMPLETE-001",
		IsActive:    true,
	}
	err = CreateProduct(db, product)
	if err != nil {
		t.Fatalf("CreateProduct failed: %v", err)
	}

	// Get products by category
	products, err := GetProductsByCategory(db, 1)
	if err != nil {
		t.Fatalf("GetProductsByCategory failed: %v", err)
	}
	if len(products) == 0 {
		t.Error("Expected to find products")
	}

	// Update stock
	err = UpdateProductStock(db, product.ID, 150)
	if err != nil {
		t.Fatalf("UpdateProductStock failed: %v", err)
	}

	// Verify update
	var updated Product
	db.First(&updated, product.ID)
	if updated.Stock != 150 {
		t.Errorf("Expected stock 150, got %d", updated.Stock)
	}
}

