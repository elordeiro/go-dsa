package pq_test

import (
	"fmt"

	"github.com/elordeiro/goext/containers/pq"
	"github.com/elordeiro/goext/containers/tuples"
)

func ExampleNewMaxPQ() {
	pq := pq.NewMaxPQ(7, 3, 1, 5)
	for v := range pq.All() {
		fmt.Println(v)
	}
	// Output:
	// 7
	// 5
	// 3
	// 1
}

func ExampleNewMinPQ() {
	pq := pq.NewMinPQ(7, 3, 1, 5)
	for v := range pq.All() {
		fmt.Println(v)
	}
	// Output:
	// 1
	// 3
	// 5
	// 7
}

func ExampleNewPQFunc() {
	pq := pq.NewPQFunc(func(item1, item2 tuples.Pair[string, int]) bool {
		return item1.Right() > item2.Right()
	})

	pq.Push(tuples.NewPair("banana", 7))
	pq.Push(tuples.NewPair("orange", 3))
	pq.Push(tuples.NewPair("apple", 1))
	pq.Push(tuples.NewPair("grape", 5))

	for v := range pq.All() {
		fmt.Println(v)
	}
	// Output:
	// (banana 7)
	// (grape 5)
	// (orange 3)
	// (apple 1)
}

func ExamplePQ_Len() {
	pq := pq.NewMaxPQ(7, 3, 1, 5)
	fmt.Println(pq.Len())
	// Output: 4
}

func ExamplePQ_Push() {
	pq := pq.NewMaxPQ(7, 3, 1, 5)
	pq.Push(9)
	for v := range pq.All() {
		fmt.Println(v)
	}
	// Output:
	// 9
	// 7
	// 5
	// 3
	// 1
}

func ExamplePQ_Pop() {
	pq := pq.NewMaxPQ(7, 3, 1, 5)
	fmt.Println(pq.Pop())
	fmt.Println(pq.Pop())
	fmt.Println(pq.Pop())
	fmt.Println(pq.Pop())
	// Output:
	// 7
	// 5
	// 3
	// 1
}

func ExamplePQ_Top() {
	pq := pq.NewMaxPQ(7, 3, 1, 5)
	fmt.Println(pq.Top())
	pq.Pop()
	fmt.Println(pq.Top())
	// Output:
	// 7
	// 5
}

func ExamplePQ_IsEmpty() {
	pq := pq.NewMaxPQ[int]()
	fmt.Println(pq.IsEmpty())
	pq.Push(1)
	fmt.Println(pq.IsEmpty())
	// Output:
	// true
	// false
}

func ExamplePQ_All() {
	pq := pq.NewMaxPQ(7, 3, 1, 5)
	for v := range pq.All() {
		fmt.Println(v)
	}
	// Output:
	// 7
	// 5
	// 3
	// 1
}

func ExamplePQ_Remove() {
	pq := pq.NewMaxPQ(7, 3, 1, 5)
	fmt.Println(pq.Remove(1))
	for v := range pq.All() {
		fmt.Println(v)
	}
	// Output:
	// 5
	// 7
	// 3
	// 1
}

func ExamplePQ_Update() {
	pq := pq.NewMaxPQ(9, 3, 5, 7)
	pq.Update(0, 1)
	for v := range pq.All() {
		fmt.Println(v)
	}
	// Output:
	// 7
	// 5
	// 3
	// 1
}

func ExamplePQ_String() {
	pq := pq.NewMaxPQ(7, 3, 1, 5)
	fmt.Println(pq)
	// Output: ![7 5 1 3]
}
