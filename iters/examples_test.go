package iters_test

import (
	"fmt"
	"strings"

	"slices"

	it "github.com/elordeiro/go/iters"
)

func ExampleIterable() {
	slice := it.Iterable([]int{1, 2, 3, 4, 5})
	fmt.Println(slice)
	// Output: ItSlice[int][1 2 3 4 5]
}

func ExampleIterable_custom() {
	type IntSlice []int
	slice := it.Iterable(IntSlice{1, 2, 3, 4, 5})
	fmt.Println(slice)
	// Output: ItSlice[int][1 2 3 4 5]
}

func ExampleIterable2() {
	m := it.Iterable2(map[string]int{"a": 1, "b": 2, "c": 3})
	for k, v := range m {
		fmt.Printf("%s:%d\n", k, v)
	}
	// Unordered Output: a:1
	// b:2
	// c:3
}

func ExampleIterable2_custom() {
	type StringIntMap map[string]int
	m := it.Iterable2(StringIntMap{"a": 1, "b": 2, "c": 3})
	for k, v := range m {
		fmt.Printf("%s:%d\n", k, v)
	}
	fmt.Println()
	// Unordered Output: a:1
	// b:2
	// c:3
}

func ExampleSeq_Len() {
	seq := it.Range(5)
	fmt.Println(seq.Len())
	// Output: 5
}

func ExampleSeq_Values() {
	slice := []int{1, 2, 3, 4, 5}
	iterSeq2 := slices.All(slice)
	seq := it.Surf2(iterSeq2).Values()
	fmt.Println(seq)
	// Output: Seq[int][1 2 3 4 5]
}

func ExampleSeq_All() {
	slice := []int{1, 2, 3, 4, 5}
	iterSeq2 := slices.All(slice)
	seq2 := it.Surf2(iterSeq2)
	fmt.Println(seq2)
	// Output: Seq2[int,int][0:1 1:2 2:3 3:4 4:5]
}

func ExampleSeq2_Len() {
	seq2 := it.Range(5).All()
	fmt.Println(seq2.Len())
	// Output: 5
}

func ExampleSeq2_Keys() {
	slice := it.Iterable([]it.Pair[string, int]{{"a", 1}, {"b", 2}, {"c", 3}})
	seq2 := it.Split(slice)
	seq := seq2.Keys()
	fmt.Println(seq)
	// Output: Seq[string][a b c]
}

func ExampleSeq2_Values() {
	slice := it.Iterable([]it.Pair[string, int]{{"a", 1}, {"b", 2}, {"c", 3}})
	seq2 := it.Split(slice)
	seq := seq2.Values()
	fmt.Println(seq)
	// Output: Seq[int][1 2 3]
}

func ExampleSeq2_All() {
	slice := it.Iterable([]it.Pair[string, int]{{"a", 1}, {"b", 2}, {"c", 3}})
	seq2 := it.Split(slice).All()
	fmt.Println(seq2)
	// Output: Seq2[string,int][a:1 b:2 c:3]
}

func ExampleItSlice_Len() {
	slice := it.Iterable([]int{1, 2, 3, 4, 5})
	fmt.Println(slice.Len())
	// Output: 5
}

func ExampleItSlice_Values() {
	slice := it.Iterable([]int{1, 2, 3, 4, 5})
	fmt.Println(slice.Values())
	// Output: Seq[int][1 2 3 4 5]
}

func ExampleItSlice_All() {
	slice := it.Iterable([]int{1, 2, 3, 4, 5})
	fmt.Println(slice.All())
	// Output: Seq2[int,int][0:1 1:2 2:3 3:4 4:5]
}

func ExampleItMap_Len() {
	m := it.Iterable2(map[string]int{"a": 1, "b": 2, "c": 3})
	fmt.Println(m.Len())
	// Output: 3
}

func ExampleItMap_Keys() {
	m := it.Iterable2(map[string]int{"a": 1, "b": 2, "c": 3})
	slice := slices.Collect(m.Keys().Sink())
	// map order is not guaranteed
	slices.Sort(slice)
	fmt.Println(slice)
	// Output: [a b c]
}

func ExampleItMap_Values() {
	m := it.Iterable2(map[string]int{"a": 1, "b": 2, "c": 3})
	slice := slices.Collect(m.Values().Sink())
	// map order is not guaranteed
	slices.Sort(slice)
	fmt.Println(slice)
	// Output: [1 2 3]
}

func ExampleItMap_All() {
	m := it.Iterable2(map[string]int{"a": 1, "b": 2, "c": 3})
	all := m.All()
	keys := slices.Collect(all.Keys().Sink())
	// map order is not guaranteed
	slices.Sort(keys)
	values := slices.Collect(all.Values().Sink())
	// map order is not guaranteed
	slices.Sort(values)
	fmt.Println(keys, values)
	// Output: [a b c] [1 2 3]
}

func ExampleSurf() {
	slice := []int{1, 2, 3, 4, 5}
	iterSeq := slices.Values(slice)
	seq := it.Surf(iterSeq)
	fmt.Println(seq)
	// Output: Seq[int][1 2 3 4 5]
}

func ExampleSurf2() {
	slice := []int{1, 2, 3, 4, 5}
	iterSeq2 := slices.All(slice)
	seq2 := it.Surf2(iterSeq2)
	fmt.Println(seq2)
	// Output: Seq2[int,int][0:1 1:2 2:3 3:4 4:5]
}

func ExampleSeq_Sink() {
	slice := it.Iterable([]int{1, 2, 3, 4, 5})
	seq := slice.Values()
	iterSeq := seq.Sink()
	for v := range iterSeq {
		fmt.Print(v, " ")
	}
	fmt.Println()
	// Output: 1 2 3 4 5
}

func ExampleSeq2_Sink() {
	slice := it.Iterable([]it.Pair[int, int]{{1, 2}, {3, 4}, {5, 6}})
	seq2 := it.Split(slice)
	iterSeq2 := seq2.Sink()
	for k, v := range iterSeq2 {
		fmt.Printf("%d:%d ", k, v)
	}
	fmt.Println()
	// Output: 1:2 3:4 5:6
}

func ExampleRange() {
	for v := range it.Range(5) {
		fmt.Print(v, " ")
	}
	fmt.Println()
	// Output: 0 1 2 3 4
}

func ExampleRange_startEnd() {
	for v := range it.Range(2, 5) {
		fmt.Print(v, " ")
	}
	fmt.Println()
	// Output: 2 3 4
}

func ExampleRange_startEndStep() {
	for v := range it.Range(2, 10, 2) {
		fmt.Print(v, " ")
	}
	fmt.Println()
	// Output: 2 4 6 8
}

func ExampleRange_reverse() {
	for v := range it.Range(5, 0) {
		fmt.Print(v, " ")
	}
	fmt.Println()
	// Output: 5 4 3 2 1
}

func ExampleCount() {
	max := 5
	for v := range it.Count(0) {
		if v >= max {
			break
		}
		fmt.Print(v, " ")
	}
	fmt.Println()
	// Output: 0 1 2 3 4
}

func ExampleCountDown() {
	for v := range it.CountDown(5) {
		if v <= 0 {
			break
		}
		fmt.Print(v, " ")
	}
	fmt.Println()
	// Output: 5 4 3 2 1
}

func ExampleEnumerate() {
	slice := it.Iterable([]string{"a", "b", "c"})
	for i, v := range it.Enumerate(1, slice) {
		fmt.Printf("%d:%s ", i, v)
	}
	fmt.Println()
	// Output: 1:a 2:b 3:c
}

func ExampleEnumerate2() {
	seq2 := it.Split(it.Iterable([]it.Pair[string, string]{{"x", "r"}, {"y", "g"}, {"z", "b"}}))
	for i, v := range it.Enumerate2(1, seq2) {
		fmt.Printf("%d:%s ", i, v)
	}
	fmt.Println()
	// Output: 1:x:r 2:y:g 3:z:b
}

func ExampleZip() {
	seq1 := it.Iterable([]int{1, 2, 3})
	seq2 := it.Iterable([]string{"a", "b", "c"})
	for v1, v2 := range it.Zip(seq1, seq2) {
		fmt.Printf("%d:%s ", v1, v2)
	}
	fmt.Println()
	// Output: 1:a 2:b 3:c
}

func ExampleZip2() {
	seq1 := it.Split(it.Iterable([]it.Pair[int, int]{{1, 2}, {3, 4}, {5, 6}}))
	seq2 := it.Split(it.Iterable([]it.Pair[string, string]{{"a", "b"}, {"c", "d"}, {"e", "f"}}).Values())
	for v1, v2 := range it.Zip2(seq1, seq2) {
		fmt.Printf("{%d %d}:{%s %s} ", v1.Key, v1.Value, v2.Key, v2.Value)
	}
	fmt.Println()
	// Output: {1 2}:{a b} {3 4}:{c d} {5 6}:{e f}
}

func ExampleRepeat() {
	for v := range it.Repeat(3, 5) {
		fmt.Print(v, " ")
	}
	fmt.Println()
	// Output: 3 3 3 3 3
}

func ExampleRepeat2() {
	for k, v := range it.Repeat2(3, 4, 5) {
		fmt.Print(k, ":", v, " ")
	}
	fmt.Println()
	// Output: 3:4 3:4 3:4 3:4 3:4
}

func ExampleCycle() {
	max := 5
	i := 0
	for v := range it.Cycle(it.Range(3)) {
		if i >= max {
			break
		}
		fmt.Print(v, " ")
		i++
	}
	fmt.Println()
	// Output: 0 1 2 0 1
}

func ExampleCycle2() {
	max := 5
	i := 0
	for k, v := range it.Cycle2(it.Range(3).All()) {
		if i >= max {
			break
		}
		fmt.Printf("%d:%d ", k, v)
		i++
	}
	fmt.Println()
	// Output: 0:0 1:1 2:2 0:0 1:1
}

func ExampleChain() {
	for v := range it.Chain(it.Range(2), it.Range(2, 5), it.Range(5, 10, 2)) {
		fmt.Print(v, " ")
	}
	fmt.Println()
	// Output: 0 1 2 3 4 5 7 9
}

func ExampleChain_second() {
	slice1 := it.Iterable([]int{0, 1, 2, 3})
	slice2 := it.Iterable([]int{4, 5, 6, 7})
	for v := range it.Chain(slice1, slice2) {
		fmt.Print(v, " ")
	}
	fmt.Println()
	// Output: 0 1 2 3 4 5 6 7
}

func ExampleChain2() {
	seq1 := it.Split(it.Iterable([]it.Pair[int, int]{{1, 2}, {3, 4}}).Values())
	seq2 := it.Split(it.Iterable([]it.Pair[int, int]{{5, 6}, {7, 8}}).Values())
	for k, v := range it.Chain2(seq1, seq2) {
		fmt.Printf("%d:%d ", k, v)
	}
	fmt.Println()
	// Output: 1:2 3:4 5:6 7:8
}

func ExampleBackwards() {
	for v := range it.Backwards(it.Range(5)) {
		fmt.Print(v, " ")
	}
	fmt.Println()
	// Output: 4 3 2 1 0
}

func ExampleBackwards2() {
	seq2 := it.Range(5).All()
	for k, v := range it.Backwards2(seq2) {
		fmt.Printf("%d:%d ", k, v)
	}
	fmt.Println()
	// Output: 4:4 3:3 2:2 1:1 0:0
}

func ExampleTake() {
	for v := range it.Take(9, it.Cycle(it.Range(0, 3))) {
		fmt.Print(v, " ")
	}
	fmt.Println()
	// Output: 0 1 2 0 1 2 0 1 2
}

func ExampleTake_less() {
	for v := range it.Take(2, it.Cycle(it.Range(0, 3))) {
		fmt.Print(v, " ")
	}
	fmt.Println()
	// Output: 0 1
}

func ExampleTake2() {
	seq2 := it.Cycle2(it.Range(0, 3).All())
	for k, v := range it.Take2(9, seq2) {
		fmt.Printf("%d:%d ", k, v)
	}
	fmt.Println()
	// Output: 0:0 1:1 2:2 0:0 1:1 2:2 0:0 1:1 2:2
}

func ExampleTake2_less() {
	seq2 := it.Cycle2(it.Range(0, 3).All())
	for k, v := range it.Take2(2, seq2) {
		fmt.Printf("%d:%d ", k, v)
	}
	fmt.Println()
	// Output: 0:0 1:1
}

func ExampleDrop() {
	for v := range it.Drop(3, it.Range(10)) {
		fmt.Print(v, " ")
	}
	fmt.Println()
	// Output: 3 4 5 6 7 8 9
}

func ExampleDrop_less() {
	for v := range it.Drop(10, it.Range(3)) {
		fmt.Print(v, " ")
	}
	fmt.Println()
	// Output:
}

func ExampleDrop2() {
	seq2 := it.Range(10).All()
	for k, v := range it.Drop2(3, seq2) {
		fmt.Printf("%d:%d ", k, v)
	}
	fmt.Println()
	// Output: 3:3 4:4 5:5 6:6 7:7 8:8 9:9
}

func ExampleDrop2_less() {
	seq2 := it.Range(3).All()
	for k, v := range it.Drop2(10, seq2) {
		fmt.Printf("%d:%d ", k, v)
	}
	fmt.Println()
	// Output:
}

func ExampleRotate() {
	for v := range it.Rotate(2, it.Range(5)) {
		fmt.Print(v, " ")
	}
	fmt.Println()
	// Output: 2 3 4 0 1
}

func ExampleRotate_less() {
	for v := range it.Rotate(-2, it.Range(5)) {
		fmt.Print(v, " ")
	}
	fmt.Println()
	// Output: 3 4 0 1 2
}

func ExampleRotate2() {
	seq2 := it.Range(5).All()
	for k, v := range it.Rotate2(2, seq2) {
		fmt.Printf("%d:%d ", k, v)
	}
	fmt.Println()
	// Output: 2:2 3:3 4:4 0:0 1:1
}

func ExampleRotate2_less() {
	seq2 := it.Range(5).All()
	for k, v := range it.Rotate2(-2, seq2) {
		fmt.Printf("%d:%d ", k, v)
	}
	fmt.Println()
	// Output: 3:3 4:4 0:0 1:1 2:2
}

func ExampleForEach() {
	slice := it.Iterable([]int{1, 2, 3, 4, 5})
	it.ForEach(slice, func(v int) {
		fmt.Print(2*v, " ")
	})
	fmt.Println()
	// Output: 2 4 6 8 10
}

func ExampleForEach2() {
	seq2 := it.Split(it.Iterable([]it.Pair[int, int]{{1, 2}, {3, 4}, {5, 6}}))
	it.ForEach2(seq2, func(k, v int) {
		fmt.Printf("%d:%d ", k, v*2)
	})
	fmt.Println()
	// Output: 1:4 3:8 5:12
}

func ExampleFilter() {
	isEven := func(v int) bool { return v%2 == 0 }
	for v := range it.Filter(it.Range(5), isEven) {
		fmt.Print(v, " ")
	}
	fmt.Println()
	// Output: 0 2 4
}

func ExampleFilter_inline() {
	for v := range it.Filter(it.Range(5), func(v int) bool { return v%2 == 0 }) {
		fmt.Print(v, " ")
	}
	fmt.Println()
	// Output: 0 2 4
}

func ExampleFilter2() {
	isEven := func(k, v int) bool { return v%2 == 0 }
	seq2 := it.Range(5).All()
	for k, v := range it.Filter2(seq2, isEven) {
		fmt.Printf("%d:%d ", k, v)
	}
	fmt.Println()
	// Output: 0:0 2:2 4:4
}

func ExampleFilter2_inline() {
	seq2 := it.Range(5).All()
	for k, v := range it.Filter2(seq2, func(k, v int) bool { return v%2 == 0 }) {
		fmt.Printf("%d:%d ", k, v)
	}
	fmt.Println()
	// Output: 0:0 2:2 4:4
}

func ExampleMap() {
	double := func(v int) int { return v * 2 }
	for v := range it.Map(it.Range(5), double) {
		fmt.Print(v, " ")
	}
	fmt.Println()
	// Output: 0 2 4 6 8
}

func ExampleMap_inline() {
	for v := range it.Map(it.Range(5), func(v int) int { return v * v }) {
		fmt.Print(v, " ")
	}
	fmt.Println()
	// Output: 0 1 4 9 16
}

func ExampleMap2() {
	double := func(k, v int) (int, int) { return k, v * 2 }
	seq2 := it.Range(5).All()
	for k, v := range it.Map2(seq2, double) {
		fmt.Printf("%d:%d ", k, v)
	}
	fmt.Println()
	// Output: 0:0 1:2 2:4 3:6 4:8
}

func ExampleMap2_inline() {
	seq2 := it.Range(5).All()
	for k, v := range it.Map2(seq2, func(k, v int) (int, int) { return k, v * v }) {
		fmt.Printf("%d:%d ", k, v)
	}
	fmt.Println()
	// Output: 0:0 1:1 2:4 3:9 4:16
}

func ExampleReduce() {
	sum := func(acc, v int) int { return acc + v }
	result := it.Reduce(it.Range(5), sum)
	fmt.Println(result)
	// Output: 10
}

func ExampleReduce_start() {
	sum := func(acc, v int) int { return acc + v }
	result := it.Reduce(it.Range(5), sum, 10)
	fmt.Println(result)
	// Output: 20
}

func ExampleReduce2() {
	sum := func(acc, v int) int { return acc + v }
	seq2 := it.Range(5).All()
	result := it.Reduce2(seq2, sum)
	fmt.Println(result)
	// Output: 10
}

func ExampleReduce2_start() {
	sum := func(acc, v int) int { return acc + v }
	seq2 := it.Range(5).All()
	result := it.Reduce2(seq2, sum, 10)
	fmt.Println(result)
	// Output: 20
}

func ExampleTakeWhile() {
	isLessThan5 := func(v int) bool { return v < 5 }
	for v := range it.TakeWhile(it.Range(10), isLessThan5) {
		fmt.Print(v, " ")
	}
	fmt.Println()
	// Output: 0 1 2 3 4
}

func ExampleTakeWhile2() {
	isLessThan5 := func(k, v int) bool { return v < 5 }
	seq2 := it.Range(10).All()
	for k, v := range it.TakeWhile2(seq2, isLessThan5) {
		fmt.Printf("%d:%d ", k, v)
	}
	fmt.Println()
	// Output: 0:0 1:1 2:2 3:3 4:4
}

func ExampleDropWhile() {
	isLessThan5 := func(v int) bool { return v < 5 }
	for v := range it.DropWhile(it.Range(10), isLessThan5) {
		fmt.Print(v, " ")
	}
	fmt.Println()
	// Output: 5 6 7 8 9
}

func ExampleDropWhile2() {
	isLessThan5 := func(k, v int) bool { return v < 5 }
	seq2 := it.Range(10).All()
	for k, v := range it.DropWhile2(seq2, isLessThan5) {
		fmt.Printf("%d:%d ", k, v)
	}
	fmt.Println()
	// Output: 5:5 6:6 7:7 8:8 9:9
}

func ExampleWith() {
	slice := it.Iterable([]int{1, 2, 3, 4, 5})
	result := 0
	process := func(v int) {
		if v%2 != 0 {
			result += v
		}
	}
	for range it.With(slice, process) {
	}
	fmt.Println(result)
	// Output: 9
}

func ExampleWith2() {
	slice := it.Iterable([]it.Pair[string, int]{{"alice", 1}, {"bob", 2}, {"charlie", 3}, {"david", 4}, {"eve", 5}})
	seq2 := it.Split(slice)
	result := ""
	process := func(k string, v int) {
		if v > 3 {
			result += k + " "
		}
	}
	for range it.With2(seq2, process) {
	}
	fmt.Println(result)
	// Output: david eve
}

func ExampleOnEmpty() {
	result := ""
	i := 31
	Else := func() { result = fmt.Sprint(i, " is prime") }
	for range it.OnEmpty(it.Range(2, i), Else) {
		if i%2 == 0 {
			result = fmt.Sprint(i, " is not prime")
			break
		}
	}
	fmt.Println(result)
	// Output: 31 is prime
}

func ExampleOnEmpty_true() {
	slice := it.Iterable([]int{1, 2, 3, 4, 5})
	Else := func() { slice = []int{100} }
	for i, v := range it.Enumerate(0, it.OnEmpty(slice, Else)) {
		if v == 10 {
			break
		}
		slice[i] = v * i
	}
	fmt.Println(slice)
	// Output: ItSlice[int][100]
}

func ExampleOnEmpty2() {
	slice := it.Iterable([]it.Pair[int, int]{{1, 2}, {3, 4}, {5, 6}})
	seq2 := it.Split(slice)
	Else := func() { slice = []it.Pair[int, int]{{100, 200}} }
	for i, p := range it.Enumerate2(0, it.OnEmpty2(seq2, Else)) {
		if p.Key == 5 {
			break
		}
		slice[i] = it.Pair[int, int]{p.Key, p.Value * i}
	}
	fmt.Println(slice)
	// Output: ItSlice[iters.Pair[int,int]][1:0 3:4 5:6]
}

func ExampleOnEmpty2_true() {
	slice := it.Iterable([]it.Pair[int, int]{{1, 2}, {3, 4}, {5, 6}})
	seq2 := it.Split(slice)
	Else := func() { slice = []it.Pair[int, int]{{100, 200}} }
	for i, p := range it.Enumerate2(0, it.OnEmpty2(seq2, Else)) {
		if p.Value == 10 {
			break
		}
		slice[i] = it.Pair[int, int]{p.Key, p.Value * i}
	}
	fmt.Println(slice)
	// Output: ItSlice[iters.Pair[int,int]][100:200]
}

func ExampleSum() {
	slice := it.Iterable([]int{1, 2, 3, 4, 5})
	fmt.Println(it.Sum(slice.Values()))
	// Output: 15
}

func ExampleProduct() {
	slice := it.Iterable([]int{1, 2, 3, 4, 5})
	fmt.Println(it.Product(slice.Values()))
	// Output: 120
}

func ExampleMax() {
	slice := it.Iterable([]int{1, 2, 3, 4, 5})
	fmt.Println(it.Max(slice.Values()))
	// Output: 5
}

func ExampleMin() {
	slice := it.Iterable([]int{1, 2, 3, 4, 5})
	fmt.Println(it.Min(slice.Values()))
	// Output: 1
}

func ExampleAll() {
	slice := it.Iterable([]bool{true, true, true, true, true})
	fmt.Println(it.All(slice.Values()))
	// Output: true
}

func ExampleAny() {
	slice := it.Iterable([]bool{true, false, true, false, true})
	fmt.Println(it.Any(slice.Values()))
	// Output: true
}

func ExampleNone() {
	slice := it.Iterable([]bool{true, false, true, false, true})
	fmt.Println(it.None(slice.Values()))
	// Output: false
}

func ExampleSplit() {
	slice := it.Iterable([]it.Pair[int, string]{{1, "a"}, {2, "b"}, {3, "c"}})
	seq2 := it.Split(slice)
	for k, v := range seq2 {
		fmt.Printf("%d:%s ", k, v)
	}
	fmt.Println()
	// Output: 1:a 2:b 3:c
}

func ExampleSwapKV() {
	seq2 := it.Split(it.Iterable([]it.Pair[string, int]{{"a", 1}, {"b", 2}, {"c", 3}}).Values())
	for k, v := range it.SwapKV(seq2) {
		fmt.Printf("%d:%s ", k, v)
	}
	fmt.Println()
	// Output: 1:a 2:b 3:c
}

func ExampleSeq_String() {
	seq := it.Iterable([]int{1, 2, 3, 4, 5})
	fmt.Println(seq.Values())
	// Output: Seq[int][1 2 3 4 5]
}

func ExampleSeq2_String() {
	seq := it.Iterable([]it.Pair[int, int]{{1, 2}, {3, 4}, {5, 6}}).Values()
	fmt.Println(it.Split(seq))
	// Output: Seq2[int,int][1:2 3:4 5:6]
}

func ExampleItSlice_String() {
	slice := it.Iterable([]int{1, 2, 3, 4, 5})
	fmt.Println(slice)
	// Output: ItSlice[int][1 2 3 4 5]
}

func ExampleItMap_String() {
	m := it.Iterable2(map[string]int{"a": 1, "b": 2, "c": 3})
	mapStr := fmt.Sprint(m)
	typeInfoIdx := strings.Index(mapStr, "]") + 2
	typeInfo := mapStr[:typeInfoIdx]
	fmt.Println(typeInfo)
	parts := strings.Split(mapStr[typeInfoIdx:len(mapStr)-1], " ")
	for _, part := range parts {
		fmt.Println(part)
	}
	fmt.Println("]")
	// Unordered output: ItMap[string,int][
	// a:1
	// b:2
	// c:3
	// ]
}

func ExamplePair_String() {
	seq := it.Iterable([]it.Pair[int, int]{{1, 2}, {3, 4}, {5, 6}})
	fmt.Println(seq)
	// Output: ItSlice[iters.Pair[int,int]][1:2 3:4 5:6]
}
