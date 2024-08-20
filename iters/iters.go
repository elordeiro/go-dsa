// Iters package provides a set of functions that operate on iterators.
package iters

import (
	"iter"
	"log"
)

// ----------------------------------------------------------------------------
// Converters
// ----------------------------------------------------------------------------

// Seq1 converts an iterator over index-value pairs to an iterator over values.
func Seq1[E any, K comparable](iterator iter.Seq2[K, E]) iter.Seq[E] {
	return func(yield func(E) bool) {
		for _, v := range iterator {
			if !yield(v) {
				return
			}
		}
	}
}

// Seq2 converts an iterator over values to an iterator over index-value pairs.
// For indices that don't start at 0, use Enumerate.
func Seq2[V any](iterator iter.Seq[V]) iter.Seq2[int, V] {
	return func(yield func(int, V) bool) {
		next, stop := iter.Pull(iterator)
		defer stop()
		for i := 0; ; i++ {
			v, ok := next()
			if !ok || !yield(i, v) {
				return
			}
		}
	}
}

// ----------------------------------------------------------------------------

// ----------------------------------------------------------------------------
// int type iterators
// ----------------------------------------------------------------------------

// Range returns an iterator over a range of integers.
// If only one argument is provided, it is the end of the range.
// If two arguments are provided, they are the start and end of the range.
// If three arguments are provided, they are the start, end, and step of the range.
func Range(vals ...int) iter.Seq[int] {
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
func Count(i int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for ; ; i++ {
			if !yield(i) {
				return
			}
		}
	}
}

// ----------------------------------------------------------------------------
// Enumerate, Zip, and Seq2 are the only functions that return Seq2 iterators
// ----------------------------------------------------------------------------

// Enumerate returns an iterator over index-value pairs in the slice.
// The start argument specifies the starting index.
func Enumerate[E any](start int, iterator iter.Seq[E]) iter.Seq2[int, E] {
	return func(yield func(int, E) bool) {
		next, stop := iter.Pull(iterator)
		defer stop()
		for i := start; ; i++ {
			v, ok := next()
			if !ok || !yield(i, v) {
				return
			}
		}
	}
}

// Zip returns an iterator over pairs of values from two sequences.
// The iteration stops when either of the sequences is exhausted.
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

// ----------------------------------------------------------------------------
// Common built-in functions
// ----------------------------------------------------------------------------

// Filter returns an iterator over values that satisfy the filter function.
func Filter[E any](iterator iter.Seq[E], filterFunc func(E) bool) iter.Seq[E] {
	return func(yield func(E) bool) {
		for v := range iterator {
			if filterFunc(v) && !yield(v) {
				return
			}
		}
	}
}

// Map returns an iterator over values that are transformed by the map function.
func Map[E any](iterator iter.Seq[E], mapFunc func(E) E) iter.Seq[E] {
	return func(yield func(E) bool) {
		for v := range iterator {
			if !yield(mapFunc(v)) {
				return
			}
		}
	}
}

// Reduce returns a single value that is the result of applying the reduce function to all values in the iterator.
func Reduce[E any](iterator iter.Seq[E], reduceFunc func(E, E) E) E {
	var acc E
	for v := range iterator {
		acc = reduceFunc(acc, v)
	}
	return acc
}

// ----------------------------------------------------------------------------
// Common itertools functions
// ----------------------------------------------------------------------------

// Cycle returns an iterator that cycles through the values of the input iterator.
func Cycle[E any](iterator iter.Seq[E]) iter.Seq[E] {
	return func(yield func(E) bool) {
		for {
			for v := range iterator {
				if !yield(v) {
					return
				}
			}

		}
	}
}

// Repeat returns an iterator that yields the same value n times.
// If n is not provided, the iterator is infinite.
func Repeat[E any](val E, end ...int) iter.Seq[E] {
	stop := -1
	if len(end) > 0 {
		stop = end[0]
	}
	return func(yield func(E) bool) {
		for i := 0; i != stop; i++ {
			if !yield(val) {
				return
			}
		}
	}
}

// Chain returns an iterator that chains the values of multiple input iterators.
func Chain[E any](iterators ...iter.Seq[E]) iter.Seq[E] {
	return func(yield func(E) bool) {
		for _, iterator := range iterators {
			for v := range iterator {
				if !yield(v) {
					return
				}
			}
		}
	}
}

// Take returns an iterator that yields the first n values of the input iterator.
func Take[E any](n int, iterator iter.Seq[E]) iter.Seq[E] {
	return func(yield func(E) bool) {
		i := 0
		for v := range iterator {
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

// Drop returns an iterator that skips the first n values of the input iterator.
func Drop[E any](n int, iterator iter.Seq[E]) iter.Seq[E] {
	return func(yield func(E) bool) {
		i := 0
		for v := range iterator {
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

// ----------------------------------------------------------------------------
// Common functool functions
// ----------------------------------------------------------------------------

// TakeWhile returns an iterator that yields values from the input iterator until the predicate is false.
func TakeWhile[E any](iterator iter.Seq[E], predicate func(E) bool) iter.Seq[E] {
	return func(yield func(E) bool) {
		for v := range iterator {
			if !predicate(v) || !yield(v) {
				return
			}
		}
	}
}

// DropWhile returns an iterator that skips values from the input iterator until the predicate is false.
func DropWhile[E any](iterator iter.Seq[E], predicate func(E) bool) iter.Seq[E] {
	return func(yield func(E) bool) {
		dropping := true
		for v := range iterator {
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

// With returns an iterator that calls a function on each iteration before yielding it.
func With[E any](iterator iter.Seq[E], process func()) iter.Seq[E] {
	return func(yield func(E) bool) {
		for v := range iterator {
			process()
			if !yield(v) {
				return
			}
		}
	}
}

// Else returns an iterator that calls an else function only if the iterator is not exhausted.
// Similar to a for ... else block in Python.
func Else[E any](iterator iter.Seq[E], callback func()) iter.Seq[E] {
	return func(yield func(E) bool) {
		for v := range iterator {
			if !yield(v) {
				return
			}
		}
		callback()
	}
}

// ----------------------------------------------------------------------------
