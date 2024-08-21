package iters_test

import (
	"fmt"
	"slices"

	itr "github.com/elordeiro/go/iters"
)

func ExampleSeq() {
	slice := []int{1, 2, 3, 4, 5}
	// seq.Seq2[K, E]
	seq2 := slices.All(slice)
	// seq.Seq[E]
	seq1 := itr.ToSeq(seq2)
	for v := range seq1 {
		fmt.Print(v, " ")
	}
	fmt.Println()
	// Output: 1 2 3 4 5
}

func ExampleSeq2() {
	slice := []int{1, 2, 3, 4, 5}
	// seq.Seq[E]
	seq1 := slices.Values(slice)
	// seq.Seq2[int, E]
	seq2 := itr.ToSeq2(seq1)
	for i, v := range seq2 {
		fmt.Printf("%d:%d ", i, v)
	}
	fmt.Println()
	// Output: 0:1 1:2 2:3 3:4 4:5
}

func ExampleRange() {
	for v := range itr.Range(5) {
		fmt.Print(v, " ")
	}
	fmt.Println()
	// Output: 0 1 2 3 4
}

func ExampleRange_startEnd() {
	for v := range itr.Range(2, 5) {
		fmt.Print(v, " ")
	}
	fmt.Println()
	// Output: 2 3 4
}

func ExampleRange_startEndStep() {
	for v := range itr.Range(2, 10, 2) {
		fmt.Print(v, " ")
	}
	fmt.Println()
	// Output: 2 4 6 8
}

func ExampleRange_reverse() {
	for v := range itr.Range(5, 0) {
		fmt.Print(v, " ")
	}
	fmt.Println()
	// Output: 5 4 3 2 1
}

func ExampleCount() {
	max := 5
	for v := range itr.Count(0) {
		if v >= max {
			break
		}
		fmt.Print(v, " ")
	}
	fmt.Println()
	// Output: 0 1 2 3 4
}

func ExampleEnumerate() {
	slice := []int{1, 2, 3, 4, 5}
	seq := slices.Values(slice)
	for i, v := range itr.Enumerate(1, seq) {
		fmt.Printf("%d:%d ", i, v)
	}
	fmt.Println()
	// Output: 1:1 2:2 3:3 4:4 5:5
}

func ExampleZip() {
	slice1 := []int{1, 2, 3}
	slice2 := []string{"a", "b", "c"}
	seq1 := slices.Values(slice1)
	seq2 := slices.Values(slice2)
	for v1, v2 := range itr.Zip(seq1, seq2) {
		fmt.Printf("%d:%s ", v1, v2)
	}
	fmt.Println()
	// Output: 1:a 2:b 3:c
}

func ExampleFilter() {
	isEven := func(v int) bool { return v%2 == 0 }
	for v := range itr.Filter(itr.Range(5), isEven) {
		fmt.Print(v, " ")
	}
	fmt.Println()
	// Output: 0 2 4
}

func ExampleFilter_inline() {
	for v := range itr.Filter(itr.Range(5), func(v int) bool { return v%2 == 0 }) {
		fmt.Print(v, " ")
	}
	fmt.Println()
	// Output: 0 2 4
}

func ExampleMap() {
	double := func(v int) int { return v * 2 }
	for v := range itr.Map(itr.Range(5), double) {
		fmt.Print(v, " ")
	}
	fmt.Println()
	// Output: 0 2 4 6 8
}

func ExampleMap_inline() {
	for v := range itr.Map(itr.Range(5), func(v int) int { return v * v }) {
		fmt.Print(v, " ")
	}
	fmt.Println()
	// Output: 0 1 4 9 16
}

func ExampleReduce() {
	sum := func(acc, v int) int { return acc + v }
	result := itr.Reduce(itr.Range(5), sum)
	fmt.Println(result)
	// Output: 10
}

func ExampleCycle() {
	max := 5
	for i, v := range itr.ToSeq2(itr.Cycle(itr.Range(3))) {
		if i >= max {
			break
		}
		fmt.Print(v, " ")
	}
	fmt.Println()
	// Output: 0 1 2 0 1
}

func ExampleRepeat() {
	max := 5
	for i, v := range itr.ToSeq2(itr.Repeat(3)) {
		if i >= max {
			break
		}
		fmt.Print(v, " ")
	}
	fmt.Println()
	// Output: 3 3 3 3 3
}

func ExampleRepeat_end() {
	for v := range itr.Repeat(3, 5) {
		fmt.Print(v, " ")
	}
	fmt.Println()
	// Output: 3 3 3 3 3
}

func ExampleChain() {
	for v := range itr.Chain(itr.Range(2), itr.Range(2, 5), itr.Range(5, 10, 2)) {
		fmt.Print(v, " ")
	}
	fmt.Println()
	// Output: 0 1 2 3 4 5 7 9
}

func ExampleTake() {
	for v := range itr.Take(9, itr.Cycle(itr.Range(0, 3))) {
		fmt.Print(v, " ")
	}
	fmt.Println()
	// Output: 0 1 2 0 1 2 0 1 2
}

func ExampleTake_less() {
	for v := range itr.Take(2, itr.Cycle(itr.Range(0, 3))) {
		fmt.Print(v, " ")
	}
	fmt.Println()
	// Output: 0 1
}

func ExampleDrop() {
	for v := range itr.Drop(3, itr.Range(10)) {
		fmt.Print(v, " ")
	}
	fmt.Println()
	// Output: 3 4 5 6 7 8 9
}

func ExampleDrop_less() {
	for v := range itr.Drop(10, itr.Range(3)) {
		fmt.Print(v, " ")
	}
	fmt.Println()
	// Output:
}

func ExampleTakeWhile() {
	isLessThan5 := func(v int) bool { return v < 5 }
	for v := range itr.TakeWhile(itr.Range(10), isLessThan5) {
		fmt.Print(v, " ")
	}
	fmt.Println()
	// Output: 0 1 2 3 4
}

func ExampleDropWhile() {
	isLessThan5 := func(v int) bool { return v < 5 }
	for v := range itr.DropWhile(itr.Range(10), isLessThan5) {
		fmt.Print(v, " ")
	}
	fmt.Println()
	// Output: 5 6 7 8 9
}

func ExampleWith() {
	slice := []int{1, 2, 3, 4, 5}
	process := func() { slice = slice[:len(slice)-1] }
	for range itr.With(itr.Range(5), process) {
		// loop body
	}
	fmt.Println(slice)
	// Output: []
}

func ExampleElse() {
	slice := []int{1, 2, 3, 4, 5}
	Else := func() { slice = []int{100} }
	for i, v := range itr.ToSeq2(itr.Else(slices.Values(slice), Else)) {
		if v == 5 {
			break
		}
		slice[i] = v * i
	}
	fmt.Println(slice)
	// Output: [0 2 6 12 5]
}

func ExampleElse_true() {
	slice := []int{1, 2, 3, 4, 5}
	Else := func() { slice = []int{100} }
	for i, v := range itr.ToSeq2(itr.Else(slices.Values(slice), Else)) {
		if v == 10 {
			break
		}
		slice[i] = v * i
	}
	fmt.Println(slice)
	// Output: [100]
}
