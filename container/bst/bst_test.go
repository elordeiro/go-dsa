package bst_test

import (
	"testing"

	bt "github.com/elordeiro/go/container/bst"
)

func TestNewBst(t *testing.T) {
	tests := []struct {
		nums     []int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{[]int{2, 1, 3}, []int{1, 2, 3}},
		{[]int{5, 3, 7, 2, 4, 6, 8, 1, 9}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{[]int{1, 2}, []int{1, 2}},
		{[]int{2, 1}, []int{1, 2}},
		{[]int{1}, []int{1}},
		{[]int{}, []int{}},
	}

	for _, test := range tests {
		b := bt.NewBst(test.nums...)
		for i, v := range b.Enumerate(0, b.Inorder()) {
			if v != test.expected[i] {
				t.Errorf("Expected %v, but got %v", test.expected, b.Enumerate(0, b.Inorder()))
			}
		}
	}
}

func TestInsert(t *testing.T) {
	tests := []struct {
		nums     []int
		insert   int
		expected []int
	}{
		{[]int{1, 2, 3, 4}, 0, []int{0, 1, 2, 3, 4}},
		{[]int{1, 2, 3, 4}, 1, []int{1, 1, 2, 3, 4}},
		{[]int{1, 2, 3, 4}, 2, []int{1, 2, 2, 3, 4}},
		{[]int{1, 2, 3, 4}, 3, []int{1, 2, 3, 3, 4}},
		{[]int{1, 2, 3, 4}, 4, []int{1, 2, 3, 4, 4}},
		{[]int{1, 2, 3, 4}, 5, []int{1, 2, 3, 4, 5}},
	}

	for _, test := range tests {
		b := bt.NewBst(test.nums...)
		b = b.Insert(test.insert)
		for i, v := range b.Enumerate(0, b.Inorder()) {
			if v != test.expected[i] {
				t.Errorf("Expected %v, but got %v", test.expected, b.Enumerate(0, b.Inorder()))
			}
		}
	}
}

func TestDelete(t *testing.T) {
	tests := []struct {
		nums     []int
		delete   int
		expected []int
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

	for _, test := range tests {
		b := bt.NewBst(test.nums...)
		b = b.Delete(test.delete)
		for i, v := range b.Enumerate(0, b.Inorder()) {
			if v != test.expected[i] {
				t.Errorf("Expected %v, but got %v", test.expected, b.Enumerate(0, b.Inorder()))
			}
		}
	}
}
