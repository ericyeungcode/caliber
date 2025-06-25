package slicex

import "github.com/ericyeungcode/caliber/common"

func Transform[T, R any](xs []T, f func(T) R) []R {
	ys := make([]R, len(xs))
	for i, x := range xs {
		ys[i] = f(x)
	}
	return ys
}

func TransformIf[T, R any](xs []T, transFunc func(T) R, predictFunc func(T) bool, size ...int) []R {
	capSize := 0
	if len(size) > 0 {
		capSize = size[0]
	}

	ys := make([]R, 0, capSize)
	for _, x := range xs {
		if predictFunc(x) {
			ys = append(ys, transFunc(x))
		}
	}
	return ys
}

func MapToList[K comparable, V any, R any](m map[K]V, f func(K, V) R) []R {
	ys := make([]R, 0, len(m))
	for key, val := range m {
		ys = append(ys, f(key, val))
	}
	return ys
}

func ListToMap[T any, K comparable, V any](list []T, f func(T) (K, V)) map[K]V {
	ret := make(map[K]V, len(list))
	for _, val := range list {
		k, v := f(val)
		ret[k] = v
	}
	return ret
}

func Filter[T any](slice []T, f func(T) bool, size ...int) []T {
	var n []T
	if len(size) > 0 {
		n = make([]T, 0, size[0]) // pre allocate
	}

	for _, e := range slice {
		if f(e) {
			n = append(n, e)
		}
	}
	return n
}

func FirstItem[K comparable, V any](m map[K]V) (k K, v V, ok bool) {
	for key, value := range m {
		k, v, ok = key, value, true
		break
	}
	return k, v, ok
}

func Search[T any](collection []T, predicate func(item T) bool) (T, bool) {
	for _, v := range collection {
		if predicate(v) {
			return v, true
		}
	}

	var result T
	return result, false
}

func Unique[T comparable](sliceList []T, size ...int) []T {
	preAllocSize := 0
	if len(size) > 0 {
		preAllocSize = size[0]
	}
	allKeys := make(map[T]bool, preAllocSize)
	list := make([]T, 0, preAllocSize)

	for _, item := range sliceList {
		if _, found := allKeys[item]; !found {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func CapSliceCount[T any](items []T, count int) []T {
	if len(items) > count {
		items = items[:count]
	}
	return items
}

func CapStr(v string, maxSize int) string {
	if len(v) > maxSize {
		return v[:maxSize]
	}
	return v
}

func ChunkSlice[T any](items []T, chunkSize int) (chunks [][]T) {
	common.Assert(chunkSize > 0, "ChunkSlice: chunkSize should be > 0")

	count := len(items)
	for i := 0; i < count; i += chunkSize {
		end := i + chunkSize
		if end > count {
			end = count
		}
		chunks = append(chunks, items[i:end])
	}
	return
}
