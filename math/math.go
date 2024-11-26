package math

import (
	"iter"
	"math"

	"github.com/elordeiro/goext/constraints"
	"github.com/elordeiro/goext/types"
)

// Sum returns the sum of all values in the iterator.
func Sum[V constraints.Ordered](seq iter.Seq[V]) V {
	var sum V
	for v := range seq {
		sum += v
	}
	return sum
}

// Product returns the product of all values in the iterator.
func Product[V constraints.Integer](seq iter.Seq[V]) V {
	var prod V = V(1)
	for v := range seq {
		prod *= v
	}
	return prod
}

// Min returns the minimum value in the iterator.
func Min[V constraints.Ordered](seq iter.Seq[V]) V {
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
func Max[V constraints.Ordered](seq iter.Seq[V]) V {
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

// Abs returns the absolute value of a number.
func Abs[N constraints.Number](n N) N {
	if n < 0 {
		return -n
	}
	return n
}

// DivMod returns the quotient and remainder of a division.
func DivMod[N constraints.Integer](a, b N) (N, N) {
	return a / b, a % b
}

// Inf returns the maximum value of a number type.
func Inf[N constraints.Number]() N {
	switch max := any(types.Zero[N]()).(type) {
	case int:
		max = math.MaxInt
		return N(max)
	case int8:
		max = math.MaxInt8
		return N(max)
	case int16:
		max = math.MaxInt16
		return N(max)
	case int32:
		max = math.MaxInt32
		return N(max)
	case int64:
		max = math.MaxInt64
		return N(max)
	case uint:
		max = math.MaxUint
		return N(max)
	case uint8:
		max = math.MaxUint8
		return N(max)
	case uint16:
		max = math.MaxUint16
		return N(max)
	case uint32:
		max = math.MaxUint32
		return N(max)
	case uint64:
		max = math.MaxUint64
		return N(max)
	case uintptr:
		max = math.MaxUint64
		return N(max)
	case float32:
		max = math.MaxFloat32
		return N(max)
	case float64:
		max = math.MaxFloat64
		return N(max)
	default:
		panic("unsupported type")
	}
}
