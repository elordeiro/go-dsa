package search_test

import (
	"fmt"
	"slices"

	"github.com/elordeiro/goext/algorithms/search"
	"github.com/elordeiro/goext/containers/tuples"
	"github.com/elordeiro/goext/seq2s/transform2"
	"github.com/elordeiro/goext/seqs/transform"
)

func ExampleAll() {
	seq := slices.Values([]bool{true, true, true, true, true})
	fmt.Println(search.All(seq))
	// Output: true
}

func ExampleAny() {
	seq := slices.Values([]bool{true, false, true, false, true})
	fmt.Println(search.Any(seq))
	// Output: true
}

func ExampleNone() {
	seq := slices.Values([]bool{true, false, true, false, true})
	fmt.Println(search.None(seq))
	// Output: false
}

func ExampleUnpair() {
	seq := slices.Values(tuples.Pairs(
		tuples.NewPair(1, "a"),
		tuples.NewPair(2, "b"),
		tuples.NewPair(3, "c")))
	seq2 := transform.Unpair(seq)
	for k, v := range seq2 {
		fmt.Printf("%d:%s ", k, v)
	}
	fmt.Println()
	// Output: 1:a 2:b 3:c
}

func ExampleSwapKV() {
	seq2 := transform.Unpair(
		slices.Values(
			tuples.Pairs(
				tuples.NewPair("a", 1),
				tuples.NewPair("b", 2),
				tuples.NewPair("c", 3),
			),
		),
	)
	for k, v := range transform2.SwapKV(seq2) {
		fmt.Printf("%d:%s ", k, v)
	}
	fmt.Println()
	// Output: 1:a 2:b 3:c
}
