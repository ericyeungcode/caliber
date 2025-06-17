package caliber

type Pair[T1 any, T2 any] struct {
	First  T1
	Second T2
}

// p := NewPair("Dave", 40) // type inferred, no need to specify types like &Pair[string, int]{"Dave", 40}
func NewPair[T1, T2 any](first T1, second T2) Pair[T1, T2] {
	return Pair[T1, T2]{First: first, Second: second}
}
