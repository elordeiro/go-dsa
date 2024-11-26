package avl_test

import (
	"slices"
	"testing"

	"github.com/elordeiro/goext/containers/avl"
	"github.com/elordeiro/goext/seqs"
)

func TestNew(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{[]int{2, 1, 3}, []int{1, 2, 3}},
		{[]int{5, 3, 7, 2, 4, 6, 8, 1, 9}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{[]int{1, 2}, []int{1, 2}},
		{[]int{2, 1}, []int{1, 2}},
		{[]int{1}, []int{1}},
		{[]int{}, []int{}},
	}

	for _, tc := range tests {
		b := avl.New(tc.nums...)
		want := slices.Values(tc.want)
		got := b.Inorder()
		if !seqs.Equal(got, want) {
			t.Errorf("New(%v) = %v; want %v", tc.nums, seqs.String(got), seqs.String(want))
		}
	}
}

func TestInsert(t *testing.T) {
	tests := []struct {
		nums []int
		i    int
		want []int
	}{
		{[]int{1, 2, 3, 4}, 0, []int{0, 1, 2, 3, 4}},
		{[]int{0, 2, 3, 4}, 1, []int{0, 1, 2, 3, 4}},
		{[]int{0, 1, 3, 4}, 2, []int{0, 1, 2, 3, 4}},
		{[]int{0, 1, 2, 4}, 3, []int{0, 1, 2, 3, 4}},
		{[]int{0, 1, 2, 3}, 4, []int{0, 1, 2, 3, 4}},
		{[]int{0, 1, 2, 3, 4}, 5, []int{0, 1, 2, 3, 4, 5}},
		{[]int{0, 1, 2, 3, 4}, 0, []int{0, 1, 2, 3, 4}},
		{[]int{0, 1, 2, 3, 4}, 1, []int{0, 1, 2, 3, 4}},
		{[]int{0, 1, 2, 3, 4}, 2, []int{0, 1, 2, 3, 4}},
		{[]int{0, 1, 2, 3, 4}, 3, []int{0, 1, 2, 3, 4}},
		{[]int{0, 1, 2, 3, 4}, 4, []int{0, 1, 2, 3, 4}},
	}

	for _, tc := range tests {
		b := avl.New(tc.nums...)
		b = b.Insert(tc.i)
		want := slices.Values(tc.want)
		got := b.Inorder()
		if !seqs.Equal(got, want) {
			t.Errorf("Insert(%v) = %v; want %v", tc.i, seqs.String(got), seqs.String(want))
		}
	}
}

func TestInsertRotations(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{[]int{1, 2, 3}, []int{2, 1, 3}},
		{[]int{3, 2, 1}, []int{2, 1, 3}},
		{[]int{3, 1, 2}, []int{2, 1, 3}},
		{[]int{2, 3, 1}, []int{2, 1, 3}},
	}

	for _, tc := range tests {
		b := avl.New(tc.nums...)
		want := slices.Values(tc.want)
		got := b.Levelorder()
		if !seqs.Equal(got, want) {
			t.Errorf("Levelorder() = %v; want %v", seqs.String(got), seqs.String(want))
		}
	}
}

func TestDelete(t *testing.T) {
	tests := []struct {
		nums []int
		i    int
		want []int
	}{
		{[]int{1, 2, 3, 4, 5}, 0, []int{1, 2, 3, 4, 5}},
		{[]int{1, 2, 3, 4, 5}, 1, []int{2, 3, 4, 5}},
		{[]int{1, 2, 3, 4, 5}, 2, []int{1, 3, 4, 5}},
		{[]int{1, 2, 3, 4, 5}, 3, []int{1, 2, 4, 5}},
		{[]int{1, 2, 3, 4, 5}, 4, []int{1, 2, 3, 5}},
		{[]int{1, 2, 3, 4, 5}, 5, []int{1, 2, 3, 4}},
		{[]int{1, 2}, 1, []int{2}},
		{[]int{1, 2}, 2, []int{1}},
		{[]int{1}, 1, []int{}},
		{[]int{}, 0, []int{}},
	}

	for _, tc := range tests {
		b := avl.New(tc.nums...)
		b = b.Delete(tc.i)
		want := slices.Values(tc.want)
		got := b.Inorder()
		if !seqs.Equal(got, want) {
			t.Errorf("Delete(%v) = %v; want %v", tc.i, seqs.String(got), seqs.String(want))
		}
	}
}

func TestDeleteRotations(t *testing.T) {
	tests := []struct {
		nums []int
		i    int
		want []int
	}{
		{[]int{3, 2, 5, 1, 4, 7, 6, 8}, 1, []int{5, 3, 7, 2, 4, 6, 8}},
		{[]int{6, 4, 7, 2, 5, 8, 1, 3}, 8, []int{4, 2, 6, 1, 3, 5, 7}},
		{[]int{6, 2, 7, 1, 4, 8, 3, 5}, 8, []int{4, 2, 6, 1, 3, 5, 7}},
		{[]int{3, 2, 7, 1, 5, 8, 4, 6}, 1, []int{5, 3, 7, 2, 4, 6, 8}},
	}

	for _, tc := range tests {
		b := avl.New(tc.nums...)
		want := slices.Values(tc.want)

		b = b.Delete(tc.i)

		got := b.Levelorder()
		if !seqs.Equal(got, want) {
			t.Errorf("Levelorder() = %v; want %v", seqs.String(got), seqs.String(want))
		}
	}
}

func TestSearch(t *testing.T) {
	tests := []struct {
		nums   []int
		search int
		want   *avl.Tree[int]
	}{
		{[]int{1, 2, 3, 4, 5}, 0, nil},
		{[]int{1, 2, 3, 4, 5}, 1, avl.New(1)},
		{[]int{1, 2, 3, 4, 5}, 2, avl.New(2)},
		{[]int{1, 2, 3, 4, 5}, 3, avl.New(3)},
		{[]int{1, 2, 3, 4, 5}, 4, avl.New(4)},
		{[]int{1, 2, 3, 4, 5}, 5, avl.New(5)},
	}

	for _, test := range tests {
		b := avl.New(test.nums...)
		got := b.Search(test.search)
		if got == nil {
			if test.want != nil {
				t.Errorf("Search(%v) = nil; want %v", test.search, test.want)
			}
			continue
		}
		if got.Value() != test.want.Value() {
			t.Errorf("Search(%v) = %v; want %v", test.search, got, test.want)
		}
	}
}

func TestMax(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, 2, 3, 4, 5}, 5},
		{[]int{5, 4, 3, 2, 1}, 5},
		{[]int{3, 1, 2}, 3},
		{[]int{3, 2, 1}, 3},
		{[]int{3, 1, 2}, 3},
		{[]int{2, 3, 1}, 3},
		{[]int{}, 0},
	}

	for _, tc := range tests {
		b := avl.New(tc.nums...)
		max := b.Max()
		if max == nil {
			if tc.want != 0 {
				t.Errorf("Max() = nil; want %v", tc.want)
			}
			continue
		}
		got := b.Max().Value()
		if got != tc.want {
			t.Errorf("Max() = %v; want %v", got, tc.want)
		}
	}
}

func TestMin(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, 2, 3, 4, 5}, 1},
		{[]int{5, 4, 3, 2, 1}, 1},
		{[]int{3, 1, 2}, 1},
		{[]int{3, 2, 1}, 1},
		{[]int{3, 1, 2}, 1},
		{[]int{2, 3, 1}, 1},
		{[]int{}, 0},
	}

	for _, tc := range tests {
		b := avl.New(tc.nums...)
		min := b.Min()
		if min == nil {
			if tc.want != 0 {
				t.Errorf("Min() = nil; want %v", tc.want)
			}
			continue
		}
		got := b.Min().Value()
		if got != tc.want {
			t.Errorf("Min() = %v; want %v", got, tc.want)
		}
	}
}

func TestLeft(t *testing.T) {
	tests := []struct {
		nums []int
		want *avl.Tree[int]
	}{
		{[]int{3, 1, 2}, avl.New(1)},
		{[]int{3, 2, 1}, avl.New(1)},
		{[]int{2, 3, 1}, avl.New(1)},
		{[]int{}, nil},
	}

	for _, tc := range tests {
		b := avl.New(tc.nums...)
		want := tc.want
		if want == nil {
			if b.Left() != nil {
				t.Errorf("Left() = %v; want nil", b.Left())
			}
			continue
		}
		got := b.Left()
		if got.Value() != want.Value() {
			t.Errorf("Left() = %v; want %v", got, want)
		}
	}
}

func TestRight(t *testing.T) {
	tests := []struct {
		nums []int
		want *avl.Tree[int]
	}{
		{[]int{1, 2, 3}, avl.New(3)},
		{[]int{3, 2, 1}, avl.New(3)},
		{[]int{3, 1, 2}, avl.New(3)},
		{[]int{}, nil},
	}

	for _, tc := range tests {
		b := avl.New(tc.nums...)
		want := tc.want
		if want == nil {
			if b.Right() != nil {
				t.Errorf("Right() = %v; want nil", b.Right())
			}
			continue
		}
		got := b.Right()
		if got.Value() != want.Value() {
			t.Errorf("Right() = %v; want %v", got, want)
		}
	}
}

func TestValue(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, 2, 3}, 2},
		{[]int{3, 2, 1}, 2},
		{[]int{3, 1, 2}, 2},
		{[]int{2, 3, 1}, 2},
		{[]int{}, -1},
	}

	for _, tc := range tests {
		b := avl.New(tc.nums...)
		if tc.want == -1 {
			defer func() {
				if r := recover(); r == nil {
					t.Errorf("Value() = %v; want panic", b.Value())
				}
			}()
			b.Value()
			continue
		}
		got := b.Value()
		if got != tc.want {
			t.Errorf("Value() = %v; want %v", got, tc.want)
		}
	}
}
