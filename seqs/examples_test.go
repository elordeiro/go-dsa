package seqs_test

import (
	"fmt"

	"slices"

	"github.com/elordeiro/goext/seqs"
)

func ExampleEqual() {
	seq1 := slices.Values([]int{1, 2, 3})
	seq2 := slices.Values([]int{1, 2, 3})
	fmt.Println(seqs.Equal(seq1, seq2))
	// Output: true
}

func ExampleEqualFunc() {
	seq1 := slices.Values([]int{1, 2, 3})
	seq2 := slices.Values([]int{1, 2, 3})
	fmt.Println(seqs.EqualFunc(seq1, seq2, func(a, b int) bool { return a == b }))
	// Output: true
}

func ExampleEqualUnordered() {
	seq1 := slices.Values([]int{1, 2, 3})
	seq2 := slices.Values([]int{3, 2, 1})
	fmt.Println(seqs.EqualUnordered(seq1, seq2))
	// Output: true
}

func ExampleLen() {
	seq := slices.Values([]int{1, 2, 3, 4, 5})
	fmt.Println(seqs.Len(seq))
	// Output: 5
}

func ExampleString() {
	seq := slices.Values([]int{1, 2, 3, 4, 5})
	fmt.Println(seqs.String(seq))
	// Output: =>[1 2 3 4 5]
}

func ExampleChain() {
	for v := range seqs.Chain(seqs.Range(2), seqs.Range(2, 5), seqs.Range(5, 10, 2)) {
		fmt.Print(v, " ")
	}
	fmt.Println()
	// Output: 0 1 2 3 4 5 7 9
}

func ExampleChain_second() {
	slice1 := slices.Values([]int{0, 1, 2, 3})
	slice2 := slices.Values([]int{4, 5, 6, 7})
	for v := range seqs.Chain(slice1, slice2) {
		fmt.Print(v, " ")
	}
	fmt.Println()
	// Output: 0 1 2 3 4 5 6 7
}

func ExampleCollect() {
	seq := seqs.Range(5)
	fmt.Println(seqs.Collect(seq))
	// Output: [0 1 2 3 4]
}

func ExampleCount() {
	max := 5
	for v := range seqs.Count(0) {
		if v >= max {
			break
		}
		fmt.Print(v, " ")
	}
	fmt.Println()
	// Output: 0 1 2 3 4
}

func ExampleCount_down() {
	for v := range seqs.Count(5, -1) {
		if v <= 0 {
			break
		}
		fmt.Print(v, " ")
	}
	fmt.Println()
	// Output: 5 4 3 2 1
}

func ExampleCycle() {
	max := 5
	i := 0
	for v := range seqs.Cycle(seqs.Range(3)) {
		if i >= max {
			break
		}
		fmt.Print(v, " ")
		i++
	}
	fmt.Println()
	// Output: 0 1 2 0 1
}

func ExampleEnumerate() {
	slice := slices.Values([]string{"a", "b", "c"})
	for i, v := range seqs.Enumerate(1, slice) {
		fmt.Printf("%d:%s ", i, v)
	}
	fmt.Println()
	// Output: 1:a 2:b 3:c
}

func ExampleRange() {
	for v := range seqs.Range(5) {
		fmt.Print(v, " ")
	}
	fmt.Println()
	// Output: 0 1 2 3 4
}

func ExampleRange_startEnd() {
	for v := range seqs.Range(2, 5) {
		fmt.Print(v, " ")
	}
	fmt.Println()
	// Output: 2 3 4
}

func ExampleRange_startEndStep() {
	for v := range seqs.Range(2, 10, 2) {
		fmt.Print(v, " ")
	}
	fmt.Println()
	// Output: 2 4 6 8
}

func ExampleRange_reverse() {
	for v := range seqs.Range(5, 0) {
		fmt.Print(v, " ")
	}
	fmt.Println()
	// Output: 5 4 3 2 1
}

func ExampleRepeat() {
	for v := range seqs.Repeat(3, 5) {
		fmt.Print(v, " ")
	}
	fmt.Println()
	// Output: 3 3 3 3 3
}

func ExampleSeqRange() {
	seq := slices.Values([]int{0, 1, 2, 3, 4, 5})
	fmt.Println(seqs.String(seqs.SeqRange(0, 3, seq)))
	// Output: =>[0 1 2]
}

func ExampleZip() {
	seq1 := slices.Values([]int{1, 2, 3})
	seq2 := slices.Values([]string{"a", "b", "c"})
	for v1, v2 := range seqs.Zip(seq1, seq2) {
		fmt.Printf("%d:%s ", v1, v2)
	}
	fmt.Println()
	// Output: 1:a 2:b 3:c
}
