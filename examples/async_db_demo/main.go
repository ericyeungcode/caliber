package main

import (
	"fmt"
	"log"
	"time"

	"github.com/ericyeungcode/caliber/db_utils"
	"gorm.io/gorm"
)

// Example User model
type User struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"size:255"`
	Age  int
}

// Example Order model
type Order struct {
	ID     uint    `gorm:"primaryKey"`
	UserID uint    `gorm:"index"`
	Amount float64 `gorm:"type:decimal(10,2)"`
	Status string  `gorm:"size:50"`
}

func main() {
	// Connect to database (replace with your actual connection details)
	db := connectToDatabase()

	// Example 1: Simple async query
	fmt.Println("=== Example 1: Simple async query ===")
	simpleAsyncQuery(db)

	// Example 2: Async query with timeout
	fmt.Println("\n=== Example 2: Async query with timeout ===")
	asyncQueryWithTimeout(db)

	// Example 3: Multiple concurrent queries
	fmt.Println("\n=== Example 3: Multiple concurrent queries ===")
	multipleConcurrentQueries(db)

	// Example 4: Async query with conditions
	fmt.Println("\n=== Example 4: Async query with conditions ===")
	asyncQueryWithConditions(db)
}

func simpleAsyncQuery(db *gorm.DB) {
	var users []User

	// Start async query
	resultChan := db_utils.AsyncFetchDb(db, &users)

	// Do other work while query is running
	fmt.Println("Query is running in background...")
	time.Sleep(100 * time.Millisecond) // Simulate other work

	// Receive result with 5 second timeout
	result := db_utils.RecvAsyncResult(resultChan, 5)

	if result.Err != nil {
		log.Printf("Query failed: %v", result.Err)
		return
	}

	// Type assert the result
	if userData, ok := result.Data.(*[]User); ok {
		fmt.Printf("Found %d users\n", len(*userData))
		for _, user := range *userData {
			fmt.Printf("  - ID: %d, Name: %s, Age: %d\n", user.ID, user.Name, user.Age)
		}
	}
}

func asyncQueryWithTimeout(db *gorm.DB) {
	var orders []Order

	// Start async query
	resultChan := db_utils.AsyncFetchDb(db, &orders)

	// Try to receive with very short timeout (should timeout)
	result := db_utils.RecvAsyncResult(resultChan, 1) // 1 second timeout

	if result.Err != nil {
		fmt.Printf("Query timed out or failed: %v\n", result.Err)
		return
	}

	fmt.Println("Query completed successfully")
}

func multipleConcurrentQueries(db *gorm.DB) {
	// Start multiple async queries
	usersChan := db_utils.AsyncFetchDb(db, &[]User{})
	ordersChan := db_utils.AsyncFetchDb(db, &[]Order{})

	// Wait for both results
	userResult := db_utils.RecvAsyncResult(usersChan, 5)
	orderResult := db_utils.RecvAsyncResult(ordersChan, 5)

	// Process results
	if userResult.Err == nil {
		if userData, ok := userResult.Data.(*[]User); ok {
			fmt.Printf("Users query completed: %d users found\n", len(*userData))
		}
	}

	if orderResult.Err == nil {
		if orderData, ok := orderResult.Data.(*[]Order); ok {
			fmt.Printf("Orders query completed: %d orders found\n", len(*orderData))
		}
	}
}

func asyncQueryWithConditions(db *gorm.DB) {
	var users []User

	// Build query with conditions
	query := db.Where("age > ?", 25).Order("name ASC")

	// Start async query with conditions
	resultChan := db_utils.AsyncFetchDb(query, &users)

	// Receive result
	result := db_utils.RecvAsyncResult(resultChan, 5)

	if result.Err != nil {
		log.Printf("Query failed: %v", result.Err)
		return
	}

	if userData, ok := result.Data.(*[]User); ok {
		fmt.Printf("Found %d users over 25 years old\n", len(*userData))
		for _, user := range *userData {
			fmt.Printf("  - %s (Age: %d)\n", user.Name, user.Age)
		}
	}
}

// Helper function to connect to database (replace with your actual connection)
func connectToDatabase() *gorm.DB {
	// This is a placeholder - replace with your actual database connection
	// Example:
	// dsn := "user:password@tcp(localhost:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// if err != nil {
	//     log.Fatal(err)
	// }
	// return db

	// For demo purposes, return nil (you'll need to implement actual connection)
	fmt.Println("Note: Replace connectToDatabase() with your actual database connection")
	return nil
}
