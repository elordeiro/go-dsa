// Seqs package provides a set of functions that operate on iterators.
// Some of the functions in this package also have a counterpart in the
// seq2s package. For functions that transform the input iterator, check out
// the seqs/transform package.
package seqs

import (
	"fmt"
	"iter"
	"log"
	"strings"

	"github.com/elordeiro/goext/constraints"
)

func NilSeq[V any](yield func(V) bool) {}

// Equal returns true if two Seq[V] are equal. The values must be of comparable type.
// For non-comparable types, use EqualFunc().
func Equal[V comparable](seq1, seq2 iter.Seq[V]) bool {
	p1, stop := iter.Pull(seq1)
	defer stop()
	p2, stop := iter.Pull(seq2)
	defer stop()

	v1, ok1 := p1()
	v2, ok2 := p2()
	for ok1 && ok2 {
		if v1 != v2 {
			return false
		}
		v1, ok1 = p1()
		v2, ok2 = p2()
	}
	return ok1 == ok2
}

// EqualFunc returns true if two Seq[V] are equal. The values are compared using the
// provided function.
func EqualFunc[V1, V2 any](seq1 iter.Seq[V1], seq2 iter.Seq[V2], eq func(V1, V2) bool) bool {
	p1, stop := iter.Pull(seq1)
	defer stop()
	p2, stop := iter.Pull(seq2)
	defer stop()

	v1, ok1 := p1()
	v2, ok2 := p2()
	for ok1 && ok2 {
		if !eq(v1, v2) {
			return false
		}
		v1, ok1 = p1()
		v2, ok2 = p2()
	}
	return ok1 == ok2
}

// EqualUnordered returns true if two Seq[V] are equal. The values must be of comparable type.
// The values are compared in an unordered fashion.
func EqualUnordered[V comparable](seq1, seq2 iter.Seq[V]) bool {
	m1 := make(map[V]struct{})
	for v := range seq1 {
		m1[v] = struct{}{}
	}

	m2 := make(map[V]struct{})
	for v := range seq2 {
		m2[v] = struct{}{}
	}

	if len(m1) != len(m2) {
		return false
	}

	for v := range m1 {
		if _, ok := m2[v]; !ok {
			return false
		}
	}

	return true
}

func IsEmpty[V any](seq iter.Seq[V]) bool {
	for range seq {
		return false
	}
	return true
}

// Len returns the length of a Seq[V].
// Note that the function consumes the sequence, so if the seq passed is an one-time
// iterator, it will be consumed and cannot be used again.
func Len[V any](seq iter.Seq[V]) int {
	total := 0
	for range seq {
		total++
	}
	return total
}

// String returns a string representation of a Seq[V].
// Note that the function consumes the iterator.
func String[V any](seq iter.Seq[V]) string {
	var sb strings.Builder
	sb.WriteString("=>[")
	first := true
	for v := range seq {
		if first {
			first = false
		} else {
			sb.WriteByte(' ')
		}
		sb.WriteString(fmt.Sprint(v))
	}
	sb.WriteByte(']')
	return sb.String()
}

// Chain returns a Seq[V] that chains the values of multiple input iterators
// into a single iterator.
func Chain[V any](seqs ...iter.Seq[V]) iter.Seq[V] {
	return func(yield func(V) bool) {
		for _, seq := range seqs {
			for v := range seq {
				if !yield(v) {
					return
				}
			}
		}
	}
}

// Collect returns a slice of type V with all the values of a Seq[V].
func Collect[V any](seq iter.Seq[V]) []V {
	var vals []V
	for v := range seq {
		vals = append(vals, v)
	}
	return vals
}

// Count returns Seq[Integer] that counts up from a given number.
// The step argument is optional and defaults to 1. For a decreasing count,
// provide a negative step. If the step is 0, the function yields the same
// value indefinitely.
// The iterator is infinite unless it is stopped by the caller.
func Count[I constraints.Integer](start I, step ...I) iter.Seq[I] {
	return func(yield func(I) bool) {
		var s I = 1
		if len(step) != 0 {
			s = step[0]
		}
		for ; ; start += s {
			if !yield(start) {
				return
			}
		}
	}
}

// Cycle returns a Seq[V] that cycles through the values of the input iterator.
// The iterator is infinite unless it is stopped by the caller.
func Cycle[V any](seq iter.Seq[V]) iter.Seq[V] {
	return func(yield func(V) bool) {
		for {
			for v := range seq {
				if !yield(v) {
					return
				}
			}
		}
	}
}

// Enumerate returns a Seq2[Integer, V] over index-value pairs in the Iterable.
// The start argument specifies the starting index.
func Enumerate[I constraints.Integer, V any](start I, seq iter.Seq[V]) iter.Seq2[I, V] {
	return func(yield func(I, V) bool) {
		i := start
		for v := range seq {
			if !yield(i, v) {
				return
			}
			i++
		}
	}
}

// MultiUse takes a single use Seq[V] and returns a multi-use Seq[V].
func MultiUse[V any](seq iter.Seq[V]) iter.Seq[V] {
	vals := Collect(seq)
	return func(yield func(V) bool) {
		for _, v := range vals {
			if !yield(v) {
				return
			}
		}
	}
}

// Range returns a Seq[Signed]. It is similar to:
//
//	for i := start; i < end; i += step.
//
//	If only one argument is provided, range is [0, end).
//	If two arguments are provided, range is [start, end).
//	If three arguments are provided, range is [start, end) with the given step.
//	If an infinite loop is detected, the function logs a message and returns an empty seq.
func Range[S constraints.Signed](val S, vals ...S) iter.Seq[S] {
	var start, end S
	var step S = 1

	switch len(vals) {
	case 0:
		start, end = 0, val
	case 1:
		start, end = val, vals[0]
		if start > end {
			step = -1
		}
	case 2:
		start, end, step = val, vals[0], vals[1]
	default:
		start, end, step = val, vals[0], vals[1]
		log.Println("ignoring extra arguments in Range()")
	}

	if step == 0 {
		log.Println("infinite loop in Range(); step == 0")
		return NilSeq
	}

	if start > end && step >= 0 {
		log.Println("empty iterator in Range(); start > end && step >= 0")
	}

	if start < end && step < 0 {
		log.Println("empty iterator in Range(); start < end && step < 0")
	}

	var cond func(S) bool
	if step < 0 {
		cond = func(i S) bool { return i > end }
	} else {
		cond = func(i S) bool { return i < end }
	}

	return func(yield func(S) bool) {
		for i := start; cond(i); i += step {
			if !yield(i) {
				return
			}
		}
	}
}

// Repeat returns a Seq[V] that yields the same value n times.
// If n is negative, the iterator is infinite.
func Repeat[I constraints.Integer, V any](val V, n I) iter.Seq[V] {
	return func(yield func(V) bool) {
		for i := I(0); i != n; i++ {
			if !yield(val) {
				return
			}
		}
	}
}

// SeqRange returns a Seq[V] that yields the values of a Seq[V] between [start, end).
//
//	If start is 0, the function is equivalent to a Take(n, seq), where n = end.
//	If start > end, the function is equivalent to a Drop(n, seq), where n = start.
func SeqRange[I constraints.Integer, V any](start, end I, seq iter.Seq[V]) iter.Seq[V] {
	return func(yield func(V) bool) {
		var i I = 0
		for v := range seq {
			if i < start {
				i++
				continue
			}
			if i == end {
				return
			}
			if !yield(v) {
				return
			}
			i++
		}
	}
}

// Zip returns a Seq2[V1, V2] over values from a Seq[V1] and a Seq[V2].
// The iteration stops when either of the sequences are exhausted.
// In other words, the length of the resulting sequence is the minimum
// of the lengths of the input sequences.
func Zip[V1, V2 any](seq1 iter.Seq[V1], seq2 iter.Seq[V2]) iter.Seq2[V1, V2] {
	return func(yield func(V1, V2) bool) {
		p1, stop := iter.Pull(seq1)
		defer stop()
		p2, stop := iter.Pull(seq2)
		defer stop()
		for {
			v1, ok1 := p1()
			v2, ok2 := p2()
			if (!ok1 || !ok2) || !yield(v1, v2) {
				return
			}
		}
	}
}
