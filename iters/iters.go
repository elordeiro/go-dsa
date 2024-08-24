// Iters package provides a set of functions that operate on iterators.
package iters

import (
	"fmt"
	"iter"
	"log"
	"maps"
	"reflect"
	"slices"

	"golang.org/x/exp/constraints"
)

type (
	Seq[V any]                 iter.Seq[V]     // Same as iter.Seq[V] but accessible in this package
	Seq2[K, V any]             iter.Seq2[K, V] // Same as iter.Seq2[K, V] but accessible in this package
	ItSlice[V any]             []V             // Generic slice type
	ItMap[K comparable, V any] map[K]V         // Generic map type

	Signed  constraints.Signed  // Signed constraint
	Integer constraints.Integer // Integer constraint
	Ord     constraints.Ordered // Ordered constraint
)

// Pair is a key-value pair for Seq2 iterators that need to return more than 2 values.
type Pair[K, V any] struct {
	Key   K
	Value V
}

// IIterable is an interface that defines the methods for a sequence of values or index-value pairs.
type IIterable[V any] interface {
	Len() int
	Values() Seq[V]
	All() Seq2[int, V]
}

// IIterable2 is an interface that defines the methods of Iterable and adds a method to get the keys.
type IIterable2[K, V any] interface {
	Len() int
	Keys() Seq[K]
	Values() Seq[V]
	All() Seq2[K, V]
}

// ----------------------------------------------------------------------------
// Constructors
// ----------------------------------------------------------------------------

// Iterable returns a new ItSlice. ItSlices can be passed to any function that
// also accepts Seq[V] iterators, while preserving the underlying slice, which
// means they can still be indexed and sliced.
func Iterable[S ~[]V, V any](s S) ItSlice[V] {
	return ItSlice[V](s)
}

// Iterable2 returns a new ItMap. ItMaps can be passed to any function that
// also accepts Seq2[K, V] iterators, while preserving the underlying map, which
// means they can still be indexed with keys.
func Iterable2[M ~map[K]V, K comparable, V any](m M) ItMap[K, V] {
	return ItMap[K, V](m)
}

// ----------------------------------------------------------------------------
// Interface implementations
// ----------------------------------------------------------------------------

// Len returns the length of the Seq.
func (s Seq[V]) Len() int {
	count := 0
	for range s {
		count++
	}
	return count
}

// Values simply returns the Seq.
func (s Seq[V]) Values() Seq[V] {
	return s
}

// All returns a Seq2[int, V] iterator over index-value pairs of the Seq.
func (s Seq[V]) All() Seq2[int, V] {
	return func(yield func(int, V) bool) {
		i := 0
		for v := range s {
			if !yield(i, v) {
				return
			}
			i++
		}
	}
}

// Len returns the length of the Seq2.
func (s Seq2[K, V]) Len() int {
	count := 0
	for range s {
		count++
	}
	return count
}

// Keys returns a Seq iterator over the keys of the Seq2.
func (s Seq2[K, V]) Keys() Seq[K] {
	return func(yield func(K) bool) {
		for k := range s {
			if !yield(k) {
				return
			}
		}
	}
}

// Values returns a Seq iterator over the values of the Seq2.
func (s Seq2[K, V]) Values() Seq[V] {
	return func(yield func(V) bool) {
		for _, v := range s {
			if !yield(v) {
				return
			}
		}
	}
}

// All returns a Seq2[K, V] iterator over the index-value pairs of the Seq2.
func (s Seq2[K, V]) All() Seq2[K, V] {
	return s
}

// Len returns the length of the underlying slice.
func (s ItSlice[V]) Len() int {
	return len(s)
}

// Values returns a Seq[V] iterator over the values of the slice.
func (s ItSlice[V]) Values() Seq[V] {
	return Surf(slices.Values(s))
}

// All returns a Seq2[int, V] iterator over index-value pairs of the slice.
func (s ItSlice[V]) All() Seq2[int, V] {
	return Surf2(slices.All(s))
}

// Len returns the length of the underlying map.
func (m ItMap[K, V]) Len() int {
	return len(m)
}

// Keys returns a Seq[K] iterator over the keys of the map.
func (m ItMap[K, V]) Keys() Seq[K] {
	return Surf(maps.Keys(m))
}

// Values returns a Seq[V] iterator over the values of the map.
func (m ItMap[K, V]) Values() Seq[V] {
	return Surf(maps.Values(m))
}

// All returns a Seq2[K, V] iterator over the key-value pairs of the map.
func (m ItMap[K, V]) All() Seq2[K, V] {
	return Surf2(maps.All(m))
}

// ----------------------------------------------------------------------------
// Iter utils
// ----------------------------------------------------------------------------

// Surf Surfaces a Seq[V] iterator.
// It takes an iter.Seq[V] iterator and returns a Seq[V] iterator.
func Surf[V any](seq iter.Seq[V]) Seq[V] {
	return Seq[V](seq)
}

// Surf2 Surfaces a Seq2[K, V] iterator.
// It takes an iter.Seq2[K, V] iterator and returns a Seq2Seq2[K, V] iterator.
func Surf2[K, V any](seq iter.Seq2[K, V]) Seq2[K, V] {
	return Seq2[K, V](seq)
}

// Sink sinks a Seq[V] iterator
// Returns an iter.Seq[V] iterator.
func (seq Seq[V]) Sink() iter.Seq[V] {
	return iter.Seq[V](seq)
}

// Sink2 sinks a Seq2[K, V] iterator
// Returns an iter.Seq2[K, V] iterator.
func (seq Seq2[K, V]) Sink() iter.Seq2[K, V] {
	return iter.Seq2[K, V](seq)
}

// ----------------------------------------------------------------------------
// Int iterators
// ----------------------------------------------------------------------------

// Range returns an iterator over a range of integers.
// If only one argument is provided, it is the end of the range.
// If two arguments are provided, they are the start and end of the range.
// If three arguments are provided, they are the start, end, and step of the range.
// If an infinite loop is detected, the function panics.
func Range[S Signed](val S, vals ...S) Seq[S] {
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
		log.Println("Ignoring extra arguments in Range function")
	}

	if step == 0 {
		log.Panic("Infinite loop in Range(). Step == 0")
	}

	if start > end && step >= 0 {
		log.Panic("Infinite loop in Range(). Start > end && step >= 0")
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

// Count returns an iterator that counts up from a given number.
// The iterator is infinite unless it is stopped by the caller.
func Count[I Integer](i I) Seq[I] {
	return func(yield func(I) bool) {
		for ; ; i++ {
			if !yield(i) {
				return
			}
		}
	}
}

// CountDown returns an iterator that counts down from a given number.
// The iterator is infinite unless it is stopped by the caller.
func CountDown[I Integer](i I) Seq[I] {
	return func(yield func(I) bool) {
		for ; ; i-- {
			if !yield(i) {
				return
			}
		}
	}
}

// ----------------------------------------------------------------------------
// Itertools functions
// ----------------------------------------------------------------------------

// Enumerate returns a Seq2[int, V] iterator over index-value pairs in the Iterable.
// The start argument specifies the starting index.
func Enumerate[I Integer, V any](start I, seq IIterable[V]) Seq2[I, V] {
	return func(yield func(I, V) bool) {
		i := start
		for v := range seq.Values() {
			if !yield(i, v) {
				return
			}
			i++
		}
	}
}

// Enumerate2 returns a Seq2[int, Pair] iterator over index-pair pairs in the Iterable.
// A pair is a key-value pair for Seq2 iterators that need to return more than 2 values.
// The start argument specifies the starting index.
func Enumerate2[I Integer, K, V any](start I, seq IIterable2[K, V]) Seq2[I, Pair[K, V]] {
	return func(yield func(I, Pair[K, V]) bool) {
		i := start
		for k, v := range seq.All() {
			if !yield(i, Pair[K, V]{k, v}) {
				return
			}
			i++
		}
	}
}

// Zip returns a Seq2[V1, V2] iterator over values from two sequences.
// The iteration stops when either of the sequences are exhausted.
// In other words, the length of the resulting sequence is the minimum
// of the lengths of the input sequences.
func Zip[V1, V2 any](seq1 IIterable[V1], seq2 IIterable[V2]) Seq2[V1, V2] {
	return func(yield func(V1, V2) bool) {
		p1, stop := iter.Pull(seq1.Values().Sink())
		defer stop()
		p2, stop := iter.Pull(seq2.Values().Sink())
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

// Zip2 returns a Seq2[Pair, Pair] iterator over values from two sequences.
// The iteration stops when either of the sequences are exhausted.
// In other words, the length of the resulting sequence is the minimum
// of the lengths of the input sequences.
func Zip2[K1, V1, K2, V2 any](seq1 IIterable2[K1, V1], seq2 IIterable2[K2, V2]) Seq2[Pair[K1, V1], Pair[K2, V2]] {
	return func(yield func(Pair[K1, V1], Pair[K2, V2]) bool) {
		p1, stop := iter.Pull2(seq1.All().Sink())
		defer stop()
		p2, stop := iter.Pull2(seq2.All().Sink())
		defer stop()
		for {
			k1, v1, ok1 := p1()
			k2, v2, ok2 := p2()
			if (!ok1 || !ok2) || !yield(Pair[K1, V1]{k1, v1}, Pair[K2, V2]{k2, v2}) {
				return
			}
		}
	}
}

// Repeat returns a Seq[V] iterator that yields the same value n times.
// If n is negative, the iterator is infinite.
func Repeat[I Integer, V any](val V, n I) Seq[V] {
	return func(yield func(V) bool) {
		for i := I(0); i != n; i++ {
			if !yield(val) {
				return
			}
		}
	}
}

// Reapeat2 returns a Seq2[K, V] iterator that yields the same key-value pair n times.
// If n is negative, the iterator is infinite.
func Repeat2[I Integer, K, V any](key K, val V, n I) Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for i := I(0); i != n; i++ {
			if !yield(key, val) {
				return
			}
		}
	}
}

// Cycle returns a Seq[V] iterator that cycles through the values of the input iterator.
// The iterator is infinite unless it is stopped by the caller.
func Cycle[V any](seq IIterable[V]) Seq[V] {
	return func(yield func(V) bool) {
		for {
			for v := range seq.Values() {
				if !yield(v) {
					return
				}
			}
		}
	}
}

// Cycle returns a Seq2[K, V] iterator that cycles through the values of the input iterator.
func Cycle2[K, V any](seq IIterable2[K, V]) Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for {
			for k, v := range seq.All() {
				if !yield(k, v) {
					return
				}
			}
		}
	}
}

// Chain returns a Seq[V] iterator that chains the values of multiple input iterators.
func Chain[V any](seqs ...IIterable[V]) Seq[V] {
	return func(yield func(V) bool) {
		for _, seq := range seqs {
			for v := range seq.Values() {
				if !yield(v) {
					return
				}
			}
		}
	}
}

// Chain2 returns a Seq2[K, V] iterator that chains the values of multiple input iterators.
func Chain2[K, V any](seqs ...IIterable2[K, V]) Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for _, seq := range seqs {
			for k, v := range seq.All() {
				if !yield(k, v) {
					return
				}
			}
		}
	}
}

// Backwards returns a Seq[V] iterator that iterates over the values in reverse order.
// If seq is an iterator, it takes O(n) time and space to collect all the values.
func Backwards[V any](seq IIterable[V]) Seq[V] {
	return func(yield func(V) bool) {
		var vals []V
		if _, ok := seq.(Seq[V]); ok {
			vals = slices.Collect(seq.Values().Sink())
		} else {
			vals = seq.(ItSlice[V])
		}
		for i := len(vals) - 1; i >= 0; i-- {
			if !yield(vals[i]) {
				return
			}
		}
	}
}

// Backwards returns a Seq2[K, V] iterator that iterates over the values in reverse order.
// If seq is an iterator, it takes O(n) time and space to collect all the values.
// If seq is ItMap, the function return the Seq2 from All() since map ordering is not guaranteed
func Backwards2[K, V any](seq IIterable2[K, V]) Seq2[K, V] {
	if _, ok := seq.(Seq2[K, V]); !ok {
		return seq.All()
	}
	return func(yield func(K, V) bool) {
		vals := []Pair[K, V]{}
		for k, v := range seq.All() {
			vals = append(vals, Pair[K, V]{k, v})
		}
		for i := len(vals) - 1; i >= 0; i-- {
			p := vals[i]
			if !yield(p.Key, p.Value) {
				return
			}
		}
	}
}

// Take returns a Seq[V] iterator that yields the first n values of the input iterator.
func Take[I Integer, V any](n I, seq IIterable[V]) Seq[V] {
	return func(yield func(V) bool) {
		var i I = 0
		for v := range seq.Values() {
			if i == n {
				return
			}
			if !yield(v) {
				return
			}
			i++
		}
	}
}

// Take2 returns a Seq2[K, V] iterator that yields the first n values of the input iterator.
func Take2[I Integer, K, V any](n I, seq IIterable2[K, V]) Seq2[K, V] {
	return func(yield func(K, V) bool) {
		var i I = 0
		for k, v := range seq.All() {
			if i == n {
				return
			}
			if !yield(k, v) {
				return
			}
			i++
		}
	}
}

// Drop returns a Seq[V] iterator that skips the first n values of the input iterator.
func Drop[I Integer, V any](n I, seq IIterable[V]) Seq[V] {
	return func(yield func(V) bool) {
		var i I = 0
		for v := range seq.Values() {
			if i < n {
				i++
				continue
			}
			if !yield(v) {
				return
			}
			i++
		}
	}
}

// Drop returns a Seq2[K, V] iterator that skips the first n values of the input iterator.
func Drop2[I Integer, K, V any](n I, seq IIterable2[K, V]) Seq2[K, V] {
	return func(yield func(K, V) bool) {
		var i I = 0
		for k, v := range seq.All() {
			if i < n {
				i++
				continue
			}
			if !yield(k, v) {
				return
			}
			i++
		}
	}
}

// TakeBetween returns a Seq[V] iterator that yields values between start and end.
// The start and end values are inclusive.
func TakeBetween[I Integer, V any](start, end I, seq IIterable[V]) Seq[V] {
	return func(yield func(V) bool) {
		var i I = 0
		for v := range seq.Values() {
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

// TakeBetween returns a Seq2[K, V] iterator that yields values between start and end.
// The start and end values are inclusive.
func TakeBetween2[I Integer, K, V any](start, end I, seq IIterable2[K, V]) Seq2[K, V] {
	return func(yield func(K, V) bool) {
		var i I = 0
		for k, v := range seq.All() {
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

// Rotate returns a Seq[V] iterator that rotates the values
// of the input iterator by n steps. A positive n rotates the values
// to the left, and a negative n rotates the values to the right.
// If seq is an iterator, it is advised to pass the length of it as
// the third argument to avoid a call to Len() which is an O(n) operation.
func Rotate[I Integer, V any](n I, seq IIterable[V], len ...I) Seq[V] {
	var length I
	if len == nil {
		length = I(seq.Len())
	}
	n = n % length
	if n == 0 {
		return seq.Values()
	}
	if n < 0 {
		return Chain(Drop(length+n, seq), Take(length+n, seq))
	} else {
		return Chain(Drop(n, seq), Take(n, seq))
	}
}

// Rotate returns a Seq2[K, V] iterator that rotates the values
// of the input iterator by n steps. A positive n rotates the values
// to the left, and a negative n rotates the values to the right.
// If seq is an iterator, it is advised to pass the length of it as
// the third argument to avoid a call to Len() which is an O(n) operation.
func Rotate2[I Integer, K, V any](n I, seq IIterable2[K, V], len ...I) Seq2[K, V] {
	var length I
	if len == nil {
		length = I(seq.Len())
	}
	if n == 0 || length == 0 || n%length == 0 {
		return seq.All()
	}
	n = n % length
	if n < 0 {
		return Chain2(Drop2(length+n, seq), Take2(length+n, seq))
	} else {
		return Chain2(Drop2(n, seq), Take2(n, seq))
	}
}

// ----------------------------------------------------------------------------
// Functools functions
// ----------------------------------------------------------------------------

// ForEach applies a function to each value in the iterator.
// This function is not lazy and will consume the entire iterator.
// Note that break, continue, and return statements are not supported in the function.
func ForEach[V any](seq IIterable[V], do func(V)) {
	for v := range seq.Values() {
		do(v)
	}
}

// ForEach2 applies a function to each key-value pair in the iterator.
// This function is not lazy and will consume the entire iterator.
// Note that break, continue, and return statements are not supported in the function.
func ForEach2[K, V any](seq IIterable2[K, V], do func(K, V)) {
	for k, v := range seq.All() {
		do(k, v)
	}
}

// Filter returns a Seq[V] iterator over values that satisfy the filter function.
func Filter[V any](seq IIterable[V], filter func(V) bool) Seq[V] {
	return func(yield func(V) bool) {
		for v := range seq.Values() {
			if filter(v) && !yield(v) {
				return
			}
		}
	}
}

// Filter returns a Seq2[K, V] iterator over key-value pairs that satisfy the filter function.
func Filter2[K, V any](seq IIterable2[K, V], filter func(K, V) bool) Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k, v := range seq.All() {
			if filter(k, v) && !yield(k, v) {
				return
			}
		}
	}
}

// Map returns a Seq[V] iterator over values that are transformed by the map function.
func Map[V any](seq IIterable[V], mapf func(V) V) Seq[V] {
	return func(yield func(V) bool) {
		for v := range seq.Values() {
			if !yield(mapf(v)) {
				return
			}
		}
	}
}

// Map returns a Seq2[K, V] iterator over key-value pairs that are transformed by the map function.
func Map2[K, V any](seq IIterable2[K, V], mapf func(K, V) (K, V)) Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k, v := range seq.All() {
			if !yield(mapf(k, v)) {
				return
			}
		}
	}
}

// Reduce returns a single value that is the result of applying the reduce function to all values in the iterator.
func Reduce[V any](seq IIterable[V], reduce func(V, V) V, start ...V) V {
	var acc V
	if len(start) > 0 {
		acc = start[0]
	}
	for v := range seq.Values() {
		acc = reduce(acc, v)
	}
	return acc
}

// Reduce returns a single value that is the result of applying the reduce function to all values in the iterator.
func Reduce2[K, V any](seq IIterable2[K, V], reduce func(V, V) V, start ...V) V {
	var acc V
	if len(start) > 0 {
		acc = start[0]
	}
	for v := range seq.Values() {
		acc = reduce(acc, v)
	}
	return acc
}

// TakeWhile returns a Seq[V] iterator that yields values from the input iterator until the predicate is false.
func TakeWhile[V any](seq IIterable[V], predicate func(V) bool) Seq[V] {
	return func(yield func(V) bool) {
		for v := range seq.Values() {
			if !predicate(v) || !yield(v) {
				return
			}
		}
	}
}

// TakeWhile returns a Seq2[K, V] iterator that yields values from the input iterator until the predicate is false.
func TakeWhile2[K, V any](seq IIterable2[K, V], predicate func(K, V) bool) Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k, v := range seq.All() {
			if !predicate(k, v) || !yield(k, v) {
				return
			}
		}
	}
}

// DropWhile returns a Seq[V] iterator that skips values from the input iterator until the predicate is false.
func DropWhile[V any](seq IIterable[V], predicate func(V) bool) Seq[V] {
	return func(yield func(V) bool) {
		dropping := true
		for v := range seq.Values() {
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

// DropWhile returns a Seq2[K, V] iterator that skips values from the input iterator until the predicate is false.
func DropWhile2[K, V any](seq IIterable2[K, V], predicate func(K, V) bool) Seq2[K, V] {
	return func(yield func(K, V) bool) {
		dropping := true
		for k, v := range seq.All() {
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

// With returns a Seq[V] iterator that calls a function with each element before yielding it.
func With[V any](seq IIterable[V], process func(V)) Seq[V] {
	return func(yield func(V) bool) {
		for v := range seq.Values() {
			process(v)
			if !yield(v) {
				return
			}
		}
	}
}

// With returns a Seq2[K, V] iterator that calls a function on each iteration before yielding it.
func With2[K, V any](seq IIterable2[K, V], process func(K, V)) Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k, v := range seq.All() {
			process(k, v)
			if !yield(k, v) {
				return
			}
		}
	}
}

// OnEmpty returns a Seq[V] iterator that calls an else function only if the iterator is exhausted.
// Similar to a for ... else block in Python.
func OnEmpty[V any](seq IIterable[V], callback func()) Seq[V] {
	return func(yield func(V) bool) {
		for v := range seq.Values() {
			if !yield(v) {
				return
			}
		}
		callback()
	}
}

// OnEmpty returns a Seq2[K, V] iterator that calls an else function only if the iterator is not exhausted.
// Similar to a for ... else block in Python.
func OnEmpty2[K, V any](seq IIterable2[K, V], callback func()) Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k, v := range seq.All() {
			if !yield(k, v) {
				return
			}
		}
		callback()
	}
}

// ----------------------------------------------------------------------------
// Seq only functions
// ----------------------------------------------------------------------------

// Sum returns the sum of all values in the iterator.
func Sum[V Ord](seq Seq[V]) V {
	var sum V
	for v := range seq {
		sum += v
	}
	return sum
}

// Product returns the product of all values in the iterator.
func Product[V Integer](seq Seq[V]) V {
	var prod V = V(1)
	for v := range seq {
		prod *= v
	}
	return prod
}

// Min returns the minimum value in the iterator.
func Min[V Ord](seq Seq[V]) V {
	var min V
	first := true
	for v := range seq {
		if first {
			min = v
			first = false
		} else if v < min {
			min = v
		}
	}
	return min
}

// Max returns the maximum value in the iterator.
func Max[V Ord](seq Seq[V]) V {
	var max V
	first := true
	for v := range seq {
		if first {
			max = v
			first = false
		} else if v > max {
			max = v
		}
	}
	return max
}

// All returns true if all values in the iterator satisfy the predicate.
// If no predicate is provided, it defaults to checking if a value is a zero
// value for its type.
func All[V any](seq Seq[V], predicate ...func(V) bool) bool {
	pred := predZero[V]
	if len(predicate) > 0 {
		pred = predicate[0]
	}
	for v := range seq {
		if !pred(v) {
			return false
		}
	}
	return true
}

// Any returns true if any value in the iterator satisfies the predicate.
// If no predicate is provided, it defaults to checking if a value is a zero
// value for its type.
func Any[V any](seq Seq[V], predicate ...func(V) bool) bool {
	pred := predZero[V]
	if len(predicate) > 0 {
		pred = predicate[0]
	}
	for v := range seq {
		if pred(v) {
			return true
		}
	}
	return false
}

// None returns true if no value in the iterator satisfies the predicate.
// If no predicate is provided, it defaults to checking if a value is a zero
// value for its type.
func None[V any](seq Seq[V], predicate ...func(V) bool) bool {
	pred := predZero[V]
	if len(predicate) > 0 {
		pred = predicate[0]
	}
	for v := range seq {
		if pred(v) {
			return false
		}
	}
	return true
}

// ----------------------------------------------------------------------------
// Seq2 only functions
// ----------------------------------------------------------------------------

// Split converts an iterator over pair values to an iterator over key-values.
func Split[K, V any](seq IIterable[Pair[K, V]]) Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for p := range seq.Values() {
			if !yield(p.Key, p.Value) {
				return
			}
		}
	}
}

// SwapKV returns a Seq2[V, K] iterator that swaps the keys and values of the input iterator.
func SwapKV[K, V any](seq IIterable2[K, V]) Seq2[V, K] {
	return func(yield func(V, K) bool) {
		for k, v := range seq.All() {
			if !yield(v, k) {
				return
			}
		}
	}
}

// ----------------------------------------------------------------------------
// Stringer implementations
// ----------------------------------------------------------------------------

// String returns a string representation of a Seq[V] iterator.
func (s Seq[V]) String() string {
	str := fmt.Sprintf("Seq[%s][", seqString(s))
	for v := range s {
		str += fmt.Sprintf("%v ", v)
	}
	if str[len(str)-1] == ' ' {
		str = str[:len(str)-1]
	}
	str += "]"
	return str
}

// String returns a string representation of a Seq2[K, V] iterator.
func (s Seq2[K, V]) String() string {
	str := fmt.Sprintf("Seq2[%s][", seq2String(s))
	for k, v := range s {
		str += fmt.Sprintf("%v:%v ", k, v)
	}
	if str[len(str)-1] == ' ' {
		str = str[:len(str)-1]
	}
	str += "]"
	return str
}

// String returns a string representation of a ItSlice[V] iterator.
func (s ItSlice[V]) String() string {
	str := fmt.Sprintf("ItSlice[%s][", reflect.TypeOf(s).Elem())
	for _, v := range s {
		str += fmt.Sprintf("%v ", v)
	}
	if str[len(str)-1] == ' ' {
		str = str[:len(str)-1]
	}
	str += "]"
	return str
}

// String returns a string representation of a ItMap[K, V] iterator.
func (m ItMap[K, V]) String() string {
	str := fmt.Sprintf("ItMap[%s,%s][", reflect.TypeOf(m).Key(), reflect.TypeOf(m).Elem())
	for k, v := range m {
		str += fmt.Sprintf("%v:%v ", k, v)
	}
	if str[len(str)-1] == ' ' {
		str = str[:len(str)-1]
	}
	str += "]"
	return str
}

// String returns a string representation of a Pair.
func (p Pair[K, V]) String() string {
	return fmt.Sprintf("%v:%v", p.Key, p.Value)
}

// ----------------------------------------------------------------------------
// Private functions
// ----------------------------------------------------------------------------

// seqString returns a string representation of the type V from Seq[V].
func seqString(f any) string {
	return reflect.TypeOf(f).In(0).In(0).String()
}

// seq2String returns a string representation of the types K,V from Seq2[K, V].
func seq2String(f any) string {
	k := reflect.TypeOf(f).In(0).In(0)
	v := reflect.TypeOf(f).In(0).In(1)
	return fmt.Sprintf("%s,%s", k, v)
}

// predZero returns true if the value is not a zero value.
func predZero[V any](val V) bool {
	value := reflect.ValueOf(val)
	return !value.IsZero()
}
