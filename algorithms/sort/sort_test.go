package sort_test

import (
	"math/rand/v2"
	"slices"
	"testing"

	"github.com/elordeiro/goext/algorithms/sort"
)

func TestBubbleSort(t *testing.T) {
	r := rand.New(rand.NewPCG(1, 2))
	slice := []int{}
	for range 100 {
		slice = slice[:0]
		for range 20 {
			slice = append(slice, r.IntN(100))
		}
		sort.BubbleSort(slice)
		if !slices.IsSorted(slice) {
			t.Errorf("BubbleSort Err: got: %v", slice)
		}
	}
}

func TestQuickSort(t *testing.T) {
	r := rand.New(rand.NewPCG(1, 2))
	slice := []int{}
	for range 100 {
		slice = slice[:0]
		for range 20 {
			slice = append(slice, r.IntN(100))
		}
		sort.Quicksort(slice)
		if !slices.IsSorted(slice) {
			t.Errorf("QuickSort Err: got: %v", slice)
		}
	}
}

func TestQuickSort2(t *testing.T) {
	r := rand.New(rand.NewPCG(1, 2))
	slice := []int{}
	for range 100 {
		slice = slice[:0]
		for range 20 {
			slice = append(slice, r.IntN(100))
		}
		sort.Quicksort2(slice)
		if !slices.IsSorted(slice) {
			t.Errorf("QuickSort2 Err: got: %v", slice)
		}
	}
}

func TestMergeSort(t *testing.T) {
	r := rand.New(rand.NewPCG(1, 2))
	slice := []int{}
	for range 100 {
		slice = slice[:0]
		for range 20 {
			slice = append(slice, r.IntN(100))
		}
		sort.MergeSort(&slice)
		if !slices.IsSorted(slice) {
			t.Errorf("MergeSort Err: got: %v", slice)
		}
	}
}
