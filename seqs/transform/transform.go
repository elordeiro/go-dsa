// Package transform provides functions that transform seq[V] iterators in various ways.
package transform

import (
	"iter"
	"slices"

	"github.com/elordeiro/goext/constraints"
	"github.com/elordeiro/goext/containers/tuples/ituples"
	"github.com/elordeiro/goext/seqs"
)

// Backwards returns a Seq[V] that iterates over the values in reverse order.
// It takes O(n) time and space to collect all the values.
func Backwards[V any](seq iter.Seq[V]) iter.Seq[V] {
	return func(yield func(V) bool) {
		vals := slices.Collect(seq)
		for i := len(vals) - 1; i >= 0; i-- {
			if !yield(vals[i]) {
				return
			}
		}
	}
}

// DropWhile returns a Seq[V] that drops values from the input iterator until the predicate is false.
func DropWhile[V any](seq iter.Seq[V], predicate func(V) bool) iter.Seq[V] {
	return func(yield func(V) bool) {
		dropping := true
		for v := range seq {
			if dropping && predicate(v) {
				continue
			}
			dropping = false
			if !yield(v) {
				return
			}
		}
	}
}

// Filter returns a Seq[V]  over values that satisfy the filter function.
func Filter[V any](seq iter.Seq[V], filter func(V) bool) iter.Seq[V] {
	return func(yield func(V) bool) {
		for v := range seq {
			if filter(v) && !yield(v) {
				return
			}
		}
	}
}

// ForEach applies a function to each value in the iterator.
// This function is not lazy and will consume the entire iterator.
func ForEach[V any](seq iter.Seq[V], do func(V)) {
	for v := range seq {
		do(v)
	}
}

// Map returns a Seq[V] over values that are transformed by the map function.
func Map[V any](seq iter.Seq[V], transform func(V) V) iter.Seq[V] {
	return func(yield func(V) bool) {
		for v := range seq {
			if !yield(transform(v)) {
				return
			}
		}
	}
}

// OnEmpty returns a Seq[V] that calls an else function only if the iterator is exhausted.
func OnEmpty[V any](seq iter.Seq[V], callback func()) iter.Seq[V] {
	return func(yield func(V) bool) {
		for v := range seq {
			if !yield(v) {
				return
			}
		}
		callback()
	}
}

// Reduce returns a single value that is the result of applying the reduce function to
// all values in the iterator. The function alse accepts an optional starting value for
// the accumulator.
func Reduce[V any, A any](seq iter.Seq[V], reduce func(A, V) A, start ...A) A {
	var acc A
	if len(start) > 0 {
		acc = start[0]
	}
	for v := range seq {
		acc = reduce(acc, v)
	}
	return acc
}

// Rotate returns a Seq[V] with the values of the input rotated left by n steps.
// If n < 1, the function returns the same sequence. If n > Len(seq), the resulting
// sequence will be the same as the input sequence.
func Rotate[I constraints.Integer, V any](n I, seq iter.Seq[V]) iter.Seq[V] {
	if n < 1 {
		return seq
	}
	return seqs.Chain(seqs.SeqRange(n, 0, seq), seqs.SeqRange(0, n, seq))
}

// TakeWhile returns a Seq[V] that yields values from the input iterator until the predicate is false.
func TakeWhile[V any](seq iter.Seq[V], predicate func(V) bool) iter.Seq[V] {
	return func(yield func(V) bool) {
		for v := range seq {
			if !predicate(v) || !yield(v) {
				return
			}
		}
	}
}

// Unpair converts a Seq[Pair[L, R]] values to a Seq2[L, R].
func Unpair[Pair ituples.Pair[L, R], L, R any](seq iter.Seq[Pair]) iter.Seq2[L, R] {
	return func(yield func(L, R) bool) {
		for p := range seq {
			if !yield(p.Left(), p.Right()) {
				return
			}
		}
	}
}

// With returns a Seq[V] that calls a function with each element before yielding it.
func With[V any](seq iter.Seq[V], process func(V)) iter.Seq[V] {
	return func(yield func(V) bool) {
		for v := range seq {
			process(v)
			if !yield(v) {
				return
			}
		}
	}
}
