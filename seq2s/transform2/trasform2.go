// Package transform2 provides functions that transform seq2[K, V] iterators in various ways.
package transform2

import (
	"iter"

	"github.com/elordeiro/goext/constraints"
	"github.com/elordeiro/goext/seq2s"
)

// Backwards returns a Seq2[K, V] with the values of the input iterator reversed.
// It takes O(n) time and space to collect all the values.
func Backwards[K, V any](seq iter.Seq2[K, V]) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		vals := seq2s.CollectPairs(seq)
		for i := len(vals) - 1; i >= 0; i-- {
			if !yield(vals[i].Left(), vals[i].Right()) {
				return
			}
		}
	}
}

// DropWhile returns a Seq2[K, V] that drops values from the input iterator until the predicate is false.
func DropWhile[K, V any](seq iter.Seq2[K, V], predicate func(K, V) bool) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		dropping := true
		for k, v := range seq {
			if dropping && predicate(k, v) {
				continue
			}
			dropping = false
			if !yield(k, v) {
				return
			}
		}
	}
}

// Filter returns a Seq2[K, V] over key-value pairs that satisfy the filter function.
func Filter[K, V any](seq iter.Seq2[K, V], filter func(K, V) bool) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k, v := range seq {
			if filter(k, v) && !yield(k, v) {
				return
			}
		}
	}
}

// ForEach applies a function to each key-value pair in the iterator.
// This function is not lazy and will consume the entire iterator.
func ForEach[K, V any](seq iter.Seq2[K, V], do func(K, V)) {
	for k, v := range seq {
		do(k, v)
	}
}

// Map returns a Seq2[K, V] over key-value pairs that are transformed by the map function.
func Map[K, V any](seq iter.Seq2[K, V], transform func(K, V) (K, V)) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k, v := range seq {
			if !yield(transform(k, v)) {
				return
			}
		}
	}
}

// OnEmpty returns a Seq2[K, V] that calls an else function only if the iterator is exhausted.
func OnEmpty[K, V any](seq iter.Seq2[K, V], callback func()) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k, v := range seq {
			if !yield(k, v) {
				return
			}
		}
		callback()
	}
}

// Reduce returns a single value that is the result of applying the reduce function to
// all values in the iterator. The function also accepts an optional starting value for
// the accumulator.
func Reduce[K, V, A any](seq iter.Seq2[K, V], reduce func(A, V) A, start ...A) A {
	var acc A
	if len(start) > 0 {
		acc = start[0]
	}
	for _, v := range seq {
		acc = reduce(acc, v)
	}
	return acc
}

// Rotate returns a Seq[V] with the values of the input rotated left by n steps.
// If n < 1 the function returns the same sequence. If n > Len(seq), the resulting
// sequence will be the same as the input sequence.
func Rotate[I constraints.Integer, K, V any](n I, seq iter.Seq2[K, V]) iter.Seq2[K, V] {
	if n < 1 {
		return seq
	}
	return seq2s.Chain(seq2s.SeqRange(n, 0, seq), seq2s.SeqRange(0, n, seq))
}

// SwapKV returns a Seq2[V, K] that swaps the keys and values of the input iterator.
func SwapKV[K, V any](seq iter.Seq2[K, V]) iter.Seq2[V, K] {
	return func(yield func(V, K) bool) {
		for k, v := range seq {
			if !yield(v, k) {
				return
			}
		}
	}
}

// TakeWhile returns a Seq2[K, V] that yields values from the input iterator until the predicate is false.
func TakeWhile[K, V any](seq iter.Seq2[K, V], predicate func(K, V) bool) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k, v := range seq {
			if !predicate(k, v) || !yield(k, v) {
				return
			}
		}
	}
}

// With returns a Seq2[K, V] that calls a function on each iteration before yielding it.
func With[K, V any](seq iter.Seq2[K, V], process func(K, V)) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k, v := range seq {
			process(k, v)
			if !yield(k, v) {
				return
			}
		}
	}
}
