package list_test

import (
	"testing"

	"github.com/elordeiro/go/container/list"
)

func TestNewList(t *testing.T) {
	tests := []struct {
		nums []int
	}{
		{[]int{1, 2, 3}},
		{[]int{1, 2}},
		{[]int{1}},
		{[]int{}},
	}

	for _, test := range tests {
		testname := "Create List"
		t.Run(testname, func(t *testing.T) {
			list := list.NewList(test.nums...)
			for _, val := range test.nums {
				if val != list.Val {
					t.Errorf("Actual   : %v", list.Val)
					t.Errorf("Expected : %v", val)
				}
				list = list.Next
			}
		})
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
		testname := "Append Value"
		t.Run(testname, func(t *testing.T) {
			list := list.NewList(test.nums...)
			list = list.Append(test.val)
			for _, val := range test.nums {
				if val != list.Val {
					t.Errorf("Actual   : %v", list.Val)
					t.Errorf("Expected : %v", val)
				}
				list = list.Next
			}
			if test.val != list.Val {
				t.Errorf("Actual   : %v", list.Val)
				t.Errorf("Expected : %v", test.val)
			}
		})
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
		testname := "Prepend Value"
		t.Run(testname, func(t *testing.T) {
			list := list.NewList(test.nums...)
			list = list.Prepend(test.val)
			if test.val != list.Val {
				t.Errorf("Actual   : %v", list.Val)
				t.Errorf("Expected : %v", test.val)
			}
			list = list.Next
			for _, val := range test.nums {
				if val != list.Val {
					t.Errorf("Actual   : %v", list.Val)
					t.Errorf("Expected : %v", val)
				}
				list = list.Next
			}
		})
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
		testname := "Remove Value"
		t.Run(testname, func(t *testing.T) {
			l := list.NewList(test.nums...)
			l = list.Remove(l, test.val)
			for _, val := range test.expected {
				if val != l.Val {
					t.Errorf("Actual   : %v", l.Val)
					t.Errorf("Expected : %v", val)
				}
				l = l.Next
			}
		})
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
		testname := "Remove Multiple Values"
		t.Run(testname, func(t *testing.T) {
			l := list.NewList(test.nums...)
			l = list.RemoveAll(l, test.val)
			for _, val := range test.expected {
				if val != l.Val {
					t.Errorf("Actual   : %v", l.Val)
					t.Errorf("Expected : %v", val)
				}
				l = l.Next
			}
		})
	}
}

func TestCompare(t *testing.T) {
	tests := []struct {
		nums1 []int
		nums2 []int
	}{
		{[]int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{1, 2}, []int{1, 2}},
		{[]int{1}, []int{1}},
		{[]int{}, []int{}},
	}

	for _, test := range tests {
		list1, list2 := list.NewList(test.nums1...), list.NewList(test.nums2...)
		testname := "Compare Lists"
		t.Run(testname, func(t *testing.T) {
			if !list.Compare(list1, list2) {
				t.Errorf("Actual   : %v", list1)
				t.Errorf("Expected : %v", list2)
			}
		})
	}
}

func TestCompareDifferent(t *testing.T) {
	tests := []struct {
		nums1 []int
		nums2 []int
	}{
		{[]int{1, 2, 3}, []int{1, 2, 4}},
		{[]int{1, 2}, []int{1, 3}},
		{[]int{1}, []int{2}},
		{[]int{}, []int{1}},
	}

	for _, test := range tests {
		list1, list2 := list.NewList(test.nums1...), list.NewList(test.nums2...)
		testname := "Compare Different Lists"
		t.Run(testname, func(t *testing.T) {
			if list.Compare(list1, list2) {
				t.Errorf("Actual   : %v", list1)
				t.Errorf("Expected : %v", list2)
			}
		})
	}
}

func TestEnumerate(t *testing.T) {
	type testStruct struct {
		index int
		val   string
	}
	tests := []struct {
		nums     []string
		expected []testStruct
	}{
		{[]string{"a", "b", "c"}, []testStruct{{0, "a"}, {1, "b"}, {2, "c"}}},
		{[]string{"a", "b"}, []testStruct{{0, "a"}, {1, "b"}}},
		{[]string{"a"}, []testStruct{{0, "a"}}},
		{[]string{}, []testStruct{}},
	}

	for _, test := range tests {
		testname := "Enumerate List"
		t.Run(testname, func(t *testing.T) {
			list := list.NewList(test.nums...)
			for i, val := range list.Enumerate(0) {
				if val != test.expected[i].val {
					t.Errorf("Actual   : %v", val)
					t.Errorf("Expected : %v", test.expected[i])
				}
			}
		})
	}
}
