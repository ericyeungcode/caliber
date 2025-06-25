package atomicx

import "sync/atomic"

type AtomicValue[T any] struct {
	v atomic.Value
}

// Store stores a value atomically.
func (a *AtomicValue[T]) Store(val T) {
	a.v.Store(val)
}

// Load retrieves the value atomically, returns zero if not set.
func (a *AtomicValue[T]) Load() T {
	val := a.v.Load()
	if val == nil {
		var zero T
		return zero
	}
	return val.(T)
}

// Swap replaces the current value and returns the old value.
func (a *AtomicValue[T]) Swap(newVal T) (oldVal T) {
	old := a.v.Swap(newVal)
	if old == nil {
		var zero T
		return zero
	}
	return old.(T)
}

// CompareAndSwap sets the value to new if the current value is equal to old.
func (a *AtomicValue[T]) CompareAndSwap(oldVal, newVal T) bool {
	return a.v.CompareAndSwap(oldVal, newVal)
}
