package common

func NewValuePtr[T any](v T) *T {
	return &v
}
