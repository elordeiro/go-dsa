// Package ituples provides interfaces for tuples that can be used in various algorithms and data structures.
package ituples

import "github.com/elordeiro/goext/constraints"

// Pair is an interface for a pair of values.
type Pair[L any, R any] interface {
	Left() L
	Right() R
}

// Edge is an interface for a weighted edge.
type Edge[V comparable, N constraints.Number] interface {
	Src() V
	Dst() V
	Weight() N
}

// Cell is an interface for a grid cell.
type Cell[N constraints.Number] interface {
	Row() N
	Col() N
}

// Point2d is an interface for a 2D-point.
type Point2d[N constraints.Number] interface {
	X() N
	Y() N
}

// Point3d is an interface for a 3D-point.
type Point3d[N constraints.Number] interface {
	Point2d[N]
	Z() N
}
