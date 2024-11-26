package sort_test

import (
	"math/rand/v2"
	"testing"

	"github.com/elordeiro/goext/algorithms/sort"
)

var Trials = 8388608

// generateRandomSlice generates a slice of random integers of a given size
func generateRandomSlice(size int, done chan bool, slice *[]int) {
	r := rand.New(rand.NewPCG(1, 2))
	*slice = make([]int, size)
	for i := range *slice {
		(*slice)[i] = r.IntN(size)
	}
	done <- true
}

// func BenchmarkBubbleSort(b *testing.B) {
// 	size := Trials
// 	data := generateRandomSlice(size)
// 	b.ResetTimer()

// 	for i := 0; i < b.N; i++ {
// 		cpy := append([]int(nil), data...)
// 		algs.BubbleSort(cpy)
// 	}
// }

func BenchmarkQuickSort(b *testing.B) {
	size := Trials
	var slice1 []int
	var slice2 []int
	done := make(chan bool, 1)
	generateRandomSlice(size, done, &slice1)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		go generateRandomSlice(size, done, &slice2)
		sort.Quicksort(slice1)
		// algs.Quicksort(cpy)
		<-done
		slice1, slice2 = slice2, slice1
	}
}

func BenchmarkQuickSortExp(b *testing.B) {
	size := 512
	var slice1 []int
	var slice2 []int
	done := make(chan bool, 1)
	generateRandomSlice(size, done, &slice1)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for ; size < 10_000_000; i++ {
			size *= 2
			go generateRandomSlice(size, done, &slice2)
			sort.Quicksort(slice1)
			// algs.Quicksort(cpy)
			<-done
			slice1, slice2 = slice2, slice1
		}
	}
}

func BenchmarkQuickSort2(b *testing.B) {
	size := Trials
	var slice1 []int
	var slice2 []int
	done := make(chan bool, 1)
	generateRandomSlice(size, done, &slice1)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		go generateRandomSlice(size, done, &slice2)
		sort.Quicksort2(slice1)
		// algs.Quicksort2(cpy)
		<-done
		slice1, slice2 = slice2, slice1
	}
}

func BenchmarkQuickSort2Exp(b *testing.B) {
	size := 512
	var slice1 []int
	var slice2 []int
	done := make(chan bool, 1)
	generateRandomSlice(size, done, &slice1)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for ; size < 10_000_000; i++ {
			size *= 2
			go generateRandomSlice(size, done, &slice2)
			sort.Quicksort2(slice1)
			// algs.Quicksort(cpy)
			<-done
			slice1, slice2 = slice2, slice1
		}
	}
}

func BenchmarkMergeSort(b *testing.B) {
	size := Trials
	var slice1 []int
	var slice2 []int
	done := make(chan bool, 1)
	generateRandomSlice(size, done, &slice1)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		go generateRandomSlice(size, done, &slice2)
		sort.MergeSort(&slice1)
		// algs.MergeSort(&cpy)
		<-done
		slice1, slice2 = slice2, slice1
	}
}

func BenchmarkMergeSortExp(b *testing.B) {
	size := 512
	var slice1 []int
	var slice2 []int
	done := make(chan bool, 1)
	generateRandomSlice(size, done, &slice1)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for ; size < 10_000_000; i++ {
			size *= 2
			go generateRandomSlice(size, done, &slice2)
			sort.MergeSort(&slice1)
			// algs.Quicksort(cpy)
			<-done
			slice1, slice2 = slice2, slice1
		}
	}
}
