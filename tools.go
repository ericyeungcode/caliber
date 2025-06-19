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

func Must1[T any](arg T, err error) T {
	if err != nil {
		panic(err)
	}
	return arg
}

func Must2[T any, T2 any](arg T, arg2 T2, err error) (T, T2) {
	if err != nil {
		panic(err)
	}
	return arg, arg2
}
