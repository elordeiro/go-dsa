// Iters package provides a set of functions that operate on iterators.
package iters

import (
	"fmt"
	"iter"
	"log"
	"maps"
	"slices"
)

type (
	Seq[V any]               iter.Seq[V]     // Same as iter.Seq[V] but accessible in this package
	Seq2[K, V any]           iter.Seq2[K, V] // Same as iter.Seq2[K, V] but accessible in this package
	Slice[V any]             []V             // Generic slice type
	Map[K comparable, V any] map[K]V         // Generic map type
)

// MapPair is a key-value pair for maps. Key must be comparable.
type MapPair[K comparable, V any] struct {
	Key   K
	Value V
}

// Pair is a key-value pair for Seq2 iterators that need to return more than 2 values.
type Pair struct {
	Key   any
	Value any
}

// Iterable is an interface that defines the methods for a sequence of values or index-value pairs.
type Iterable[V any] interface {
	Values() Seq[V]
	All() Seq2[int, V]
}

// Iterable2 is an interface that defines the methods of Iterable and adds a method to get the keys.
type Iterable2[K, V any] interface {
	Keys() Seq[K]
	Values() Seq[V]
	All() Seq2[K, V]
}

// type Iterable[K, V any] interface {
// 	IIterable[K, V]
// }

// type Iterable2[K, V any] interface {
// 	IIterable2[K, V]
// }

// ----------------------------------------------------------------------------
// Constructors
// ----------------------------------------------------------------------------

// NewSlice returns a new Slice with the given values.
// Slices can then call other functions in this package using dot notation.
// Example: itr.NewSlice(1, 2, 3).Values() // returns an iterator over the values
func NewSlice[V any](vals ...V) Slice[V] {
	return append([]V{}, vals...)
}

// NewMap returns a new Map with the given key-value pairs.
// Maps can then call other functions in this package using dot notation.
// Example: itr.NewMap(itr.MapPair{Key: 1, Value: "one"}).Values() // returns an iterator over the values
func NewMap[K comparable, V any](pairs ...MapPair[K, V]) Map[K, V] {
	m := make(map[K]V)
	for _, p := range pairs {
		m[p.Key] = p.Value
	}
	return m
}

// ----------------------------------------------------------------------------
// Interface Implementations
// ----------------------------------------------------------------------------

// Seq[V] -----------------------------

// Values returns an iterator over the values of the Seq.
func (s Seq[V]) Values() Seq[V] {
	return s
}

// All returns an iterator over the index-value pairs of the Seq.
func (s Seq[V]) All() Seq2[int, V] {
	return All(s)
}

// Seq2[K, V] -------------------------

// Keys returns an iterator over the keys of the Seq2.
func (s Seq2[K, V]) Keys() Seq[K] {
	return Keys(s)
}

// Values returns an iterator over the values of the Seq2.
func (s Seq2[K, V]) Values() Seq[V] {
	return Values(s)
}

// All returns an iterator over the key-value pairs of the Seq2.
func (s Seq2[K, V]) All() Seq2[K, V] {
	return s
}

// Slice[V] ---------------------------

// Values returns an iterator over the values of the Slice.
func (s Slice[V]) Values() Seq[V] {
	return Surf(slices.Values(s))
}

// All returns an iterator over the index-value pairs of the Slice.
func (s Slice[V]) All() Seq2[int, V] {
	return Surf2(slices.All(s))
}

// Map[K, V] --------------------------

// Keys returns an iterator over the keys of the Map.
func (m Map[K, V]) Keys() Seq[K] {
	return Surf(maps.Keys(m))
}

// Values returns an iterator over the values of the Map.
func (m Map[K, V]) Values() Seq[V] {
	return Surf(maps.Values(m))
}

// All returns an iterator over the key-value pairs of the Map.
func (m Map[K, V]) All() Seq2[K, V] {
	return Surf2(maps.All(m))
}

// ----------------------------------------------------------------------------
// Iterator Converters
// ----------------------------------------------------------------------------

// Keys converts an iterator over key-value pairs to an iterator over keys.
func Keys[K, V any](seq Seq2[K, V]) Seq[K] {
	return func(yield func(K) bool) {
		for k := range seq {
			if !yield(k) {
				return
			}
		}
	}
}

// Values converts an iterator over key-value pairs to an iterator over values.
func Values[K, V any](seq Seq2[K, V]) Seq[V] {
	return func(yield func(V) bool) {
		for _, v := range seq {
			if !yield(v) {
				return
			}
		}
	}
}

// All converts an iterator over values to an iterator over index-value pairs.
// For indices that don't start at 0, use Enumerate.
func All[V any](seq Seq[V]) Seq2[int, V] {
	return func(yield func(int, V) bool) {
		next, stop := iter.Pull(seq.Sink())
		defer stop()
		for i := 0; ; i++ {
			v, ok := next()
			if !ok || !yield(i, v) {
				return
			}
		}
	}
}

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
// It takes a Seq[V] iterator and returns an iter.Seq[V] iterator.
func Sink[V any](seq Seq[V]) iter.Seq[V] {
	return iter.Seq[V](seq)
}

// Sink sinks a Seq[V] iterator
// Returns an iter.Seq[V] iterator.
func (seq Seq[V]) Sink() iter.Seq[V] {
	return iter.Seq[V](seq)
}

// Sink2 sinks a Seq2[K, V] iterator
// It takes a Seq2[K, V] iterator and returns an iter.Seq2[K, V] iterator.
func Sink2[K, V any](seq Seq2[K, V]) iter.Seq2[K, V] {
	return iter.Seq2[K, V](seq)
}

// Sink2 sinks a Seq2[K, V] iterator
// Returns an iter.Seq2[K, V] iterator.
func (seq Seq2[K, V]) Sink2() iter.Seq2[K, V] {
	return iter.Seq2[K, V](seq)
}

// ----------------------------------------------------------------------------

// ----------------------------------------------------------------------------
// Possibly infinite iterators
// ----------------------------------------------------------------------------

// Range returns an iterator over a range of integers.
// If only one argument is provided, it is the end of the range.
// If two arguments are provided, they are the start and end of the range.
// If three arguments are provided, they are the start, end, and step of the range.
func Range(vals ...int) Seq[int] {
	start, end, step := 0, 0, 1
	switch len(vals) {
	case 0:
		log.Println("Empty iterator! Range requires at least 1 argument")
	case 1:
		end = vals[0]
	case 2:
		start, end = vals[0], vals[1]
		if start > end {
			step = -1
		}
	default:
		log.Println("Too many args to Range Func! Range requires at most 3 arguments")
		fallthrough
	case 3:
		start, end, step = vals[0], vals[1], vals[2]
		if start > end && step > 0 {
			log.Println("Infinite range iterator! Start > End and Step > 0")
			log.Println("Step may be omitted and will be set to -1 for reverse ranges")
		}
	}

	if step > 0 {
		return func(yield func(int) bool) {
			for i := start; i < end; i += step {
				if !yield(i) {
					return
				}
			}
		}
	} else {
		return func(yield func(int) bool) {
			for i := start; i > end; i += step {
				if !yield(i) {
					return
				}
			}
		}
	}
}

// Count returns an iterator that counts up from a given integer.
// The iterator is infinite unless it is stopped by the caller.
func Count(i int) Seq[int] {
	return func(yield func(int) bool) {
		for ; ; i++ {
			if !yield(i) {
				return
			}
		}
	}
}

// Repeat returns a Seq[V] iterator that yields the same value n times.
// If n is not provided, the iterator is infinite.
func Repeat[V any](val V, end ...int) Seq[V] {
	stop := -1
	if len(end) > 0 {
		stop = end[0]
	}
	return func(yield func(V) bool) {
		for i := 0; i != stop; i++ {
			if !yield(val) {
				return
			}
		}
	}
}

// Reapeat2 returns a Seq2[K, V] iterator that yields the same key-value pair n times.
// If n is not provided, the iterator is infinite.
func Repeat2[K, V any](key K, val V, end ...int) Seq2[K, V] {
	stop := -1
	if len(end) > 0 {
		stop = end[0]
	}
	return func(yield func(K, V) bool) {
		for i := 0; i != stop; i++ {
			if !yield(key, val) {
				return
			}
		}
	}
}

// ----------------------------------------------------------------------------
// Iterator functions
// Enumerate, Zip and Chain show how to receive an iterable type.
// These functions can receive either a Map[K]V, a Slice[V], a Seq[V] or a Seq2[K, V].
// Note: Slice[V] and Map[K]V are custom types defined in this package, not the built-in types.
// ----------------------------------------------------------------------------

// Enumerate returns an a Seq2[int, V] iterator over index-value pairs in the Iterable.
// The start argument specifies the starting index.
func Enumerate[V any](iter Iterable[V], start int) Seq2[int, V] {
	return func(yield func(int, V) bool) {
		i := start
		for v := range iter.Values() {
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
func Enumerate2[K, V any](iter Iterable2[K, V], start int) Seq2[int, Pair] {
	return func(yield func(int, Pair) bool) {
		i := start
		for k, v := range iter.All() {
			if !yield(i, Pair{k, v}) {
				return
			}
			i++
		}
	}
}

// Zip returns an Seq2[V1, V2] iterator over values from two sequences.
// The iteration stops when either of the sequences is exhausted.
// In other words, the length of the resulting sequence is the minimum
// of the lengths of the input sequences.
func Zip[V1, V2 any](seq1 Iterable[V1], seq2 Iterable[V2]) Seq2[V1, V2] {
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
// The iteration stops when either of the sequences is exhausted.
// In other words, the length of the resulting sequence is the minimum
// of the lengths of the input sequences.

func Zip2[K1, V1, K2, V2 any](seq1 Iterable2[K1, V1], seq2 Iterable2[K2, V2]) Seq2[Pair, Pair] {
	return func(yield func(Pair, Pair) bool) {
		p1, stop := iter.Pull2(seq1.All().Sink2())
		defer stop()
		p2, stop := iter.Pull2(seq2.All().Sink2())
		defer stop()
		for {
			k1, v1, ok1 := p1()
			k2, v2, ok2 := p2()
			if (!ok1 || !ok2) || !yield(Pair{k1, v1}, Pair{k2, v2}) {
				return
			}
		}
	}
}

// Chain returns a Seq[V] iterator that chains the values of multiple input iterators.
func Chain[Iter Iterable[V], V any](seqs ...Iter) Seq[V] {
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
func Chain2[Iter Iterable2[K, V], K, V any](seqs ...Iter) Seq2[K, V] {
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

// ----------------------------------------------------------------------------
// Iterator Methods
// ----------------------------------------------------------------------------

// Enumerate returns a Seq2[int, V] iterator over index-value pairs in the Iterable.
// The start argument specifies the starting index.
func (seq Seq[V]) Enumerate(start int) Seq2[int, V] {
	return func(yield func(int, V) bool) {
		i := start
		for v := range seq.Values() {
			if !yield(i, v) {
				return
			}
			i++
		}
	}
}

// Enumerate returns a Seq2[int, Pair] iterator over index-pair pairs in the Iterable.
// A pair is a key-value pair for Seq2 iterators that need to return more than 2 values.
// The start argument specifies the starting index.
func (seq Seq2[K, V]) Enumerate(start int) Seq2[int, Pair] {
	return func(yield func(int, Pair) bool) {
		i := start
		for k, v := range seq.All() {
			if !yield(i, Pair{k, v}) {
				return
			}
			i++
		}
	}
}

// ----------------------------------------------------------------------------
// Common built-in methods
// ----------------------------------------------------------------------------

// Filter returns a Seq[V] iterator over values that satisfy the filter function.
func (seq Seq[V]) Filter(filterFunc func(V) bool) Seq[V] {
	return func(yield func(V) bool) {
		for v := range seq {
			if filterFunc(v) && !yield(v) {
				return
			}
		}
	}
}

// Filter returns a Seq2[K, V] iterator over key-value pairs that satisfy the filter function.
func (seq Seq2[K, V]) Filter(filterFunc func(K, V) bool) Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k, v := range seq {
			if filterFunc(k, v) && !yield(k, v) {
				return
			}
		}
	}
}

// Map returns a Seq[V] iterator over values that are transformed by the map function.
func (seq Seq[V]) Map(mapFunc func(V) V) Seq[V] {
	return func(yield func(V) bool) {
		for v := range seq {
			if !yield(mapFunc(v)) {
				return
			}
		}
	}
}

// Map returns a Seq2[K, V] iterator over key-value pairs that are transformed by the map function.
func (seq Seq2[K, V]) Map(mapFunc func(K, V) Pair) Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k, v := range seq {
			p := mapFunc(k, v)
			if !yield(p.Key.(K), p.Value.(V)) {
				return
			}
		}
	}
}

// Reduce returns a single value that is the result of applying the reduce function to all values in the iterator.
func (seq Seq[V]) Reduce(reduceFunc func(V, V) V) V {
	var acc V
	for v := range seq {
		acc = reduceFunc(acc, v)
	}
	return acc
}

// Reduce returns a single value that is the result of applying the reduce function to all values in the iterator.
func (seq Seq2[K, V]) Reduce(reduceFunc func(V, V) V) V {
	var acc V
	for _, v := range seq {
		acc = reduceFunc(acc, v)
	}
	return acc
}

// Reduce2 returns a single value that is the result of applying the reduce function to all keys in the iterator.
func (seq Seq2[K, V]) Reduce2(reduceFunc func(K, K) K) K {
	var acc K
	for k := range seq {
		acc = reduceFunc(acc, k)
	}
	return acc
}

// ReduceAll returns a single value that is the result of applying the reduce function
// to all keys and values in the iterator.
// The accumulator is the same type as the Values.
// For an accumulator of a type K, use Reduce2.
func (seq Seq2[K, V]) ReduceAll(reduceFunc func(V, K, V) V) V {
	var acc V
	for k, v := range seq {
		acc = reduceFunc(acc, k, v)
	}
	return acc
}

// ReduceAll2 returns a single value that is the result of applying the reduce function
// to all keys and values in the iterator.
// The accumulator is the same type as the Keys.
// For an accumulator of a type V, use Reduce.
func (seq Seq2[K, V]) ReduceAll2(reduceFunc func(K, K, V) K) K {
	var acc K
	for k, v := range seq {
		acc = reduceFunc(acc, k, v)
	}
	return acc
}

// ----------------------------------------------------------------------------
// Common itertools functions
// ----------------------------------------------------------------------------

// Cycle returns a Seq[V] iterator that cycles through the values of the input iterator.
func (seq Seq[V]) Cycle() Seq[V] {
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

// Cycle returns a Seq2[K, V] iterator that cycles through the values of the input iterator.
func (seq Seq2[K, V]) Cycle() Seq2[K, V] {
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

// Chain returns Seq[V] iterator that chains the values of multiple input iterators.
func (seq Seq[V]) Chain(seqs ...Seq[V]) Seq[V] {
	return func(yield func(V) bool) {
		for v := range seq {
			if !yield(v) {
				return
			}
		}
		for _, seq := range seqs {
			for v := range seq {
				if !yield(v) {
					return
				}
			}
		}
	}
}

// Chain returns a Seq2[K, V] iterator that chains the values of multiple input iterators.
func (seq Seq2[K, V]) Chain(seqs ...Seq2[K, V]) Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k, v := range seq {
			if !yield(k, v) {
				return
			}
		}
		for _, seq := range seqs {
			for k, v := range seq {
				if !yield(k, v) {
					return
				}
			}
		}
	}
}

// Take returns a Seq[V] iterator that yields the first n values of the input iterator.
func (seq Seq[V]) Take(n int) Seq[V] {
	return func(yield func(V) bool) {
		i := 0
		for v := range seq {
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

// Take returns a Seq2[K, V] iterator that yields the first n values of the input iterator.
func (seq Seq2[K, V]) Take(n int) Seq2[K, V] {
	return func(yield func(K, V) bool) {
		i := 0
		for k, v := range seq {
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
func (seq Seq[V]) Drop(n int) Seq[V] {
	return func(yield func(V) bool) {
		i := 0
		for v := range seq {
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
func (seq Seq2[K, V]) Drop(n int) Seq2[K, V] {
	return func(yield func(K, V) bool) {
		i := 0
		for k, v := range seq {
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

// ----------------------------------------------------------------------------
// Common functool functions
// ----------------------------------------------------------------------------

// TakeWhile returns a Seq[V] iterator that yields values from the input iterator until the predicate is false.
func (seq Seq[V]) TakeWhile(predicate func(V) bool) Seq[V] {
	return func(yield func(V) bool) {
		for v := range seq {
			if !predicate(v) || !yield(v) {
				return
			}
		}
	}
}

// TakeWhile returns a Seq2[K, V] iterator that yields values from the input iterator until the predicate is false.
func (seq Seq2[K, V]) TakeWhile(predicate func(K, V) bool) Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k, v := range seq {
			if !predicate(k, v) || !yield(k, v) {
				return
			}
		}
	}
}

// DropWhile returns a Seq[V] iterator that skips values from the input iterator until the predicate is false.
func (seq Seq[V]) DropWhile(predicate func(V) bool) Seq[V] {
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

// DropWhile returns a Seq2[K, V] iterator that skips values from the input iterator until the predicate is false.
func (seq Seq2[K, V]) DropWhile(predicate func(K, V) bool) Seq2[K, V] {
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

// With returns a Seq[V] iterator that calls a function on each iteration before yielding it.
func (seq Seq[V]) With(process func()) Seq[V] {
	return func(yield func(V) bool) {
		for v := range seq {
			process()
			if !yield(v) {
				return
			}
		}
	}
}

// With returns a Seq2[K, V] iterator that calls a function on each iteration before yielding it.
func (seq Seq2[K, V]) With(process func()) Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k, v := range seq {
			process()
			if !yield(k, v) {
				return
			}
		}
	}
}

// Else returns a Seq[V] iterator that calls an else function only if the iterator is not exhausted.
// Similar to a for ... else block in Python.
func (seq Seq[V]) Else(callback func()) Seq[V] {
	return func(yield func(V) bool) {
		for v := range seq {
			if !yield(v) {
				return
			}
		}
		callback()
	}
}

// Else returns a Seq2[K, V] iterator that calls an else function only if the iterator is not exhausted.
// Similar to a for ... else block in Python.
func (seq Seq2[K, V]) Else(callback func()) Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k, v := range seq {
			if !yield(k, v) {
				return
			}
		}
		callback()
	}
}

// ----------------------------------------------------------------------------
// Stringer implementations
// ----------------------------------------------------------------------------

// String returns a string representation of the Seq[V].
func (s Seq[V]) String() string {
	str := "Seq[V]["
	for v := range s {
		str += fmt.Sprintf("%v ", v)
	}
	str = str[:len(str)-1] + "]"
	return str
}

// String returns a string representation of the Seq2[K, V].
func (s Seq2[K, V]) String() string {
	str := "Seq2[K,V]["
	for k, v := range s {
		str += fmt.Sprintf("%v:%v ", k, v)
	}
	str = str[:len(str)-1] + "]"
	return str
}

// String returns a string representation of the Pair.
func (p Pair) String() string {
	return fmt.Sprintf("%v:%v", p.Key, p.Value)
}

// String returns a string representation of the MapPair[K, V].
func (p MapPair[K, V]) String() string {
	return fmt.Sprintf("%v:%v", p.Key, p.Value)
}
