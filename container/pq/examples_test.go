package pq_test

import (
	"fmt"

	"github.com/elordeiro/go/container/pq"
)

func ExampleNewMaxHeap() {
	pq := pq.NewMaxHeap(7, 3, 1, 5)
	for v := range pq.All() {
		fmt.Println(v)
	}
	// Output:
	// 7
	// 5
	// 3
	// 1
}

func ExampleNewMinHeap() {
	pq := pq.NewMinHeap(7, 3, 1, 5)
	for v := range pq.All() {
		fmt.Println(v)
	}
	// Output:
	// 1
	// 3
	// 5
	// 7
}

func ExampleNewPqFunc() {
	pq := pq.NewPqFunc(func(item1, item2 pair) bool {
		return item1.priority > item2.priority
	})

	pq.Push(pair{"banana", 7})
	pq.Push(pair{"orange", 3})
	pq.Push(pair{"apple", 1})
	pq.Push(pair{"grape", 5})

	for v := range pq.All() {
		fmt.Println(v)
	}
	// Output:
	// {banana 7}
	// {grape 5}
	// {orange 3}
	// {apple 1}
}

func ExamplePq_Len() {
	pq := pq.NewMaxHeap(7, 3, 1, 5)
	fmt.Println(pq.Len())
	// Output: 4
}

func ExamplePq_Push() {
	pq := pq.NewMaxHeap(7, 3, 1, 5)
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

func ExamplePq_Pop() {
	pq := pq.NewMaxHeap(7, 3, 1, 5)
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

func ExamplePq_Peek() {
	pq := pq.NewMaxHeap(7, 3, 1, 5)
	fmt.Println(pq.Peek())
	pq.Pop()
	fmt.Println(pq.Peek())
	// Output:
	// 7
	// 5
}

func ExamplePq_IsEmpty() {
	pq := pq.NewMaxHeap[int]()
	fmt.Println(pq.IsEmpty())
	pq.Push(1)
	fmt.Println(pq.IsEmpty())
	// Output:
	// true
	// false
}

func ExamplePq_All() {
	pq := pq.NewMaxHeap(7, 3, 1, 5)
	for v := range pq.All() {
		fmt.Println(v)
	}
	// Output:
	// 7
	// 5
	// 3
	// 1
}

func ExamplePq_Enumerate() {
	pq := pq.NewMaxHeap(7, 3, 1, 5)
	for i, v := range pq.Enumerate(0) {
		fmt.Println(i, v)
	}
	fmt.Println(pq.IsEmpty())
	// Output:
	// 0 7
	// 1 5
	// 2 3
	// 3 1
	// true
}

func ExamplePq_Remove() {
	pq := pq.NewMaxHeap(7, 3, 1, 5)
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

func ExamplePq_Clear() {
	pq := pq.NewMaxHeap(7, 3, 1, 5)
	pq.Clear()
	fmt.Println(pq.IsEmpty())
	// Output: true
}

func ExamplePq_Update() {
	pq := pq.NewMaxHeap(9, 3, 5, 7)
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

func ExamplePq_UpdateLess() {
	pq := pq.NewMaxHeap(9, 3, 5, 7)
	pq.UpdateLess(func(item1, item2 int) bool {
		return item1 < item2
	})
	for v := range pq.All() {
		fmt.Println(v)
	}
	// Output:
	// 3
	// 5
	// 7
	// 9
}

func ExamplePq_UpdateAll() {
	pq := pq.NewMaxHeap(9, 3, 5, 7)
	pq.UpdateAll(1, 2, 3, 4)
	for v := range pq.All() {
		fmt.Println(v)
	}
	// Output:
	// 4
	// 3
	// 2
	// 1
}
