package pq_test

import (
	"slices"
	"testing"

	"github.com/elordeiro/goext/containers/pq"
	"github.com/elordeiro/goext/containers/tuples"
	"github.com/elordeiro/goext/containers/vector"
	"github.com/elordeiro/goext/seq2s"
	"github.com/elordeiro/goext/seqs"
	"github.com/elordeiro/goext/seqs/transform"
)

func TestPQAll(t *testing.T) {
	tests := []struct {
		vals []int
		want []int
	}{
		{
			[]int{7, 3, 1, 5},
			[]int{7, 5, 3, 1},
		},
		{
			[]int{1, 3, 5, 7},
			[]int{7, 5, 3, 1},
		},
		{
			[]int{},
			[]int{},
		},
		{
			nil,
			[]int{},
		},
	}

	for _, tc := range tests {
		pq := pq.NewMaxPQ(tc.vals...)
		t.Run("pq all", func(t *testing.T) {
			want := slices.Values(tc.want)
			got := pq.All()
			if !seqs.Equal(got, want) {
				t.Errorf("All() = %v, want %v", seqs.String(got), seqs.String(want))
			}
			if pq.Len() != len(tc.want) {
				t.Errorf("Len() = %v, want %v", pq.Len(), len(tc.want))
			}
		})
	}
}

func TestNewMaxPQ(t *testing.T) {
	tests := []struct {
		vals []int
		want []int
	}{
		{
			[]int{7, 3, 1, 5},
			[]int{7, 5, 3, 1},
		},
		{
			[]int{1, 3, 5, 7},
			[]int{7, 5, 3, 1},
		},
		{
			[]int{},
			[]int{},
		},
		{
			nil,
			[]int{},
		},
	}

	for _, tc := range tests {
		pq := pq.NewMaxPQ(tc.vals...)
		t.Run("max pq", func(t *testing.T) {
			want := slices.Values(tc.want)
			got := pq.All()
			if !seqs.Equal(got, want) {
				t.Errorf("All() = %v, want %v", seqs.String(got), seqs.String(want))
			}
		})
	}
}

func TestPush(t *testing.T) {
	tests := []struct {
		vals []int
		push []int
		want []int
	}{
		{
			[]int{7, 3, 1, 5},
			[]int{10},
			[]int{10, 7, 5, 3, 1},
		},
		{
			[]int{1, 3, 5, 7},
			[]int{4},
			[]int{7, 5, 4, 3, 1},
		},
		{
			[]int{1, 3, 5, 7},
			[]int{0},
			[]int{7, 5, 3, 1, 0},
		},
		{
			[]int{},
			[]int{0},
			[]int{0},
		},
		{
			[]int{},
			[]int{0, 1, 2, 3, 4},
			[]int{4, 3, 2, 1, 0},
		},
		{
			nil,
			[]int{0},
			[]int{0},
		},
	}

	for _, tc := range tests {
		pq := pq.NewMaxPQ(tc.vals...)
		pq.Push(tc.push...)
		t.Run("max pq", func(t *testing.T) {
			want := slices.Values(tc.want)
			got := pq.All()
			if !seqs.Equal(got, want) {
				t.Errorf("All() = %v, want %v", seqs.String(got), seqs.String(want))
			}
		})
	}
}

func TestMinPQ(t *testing.T) {
	tests := []struct {
		vals []int
		want []int
	}{
		{
			[]int{7, 3, 1, 5},
			[]int{1, 3, 5, 7},
		},
		{
			[]int{1, 3, 5, 7},
			[]int{1, 3, 5, 7},
		},
		{
			[]int{},
			[]int{},
		},
		{
			nil,
			[]int{},
		},
	}

	for _, tc := range tests {
		pq := pq.NewMinPQ(tc.vals...)
		t.Run("min pq", func(t *testing.T) {
			want := slices.Values(tc.want)
			got := pq.All()
			if !seqs.Equal(got, want) {
				t.Errorf("All() = %v, want %v", seqs.String(got), seqs.String(want))
			}
		})
	}
}

func TestNewPQFunc(t *testing.T) {
	tests := []struct {
		vals []tuples.Pair[string, int]
		want []string
	}{
		{
			tuples.Pairs(
				tuples.NewPair("banana", 7),
				tuples.NewPair("orange", 3),
				tuples.NewPair("apple", 1),
				tuples.NewPair("grape", 5)),
			[]string{"banana", "grape", "orange", "apple"},
		},
		{
			nil,
			[]string{},
		},
	}

	for _, tc := range tests {
		pq := pq.NewPQFunc(func(item1, item2 tuples.Pair[string, int]) bool {
			return item1.Right() > item2.Right()
		}, tc.vals...)
		t.Run("max pq func", func(t *testing.T) {
			want := slices.Values(tc.want)
			got := seq2s.Keys(transform.Unpair(pq.All()))
			if !seqs.Equal(got, want) {
				t.Errorf("All() = %v, want %v", seqs.String(got), seqs.String(want))
			}
		})
	}
}

func TestUpdate(t *testing.T) {
	tests := []struct {
		vals []tuples.Pair[string, int]
		want []string
	}{
		{
			tuples.Pairs(
				tuples.NewPair("banana", 7),
				tuples.NewPair("orange", 3),
				tuples.NewPair("apple", 1),
				tuples.NewPair("grape", 5)),
			[]string{"apple", "banana", "grape", "orange"},
		},
	}

	for _, tc := range tests {
		pq := pq.NewPQFunc(func(item1, item2 tuples.Pair[string, int]) bool {
			return item1.Right() > item2.Right()
		}, tc.vals...)
		pq.Update(2, tuples.NewPair("apple", 10))
		t.Run("update max pq", func(t *testing.T) {
			want := slices.Values(tc.want)
			got := seq2s.Keys(transform.Unpair(pq.All()))
			if !seqs.Equal(got, want) {
				t.Errorf("All() = %v, want %v", seqs.String(got), seqs.String(want))
			}
		})
	}
}

func TestRemove(t *testing.T) {
	tests := []struct {
		vals []int
		i, j int
		want []int
	}{
		{
			[]int{7, 3, 1, 5},
			0, 9,
			[]int{9, 5, 3, 1},
		},
		{
			[]int{7, 3, 1, 5},
			0, 0,
			[]int{5, 3, 1, 0},
		},
		{
			[]int{7, 3, 1, 5},
			2, 2,
			[]int{7, 5, 3, 2},
		},
	}

	for _, tc := range tests {
		pq := pq.NewMaxPQ(tc.vals...)
		pq.Remove(tc.i)
		pq.Push(tc.j)
		t.Run("remove max pq", func(t *testing.T) {
			want := slices.Values(tc.want)
			got := pq.All()
			if !seqs.Equal(got, want) {
				t.Errorf("All() = %v, want %v", seqs.String(got), seqs.String(want))
			}
		})
	}
}

// For testing NewPQFrom

type vec struct {
	vector.Vector[int]
}

func (v vec) Less(i, j int) bool {
	return v.Vector[i] > v.Vector[j]
}

func (v vec) Copy() pq.Interface[int] {
	new := make([]int, len(v.Vector))
	v.Vector.Copy(new)
	return pq.Interface[int](&vec{new})
}

func TestFrom(t *testing.T) {

	tests := []struct {
		vals1 []int
		vals2 []int
		want  []int
	}{
		{
			[]int{7, 3, 1, 5},
			[]int{},
			[]int{7, 5, 3, 1},
		},
		{
			[]int{7, 3, 1, 5},
			[]int{9, 0},
			[]int{9, 7, 5, 3, 1, 0},
		},
		{
			[]int{1, 3, 5, 7},
			[]int{},
			[]int{7, 5, 3, 1},
		},
		{
			[]int{},
			[]int{},
			[]int{},
		},
		{
			nil,
			[]int{},
			[]int{},
		},
	}

	for _, tc := range tests {
		v := &vec{tc.vals1}
		pq := pq.NewPQFrom(v, tc.vals2...)
		t.Run("pq from", func(t *testing.T) {
			want := slices.Values(tc.want)
			got := pq.All()
			if !seqs.Equal(got, want) {
				t.Errorf("All() = %v, want %v", seqs.String(got), seqs.String(want))
			}
		})
	}
}

func TestDrain(t *testing.T) {
	tests := []struct {
		vals []int
		want []int
	}{
		{
			[]int{7, 3, 1, 5},
			[]int{7, 5, 3, 1},
		},
		{
			[]int{1, 3, 5, 7},
			[]int{7, 5, 3, 1},
		},
		{
			[]int{},
			[]int{},
		},
		{
			nil,
			[]int{},
		},
	}

	for _, tc := range tests {
		pq := pq.NewMaxPQ(tc.vals...)
		t.Run("drain", func(t *testing.T) {
			want := slices.Values(tc.want)
			got := pq.Drain()
			if !seqs.Equal(got, want) {
				t.Errorf("Drain() = %v, want %v", seqs.String(got), seqs.String(want))
			}
			if pq.Len() != 0 {
				t.Errorf("Len() = %v, want 0", pq.Len())
			}
		})
	}
}
