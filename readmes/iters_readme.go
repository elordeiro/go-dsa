package main

import (
	"fmt"

	it "github.com/elordeiro/go/iters"
)

func main() {
	{
		fmt.Println("Basic iterators")

		for v := range it.Range(1, 6) {
			fmt.Print(v, " ")
		}
		fmt.Println()
		// Output: 1 2 3 4 5

		slice := it.Iterable([]bool{true, false, true, false, true})
		fmt.Println(it.Any(slice.Values()))
		// Output: true

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

	{

		fmt.Println("Chaining operations")

		square := func(x int) int { return x * x }
		for v := range it.Map(it.Range(6), square) {
			fmt.Print(v, " ")
		}
		fmt.Println()
		// Output: 1 4 9 16 25

		sum := func(a, b int) int { return a + b }
		result := it.Reduce(it.Range(1, 6), sum)
		fmt.Println(result)
		// Output: 15

		for v := range it.Take(9, it.Cycle(it.Range(0, 3))) {
			fmt.Print(v, " ")
		}
		fmt.Println()
		// Output: 0 1 2 0 1 2 0 1 2

		for v := range it.Chain(it.Range(2), it.Range(2, 5), it.Range(5, 10, 2)) {
			fmt.Print(v, " ")
		}
		fmt.Println()
		// Output: 0 1 2 3 4 5 7 9
	}

	{
		fmt.Println("Using other iterable types")

		type IntSlice []int
		slice := it.Iterable(IntSlice{1, 2, 3, 4, 5})
		fmt.Println(slice)
		// Output: ItSlice[int][1 2 3 4 5]

		slice2 := it.Iterable([]string{"a", "b", "c"})
		for i, v := range it.Enumerate(1, slice2) {
			fmt.Printf("%d:%s ", i, v)
		}
		fmt.Println()
		// Output: 1:a 2:b 3:c

		seq1 := it.Iterable([]int{1, 2, 3})
		seq2 := it.Iterable([]string{"a", "b", "c"})
		for v1, v2 := range it.Zip(seq1, seq2) {
			fmt.Printf("%d:%s ", v1, v2)
		}
		fmt.Println()
		// Output: 1:a 2:b 3:c

		slice3 := it.Iterable([]int{0, 1, 2, 3})
		slice4 := it.Iterable([]int{4, 5, 6, 7})
		for v := range it.Chain(slice3, slice4) {
			fmt.Print(v, " ")
		}
		fmt.Println()
		// Output: 0 1 2 3 4 5 6 7

		m := it.Iterable2(map[string]int{"a": 1, "b": 2, "c": 3})
		isEven := func(k string, v int) bool { return v%2 == 0 }
		for k, v := range it.Filter2(m, isEven) {
			fmt.Printf("%s:%d ", k, v)
		}
		fmt.Println()
		// Output: b:2
	}
}
