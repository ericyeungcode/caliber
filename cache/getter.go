package cache

func GetLocalMapOrCompute[K comparable, V any](localCache map[K]V, key K, f func(K) (V, error)) (val V, fromCache bool, err error) {
	// Check if key exists in the cache
	if v, ok := localCache[key]; ok {
		return v, true, nil
	}

	// Otherwise, call f to get the value
	v, err := f(key)
	if err != nil {
		var zero V
		return zero, false, err
	}

	// Store in cache
	localCache[key] = v
	return v, false, nil
}

// GetTTLMapOrCompute
// GetRedisOrCompute
