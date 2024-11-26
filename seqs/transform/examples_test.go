package transform_test

import (
	"fmt"
	"slices"

	"github.com/elordeiro/goext/containers/tuples"
	"github.com/elordeiro/goext/seq2s/transform2"
	"github.com/elordeiro/goext/seqs"
	"github.com/elordeiro/goext/seqs/transform"
)

func ExampleBackwards() {
	for v := range transform.Backwards(seqs.Range(5)) {
		fmt.Print(v, " ")
	}
	fmt.Println()
	// Output: 4 3 2 1 0
}

func ExampleDropWhile() {
	isLessThan5 := func(v int) bool { return v < 5 }
	for v := range transform.DropWhile(seqs.Range(10), isLessThan5) {
		fmt.Print(v, " ")
	}
	fmt.Println()
	// Output: 5 6 7 8 9
}

func ExampleFilter() {
	isEven := func(v int) bool { return v%2 == 0 }
	for v := range transform.Filter(seqs.Range(5), isEven) {
		fmt.Print(v, " ")
	}
	fmt.Println()
	// Output: 0 2 4
}

func ExampleFilter_inline() {
	for v := range transform.Filter(seqs.Range(5), func(v int) bool { return v%2 == 0 }) {
		fmt.Print(v, " ")
	}
	fmt.Println()
	// Output: 0 2 4
}

func ExampleForEach() {
	seq := slices.Values([]int{1, 2, 3, 4, 5})
	transform.ForEach(seq, func(v int) {
		fmt.Print(2*v, " ")
	})
	fmt.Println()
	// Output: 2 4 6 8 10
}

func ExampleMap() {
	double := func(v int) int { return v * 2 }
	for v := range transform.Map(seqs.Range(5), double) {
		fmt.Print(v, " ")
	}
	fmt.Println()
	// Output: 0 2 4 6 8
}

func ExampleMap_inline() {
	for v := range transform.Map(seqs.Range(5), func(v int) int { return v * v }) {
		fmt.Print(v, " ")
	}
	fmt.Println()
	// Output: 0 1 4 9 16
}

func ExampleOnEmpty() {
	result := ""
	i := 31
	Else := func() { result = fmt.Sprint(i, " is prime") }
	for range transform.OnEmpty(seqs.Range(2, i), Else) {
		if i%2 == 0 {
			result = fmt.Sprint(i, " is not prime")
			break
		}
	}
	fmt.Println(result)
	// Output: 31 is prime
}

func ExampleOnEmpty_true() {
	slice := []int{1, 2, 3, 4, 5}
	Else := func() { slice = []int{100} }
	for i, v := range seqs.Enumerate(0, transform.OnEmpty(slices.Values(slice), Else)) {
		if v == 10 {
			break
		}
		slice[i] = v * i
	}
	fmt.Println(slice)
	// Output: [100]
}

func ExampleReduce() {
	sum := func(acc, v int) int { return acc + v }
	result := transform.Reduce(seqs.Range(5), sum)
	fmt.Println(result)
	// Output: 10
}

func ExampleReduce_start() {
	sum := func(acc, v int) int { return acc + v }
	result := transform.Reduce(seqs.Range(5), sum, 10)
	fmt.Println(result)
	// Output: 20
}

func ExampleRotate() {
	for v := range transform.Rotate(2, seqs.Range(5)) {
		fmt.Print(v, " ")
	}
	fmt.Println()
	// Output: 2 3 4 0 1
}

func ExampleTakeWhile() {
	isLessThan5 := func(v int) bool { return v < 5 }
	for v := range transform.TakeWhile(seqs.Range(10), isLessThan5) {
		fmt.Print(v, " ")
	}
	fmt.Println()
	// Output: 0 1 2 3 4
}

func ExampleUnpair() {
	seq1 := slices.Values([]int{1, 2, 3, 4, 5})
	seq2 := slices.Values([]string{"a", "b", "c", "d", "e"})
	seq := seqs.Zip(seq1, seq2)

	var pairSlice []tuples.Pair[int, string]
	transform2.ForEach(seq, func(i int, a string) {
		pairSlice = append(pairSlice, tuples.NewPair(i, a))
	})

	for k, v := range transform.Unpair(slices.Values(pairSlice)) {
		fmt.Print(k, ":", v, " ")
	}
	fmt.Println()
	// Output: 1:a 2:b 3:c 4:d 5:e
}

func ExampleWith() {
	seq := slices.Values([]int{1, 2, 3, 4, 5})
	result := 0
	process := func(v int) {
		if v%2 != 0 {
			result += v
		}
	}
	for range transform.With(seq, process) {
	}
	fmt.Println(result)
	// Output: 9
}
