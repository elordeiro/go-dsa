package main

import (
	"fmt"

	itr "github.com/elordeiro/go/iters/iters2"
)

func main() {
	// Enumarator works with slices
	arr := itr.NewSlice(1, 2, 3, 4, 5)
	iter1 := itr.Enumerator(arr)
	fmt.Println(iter1)

	iter5 := itr.Enumerator2(iter1)
	fmt.Println(iter5)

	// Enumarator works with maps
	m := itr.NewMap([]itr.MapPair[int, string]{{1, "one"}, {2, "two"}, {3, "three"}}...)
	iter2 := itr.Enumerator2(m)
	fmt.Println(iter2)

	// Enumarator works with Seq[V]
	iter3 := itr.Enumerator(itr.Range(5))
	fmt.Println(iter3)

	// Enumarator works with Seq2[K, V]
	iter4 := itr.Enumerator(itr.ToSeq2(itr.Range(2, 5)))
	fmt.Println(iter4)

	fmt.Println()

	map1 := map[int]string{1: "one", 2: "two", 3: "three"}
	fmt.Println(map1)

	map2 := map[int]map[int]string{1: {1: "one", 2: "two", 3: "three"}}
	fmt.Println(map2)

	arr = []int{1, 2, 3, 4, 5}
	fmt.Println(arr)
}
