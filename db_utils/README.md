# Database Utilities - Async Query Functions

This package provides async database query functionality using GORM's `Find()` method.

## Functions

### `AsyncFetchDb(querier *gorm.DB, outItems any) chan *AsyncDbResult`

Executes a GORM query asynchronously and returns a channel that will receive the result.

**Parameters:**
- `querier *gorm.DB`: The GORM database query builder (can include WHERE, ORDER BY, etc.)
- `outItems any`: Pointer to a slice where results will be stored (e.g., `&[]User{}`)

**Returns:**
- `chan *AsyncDbResult`: Channel that will receive the query result

### `RecvAsyncResult(resultC chan *AsyncDbResult, waitSeconds int) *AsyncDbResult`

Receives the result from an async query with timeout.

**Parameters:**
- `resultC chan *AsyncDbResult`: The result channel from `AsyncFetchDb`
- `waitSeconds int`: Timeout in seconds

**Returns:**
- `*AsyncDbResult`: The query result or timeout error

## Types

### `AsyncDbResult`

```go
type AsyncDbResult struct {
    Data any    // The query result data
    Err  error  // Any error that occurred
}
```

## Usage Examples

### Basic Async Query

```go
var users []User

// Start async query
resultChan := db_utils.AsyncFetchDb(db, &users)

// Do other work while query runs
fmt.Println("Query is running in background...")

// Receive result with 5 second timeout
result := db_utils.RecvAsyncResult(resultChan, 5)

if result.Err != nil {
    log.Printf("Query failed: %v", result.Err)
    return
}

// Type assert the result
if userData, ok := result.Data.(*[]User); ok {
    fmt.Printf("Found %d users\n", len(*userData))
}
```

### Async Query with Conditions

```go
var users []User

// Build query with conditions
query := db.Where("age > ?", 25).Order("name ASC")

// Start async query
resultChan := db_utils.AsyncFetchDb(query, &users)

// Receive result
result := db_utils.RecvAsyncResult(resultChan, 5)

if result.Err != nil {
    log.Printf("Query failed: %v", result.Err)
    return
}

if userData, ok := result.Data.(*[]User); ok {
    fmt.Printf("Found %d users over 25\n", len(*userData))
}
```

### Multiple Concurrent Queries

```go
// Start multiple async queries
usersChan := db_utils.AsyncFetchDb(db, &[]User{})
ordersChan := db_utils.AsyncFetchDb(db, &[]Order{})

// Wait for both results
userResult := db_utils.RecvAsyncResult(usersChan, 5)
orderResult := db_utils.RecvAsyncResult(ordersChan, 5)

// Process results
if userResult.Err == nil {
    if userData, ok := userResult.Data.(*[]User); ok {
        fmt.Printf("Users: %d found\n", len(*userData))
    }
}

if orderResult.Err == nil {
    if orderData, ok := orderResult.Data.(*[]Order); ok {
        fmt.Printf("Orders: %d found\n", len(*orderData))
    }
}
```

### Handling Timeouts

```go
var orders []Order

// Start async query
resultChan := db_utils.AsyncFetchDb(db, &orders)

// Try to receive with short timeout
result := db_utils.RecvAsyncResult(resultChan, 1) // 1 second timeout

if result.Err != nil {
    fmt.Printf("Query timed out or failed: %v\n", result.Err)
    return
}

fmt.Println("Query completed successfully")
```

## Key Features

1. **Non-blocking**: Queries run in goroutines, allowing your main thread to continue working
2. **Timeout support**: Built-in timeout mechanism to prevent hanging
3. **Error handling**: Comprehensive error reporting
4. **Type safety**: Results are properly typed and can be type-asserted
5. **Resource management**: Channels are properly closed to prevent goroutine leaks

## Best Practices

1. **Always check for errors**: Always check `result.Err` before using the data
2. **Use appropriate timeouts**: Set reasonable timeout values based on your query complexity
3. **Type assertion**: Use type assertion to safely access the result data
4. **Resource cleanup**: The channel is automatically closed, but ensure you receive the result to prevent goroutine leaks

## See Also

- Check the `examples/async_db_demo/main.go` for complete working examples
- The `db_utils/generic.go` file contains additional database utilities like `AutoQuery` and `BuildStmt` 

[cursor generated]