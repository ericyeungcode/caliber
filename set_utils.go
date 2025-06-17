package caliber

// Set difference: a - b
func SetDiff[K comparable, V1 any, V2 any](a map[K]V1, b map[K]V2) []K {
	var diff []K
	for val := range a {
		if _, found := b[val]; !found {
			diff = append(diff, val)
		}
	}
	return diff
}

// Set difference: a - b
// input is array
func SetDiffSlice[K comparable](a, b []K) []K {
	x := ListToMap(a, func(v K) (K, struct{}) { return v, struct{}{} })
	y := ListToMap(b, func(v K) (K, struct{}) { return v, struct{}{} })
	return SetDiff(x, y)
}
