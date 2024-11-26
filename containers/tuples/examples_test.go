package tuples_test

import (
	"fmt"

	"github.com/elordeiro/goext/containers/tuples"
)

func ExampleNewPair() {
	pair := tuples.NewPair(1, 2)
	fmt.Println(pair)
	// Output: (1 2)
}

func ExamplePairs() {
	pairSlice := tuples.Pairs(
		tuples.NewPair(1, 2),
		tuples.NewPair(3, 4),
		tuples.NewPair(5, 6))
	fmt.Println(pairSlice)
	// Output: [(1 2) (3 4) (5 6)]
}

func ExamplePair_Left() {
	pair := tuples.NewPair(1, 2)
	fmt.Println(pair.Left())
	// Output: 1
}

func ExamplePair_Right() {
	pair := tuples.NewPair(1, 2)
	fmt.Println(pair.Right())
	// Output: 2
}

func ExampleNewEdge() {
	edge := tuples.NewEdge("A", "B", 5)
	fmt.Println(edge)
	// Output: (A->B 5)
}

func ExampleNewEdge_defaultWeight() {
	edge := tuples.NewEdge[string, int]("A", "B")
	fmt.Println(edge)
	// Output: (A->B 1)
}

func ExampleEdges() {
	edgeSlice := tuples.Edges(
		[3]int{1, 2, 1},
		[3]int{3, 4, 1},
		[3]int{5, 6, 1})
	fmt.Println(edgeSlice)
	// Output: [(1->2 1) (3->4 1) (5->6 1)]
}

func ExampleEdge_Src() {
	edge := tuples.NewEdge("A", "B", 5)
	fmt.Println(edge.Src())
	// Output: A
}

func ExampleEdge_Dst() {
	edge := tuples.NewEdge("A", "B", 5)
	fmt.Println(edge.Dst())
	// Output: B
}

func ExampleEdge_Weight() {
	edge := tuples.NewEdge("A", "B", 5)
	fmt.Println(edge.Weight())
	// Output: 5
}

func ExampleNewCell() {
	cell := tuples.NewCell(1, 2)
	fmt.Println(cell)
	// Output: (1, 2)
}

func ExampleCells() {
	cellSlice := tuples.Cells(
		[2]int{1, 2},
		[2]int{3, 4},
		[2]int{5, 6})
	fmt.Println(cellSlice)
	// Output: [(1, 2) (3, 4) (5, 6)]
}

func ExampleCell_Row() {
	cell := tuples.NewCell(1, 2)
	fmt.Println(cell.Row())
	// Output: 1
}

func ExampleCell_Col() {
	cell := tuples.NewCell(1, 2)
	fmt.Println(cell.Col())
	// Output: 2
}

func ExampleNewPoint2d() {
	point := tuples.NewPoint2d(1, 2)
	fmt.Println(point)
	// Output: (1, 2)
}

func ExamplePoints2d() {
	pointSlice := tuples.Points2d(
		[2]int{1, 2},
		[2]int{3, 4},
		[2]int{5, 6})
	fmt.Println(pointSlice)
	// Output: [(1, 2) (3, 4) (5, 6)]
}

func ExamplePoint2d_X() {
	point := tuples.NewPoint2d(1, 2)
	fmt.Println(point.X())
	// Output: 1
}

func ExamplePoint2d_Y() {
	point := tuples.NewPoint2d(1, 2)
	fmt.Println(point.Y())
	// Output: 2
}

func ExampleNewPoint3d() {
	point := tuples.NewPoint3d(1, 2, 3)
	fmt.Println(point)
	// Output: (1, 2, 3)
}

func ExamplePoints3d() {
	pointSlice := tuples.Points3d(
		[3]int{1, 2, 3},
		[3]int{4, 5, 6},
		[3]int{7, 8, 9})
	fmt.Println(pointSlice)
	// Output: [(1, 2, 3) (4, 5, 6) (7, 8, 9)]
}

func ExamplePoint3d_X() {
	point := tuples.NewPoint3d(1, 2, 3)
	fmt.Println(point.X())
	// Output: 1
}

func ExamplePoint3d_Y() {
	point := tuples.NewPoint3d(1, 2, 3)
	fmt.Println(point.Y())
	// Output: 2
}

func ExamplePoint3d_Z() {
	point := tuples.NewPoint3d(1, 2, 3)
	fmt.Println(point.Z())
	// Output: 3
}
