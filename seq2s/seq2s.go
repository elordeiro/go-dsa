// Seq2s package provides a set of functions that operate on iterators.
// Some of the functions in this package also have a counterpart in the
// seqs package. For functions that transform the input iterator, check out
// the seq2s/transform2 package.
package seq2s

import (
	"fmt"
	"iter"
	"strings"

	"github.com/elordeiro/goext/constraints"
	"github.com/elordeiro/goext/containers/tuples"
)

func NilSeq2[K, V any](yield func(K, V) bool) {}

// Equal returns true if two Seq2[K, V] are equal. The values must be of comparable type.
// For non-comparable types, use EqualFunc2().
func Equal[K, V comparable](seq1, seq2 iter.Seq2[K, V]) bool {
	p1, stop := iter.Pull2(seq1)
	defer stop()
	p2, stop := iter.Pull2(seq2)
	defer stop()

	k1, v1, ok1 := p1()
	k2, v2, ok2 := p2()
	for ok1 && ok2 {
		if k1 != k2 || v1 != v2 {
			return false
		}
		k1, v1, ok1 = p1()
		k2, v2, ok2 = p2()
	}
	return ok1 == ok2
}

// EqualFunc returns true if two Seq2[K, V] are equal. The values are compared using the
// provided function.
func EqualFunc[K1, K2, V1, V2 any](
	seq1 iter.Seq2[K1, V1],
	seq2 iter.Seq2[K2, V2],
	eq func(p1 tuples.Pair[K1, V1], p2 tuples.Pair[K2, V2]) bool) bool {
	p1, stop := iter.Pull2(seq1)
	defer stop()
	p2, stop := iter.Pull2(seq2)
	defer stop()

	k1, v1, ok1 := p1()
	k2, v2, ok2 := p2()
	for ok1 && ok2 {
		if !eq(tuples.NewPair(k1, v1), tuples.NewPair(k2, v2)) {
			return false
		}
		k1, v1, ok1 = p1()
		k2, v2, ok2 = p2()
	}
	return ok1 == ok2
}

// EqualUnordered returns true if two Seq2[K, V] are equal. The values must be of comparable type.
// The values are compared in an unordered fashion.
func EqualUnordered[K, V comparable](seq1, seq2 iter.Seq2[K, V]) bool {
	m1 := make(map[K]V)
	for k, v := range seq1 {
		m1[k] = v
	}
	m2 := make(map[K]V)
	for k, v := range seq2 {
		m2[k] = v
	}
	if len(m1) != len(m2) {
		return false
	}
	for k, v1 := range m1 {
		v2, ok := m2[k]
		if !ok || v1 != v2 {
			return false
		}
	}
	return true
}

// FromSlice returns a Seq2[V, V] from a slice of 2-element arrays.
func FromSlice[V any](slice [][2]V) iter.Seq2[V, V] {
	return func(yield func(V, V) bool) {
		for _, v := range slice {
			if !yield(v[0], v[1]) {
				break
			}
		}
	}
}

// Keys returns a Seq[K] over the keys of a Seq2[K, V].
func Keys[K, V any](seq iter.Seq2[K, V]) iter.Seq[K] {
	return func(yield func(K) bool) {
		for k := range seq {
			if !yield(k) {
				return
			}
		}
	}
}

// Len returns the length of a Seq2[K, V].
// Note that the function consumes the sequence, so if the seq passed is an one-time
// iterator, it will be consumed and cannot be used again.
func Len[K, V any](seq iter.Seq2[K, V]) int {
	total := 0
	for range seq {
		total++
	}
	return total
}

// String returns a string representation of a Seq2[K, V].
// Note that the function consumes the iterator.
func String[V, K any](seq iter.Seq2[K, V]) string {
	var sb strings.Builder
	sb.WriteString("=>[")
	first := true
	for k, v := range seq {
		if first {
			first = false
		} else {
			sb.WriteByte(' ')
		}
		sb.WriteString(fmt.Sprintf("(%v %v)", k, v))
	}
	sb.WriteByte(']')
	return sb.String()
}

// Values returns a Seq[V] over the values of a Seq2[K, V].
func Values[K, V any](seq iter.Seq2[K, V]) iter.Seq[V] {
	return func(yield func(V) bool) {
		for _, v := range seq {
			if !yield(v) {
				return
			}
		}
	}
}

// Chain returns a Seq2[K, V] that chains the values of multiple input iterators
// into a single iterator.
func Chain[K, V any](seqs ...iter.Seq2[K, V]) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for _, seq := range seqs {
			for k, v := range seq {
				if !yield(k, v) {
					return
				}
			}
		}
	}
}

// Collect returns a slice of 2-element slices of type V with all the values of a Seq2[K, V].
// Note that both the keys and values in the sequence must be of the same type. For different
// types, use CollectPairs.
func Collect[V any](seq iter.Seq2[V, V]) [][2]V {
	var vals [][2]V
	for k, v := range seq {
		vals = append(vals, [2]V{k, v})
	}
	return vals
}

// CollectPairs returns a slice of type Pair[K, V] with all the values of a Seq2[K, V].
func CollectPairs[K, V any](seq iter.Seq2[K, V]) []tuples.Pair[K, V] {
	var vals []tuples.Pair[K, V]
	for k, v := range seq {
		vals = append(vals, tuples.NewPair(k, v))
	}
	return vals
}

// Cycle returns a Seq2[K, V] that cycles through the values of the input iterator.
// The iterator is infinite unless it is stopped by the caller.
func Cycle[K, V any](seq iter.Seq2[K, V]) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for {
			for k, v := range seq {
				if !yield(k, v) {
					return
				}
			}
		}
	}
}

// Enumerate returns a Seq2[Integer, Pair] over index-pair pairs in the Iterable.
// The Pair will contain the key-value with the key as the left and the value as the right.
// The start argument specifies the starting index.
func Enumerate[I constraints.Integer, K, V any](
	start I,
	seq iter.Seq2[K, V],
) iter.Seq2[I, tuples.Pair[K, V]] {
	return func(yield func(I, tuples.Pair[K, V]) bool) {
		i := start
		for k, v := range seq {
			if !yield(i, tuples.NewPair(k, v)) {
				return
			}
			i++
		}
	}
}

// MultiUse takes a single use Seq[V] and returns a multi-use Seq[V].
func MultiUse[K, V any](seq iter.Seq2[K, V]) iter.Seq2[K, V] {
	vals := CollectPairs(seq)
	return func(yield func(K, V) bool) {
		for _, p := range vals {
			if !yield(p.Left(), p.Right()) {
				return
			}
		}
	}
}

// Repeat returns a Seq2[K, V] that yields the same key-value pair n times.
// If n is negative, the iterator is infinite.
func Repeat[I constraints.Integer, K, V any](key K, val V, n I) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for i := I(0); i != n; i++ {
			if !yield(key, val) {
				return
			}
		}
	}
}

// SeqRange returns a Seq2[K, V] that yields the values of a Seq2[K, V] between [start, end).
//
//	If start is 0, the function is equivalent to a Take2(n, seq), where n is end.
//	If start > end, the function is equivalent to a Drop2(n, seq), where n is start.
func SeqRange[I constraints.Integer, K, V any](start, end I, seq iter.Seq2[K, V]) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		var i I = 0
		for k, v := range seq {
			if i < start {
				i++
				continue
			}
			if i == end {
				return
			}
			if !yield(k, v) {
				return
			}
			i++
		}
	}
}

// Zip returns a Seq2[Pair, Pair] over values from a Seq2[K1, V1] and a Seq2[K2, V2].
// The iteration stops when either of the sequences are exhausted.
// In other words, the length of the resulting sequence is the minimum
// of the lengths of the input sequences.
func Zip[K1, V1, K2, V2 any](
	seq1 iter.Seq2[K1, V1],
	seq2 iter.Seq2[K2, V2],
) iter.Seq2[tuples.Pair[K1, V1], tuples.Pair[K2, V2]] {
	return func(yield func(tuples.Pair[K1, V1], tuples.Pair[K2, V2]) bool) {
		p1, stop := iter.Pull2(seq1)
		defer stop()
		p2, stop := iter.Pull2(seq2)
		defer stop()
		for {
			k1, v1, ok1 := p1()
			k2, v2, ok2 := p2()
			if (!ok1 || !ok2) || !yield(tuples.NewPair(k1, v1), tuples.NewPair(k2, v2)) {
				return
			}
		}
	}
}
