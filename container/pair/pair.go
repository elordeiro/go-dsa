package pair

import "fmt"

type Interface[L, R any] interface {
	Left() L
	Right() R
}

type Pair[L, R any] struct {
	L L // Left
	R R // Right
}

func (p Pair[L, R]) Left() L {
	return p.L
}

func (p Pair[L, R]) Right() R {
	return p.R
}

type Pairs[L, R any] []Pair[L, R]

func NewPair[L, R any](left L, right R) Pair[L, R] {
	return Pair[L, R]{L: left, R: right}
}

func NewPairs[L, R any](pairs ...Pair[L, R]) []Pair[L, R] {
	return []Pair[L, R](pairs)
}

// String returns a string representation of a Pair.
func (p Pair[K, V]) String() string {
	return fmt.Sprintf("%v:%v", p.L, p.R)
}
