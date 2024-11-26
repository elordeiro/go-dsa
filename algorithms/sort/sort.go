package sort

import "github.com/elordeiro/goext/constraints"

type Ord constraints.Ordered

func BubbleSort[Slice ~[]O, O Ord](s Slice) {
	n := len(s)
	for n > 1 {
		newN := 0
		for i := 1; i < n; i++ {
			if s[i-1] > s[i] {
				s[i], s[i-1] = s[i-1], s[i]
				newN = i
			}
		}
		n = newN
	}
}

func Quicksort2[Slice ~[]O, O Ord](s Slice) {
	quicksort2(s, 0, len(s)-1)
}

func quicksort2[Slice ~[]O, O Ord](s Slice, lo, hi int) {
	if hi < lo {
		return
	}
	pivot := s[hi]
	i := lo - 1
	for j := lo; j < hi; j++ {
		if s[j] < pivot {
			i++
			temp := s[j]
			s[j] = s[i]
			s[i] = temp
		}
	}
	i++
	temp := s[i]
	s[i] = s[hi]
	s[hi] = temp
	quicksort(s, lo, i-1)
	quicksort(s, i+1, hi)
}

func Quicksort[Slice ~[]O, O Ord](s Slice) {
	quicksort(s, 0, len(s)-1)
}

func quicksort[Slice ~[]O, O Ord](s Slice, lo, hi int) {
	if lo < hi {
		pivot := partition(s, lo, hi)
		quicksort(s, lo, pivot-1)
		quicksort(s, pivot+1, hi)
	}
}

func partition[Slice ~[]O, O Ord](s Slice, lo, hi int) int {
	mid := (lo + hi) / 2
	if s[mid] < s[lo] {
		s[lo], s[mid] = s[mid], s[lo]
	}
	if s[hi] < s[lo] {
		s[lo], s[hi] = s[hi], s[lo]
	}
	if s[mid] < s[hi] {
		s[mid], s[hi] = s[hi], s[mid]
	}

	pivot := s[hi]
	i := lo - 1

	for j := lo; j < hi; j++ {
		if s[j] <= pivot {
			i++
			s[i], s[j] = s[j], s[i]
		}
	}

	s[i+1], s[hi] = s[hi], s[i+1]
	return i + 1
}

func MergeSort[Slice ~[]O, O Ord](s *Slice) {
	*s = mergeSort(*s)
}

func mergeSort[Slice ~[]O, O Ord](s Slice) []O {
	if len(s) <= 1 {
		return s
	}
	mid := len(s) / 2
	leftHalf := s[:mid]
	rightHalf := s[mid:]

	sortedLeft := mergeSort(leftHalf)
	sortedRight := mergeSort(rightHalf)

	return merge(sortedLeft, sortedRight)
}

func merge[Slice ~[]O, O Ord](left, right Slice) []O {
	result := []O{}
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}
