package stack_test

import (
	"slices"
	"testing"

	"github.com/elordeiro/goext/containers/stack"
	"github.com/elordeiro/goext/seqs"
)

func TestNew(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{[]int{1, 2, 3}, []int{3, 2, 1}},
		{[]int{1, 2}, []int{2, 1}},
		{[]int{1}, []int{1}},
		{[]int{}, []int{}},
	}

	for _, tc := range tests {
		s := stack.New(tc.nums...)
		got := seqs.MultiUse(s.All())
		want := slices.Values(tc.want)
		if !seqs.Equal(got, want) {
			t.Errorf("New(%v) = %v, want %v", tc.nums, seqs.String(got), seqs.String(want))
		}
	}
}

func TestPush(t *testing.T) {
	tests := []struct {
		nums []int
		vals []int
		want []int
	}{
		{[]int{1, 2, 3}, []int{4, 5, 6}, []int{6, 5, 4, 3, 2, 1}},
		{[]int{1, 2}, []int{3, 4}, []int{4, 3, 2, 1}},
		{[]int{1}, []int{2}, []int{2, 1}},
		{[]int{}, []int{1}, []int{1}},
	}

	for _, tc := range tests {
		s := stack.New(tc.nums...)
		s.Push(tc.vals...)
		got := seqs.MultiUse(s.All())
		want := slices.Values(tc.want)
		if !seqs.Equal(got, want) {
			t.Errorf("New(%v).Push(%v) = %v, want %v", tc.nums, tc.vals, seqs.String(got), seqs.String(want))
		}
	}
}

func TestPop(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, 2, 3}, 3},
		{[]int{1, 2}, 2},
		{[]int{1}, 1},
	}

	for _, tc := range tests {
		s := stack.New(tc.nums...)
		got := s.Pop()
		if got != tc.want {
			t.Errorf("New(%v).Pop() = %v, want %v", tc.nums, got, tc.want)
		}
	}
}

func TestIsEmpty(t *testing.T) {
	tests := []struct {
		nums []int
		want bool
	}{
		{[]int{1, 2, 3}, false},
		{[]int{1, 2}, false},
		{[]int{1}, false},
		{[]int{}, true},
	}

	for _, tc := range tests {
		s := stack.New(tc.nums...)
		got := s.IsEmpty()
		if got != tc.want {
			t.Errorf("New(%v).IsEmpty() = %v, want %v", tc.nums, got, tc.want)
		}
	}
}

func TestTop(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, 2, 3}, 3},
		{[]int{1, 2}, 2},
		{[]int{1}, 1},
	}

	for _, tc := range tests {
		s := stack.New(tc.nums...)
		got := s.Top()
		if got != tc.want {
			t.Errorf("New(%v).Top() = %v, want %v", tc.nums, got, tc.want)
		}
	}
}

func TestLen(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, 2, 3}, 3},
		{[]int{1, 2}, 2},
		{[]int{1}, 1},
		{[]int{}, 0},
	}

	for _, tc := range tests {
		s := stack.New(tc.nums...)
		got := s.Len()
		if got != tc.want {
			t.Errorf("New(%v).Len() = %v, want %v", tc.nums, got, tc.want)
		}
	}
}

func TestAll(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{[]int{1, 2, 3}, []int{3, 2, 1}},
		{[]int{1, 2}, []int{2, 1}},
		{[]int{1}, []int{1}},
		{[]int{}, []int{}},
	}

	for _, tc := range tests {
		s := stack.New(tc.nums...)
		got := seqs.MultiUse(s.All())
		want := slices.Values(tc.want)
		if !seqs.Equal(got, want) {
			t.Errorf("New(%v).All() = %v, want %v", tc.nums, seqs.String(got), seqs.String(want))
		}
	}
}

func TestBackwards(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{[]int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{1, 2}, []int{1, 2}},
		{[]int{1}, []int{1}},
		{[]int{}, []int{}},
	}

	for _, tc := range tests {
		s := stack.New(tc.nums...)
		got := seqs.MultiUse(s.Backwards())
		want := slices.Values(tc.want)
		if !seqs.Equal(got, want) {
			t.Errorf("New(%v).Backwards() = %v, want %v", tc.nums, seqs.String(got), seqs.String(want))
		}
	}
}
