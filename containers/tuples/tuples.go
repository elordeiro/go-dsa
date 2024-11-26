// Package tuples provides a set of convenient tuple types that can be used in various
// algorithms and data structures.
package tuples

import (
	"fmt"

	"github.com/elordeiro/goext/constraints"
)

// Pair is any pair of values.
type Pair[L, R any] struct {
	left  L
	right R
}

// NewPair returns a new pair tuple with the given values.
func NewPair[L, R any](left L, right R) Pair[L, R] {
	return Pair[L, R]{left, right}
}

// Pairs returns a new slice of Pairs with the given pairs.
// This function is useful for initializing a slice of Pairs without having to
// specify the type of the slice.
func Pairs[L, R any](pairs ...Pair[L, R]) []Pair[L, R] {
	return []Pair[L, R](pairs)
}

// Left returns the left value of a Pair.
func (p Pair[L, R]) Left() L {
	return p.left
}

// Right returns the right value of a Pair.
func (p Pair[L, R]) Right() R {
	return p.right
}

// String returns a string representation of a Pair.
func (p Pair[K, V]) String() string {
	return fmt.Sprintf("(%v %v)", p.left, p.right)
}

// ---- Edge ----

// Edge is a pair of vertices with a weight that represents a weighted Edge in a graph.
type Edge[V comparable, N constraints.Number] struct {
	src, dst V
	weight   N
}

// NewEdge returns a new edge with the given source and destination vertices and weight.
// The vertices can be of any type that supports comparison, and the weight can be any
// type that satisfies the constraints.Number interface. The weight is optional and
// defaults to 1 if not provided.
func NewEdge[V comparable, N constraints.Number](src, dst V, weight ...N) Edge[V, N] {
	if len(weight) == 0 {
		return Edge[V, N]{src, dst, 1}
	}
	return Edge[V, N]{src, dst, weight[0]}
}

// Edges returns a new slice of Edges with the given edges.
// Note that the edges are represented as arrays of 3 elements, where the first two
// elements are the source and destination vertices, and the third element is the weight.
// This is useful for initializing a slice of Edges in test cases without having to
// specify the type of the slice.
func Edges[N constraints.Number](edges ...[3]N) []Edge[N, N] {
	var slice []Edge[N, N]
	for _, edge := range edges {
		slice = append(slice, Edge[N, N]{edge[0], edge[1], edge[2]})
	}
	return slice
}

// Src returns the source vertex of an edge
func (e Edge[V, N]) Src() V {
	return e.src
}

// Dst returns the destination vertex of an edge
func (e Edge[V, N]) Dst() V {
	return e.dst
}

// Weight returns the weight of an edge
func (e Edge[V, N]) Weight() N {
	return e.weight
}

// String returns a string representation of an edge.
func (e Edge[V, N]) String() string {
	return fmt.Sprintf("(%v->%v %v)", e.src, e.dst, e.weight)
}

// ---- Cell ----

// Cell is a pair of numbers that represents a coordinate on a 2D grid.
type Cell[N constraints.Number] struct {
	row, col N
}

// NewCell returns a new cell with the given row and column.
func NewCell[N constraints.Number](row, col N) Cell[N] {
	return Cell[N]{row, col}
}

// Cells returns a new slice of Cells with.
// Note that the cells are represented as arrays of 2 elements, where the first element
// is the row and the second element is the column. This is useful for initializing a
// slice of Cells without having to specify the type of the slice.
func Cells[N constraints.Number](cells ...[2]N) []Cell[N] {
	var slice []Cell[N]
	for _, cell := range cells {
		slice = append(slice, Cell[N]{cell[0], cell[1]})
	}
	return slice
}

// Row returns the row of a cell.
func (c Cell[N]) Row() N {
	return c.row
}

// Col returns the column of a cell.
func (c Cell[N]) Col() N {
	return c.col
}

// String returns a string representation of a cell.
func (c Cell[N]) String() string {
	return fmt.Sprintf("(%v, %v)", c.Row(), c.Col())
}

// ---- Point2d ----

// Point2d is a pair of numbers that represent a point in a 2D plane.
type Point2d[N constraints.Number] struct {
	x, y N
}

// NewPoint2d returns a new point with the given x and y coordinates.
func NewPoint2d[N constraints.Number](x, y N) Point2d[N] {
	return Point2d[N]{x, y}
}

// Points2d returns a new slice of Points with the given points.
// Note that the points are represented as arrays of 2 elements, where the first element
// is the x coordinate and the second element is the y coordinate. This is useful for
// initializing a slice of Points without having to specify the type of the slice.
func Points2d[N constraints.Number](points ...[2]N) []Point2d[N] {
	var slice []Point2d[N]
	for _, point := range points {
		slice = append(slice, Point2d[N]{point[0], point[1]})
	}
	return slice
}

// X returns the x coordinate of a 2D-point.
func (p Point2d[N]) X() N {
	return p.x
}

// Y returns the y coordinate of a 2D-point.
func (p Point2d[N]) Y() N {
	return p.y
}

// String returns a string representation of a 2D-point.
func (p Point2d[N]) String() string {
	return fmt.Sprintf("(%v, %v)", p.x, p.y)
}

// ---- Point3d ----

// Point3d is a pair of numbers that represent a point in a 3D space.
type Point3d[N constraints.Number] struct {
	x, y, z N
}

// NewPoint3 returns a new 3D point with the given x, y, and z coordinates.
func NewPoint3d[N constraints.Number](x, y, z N) Point3d[N] {
	return Point3d[N]{x, y, z}
}

// Points3d returns a new slice of 3D Points with the given points.
// Note that the points are represented as arrays of 3 elements, where the first element
// is the x coordinate, the second element is the y coordinate, and the third element is
// the z coordinate. This is useful for initializing a slice of 3D Points without having
// to specify the type of the slice.
func Points3d[N constraints.Number](points ...[3]N) []Point3d[N] {
	var slice []Point3d[N]
	for _, point := range points {
		slice = append(slice, Point3d[N]{point[0], point[1], point[2]})
	}
	return slice
}

// X returns the x coordinate of a 3D-point.
func (p Point3d[N]) X() N {
	return p.x
}

// Y returns the y coordinate of a 3D-point.
func (p Point3d[N]) Y() N {
	return p.y
}

// Z returns the z coordinate of a 3D-point.
func (p Point3d[N]) Z() N {
	return p.z
}

// String returns a string representation of a 3D-point.
func (p Point3d[N]) String() string {
	return fmt.Sprintf("(%v, %v, %v)", p.x, p.y, p.z)
}
