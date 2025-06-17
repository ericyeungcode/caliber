package caliber

// func MapKeys[K comparable, V any](m map[K]V) []K {
// 	keys := make([]K, len(m))
// 	i := 0
// 	for k := range m {
// 		keys[i] = k
// 		i++
// 	}
// 	return keys
// }

// func MapValues[K comparable, V any](m map[K]V) []V {
// 	values := make([]V, len(m))
// 	i := 0
// 	for _, v := range m {
// 		values[i] = v
// 		i++
// 	}
// 	return values
// }

func MapCopy[K comparable, V any](m map[K]V) map[K]V {
	result := make(map[K]V, len(m))
	for k, v := range m {
		result[k] = v
	}
	return result
}

func MapContainsKey[K comparable, V any](m map[K]V, key K) bool {
	_, ok := m[key]
	return ok
}

func MapGetOrDefault[K comparable, V any](m map[K]V, key K, defaultVal V) V {
	val, ok := m[key]
	if ok {
		return val
	}
	return defaultVal
}

func CompareSameKeys[K comparable, V1 any, V2 any](map1 map[K]V1, map2 map[K]V2) bool {
	if len(map1) != len(map2) {
		return false
	}

	for key := range map1 {
		if _, exists := map2[key]; !exists {
			return false
		}
	}

	return true
}

func MergeMaps[K comparable, V any](map1, map2 map[K]V) map[K]V {
	// Create a new map with an initial size equal to the combined size of both maps
	merged := make(map[K]V, len(map1)+len(map2))

	// Add all elements from map1
	for k, v := range map1 {
		merged[k] = v
	}

	// Add all elements from map2, overwriting if keys overlap
	for k, v := range map2 {
		merged[k] = v
	}

	return merged
}
