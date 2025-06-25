package cache

import (
	"sync"
)

// TypedSyncMap is a statically typed wrapper around sync.Map.
type TypedSyncMap[K comparable, V any] struct {
	m sync.Map
}

// Store sets the value for a key.
func (tm *TypedSyncMap[K, V]) Store(key K, value V) {
	tm.m.Store(key, value)
}

// Load retrieves the value for a key. The second return value indicates whether the key was found.
func (tm *TypedSyncMap[K, V]) Load(key K) (V, bool) {
	value, ok := tm.m.Load(key)
	if !ok {
		var zero V
		return zero, false
	}
	return value.(V), true
}

// LoadOrStore returns the existing value for the key if present.
// Otherwise, it stores and returns the given value.
func (tm *TypedSyncMap[K, V]) LoadOrStore(key K, value V) (V, bool) {
	actual, loaded := tm.m.LoadOrStore(key, value)
	return actual.(V), loaded
}

// Delete removes the value for a key.
func (tm *TypedSyncMap[K, V]) Delete(key K) {
	tm.m.Delete(key)
}

// Range iterates over all key-value pairs in the map.
// The iteration stops if the provided function returns false.
func (tm *TypedSyncMap[K, V]) Range(f func(key K, value V) bool) {
	tm.m.Range(func(rawKey, rawValue any) bool {
		return f(rawKey.(K), rawValue.(V))
	})
}

// ToMap copies the content of the TypedMap into a standard map[K]V.
func (m *TypedSyncMap[K, V]) ToMap(size ...int) map[K]V {
	capSize := 0
	if len(size) > 0 {
		capSize = size[0]
	}

	result := make(map[K]V, capSize)
	m.Range(func(key K, value V) bool {
		result[key] = value
		return true
	})
	return result
}
