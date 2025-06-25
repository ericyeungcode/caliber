package cache

import (
	"testing"
)

func TestTypedMap_StoreAndLoad(t *testing.T) {
	m := &TypedSyncMap[string, int]{}

	// Test storing and loading a value
	m.Store("key1", 42)
	value, ok := m.Load("key1")
	if !ok {
		t.Fatalf("expected key1 to exist")
	}
	if value != 42 {
		t.Fatalf("expected value 42, got %d", value)
	}

	// Test loading a non-existent key
	_, ok = m.Load("non_existent")
	if ok {
		t.Fatalf("expected non-existent key to not exist")
	}
}

func TestTypedMap_LoadOrStore(t *testing.T) {
	m := &TypedSyncMap[string, int]{}

	// Test storing a new key
	value, loaded := m.LoadOrStore("key1", 42)
	if loaded {
		t.Fatalf("expected key1 to not be loaded")
	}
	if value != 42 {
		t.Fatalf("expected value 42, got %d", value)
	}

	// Test loading an existing key
	value, loaded = m.LoadOrStore("key1", 84)
	if !loaded {
		t.Fatalf("expected key1 to be loaded")
	}
	if value != 42 {
		t.Fatalf("expected value 42, got %d", value)
	}
}

func TestTypedMap_Delete(t *testing.T) {
	m := &TypedSyncMap[string, int]{}

	// Store and then delete a key
	m.Store("key1", 42)
	m.Delete("key1")

	// Ensure the key no longer exists
	_, ok := m.Load("key1")
	if ok {
		t.Fatalf("expected key1 to be deleted")
	}
}

func TestTypedMap_Range(t *testing.T) {
	m := &TypedSyncMap[string, int]{}

	// Store multiple keys
	m.Store("key1", 42)
	m.Store("key2", 84)
	m.Store("key3", 126)

	// Use Range to iterate over all keys
	expected := map[string]int{
		"key1": 42,
		"key2": 84,
		"key3": 126,
	}
	actual := make(map[string]int)

	m.Range(func(key string, value int) bool {
		actual[key] = value
		return true
	})

	// Verify all expected keys are present
	if len(expected) != len(actual) {
		t.Fatalf("expected %d items, got %d", len(expected), len(actual))
	}
	for k, v := range expected {
		if actual[k] != v {
			t.Fatalf("expected key %s to have value %d, got %d", k, v, actual[k])
		}
	}
}

func TestTypedMap_ConcurrentAccess(t *testing.T) {
	m := &TypedSyncMap[string, int]{}

	// Store values concurrently
	done := make(chan bool)
	for i := 0; i < 100; i++ {
		go func(i int) {
			m.Store(string(rune('A'+i%26)), i)
			done <- true
		}(i)
	}

	// Wait for all goroutines to finish
	for i := 0; i < 100; i++ {
		<-done
	}

	// Verify some keys exist
	for i := 0; i < 26; i++ {
		_, ok := m.Load(string(rune('A' + i)))
		if !ok {
			t.Fatalf("expected key %c to exist", 'A'+i)
		}
	}
}

func TestTypedMap_ToMap(t *testing.T) {
	m := &TypedSyncMap[string, int]{}

	// Store key-value pairs
	m.Store("key1", 42)
	m.Store("key2", 84)
	m.Store("key3", 126)

	// Call ToMap and validate the result
	result := m.ToMap()

	// Define the expected map
	expected := map[string]int{
		"key1": 42,
		"key2": 84,
		"key3": 126,
	}

	// Check lengths
	if len(result) != len(expected) {
		t.Fatalf("expected map length %d, got %d", len(expected), len(result))
	}

	// Verify each key-value pair
	for key, expectedValue := range expected {
		actualValue, exists := result[key]
		if !exists {
			t.Fatalf("expected key %s to exist in the map", key)
		}
		if actualValue != expectedValue {
			t.Fatalf("for key %s, expected value %d, got %d", key, expectedValue, actualValue)
		}
	}
}

func TestTypedMap_ToMapEmpty(t *testing.T) {
	m := &TypedSyncMap[string, int]{}

	// Call ToMap on an empty TypedMap
	result := m.ToMap(10)

	// Validate that the result is an empty map
	if len(result) != 0 {
		t.Fatalf("expected an empty map, got map with %d entries", len(result))
	}
}
