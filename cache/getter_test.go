package cache

import (
	"errors"
	"testing"
)

func TestGetLocalMapOrCompute(t *testing.T) {
	// Define a simple in-memory cache
	cacheMap := make(map[string]int)

	// Counter to track number of times f() is called
	callCount := 0

	// Function to compute the value
	compute := func(key string) (int, error) {
		callCount++
		if key == "fail" {
			return 0, errors.New("forced failure")
		}
		return len(key), nil
	}

	// First call: should compute and store in cache
	val, fromCache, err := GetLocalMapOrCompute(cacheMap, "hello", compute)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if fromCache {
		t.Errorf("expected fromCache=false, got true")
	}
	if val != 5 {
		t.Errorf("expected value 5, got %d", val)
	}
	if callCount != 1 {
		t.Errorf("expected callCount 1, got %d", callCount)
	}

	// Second call: should return from cache
	val, fromCache, err = GetLocalMapOrCompute(cacheMap, "hello", compute)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if !fromCache {
		t.Errorf("expected fromCache=true, got false")
	}
	if val != 5 {
		t.Errorf("expected value 5, got %d", val)
	}
	if callCount != 1 {
		t.Errorf("expected callCount to remain 1, got %d", callCount)
	}

	// Call with key that triggers error
	val, fromCache, err = GetLocalMapOrCompute(cacheMap, "fail", compute)
	if err == nil {
		t.Errorf("expected error, got nil")
	}
	if fromCache {
		t.Errorf("expected fromCache=false on error, got true")
	}
	if val != 0 {
		t.Errorf("expected zero value on error, got %d", val)
	}
}
