package caliber

import (
	"fmt"
)

func Assert(cond bool, format string, args ...any) {
	if !cond {
		panic(fmt.Sprintf(format, args...))
	}
}

func If[T any](condition bool, trueValue, falseValue T) T {
	if condition {
		return trueValue
	}
	return falseValue
}

func NewValuePtr[T any](v T) *T {
	return &v
}
