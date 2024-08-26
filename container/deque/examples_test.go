package deque_test

import (
	"fmt"

	dq "github.com/elordeiro/go/container/deque"
)

func ExampleNewDeque() {
	d := dq.NewDeque(1, 2, 3)
	d.PushFront(0)
	d.PushBack(4)
	fmt.Println(d)
	// Output: <->[0 1 2 3 4]
}

func ExampleDeque_PushFront() {
	d := dq.NewDeque(1, 2, 3)
	d.PushFront(0)
	fmt.Println(d)
	// Output: <->[0 1 2 3]
}

func ExampleDeque_PushBack() {
	d := dq.NewDeque(1, 2, 3)
	d.PushBack(4)
	fmt.Println(d)
	// Output: <->[1 2 3 4]
}

func ExampleDeque_PopFront() {
	d := dq.NewDeque(1, 2, 3)
	fmt.Println(d.PopFront())
	// Output: 1
}

func ExampleDeque_PopBack() {
	d := dq.NewDeque(1, 2, 3)
	fmt.Println(d.PopBack())
	// Output: 3
}

func ExampleDeque_PeekFront() {
	d := dq.NewDeque(1, 2, 3)
	fmt.Println(d.PeekFront())
	fmt.Println(d)
	// Output:
	// 1
	// <->[1 2 3]
}

func ExampleDeque_PeekBack() {
	d := dq.NewDeque(1, 2, 3)
	fmt.Println(d.PeekBack())
	fmt.Println(d)
	// Output:
	// 3
	// <->[1 2 3]
}

func ExampleDeque_Len() {
	d := dq.NewDeque(1, 2, 3)
	fmt.Println(d.Len())
	// Output: 3
}

func ExampleDeque_IsEmpty() {
	d := dq.NewDeque[int]()
	fmt.Println(d.IsEmpty())
	d.PushBack(1)
	fmt.Println(d.IsEmpty())
	// Output:
	// true
	// false
}

func ExampleDeque_All() {
	d := dq.NewDeque(1, 2, 3, 4)
	for v := range d.All() {
		fmt.Println(v)
	}
	// Output:
	// 1
	// 2
	// 3
	// 4
}

func ExampleDeque_Enumerate() {
	d := dq.NewDeque(1, 2, 3, 4)
	for i, v := range d.Enumerate(0) {
		fmt.Println(i, v)
	}
	// Output:
	// 0 1
	// 1 2
	// 2 3
	// 3 4
}

func ExampleDeque_Backwards() {
	d := dq.NewDeque(1, 2, 3, 4)
	for v := range d.Backwards() {
		fmt.Println(v)
	}
	// Output:
	// 4
	// 3
	// 2
	// 1
}

func ExampleDeque_EnumerateBackwards() {
	d := dq.NewDeque(1, 2, 3, 4)
	for i, v := range d.EnumerateBackwards(0) {
		fmt.Println(i, v)
	}
	// Output:
	// 0 4
	// 1 3
	// 2 2
	// 3 1
}

func ExampleDeque_String() {
	d := dq.NewDeque(1, 2, 3, 4)
	fmt.Println(d)
	// Output: <->[1 2 3 4]
}
