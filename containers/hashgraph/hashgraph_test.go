package hashgraph_test

import (
	"slices"
	"testing"

	"github.com/elordeiro/goext/containers/hashgraph"
	"github.com/elordeiro/goext/containers/tuples"
	"github.com/elordeiro/goext/seq2s"
	"github.com/elordeiro/goext/seqs"
	"github.com/elordeiro/goext/seqs/transform"
)

func graphTest(directed bool) *hashgraph.HashGraph[int, int] {
	g := hashgraph.New[int, int](directed)
	g.AddEdge(1, 2, 1)
	g.AddEdge(1, 3)
	g.AddEdge(2, 4, 1)
	g.AddEdge(2, 5)
	g.AddEdge(3, 6, 1)
	g.AddEdge(3, 7)
	return g
}

func TestAddEdge(t *testing.T) {
	g := hashgraph.New[int, int](false)
	g.AddEdge(1, 2)
	g.AddEdge(1, 3)
	g.AddEdge(2, 4)

	if got := g.VertexCount(); got != 4 {
		t.Errorf("VertexCount() = %v, want 4", got)
	}

	if got := seqs.Len(g.Edges()); got != 3 {
		t.Errorf("EdgeCount() = %v, want 3", got)
	}

	if got := g.HasEdge(2, 1); !got {
		t.Errorf("HasEdge(1, 2) = false, want true")
	}

	if got := g.HasEdge(3, 1); !got {
		t.Errorf("HasEdge(1, 3) = false, want true")
	}

	if got := g.HasEdge(4, 2); !got {
		t.Errorf("HasEdge(2, 4) = false, want true")
	}
}

func TestHasVertex(t *testing.T) {
	g := graphTest(true)

	tests := []struct {
		vertex int
		want   bool
	}{
		{1, true},
		{2, true},
		{3, true},
		{4, true},
		{5, true},
		{6, true},
		{7, true},
		{8, false},
	}

	for _, tc := range tests {
		if got := g.HasVertex(tc.vertex); got != tc.want {
			t.Errorf("HasVertex(%v) = %v, want %v", tc.vertex, got, tc.want)
		}
	}
}

func TestRemoveVertex(t *testing.T) {
	g := graphTest(true)
	for tc := range 8 {
		g.RemoveVertex(tc)
		if got := g.HasVertex(tc); got {
			t.Errorf("HasVertex(%v) = true, want false", tc)
		}
	}
}

func TestDegree(t *testing.T) {
	g := graphTest(true)

	tests := []struct {
		vertex int
		want   int
	}{
		{1, 2},
		{2, 2},
		{3, 2},
		{4, 0},
		{5, 0},
		{6, 0},
		{7, 0},
		{8, 0},
	}

	for _, tc := range tests {
		got, err := g.Degree(tc.vertex)
		if err != nil {
			errStr := "vertex not found: 8"
			if got := err.Error(); got != errStr {
				t.Errorf("Degree(%v) = %v, want %v", tc.vertex, got, errStr)
			}
			continue
		}

		if got != tc.want {
			t.Errorf("Degree(%v) = %v, want %v", tc.vertex, got, tc.want)
		}
	}
}

func TestRemoveEdge(t *testing.T) {
	g := graphTest(true)

	tests := []struct {
		src, dst int
	}{
		{1, 2},
		{1, 3},
		{2, 4},
		{2, 5},
		{3, 6},
		{3, 7},
	}

	for _, tc := range tests {
		g.RemoveEdge(tc.src, tc.dst)
		if got := g.HasEdge(tc.src, tc.dst); got {
			t.Errorf("HasEdge(%v, %v) = true, want false", tc.src, tc.dst)
		}
	}

	if got := g.VertexCount(); got != 7 {
		t.Errorf("VertexCount() = %v, want 7", got)
	}

	g = graphTest(false)

	for _, tc := range tests {
		g.RemoveEdge(tc.src, tc.dst)
		if got := g.HasEdge(tc.src, tc.dst); got {
			t.Errorf("HasEdge(%v, %v) = true, want false", tc.src, tc.dst)
		}
		if got := g.HasEdge(tc.dst, tc.src); got {
			t.Errorf("HasEdge(%v, %v) = true, want false", tc.dst, tc.src)
		}
	}

	if got := g.VertexCount(); got != 7 {
		t.Errorf("VertexCount() = %v, want 7", got)
	}
}

func TestHasEdge(t *testing.T) {
	g := graphTest(true)

	tests := []struct {
		src, dst, weight int
		want             bool
	}{
		{1, 2, 1, true},
		{1, 2, 2, false},
		{1, 3, 1, true},
		{1, 3, 2, false},
		{2, 4, 1, true},
		{2, 4, 2, false},
		{2, 5, 1, true},
		{2, 5, 2, false},
		{3, 6, 1, true},
		{3, 6, 2, false},
		{3, 7, 1, true},
		{3, 7, 2, false},
		{4, 2, 1, false},
		{5, 2, 1, false},
		{6, 3, 1, false},
		{7, 3, 1, false},
		{8, 9, 1, false},
	}

	for _, tc := range tests {
		if got := g.HasEdge(tc.src, tc.dst, tc.weight); got != tc.want {
			t.Errorf("HasEdge(%v, %v, %v) = %v, want %v", tc.src, tc.dst, tc.weight, got, tc.want)
		}
	}
}

func TestEdge(t *testing.T) {
	g := graphTest(true)

	tests := []struct {
		src, dst int
		want     tuples.Edge[int, int]
	}{
		{1, 2, tuples.NewEdge(1, 2, 1)},
		{1, 3, tuples.NewEdge(1, 3, 1)},
		{2, 4, tuples.NewEdge(2, 4, 1)},
		{2, 5, tuples.NewEdge(2, 5, 1)},
		{3, 6, tuples.NewEdge(3, 6, 1)},
		{3, 7, tuples.NewEdge(3, 7, 1)},
		{4, 2, tuples.Edge[int, int]{}},
	}

	for _, tc := range tests {
		got, err := g.Edge(tc.src, tc.dst)
		if err != nil {
			errStr := "edge not found: 4 -> 2"
			if got := err.Error(); got != errStr {
				t.Errorf("Edge(%v, %v) = %v, want %v", tc.src, tc.dst, got, errStr)
			}
			continue
		}
		if got != tc.want {
			t.Errorf("Edge(%v, %v) = %v, want %v", tc.src, tc.dst, got, tc.want)
		}
	}
}

func TestEdgeWeight(t *testing.T) {
	g := graphTest(true)

	tests := []struct {
		src, dst int
		want     int
	}{
		{1, 2, 1},
		{1, 3, 1},
		{2, 4, 1},
		{2, 5, 1},
		{3, 6, 1},
		{3, 7, 1},
		{4, 2, 0},
	}

	for _, tc := range tests {
		got, err := g.EdgeWeight(tc.src, tc.dst)
		if err != nil {
			errStr := "edge not found: 4 -> 2"
			if got := err.Error(); got != errStr {
				t.Errorf("EdgeWeight(%v, %v) = %v, want %v", tc.src, tc.dst, got, errStr)
			}
			continue
		}
		if got != tc.want {
			t.Errorf("EdgeWeight(%v, %v) = %v, want %v", tc.src, tc.dst, got, tc.want)
		}
	}
}

func TestSetEdgeWeight(t *testing.T) {
	g := graphTest(true)

	tests := []struct {
		src, dst, weight int
	}{
		{1, 2, 2},
		{1, 3, 2},
		{2, 4, 2},
		{2, 5, 2},
		{3, 6, 2},
		{3, 7, 2},
	}

	for _, tc := range tests {
		g.SetEdgeWeight(tc.src, tc.dst, tc.weight)
		got, _ := g.EdgeWeight(tc.src, tc.dst)
		if got != tc.weight {
			t.Errorf("HasEdge(%v, %v, %v) = false, want true", tc.src, tc.dst, tc.weight)
		}
	}
}

func TestEdgeByLabel(t *testing.T) {
	g := hashgraph.New[int, int](true)
	g.AddVertex(1)
	g.AddVertex(2)
	g.AddVertex(3)
	g.AddVertex(4)
	g.AddVertex(5)
	g.AddVertex(6)
	g.AddVertex(7)

	g.AddLabel(1, "a")
	g.AddLabel(2, "b")
	g.AddLabel(3, "c")
	g.AddLabel(4, "d")
	g.AddLabel(5, "e")
	g.AddLabel(6, "f")
	g.AddLabel(7, "g")

	g.AddEdgeByLabel("a", "b", 1)
	g.AddEdgeByLabel("a", "c", 2)
	g.AddEdgeByLabel("b", "d", 3)
	g.AddEdgeByLabel("b", "e", 4)
	g.AddEdgeByLabel("c", "f", 5)
	g.AddEdgeByLabel("c", "g", 6)

	vertexTests := []struct {
		label string
		want  int
	}{
		{"a", 1},
		{"b", 2},
		{"c", 3},
		{"d", 4},
		{"e", 5},
		{"f", 6},
		{"g", 7},
	}

	for _, tc := range vertexTests {
		got, _ := g.VertexByLabel(tc.label)
		if got != tc.want {
			t.Errorf("VertexByLabel(%v) = %v, want %v", tc.label, got, tc.want)
		}
	}

	got, err := g.VertexByLabel("z")
	want := "vertex not found: z"
	if err == nil {
		t.Errorf("VertexByLabel(%v) = %v, want %v",
			"z", got, want,
		)
	} else if err.Error() != want {
		t.Errorf("VertexByLabel(%v) = %v, want %v",
			"z", err.Error(), want,
		)
	}

	err = g.AddEdgeByLabel("a", "z", 7)
	if err == nil {
		t.Errorf("AddEdgeByLabel(%v, %v) = nil, want %v",
			"a", "z", want,
		)
	} else if err.Error() != want {
		t.Errorf("AddEdgeByLabel(%v, %v) = %v, want %v",
			"a", "z", err.Error(), want,
		)
	}

	err = g.AddEdgeByLabel("z", "a", 7)
	if err == nil {
		t.Errorf("AddEdgeByLabel(%v, %v) = nil, want %v",
			"z", "a", want,
		)
	} else if err.Error() != want {
		t.Errorf("AddEdgeByLabel(%v, %v) = %v, want %v",
			"z", "a", err.Error(), want,
		)
	}

	tests := []struct {
		src, dst string
		want     tuples.Edge[int, int]
	}{
		{"a", "b", tuples.NewEdge(1, 2, 1)},
		{"a", "c", tuples.NewEdge(1, 3, 2)},
		{"b", "d", tuples.NewEdge(2, 4, 3)},
		{"b", "e", tuples.NewEdge(2, 5, 4)},
		{"c", "f", tuples.NewEdge(3, 6, 5)},
		{"c", "g", tuples.NewEdge(3, 7, 6)},
		{"a", "z", tuples.Edge[int, int]{}},
		{"z", "a", tuples.Edge[int, int]{}},
	}

	for _, tc := range tests {
		got, err := g.EdgeByLabel(tc.src, tc.dst)
		if err != nil {
			errStr := "vertex not found: z"
			if got := err.Error(); got != errStr {
				t.Errorf("EdgeByLabel(%v, %v) = %v, want %v", tc.src, tc.dst, got, errStr)
			}
			continue
		}
		if got != tc.want {
			t.Errorf("EdgeByLabel(%v, %v) = %v, want %v", tc.src, tc.dst, got, tc.want)
		}
	}
}

func TestClear(t *testing.T) {
	g := graphTest(true)
	g.Clear()

	if got := g.VertexCount(); got != 0 {
		t.Errorf("VertexCount() = %v, want 0", got)
	}

	if got := seqs.Len(g.Edges()); got != 0 {
		t.Errorf("EdgeCount() = %v, want 0", got)
	}
}

func TestClone(t *testing.T) {
	g := graphTest(true)
	clone := g.Clone()

	want := g.Edges()
	got := clone.Edges()
	if !seqs.EqualUnordered(got, want) {
		t.Errorf("Edges() = %v, want %v", seqs.String(got), seqs.String(want))
	}
}

func TestTranspose(t *testing.T) {
	g := graphTest(true)
	transpose := g.Transpose()

	want := g.Edges()
	got := transform.Map(transpose.Edges(), func(e tuples.Edge[int, int]) tuples.Edge[int, int] {
		return tuples.NewEdge(e.Dst(), e.Src(), e.Weight())
	})

	if !seqs.EqualUnordered(got, want) {
		t.Errorf("Edges() = %v, want %v", seqs.String(got), seqs.String(want))
	}

}

func TestIsDirected(t *testing.T) {
	g := graphTest(true)

	if got := g.IsDirected(); !got {
		t.Errorf("IsDirected() = false, want true")
	}
}

func TestEdges(t *testing.T) {
	g := graphTest(true)

	want := slices.Values(
		tuples.Edges(
			[3]int{1, 2, 1}, [3]int{1, 3, 1}, [3]int{2, 4, 1}, [3]int{2, 5, 1}, [3]int{3, 6, 1}, [3]int{3, 7, 1},
		),
	)
	got := g.Edges()
	if !seqs.EqualUnordered(got, want) {
		t.Errorf("Edges() = %v, want %v", seqs.String(got), seqs.String(want))
	}
}

func TestVertices(t *testing.T) {
	g := graphTest(true)

	want := slices.Values([]int{1, 2, 3, 4, 5, 6, 7})
	got := g.Vertices()
	if !seqs.EqualUnordered(got, want) {
		t.Errorf("Vertices() = %v, want %v", seqs.String(got), seqs.String(want))
	}
}

func TestNeighbors(t *testing.T) {
	g := graphTest(true)

	tests := []struct {
		vertex int
		want   []int
	}{
		{1, []int{2, 3}},
		{2, []int{4, 5}},
		{3, []int{6, 7}},
		{4, []int{}},
		{5, []int{}},
		{6, []int{}},
		{7, []int{}},
		{8, []int{}},
	}

	for _, tc := range tests {
		got := seq2s.Keys(g.Neighbors(tc.vertex))
		want := slices.Values(tc.want)
		if !seqs.EqualUnordered(got, want) {
			t.Errorf("Neighbors(%v) = %v, want %v", tc.vertex, seqs.String(got), seqs.String(want))
		}
	}
}
