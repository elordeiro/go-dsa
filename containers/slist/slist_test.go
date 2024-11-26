package slist_test

import (
	"slices"
	"testing"

	"github.com/elordeiro/goext/containers/slist"
	"github.com/elordeiro/goext/seqs"
)

func TestNew(t *testing.T) {
	tests := []struct {
		nums []int
	}{
		{[]int{1, 2, 3}},
		{[]int{1, 2}},
		{[]int{1}},
		{[]int{}},
	}

	for _, tc := range tests {
		list := slist.New(tc.nums...)
		got := list.All()
		want := slices.Values(tc.nums)
		if !seqs.Equal(got, want) {
			t.Errorf("New(%v) = %v, want %v", tc.nums, seqs.String(got), want)
		}
	}
}

func TestAppend(t *testing.T) {
	tests := []struct {
		nums []int
		val  int
	}{
		{[]int{1, 2, 3}, 4},
		{[]int{1, 2}, 3},
		{[]int{1}, 2},
		{[]int{}, 1},
	}

	for _, test := range tests {
		list := slist.New(test.nums...)
		list = list.Append(test.val)
		got := list.All()
		want := slices.Values(append(test.nums, test.val))

		if !seqs.Equal(got, want) {
			t.Errorf("%v.Append(%v) = %v, want %v", test.nums, test.val, seqs.String(got), want)
		}
	}
}

func TestPrepend(t *testing.T) {
	tests := []struct {
		nums []int
		val  int
	}{
		{[]int{1, 2, 3}, 0},
		{[]int{1, 2}, 0},
		{[]int{1}, 0},
		{[]int{}, 0},
	}

	for _, test := range tests {
		list := slist.New(test.nums...)
		list = list.Prepend(test.val)
		got := list.All()
		want := slices.Values(append([]int{test.val}, test.nums...))

		if !seqs.Equal(got, want) {
			t.Errorf("%v.Prepend(%v) = %v, want %v", test.nums, test.val, seqs.String(got), want)
		}
	}
}

func TestRemove(t *testing.T) {
	tests := []struct {
		nums     []int
		val      int
		expected []int
	}{
		{[]int{1, 1, 1}, 1, []int{1, 1}},
		{[]int{1, 2, 2, 3}, 2, []int{1, 2, 3}},
		{[]int{1, 2, 3}, 2, []int{1, 3}},
		{[]int{1, 2}, 1, []int{2}},
		{[]int{1}, 1, []int{}},
		{[]int{}, 0, []int{}},
	}

	for _, test := range tests {
		list := slist.New(test.nums...)
		list = slist.Remove(list, test.val)
		got := list.All()
		want := slices.Values(test.expected)

		if !seqs.Equal(got, want) {
			t.Errorf("%v.Remove(%v) = %v, want %v", test.nums, test.val, seqs.String(got), want)
		}
	}
}

func TestRemoveAll(t *testing.T) {
	tests := []struct {
		nums     []int
		val      int
		expected []int
	}{
		{[]int{1, 2, 2, 3}, 2, []int{1, 3}},
		{[]int{1, 2, 2}, 2, []int{1}},
		{[]int{2, 2, 2}, 2, []int{}},
		{[]int{1, 2, 3}, 4, []int{1, 2, 3}},
	}

	for _, test := range tests {
		list := slist.New(test.nums...)
		list = slist.RemoveAll(list, test.val)
		got := list.All()
		want := slices.Values(test.expected)

		if !seqs.Equal(got, want) {
			t.Errorf("%v.RemoveAll(%v) = %v, want %v", test.nums, test.val, seqs.String(got), want)
		}
	}
}

func TestCompare(t *testing.T) {
	tests := []struct {
		nums1 []int
		nums2 []int
		want  bool
	}{
		{[]int{1, 2, 3}, []int{1, 2, 3}, true},
		{[]int{1, 2}, []int{1, 2}, true},
		{[]int{1}, []int{1}, true},
		{[]int{}, []int{}, true},
		{[]int{1, 2, 3}, []int{1, 2, 4}, false},
		{[]int{1, 2}, []int{1, 3}, false},
		{[]int{1}, []int{2}, false},
		{[]int{}, []int{1}, false},
	}

	for _, tc := range tests {
		list1, list2 := slist.New(tc.nums1...), slist.New(tc.nums2...)
		got := slist.Compare(list1, list2)
		if got != tc.want {
			t.Errorf("Compare(%v, %v) = %v, want %v", tc.nums1, tc.nums2, got, tc.want)
		}
	}
}
