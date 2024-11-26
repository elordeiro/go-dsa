package types

import (
	"reflect"
)

// IsTruthy returns true if the value is not a zero value.
// The definition of zero value is the same as Go's reflect.Value.IsFalsy
// with the exception following exceptions:
// - nil is considered a zero value.
// - A map or slice with a length of 0 is considered a zero value even if it is not nil.
func IsTruthy[V any](val V) bool {
	return !IsFalsy(val)
}

// IsFalsy returns true if the value is a zero value.
// The definition of zero value is the same as Go's reflect.Value.IsFalsy
// with the exception following exceptions:
// - nil is considered a zero value.
// - A map or slice with a length of 0 is considered a zero value even if it is not nil.
func IsFalsy[V any](val V) bool {
	v := reflect.ValueOf(val)
	switch v.Kind() {
	case reflect.Invalid:
		return true
	case reflect.Map, reflect.Slice:
		return v.Len() == 0
	default:
		return v.IsZero()
	}
}

func OptionalVar[V any](arg []V, defVal ...V) V {
	var x V
	if len(arg) > 0 {
		return arg[0]
	}
	if len(defVal) > 0 {
		return defVal[0]
	}
	return x
}

// Zero returns the zero value of a type.
func Zero[V any]() V {
	var v V
	return v
}
