package binarytree

import "testing"

func TestBinaryTree(t *testing.T) {
	tests := []struct {
		nums1    []int
		nums2    []int
		expected bool
	}{
		{[]int{3, 4, 5, 1, 2}, []int{4, 1, 2}, true},
		{[]int{3, 4, 5, 1, 2, null, null, null, null, 0}, []int{4, 1, 2}, false},
		{[]int{1, 1}, []int{1}, true},
		{[]int{3, 4, 5, 1, 2}, []int{3, 1, 2}, false},
		{[]int{
			-2, -3, -3, -4, -4, -4, -4, -5, -3, null, -5, -5, -5, -5, null, -4, -6, -2, null, -4, -6, -6, -4, -6, -4, -6, -4, null, -3, -5, -5, null, null, null, null, -7, -5, -7, -7, -5, -3, null, -7, -5, null, -5, -7, -5, -5, null, null, null, null, -6, -4, -6, -6, -4, -4, -8, -6, -8, -8, -6, -4, -4, -2, null, null, null, -4, -6, -4, null, -6, -4, -4, -6, -6, -7, -5, -5, null, null, -7, null, null, -3, -3, -5, null, null, null, -7, null, null, -7, -9, -9, -7, null, -5, null, null, -3, null, -1, null, null, -5, -5, -5, -3, -7, -5, null, null, -5, null, null, -5, null, -5, -8, -6, null, -6, -6, null, null, null, -2, null, null, null, null, null, -8, null, -6, -6, -10, null, -10, null, -6, null, null, -4, null, null, null, 0, -6, null, -6, null, -4, null, null, null, -8, -6, -6, -4, -6, null, null, null, null, null, -9, -7, -7, -5, null, null, null, null, null, null, -7, -7, -5, -5, null, null, -11, null, null, null, null, null, -3, -3, null, null, null, null, -7, -5, null, null, null, null, null, -5, null, -5, null, -3, null, null, null, null, -8, -6, -6, -6, null, null, null, null, -8, -6, null, -4, null, null, -12, null, -4, -4, -2, -2, null, -6, null, null, null, -6, null, null, null, null, null, null, -5, null, -7, -5, null, null, -7, -7, null, null, null, null, null, null, null, null, -3, null, null, null, null, -1, null, null, null, null, null, null, -8, -6, null, null, -8, null, null, null, -2, null, null, null, null, null, -7, null, null, null, null, -3, null, null, -4,
		}, []int{
			-3, -4, -4, null, null, -3, null, -2, null, null, -3, -4,
		}, true},
	}

	for _, test := range tests {
		nums1, nums2, expected := test.nums1, test.nums2, test.expected
		root1, root2 := NewBinaryTree(nums1), NewBinaryTree(nums2)
		testname := "Is SubTree"
		t.Run(testname, func(t *testing.T) {
			actual := IsSubtree(root1, root2)
			if actual != expected {
				t.Errorf("Actual    : %s", root1)
				t.Errorf("Expected  : %s", root2)

			}
		})

	}
}

func TestCompareTree(t *testing.T) {
	tests := []struct {
		nums1    []int
		nums2    []int
		expected bool
	}{
		{[]int{3, 4, 5, 1, 2}, []int{3, 4, 5, 1, 2}, true},
		{[]int{3, 4, 5, 1, 2, null, null, null, null, 0}, []int{3, 4, 5, 1, 2, null, null, null, null, 10}, false},
		{[]int{1, 1}, []int{1, 1}, true},
	}

	for _, test := range tests {
		testname := "Compare Trees"
		nums1, nums2, expected := test.nums1, test.nums2, test.expected
		root1, root2 := NewBinaryTree(nums1), NewBinaryTree(nums2)
		actual1 := CompareTrees(root1, root2)
		actual2 := IsSameTree(root1, root2)
		t.Run(testname, func(t *testing.T) {
			if actual1 != actual2 || actual2 != expected {
				t.Errorf("Actual    : %s", root1)
				t.Errorf("Expected  : %s", root2)
			}
		})
	}
}
