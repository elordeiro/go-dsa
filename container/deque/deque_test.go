package deque_test

import (
	"testing"

	dq "github.com/elordeiro/go/container/deque"
)

func TestDeque(t *testing.T) {
	d := dq.NewDeque[int]()
	if d.Len() != 0 {
		t.Errorf("Expected 0 len, got %v", d.Len())
	}
	d = dq.NewDeque[int](1, 2, 3, 4)
	if d.Len() != 4 {
		t.Errorf("Expected 4, got %v", d.Len())
	}
	if d.PeekFront() != 1 {
		t.Errorf("Expected 1, got %v", d.PeekFront())
	}
	if d.PeekBack() != 4 {
		t.Errorf("Expected 4, got %v", d.PeekBack())
	}
	d.PushFront(0)
	if d.Len() != 5 {
		t.Errorf("Expected 5, got %v", d.Len())
	}
	if d.PeekFront() != 0 {
		t.Errorf("Expected 0, got %v", d.PeekFront())
	}
	if d.PeekBack() != 4 {
		t.Errorf("Expected 4, got %v", d.PeekBack())
	}
	d.PushBack(5)
	if d.Len() != 6 {
		t.Errorf("Expected 6, got %v", d.Len())
	}
	if d.PeekFront() != 0 {
		t.Errorf("Expected 0, got %v", d.PeekFront())
	}
	if d.PeekBack() != 5 {
		t.Errorf("Expected 5, got %v", d.PeekBack())
	}
	if d.PopFront() != 0 {
		t.Errorf("Expected 0, got %v", d.PopFront())
	}
	if d.Len() != 5 {
		t.Errorf("Expected 5, got %v", d.Len())
	}
	if d.PopBack() != 5 {
		t.Errorf("Expected 5, got %v", d.PopBack())
	}
	if d.Len() != 4 {
		t.Errorf("Expected 4, got %v", d.Len())
	}
}

func TestDequeAll(t *testing.T) {
	type test struct {
		nums     []int
		expected []int
	}

	tests := []test{
		{[]int{1, 2, 3, 4}, []int{1, 2, 3, 4}},
		{[]int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{1, 2}, []int{1, 2}},
		{[]int{1}, []int{1}},
		{[]int{}, []int{}},
	}

	for _, tst := range tests {
		i := 0
		d := dq.NewDeque[int](tst.nums...)
		for v := range d.All() {
			if v != tst.expected[i] {
				t.Errorf("Expected %v, got %v", tst.expected[i], v)
			}
			i++
		}
	}
}

func TestDequeEnumerate(t *testing.T) {
	type test struct {
		nums     []int
		expected [][]int
	}

	tests := []test{
		{[]int{1, 2, 3, 4}, [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}}},
		{[]int{1, 2, 3}, [][]int{{0, 1}, {1, 2}, {2, 3}}},
		{[]int{1, 2}, [][]int{{0, 1}, {1, 2}}},
		{[]int{1}, [][]int{{0, 1}}},
		{[]int{}, [][]int{}},
	}

	for _, tst := range tests {
		d := dq.NewDeque[int](tst.nums...)
		for i, v := range d.Enumerate(0) {
			if v != tst.expected[i][1] {
				t.Errorf("Expected %v, got %v", tst.expected[i], v)
			}
		}
	}
}

func TestDequeString(t *testing.T) {
	type test struct {
		nums     []int
		expected string
	}

	tests := []test{
		{[]int{1, 2, 3, 4}, "<->[1 2 3 4]"},
		{[]int{1, 2, 3}, "<->[1 2 3]"},
		{[]int{1, 2}, "<->[1 2]"},
		{[]int{1}, "<->[1]"},
		{[]int{}, "<->[]"},
	}

	for _, tst := range tests {
		d := dq.NewDeque[int](tst.nums...)
		if d.String() != tst.expected {
			t.Errorf("Expected %v, got %v", tst.expected, d.String())
		}
	}
}
