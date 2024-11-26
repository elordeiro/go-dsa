package seq2s_test

import (
	"fmt"
	"slices"

	"github.com/elordeiro/goext/containers/tuples"
	"github.com/elordeiro/goext/seq2s"
	"github.com/elordeiro/goext/seqs"
	"github.com/elordeiro/goext/seqs/transform"
)

func ExampleEqual() {
	seq1 := seqs.Enumerate(0, slices.Values([]int{1, 2, 3}))
	seq2 := seqs.Enumerate(0, slices.Values([]int{1, 2, 3}))
	fmt.Println(seq2s.Equal(seq1, seq2))
	// Output: true
}

func ExampleEqualFunc() {
	seq1 := seqs.Enumerate(0, slices.Values([]int{1, 2, 3}))
	seq2 := seqs.Enumerate(0, slices.Values([]int{1, 2, 3}))
	fmt.Println(seq2s.EqualFunc(seq1, seq2, func(a, b tuples.Pair[int, int]) bool { return a == b }))
	// Output: true
}

func ExampleEqualUnordered() {
	pairs1 := tuples.Pairs(tuples.NewPair(1, 2), tuples.NewPair(2, 3), tuples.NewPair(3, 4))
	paris2 := tuples.Pairs(tuples.NewPair(3, 4), tuples.NewPair(2, 3), tuples.NewPair(1, 2))
	seq1 := transform.Unpair(slices.Values(pairs1))
	seq2 := transform.Unpair(slices.Values(paris2))
	fmt.Println(seq2s.EqualUnordered(seq1, seq2))
	// Output: true
}

func ExampleFromSlice() {
	want := seqs.Enumerate(0, seqs.Range(10, 60, 10))
	slice := [][2]int{{0, 10}, {1, 20}, {2, 30}, {3, 40}, {4, 50}}
	got := seq2s.FromSlice(slice)
	fmt.Println(seq2s.Equal(got, want))
	// Output: true
}

func ExampleKeys() {
	slice := tuples.Pairs(
		tuples.NewPair("a", 1),
		tuples.NewPair("b", 2),
		tuples.NewPair("c", 3),
	)
	seq2 := transform.Unpair(slices.Values(slice))
	seq := seq2s.Keys(seq2)
	fmt.Println(seqs.String(seq))
	// Output: =>[a b c]
}

func ExampleLen() {
	seq2 := seqs.Enumerate(0,
		slices.Values(
			[]int{1, 2, 3},
		),
	)
	fmt.Println(seq2s.Len(seq2))
	// Output: 3
}

func ExampleString() {
	seq := slices.Values(
		tuples.Pairs(
			tuples.NewPair(1, 2),
			tuples.NewPair(3, 4),
			tuples.NewPair(5, 6),
		),
	)
	fmt.Println(seq2s.String(transform.Unpair(seq)))
	// Output: =>[(1 2) (3 4) (5 6)]
}

func ExampleValues() {
	slice := tuples.Pairs(
		tuples.NewPair("a", 1),
		tuples.NewPair("b", 2),
		tuples.NewPair("c", 3),
	)
	seq2 := transform.Unpair(slices.Values(slice))
	seq := seq2s.Values(seq2)
	fmt.Println(seqs.String(seq))
	// Output: =>[1 2 3]
}

func ExampleChain() {
	seq1 := transform.Unpair(
		slices.Values(
			tuples.Pairs(
				tuples.NewPair(1, 2),
				tuples.NewPair(3, 4),
			),
		),
	)
	seq2 := transform.Unpair(
		slices.Values(
			tuples.Pairs(
				tuples.NewPair(5, 6),
				tuples.NewPair(7, 8),
			),
		),
	)
	for k, v := range seq2s.Chain(seq1, seq2) {
		fmt.Printf("%d:%d ", k, v)
	}
	fmt.Println()
	// Output: 1:2 3:4 5:6 7:8
}

func ExampleCollect() {
	seq2 := seqs.Enumerate(0, seqs.Range(5))
	fmt.Println(seq2s.Collect(seq2))
	// Output: [[0 0] [1 1] [2 2] [3 3] [4 4]]
}

func ExampleCollectPairs() {
	seq2 := seqs.Enumerate(0, seqs.Range(5))
	fmt.Println(seq2s.CollectPairs(seq2))
	// Output: [(0 0) (1 1) (2 2) (3 3) (4 4)]
}

func ExampleCycle() {
	max := 5
	i := 0
	for k, v := range seq2s.Cycle(seqs.Enumerate(0, seqs.Range(3))) {
		if i >= max {
			break
		}
		fmt.Printf("%d:%d ", k, v)
		i++
	}
	fmt.Println()
	// Output: 0:0 1:1 2:2 0:0 1:1
}

func ExampleEnumerate() {
	seq2 := transform.Unpair(
		slices.Values(
			tuples.Pairs(
				tuples.NewPair("x", "r"),
				tuples.NewPair("y", "g"),
				tuples.NewPair("z", "b"),
			),
		),
	)
	for i, v := range seq2s.Enumerate(1, seq2) {
		fmt.Printf("%d:%s ", i, v)
	}
	fmt.Println()
	// Output: 1:(x r) 2:(y g) 3:(z b)
}

func ExampleRepeat() {
	for k, v := range seq2s.Repeat(3, 4, 5) {
		fmt.Print(k, ":", v, " ")
	}
	fmt.Println()
	// Output: 3:4 3:4 3:4 3:4 3:4
}

func ExampleSeqRange() {
	seq2 := seqs.Enumerate(0, slices.Values([]int{0, 1, 2, 3, 4, 5}))
	fmt.Println(seq2s.String(seq2s.SeqRange(0, 3, seq2)))
	// Output: =>[(0 0) (1 1) (2 2)]
}

func ExampleZip() {
	seq1 := transform.Unpair(
		slices.Values(
			tuples.Pairs(
				tuples.NewPair(1, 2),
				tuples.NewPair(3, 4),
				tuples.NewPair(5, 6),
			),
		),
	)
	seq2 := transform.Unpair(
		slices.Values(
			tuples.Pairs(
				tuples.NewPair("a", "b"),
				tuples.NewPair("c", "d"),
				tuples.NewPair("e", "f"),
			),
		),
	)

	for v1, v2 := range seq2s.Zip(seq1, seq2) {
		fmt.Printf("{%d %d}:{%s %s} ", v1.Left(), v1.Right(), v2.Left(), v2.Right())
	}
	fmt.Println()
	// Output: {1 2}:{a b} {3 4}:{c d} {5 6}:{e f}
}
