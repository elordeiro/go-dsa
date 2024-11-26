package deque_test

import (
	"slices"
	"testing"

	"github.com/elordeiro/goext/containers/deque"
	"github.com/elordeiro/goext/seqs"
)

func TestNew(t *testing.T) {
	tests := []struct {
		nums []int
		len  int
		want []int
	}{
		{[]int{1, 2, 3, 4}, 4, []int{1, 2, 3, 4}},
		{[]int{1, 2, 3}, 3, []int{1, 2, 3}},
		{[]int{1, 2}, 2, []int{1, 2}},
		{[]int{1}, 1, []int{1}},
		{[]int{}, 0, []int{}},
	}

	for _, tc := range tests {
		d := deque.New(tc.nums...)
		if d.Len() != tc.len {
			t.Errorf("Len() = %v, want %v", tc.len, d.Len())
		}
		got := d.All()
		want := slices.Values(tc.want)
		if !seqs.Equal(got, want) {
			t.Errorf("All() = %v, want %v", got, want)
		}
	}
}

func TestPushFront(t *testing.T) {
	tests := []struct {
		nums []int
		len  int
		want []int
	}{
		{[]int{4, 3, 2, 1}, 4, []int{1, 2, 3, 4}},
		{[]int{3, 2, 1}, 3, []int{1, 2, 3}},
		{[]int{2, 1}, 2, []int{1, 2}},
		{[]int{1}, 1, []int{1}},
		{[]int{}, 0, []int{}},
	}

	for _, tc := range tests {
		d := deque.New[int]()
		d.PushFront(tc.nums...)
		if d.Len() != tc.len {
			t.Errorf("Len() = %v, want %v", tc.len, d.Len())
		}
		got := d.All()
		want := slices.Values(tc.want)
		if !seqs.Equal(got, want) {
			t.Errorf("All() = %v, want %v", got, want)
		}
	}
}

func TestPushBack(t *testing.T) {
	tests := []struct {
		nums []int
		len  int
		want []int
	}{
		{[]int{1, 2, 3, 4}, 4, []int{1, 2, 3, 4}},
		{[]int{1, 2, 3}, 3, []int{1, 2, 3}},
		{[]int{1, 2}, 2, []int{1, 2}},
		{[]int{1}, 1, []int{1}},
		{[]int{}, 0, []int{}},
	}

	for _, tc := range tests {
		d := deque.New[int]()
		d.PushBack(tc.nums...)
		if d.Len() != tc.len {
			t.Errorf("Len() = %v, want %v", tc.len, d.Len())
		}
		got := d.All()
		want := slices.Values(tc.want)
		if !seqs.Equal(got, want) {
			t.Errorf("All() = %v, want %v", got, want)
		}
	}
}

func TestPopFrontEmpty(t *testing.T) {
	d := deque.New(1)
	d.PopFront()
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("PopFront() did not panic")
		}
	}()
	d.PopFront()
}

func TestPopBackEmpty(t *testing.T) {
	d := deque.New(1)
	d.PopBack()
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("PopFront() did not panic")
		}
	}()
	d.PopBack()
}

func TestFrontEmpty(t *testing.T) {
	d := deque.New[int]()
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Front() did not panic")
		}
	}()
	d.Front()
}

func TestBackEmpty(t *testing.T) {
	d := deque.New[int]()
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Back() did not panic")
		}
	}()
	d.Back()
}

func TestDequeBackwards(t *testing.T) {
	type test struct {
		nums []int
		want []int
	}

	tests := []test{
		{[]int{1, 2, 3, 4}, []int{4, 3, 2, 1}},
		{[]int{1, 2, 3}, []int{3, 2, 1}},
		{[]int{1, 2}, []int{2, 1}},
		{[]int{1}, []int{1}},
		{[]int{}, []int{}},
	}

	for _, tc := range tests {
		d := deque.New(tc.nums...)
		got := d.Backwards()
		want := slices.Values(tc.want)
		if !seqs.Equal(got, want) {
			t.Errorf("Backwards() = %v, want %v", got, want)
		}
	}
}

func TestDrain(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{[]int{1, 2, 3, 4}, []int{1, 2, 3, 4}},
		{[]int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{1, 2}, []int{1, 2}},
		{[]int{1}, []int{1}},
		{[]int{}, []int{}},
	}

	for _, tc := range tests {
		d := deque.New(tc.nums...)
		got := d.Drain()
		want := slices.Values(tc.want)
		if !seqs.Equal(got, want) {
			t.Errorf("Drain() = %v, want %v", got, want)
		}
		if d.Len() != 0 {
			t.Errorf("Len() = %v, want 0", d.Len())
		}
	}
}

func TestDrainBackwards(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{[]int{1, 2, 3, 4}, []int{4, 3, 2, 1}},
		{[]int{1, 2, 3}, []int{3, 2, 1}},
		{[]int{1, 2}, []int{2, 1}},
		{[]int{1}, []int{1}},
		{[]int{}, []int{}},
	}

	for _, tc := range tests {
		d := deque.New(tc.nums...)
		got := d.DrainBackwards()
		want := slices.Values(tc.want)
		if !seqs.Equal(got, want) {
			t.Errorf("DrainBackwards() = %v, want %v", got, want)
		}
		if d.Len() != 0 {
			t.Errorf("Len() = %v, want 0", d.Len())
		}
	}
}

func TestDequeString(t *testing.T) {
	type test struct {
		nums []int
		want string
	}

	tests := []test{
		{[]int{1, 2, 3, 4}, "<1 2 3 4>"},
		{[]int{1, 2, 3}, "<1 2 3>"},
		{[]int{1, 2}, "<1 2>"},
		{[]int{1}, "<1>"},
		{[]int{}, "<>"},
	}

	for _, tc := range tests {
		d := deque.New[int](tc.nums...)
		got := d.String()
		if got != tc.want {
			t.Errorf("String() = %v, want %v", got, tc.want)
		}
	}
}
