package deque_test

import (
	"fmt"

	"github.com/elordeiro/goext/containers/deque"
)

func ExampleNew() {
	d := deque.New(1, 2, 3)
	d.PushFront(0)
	d.PushBack(4)
	fmt.Println(d)
	// Output: <0 1 2 3 4>
}

func ExampleDeque_PushFront() {
	d := deque.New(1, 2, 3)
	d.PushFront(0)
	fmt.Println(d)
	// Output: <0 1 2 3>
}

func ExampleDeque_PushBack() {
	d := deque.New(1, 2, 3)
	d.PushBack(4)
	fmt.Println(d)
	// Output: <1 2 3 4>
}

func ExampleDeque_PopFront() {
	d := deque.New(1, 2, 3)
	fmt.Println(d.PopFront())
	// Output: 1
}

func ExampleDeque_PopBack() {
	d := deque.New(1, 2, 3)
	fmt.Println(d.PopBack())
	// Output: 3
}

func ExampleDeque_Front() {
	d := deque.New(1, 2, 3)
	fmt.Println(d.Front())
	fmt.Println(d)
	// Output:
	// 1
	// <1 2 3>
}

func ExampleDeque_Back() {
	d := deque.New(1, 2, 3)
	fmt.Println(d.Back())
	fmt.Println(d)
	// Output:
	// 3
	// <1 2 3>
}

func ExampleDeque_Len() {
	d := deque.New(1, 2, 3)
	fmt.Println(d.Len())
	// Output: 3
}

func ExampleDeque_IsEmpty() {
	d := deque.New[int]()
	fmt.Println(d.IsEmpty())
	d.PushBack(1)
	fmt.Println(d.IsEmpty())
	// Output:
	// true
	// false
}

func ExampleDeque_All() {
	d := deque.New(1, 2, 3, 4)
	for v := range d.All() {
		fmt.Println(v)
	}
	// Output:
	// 1
	// 2
	// 3
	// 4
}

func ExampleDeque_Backwards() {
	d := deque.New(1, 2, 3, 4)
	for v := range d.Backwards() {
		fmt.Println(v)
	}
	// Output:
	// 4
	// 3
	// 2
	// 1
}

func ExampleDeque_String() {
	d := deque.New(1, 2, 3, 4)
	fmt.Println(d)
	// Output: <1 2 3 4>
}
