package stack_test

import (
	"fmt"

	"github.com/elordeiro/goext/containers/stack"
)

func ExampleNew() {
	s := stack.New(1, 2, 3, 4, 5)
	for v := range s.All() {
		fmt.Println(v)
	}
	// Output:
	// 5
	// 4
	// 3
	// 2
	// 1
}

func ExampleStack_Push() {
	s := stack.New(1, 2, 3)
	s.Push(4, 5, 6)
	for v := range s.All() {
		fmt.Println(v)
	}
	// Output:
	// 6
	// 5
	// 4
	// 3
	// 2
	// 1
}

func ExampleStack_Pop() {
	s := stack.New(1, 2, 3, 4, 5)
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	// Output:
	// 5
	// 4
	// 3
}

func ExampleStack_IsEmpty() {
	s := stack.New[int]()
	fmt.Println(s.IsEmpty())

	s.Push(1)
	fmt.Println(s.IsEmpty())
	// Output:
	// true
	// false
}

func ExampleStack_Top() {
	s := stack.New(1, 2, 3, 4, 5)
	fmt.Println(s.Top())
	s.Pop()
	fmt.Println(s.Top())
	// Output:
	// 5
	// 4
}

func ExampleStack_Len() {
	s := stack.New(1, 2, 3, 4, 5)
	fmt.Println(s.Len())
	// Output: 5
}

func ExampleStack_All() {
	s := stack.New(1, 2, 3, 4, 5)
	for v := range s.All() {
		fmt.Println(v)
	}
	// Output:
	// 5
	// 4
	// 3
	// 2
	// 1
}

func ExampleStack_Backwards() {
	s := stack.New(1, 2, 3, 4, 5)
	for v := range s.Backwards() {
		fmt.Println(v)
	}
	// Output:
	// 1
	// 2
	// 3
	// 4
	// 5
}

func ExampleStack_Drain() {
	s := stack.New(1, 2, 3, 4, 5)
	for v := range s.Drain() {
		fmt.Println(v)
	}
	fmt.Println(s.IsEmpty())
	// Output:
	// 5
	// 4
	// 3
	// 2
	// 1
	// true
}

func ExampleStack_String() {
	s := stack.New(1, 2, 3, 4, 5)
	fmt.Println(s.String())
	// Output: $[5 4 3 2 1]
}
