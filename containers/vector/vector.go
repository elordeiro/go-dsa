// https://go.dev/wiki/SliceTricks
//
// Package vector provides a vector implementation.
package vector

import (
	"fmt"
	"iter"

	"github.com/elordeiro/goext/types"
)

// Vector is a slice of values.
type Vector[V any] []V

// New returns a new Vector with the given values.
func New[V any](vals ...V) Vector[V] {
	return Vector[V](vals)
}

// At returns the value at index i.
func (v Vector[V]) At(i int) V {
	return v[i]
}

// Back returns the last value in the vector.
func (v Vector[V]) Back() V {
	return v[len(v)-1]
}

// Front returns the first value in the vector.
func (v Vector[V]) Front() V {
	return v[0]
}

// IsEmpty returns true if the vector is empty.
func (v Vector[V]) IsEmpty() bool {
	return len(v) == 0
}

// Len returns the length of the vector.
func (v Vector[V]) Len() int {
	return len(v)
}

// Reverse reverses the values in the vector.
func (v Vector[V]) Reverse() {
	for i := len(v)/2 - 1; i >= 0; i-- {
		j := len(v) - 1 - i
		(v)[i], (v)[j] = (v)[j], (v)[i]
	}
}

// Set sets the value at index i to val.
func (v Vector[V]) Set(i int, val V) {
	v[i] = val
}

// Swap swaps the values at indexes i and j.
func (v Vector[V]) Swap(i, j int) {
	v[i], v[j] = v[j], v[i]
}

// Clear removes all values from this vector.
func (v *Vector[V]) Clear() {
	clear(*v)
}

// Concat appends the values of other to this vector.
func (v *Vector[V]) Concat(other Vector[V]) {
	*v = append(*v, other...)
}

// Copy copies the values of src to this vector.
func (v *Vector[V]) Copy(src Vector[V]) {
	*v = make([]V, len(src))
	copy(*v, src)
}

// Cut removes the values from i to j from this vector.
func (v *Vector[V]) Cut(i, j int) {
	copy((*v)[i:], (*v)[j:])
	for k, l := len(*v)-j+i, len(*v); k < l; k++ {
		(*v)[k] = types.Zero[V]()
	}
	(*v) = (*v)[:len(*v)-j+i]
}

// Insert inserts the values vals at index i in this vector.
func (v *Vector[V]) Insert(i int, vals ...V) {
	(*v) = append(*v, make([]V, len(vals))...)
	copy((*v)[i+len(vals):], (*v)[i:])
	for j := 0; j < len(vals); j++ {
		(*v)[i] = vals[j]
		i++
	}
}

// Pop removes the last value from this vector and returns it.
func (v *Vector[V]) Pop() V {
	val := (*v)[len(*v)-1]
	(*v)[len(*v)-1] = types.Zero[V]()
	(*v) = (*v)[:len(*v)-1]
	return val
}

// PopAt removes the value at i from this vector and returns it.
func (v *Vector[V]) PopAt(i int) V {
	val := (*v)[i]
	copy((*v)[i:], (*v)[i+1:])
	(*v)[len(*v)-1] = types.Zero[V]()
	(*v) = (*v)[:len((*v))-1]
	return val
}

// Push appends the value val to this vector.
func (v *Vector[V]) Push(val V) {
	(*v) = append(*v, val)
}

// Values returns an iter.Seq[V] of all values in the vector.
func (v Vector[V]) Values() iter.Seq[V] {
	return func(yield func(V) bool) {
		for _, v := range v {
			if !yield(v) {
				return
			}
		}
	}
}

// All returns an iter.Seq2[int, V] of values and their indexes in the vector.
func (v Vector[V]) All() iter.Seq2[int, V] {
	return func(yield func(int, V) bool) {
		for i, v := range v {
			if !yield(i, v) {
				return
			}
		}
	}
}

// Backwards returns an iter.Seq[V] of all values in the vector in reverse order.
func (v Vector[V]) Backwards() iter.Seq[V] {
	return func(yield func(V) bool) {
		for i := len(v) - 1; i >= 0; i-- {
			if !yield(v[i]) {
				return
			}
		}
	}
}

// String returns a string representation of the vector.
func (v Vector[V]) String() string {
	return fmt.Sprint([]V(v))
}
