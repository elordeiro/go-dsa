package search

import (
	"iter"

	"github.com/elordeiro/goext/types"
)

// All returns true if all values in the iterator satisfy the predicate.
// If no predicate is provided, it defaults to checking if a value is a zero
// value for its type.
func All[V any](seq iter.Seq[V], predicate ...func(V) bool) bool {
	pred := types.IsTruthy[V]
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
func Any[V any](seq iter.Seq[V], predicate ...func(V) bool) bool {
	pred := types.IsTruthy[V]
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
func None[V any](seq iter.Seq[V], predicate ...func(V) bool) bool {
	pred := types.IsTruthy[V]
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
