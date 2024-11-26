package tuples_test

import (
	"slices"
	"testing"

	"github.com/elordeiro/goext/containers/tuples"
)

func TestNewPair(t *testing.T) {
	pair := tuples.NewPair(1, 2)

	if pair.Left() != 1 {
		t.Errorf("Left() = %v, want 1", pair.Left())
	}

	if pair.Right() != 2 {
		t.Errorf("Right() = %v, want 2", pair.Right())
	}
}

func TestPairs(t *testing.T) {
	want := []tuples.Pair[int, int]{tuples.NewPair(1, 2), tuples.NewPair(3, 4)}
	got := tuples.Pairs(tuples.NewPair(1, 2), tuples.NewPair(3, 4))

	if !slices.Equal(got, want) {
		t.Errorf("Pairs() = %v, want %v", got, want)
	}
}

func TestNewEdge(t *testing.T) {
	edge := tuples.NewEdge("A", "B", 5)

	if edge.Src() != "A" {
		t.Errorf("Src() = %v, want A", edge.Src())
	}

	if edge.Dst() != "B" {
		t.Errorf("Dst() = %v, want B", edge.Dst())
	}

	if edge.Weight() != 5 {
		t.Errorf("Weight() = %v, want 5", edge.Weight())
	}
}

func TestNewEdgeDefaultWeight(t *testing.T) {
	edge := tuples.NewEdge[string, int]("A", "B")

	if edge.Src() != "A" {
		t.Errorf("Src() = %v, want A", edge.Src())
	}

	if edge.Dst() != "B" {
		t.Errorf("Dst() = %v, want B", edge.Dst())
	}

	if edge.Weight() != 1 {
		t.Errorf("Weight() = %v, want 1", edge.Weight())
	}
}

func TestEdges(t *testing.T) {
	want := []tuples.Edge[int, int]{tuples.NewEdge(1, 2, 1), tuples.NewEdge(3, 4, 1)}
	got := tuples.Edges(
		[3]int{1, 2, 1},
		[3]int{3, 4, 1},
	)

	if !slices.Equal(got, want) {
		t.Errorf("Edges() = %v, want %v", got, want)
	}
}

func TestNewCell(t *testing.T) {
	cell := tuples.NewCell(1, 2)

	if cell.Row() != 1 {
		t.Errorf("Row() = %v, want 1", cell.Row())
	}

	if cell.Col() != 2 {
		t.Errorf("Col() = %v, want 2", cell.Col())
	}
}

func TestCells(t *testing.T) {
	want := []tuples.Cell[int]{tuples.NewCell(1, 2), tuples.NewCell(3, 4)}
	got := tuples.Cells([2]int{1, 2}, [2]int{3, 4})

	if !slices.Equal(got, want) {
		t.Errorf("Cells() = %v, want %v", got, want)
	}
}

func TestPoint2d(t *testing.T) {
	point := tuples.NewPoint2d(1, 2)

	if point.X() != 1 {
		t.Errorf("X() = %v, want 1", point.X())
	}

	if point.Y() != 2 {
		t.Errorf("Y() = %v, want 2", point.Y())
	}
}

func TestPoints2d(t *testing.T) {
	want := []tuples.Point2d[int]{tuples.NewPoint2d(1, 2), tuples.NewPoint2d(3, 4)}
	got := tuples.Points2d([2]int{1, 2}, [2]int{3, 4})

	if !slices.Equal(got, want) {
		t.Errorf("Points2d() = %v, want %v", got, want)
	}
}

func TestNewPoint3d(t *testing.T) {
	point := tuples.NewPoint3d(1, 2, 3)

	if point.X() != 1 {
		t.Errorf("X() = %v, want 1", point.X())
	}

	if point.Y() != 2 {
		t.Errorf("Y() = %v, want 2", point.Y())
	}

	if point.Z() != 3 {
		t.Errorf("Z() = %v, want 3", point.Z())
	}
}

func TestPoints3d(t *testing.T) {
	want := []tuples.Point3d[int]{tuples.NewPoint3d(1, 2, 3), tuples.NewPoint3d(4, 5, 6)}
	got := tuples.Points3d([3]int{1, 2, 3}, [3]int{4, 5, 6})

	if !slices.Equal(got, want) {
		t.Errorf("Points3d() = %v, want %v", got, want)
	}
}
