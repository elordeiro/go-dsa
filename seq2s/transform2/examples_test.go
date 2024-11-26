package transform2_test

import (
	"fmt"
	"slices"

	"github.com/elordeiro/goext/containers/tuples"
	"github.com/elordeiro/goext/seq2s"
	"github.com/elordeiro/goext/seq2s/transform2"
	"github.com/elordeiro/goext/seqs"
	"github.com/elordeiro/goext/seqs/transform"
)

func ExampleBackwards() {
	seq2 := seqs.Enumerate(0, seqs.Range(5))
	for k, v := range transform2.Backwards(seq2) {
		fmt.Printf("%d:%d ", k, v)
	}
	fmt.Println()
	// Output: 4:4 3:3 2:2 1:1 0:0
}

func ExampleDropWhile() {
	isLessThan5 := func(k, v int) bool { return v < 5 }
	seq2 := seqs.Enumerate(0, seqs.Range(10))
	for k, v := range transform2.DropWhile(seq2, isLessThan5) {
		fmt.Printf("%d:%d ", k, v)
	}
	fmt.Println()
	// Output: 5:5 6:6 7:7 8:8 9:9
}

func ExampleFilter() {
	isEven := func(k, v int) bool { return v%2 == 0 }
	seq2 := seqs.Enumerate(0, seqs.Range(5))
	for k, v := range transform2.Filter(seq2, isEven) {
		fmt.Printf("%d:%d ", k, v)
	}
	fmt.Println()
	// Output: 0:0 2:2 4:4
}

func ExampleFilter_inline() {
	seq2 := seqs.Enumerate(0, seqs.Range(5))
	for k, v := range transform2.Filter(seq2, func(k, v int) bool { return v%2 == 0 }) {
		fmt.Printf("%d:%d ", k, v)
	}
	fmt.Println()
	// Output: 0:0 2:2 4:4
}

func ExampleForEach() {
	seq2 := transform.Unpair(
		slices.Values(
			tuples.Pairs(
				tuples.NewPair(1, 2),
				tuples.NewPair(3, 4),
				tuples.NewPair(5, 6),
			),
		),
	)
	transform2.ForEach(seq2, func(k, v int) {
		fmt.Printf("%d:%d ", k, v*2)
	})
	fmt.Println()
	// Output: 1:4 3:8 5:12
}

func ExampleMap() {
	double := func(k, v int) (int, int) { return k, v * 2 }
	seq2 := seqs.Enumerate(0, seqs.Range(5))
	for k, v := range transform2.Map(seq2, double) {
		fmt.Printf("%d:%d ", k, v)
	}
	fmt.Println()
	// Output: 0:0 1:2 2:4 3:6 4:8
}

func ExampleMap_inline() {
	seq2 := seqs.Enumerate(0, seqs.Range(5))
	for k, v := range transform2.Map(seq2, func(k, v int) (int, int) { return k, v * v }) {
		fmt.Printf("%d:%d ", k, v)
	}
	fmt.Println()
	// Output: 0:0 1:1 2:4 3:9 4:16
}

func ExampleOnEmpty() {
	slice := tuples.Pairs(
		tuples.NewPair(1, 2),
		tuples.NewPair(3, 4),
		tuples.NewPair(5, 6),
	)
	seq := transform.Unpair(slices.Values(slice))
	Else := func() { slice = tuples.Pairs(tuples.NewPair(100, 200)) }
	for i, p := range seq2s.Enumerate(0, transform2.OnEmpty(seq, Else)) {
		if p.Left() == 5 {
			break
		}
		slice[i] = tuples.NewPair(p.Left(), p.Right()*i)
	}
	fmt.Println(slice)
	// Output: [(1 0) (3 4) (5 6)]
}

func ExampleOnEmpty_true() {
	slice := tuples.Pairs(
		tuples.NewPair(1, 2),
		tuples.NewPair(3, 4),
		tuples.NewPair(5, 6),
	)

	seq := transform.Unpair(slices.Values(slice))
	Else := func() { slice = tuples.Pairs(tuples.NewPair(100, 200)) }
	for i, p := range seq2s.Enumerate(0, transform2.OnEmpty(seq, Else)) {
		if p.Right() == 10 {
			break
		}
		slice[i] = tuples.NewPair(p.Left(), p.Right()*i)
	}
	fmt.Println(slice)
	// Output: [(100 200)]
}

func ExampleReduce() {
	sum := func(acc, v int) int { return acc + v }
	seq2 := seqs.Enumerate(0, seqs.Range(5))
	result := transform2.Reduce(seq2, sum)
	fmt.Println(result)
	// Output: 10
}

func ExampleReduce_start() {
	sum := func(acc, v int) int { return acc + v }
	seq2 := seqs.Enumerate(0, seqs.Range(5))
	result := transform2.Reduce(seq2, sum, 10)
	fmt.Println(result)
	// Output: 20
}

func ExampleRotate() {
	for k, v := range transform2.Rotate(2, seqs.Enumerate(0, seqs.Range(5))) {
		fmt.Printf("%d:%d ", k, v)
	}
	fmt.Println()
	// Output: 2:2 3:3 4:4 0:0 1:1
}

func ExampleSwapKV() {
	seq1 := slices.Values([]int{1, 2, 3, 4, 5})
	seq2 := slices.Values([]string{"a", "b", "c", "d", "e"})
	seq := seqs.Zip(seq1, seq2)
	for k, v := range transform2.SwapKV(seq) {
		fmt.Print(k, ":", v, " ")
	}
	fmt.Println()
	// Output: a:1 b:2 c:3 d:4 e:5
}

func ExampleTakeWhile() {
	isLessThan5 := func(k, v int) bool { return v < 5 }
	seq2 := seqs.Enumerate(0, seqs.Range(10))
	for k, v := range transform2.TakeWhile(seq2, isLessThan5) {
		fmt.Printf("%d:%d ", k, v)
	}
	fmt.Println()
	// Output: 0:0 1:1 2:2 3:3 4:4
}

func ExampleWith() {
	seq := slices.Values(tuples.Pairs(
		tuples.NewPair("alice", 1),
		tuples.NewPair("bob", 2),
		tuples.NewPair("charlie", 3),
		tuples.NewPair("david", 4),
		tuples.NewPair("eve", 5)))
	seq2 := transform.Unpair(seq)
	result := ""
	process := func(k string, v int) {
		if v > 3 {
			result += k + " "
		}
	}
	for range transform2.With(seq2, process) {
	}
	fmt.Println(result)
	// Output: david eve
}
