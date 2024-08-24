package stack_test

import (
	"fmt"

	stk "github.com/elordeiro/go/container/stack"
)

func ExampleStack() {
	s := stk.NewStack(1, 2, 3, 4, 5)
	fmt.Println(s)
	// Output: [5 4 3 2 1]
}

func ExampleStack_Push() {
	s := stk.NewStack(1, 2, 3, 4, 5)
	s.Push(6, 7, 8)
	fmt.Println(s)
	// Output: [8 7 6 5 4 3 2 1]
}

func ExampleStack_Pop() {
	s := stk.NewStack(1, 2, 3, 4, 5)
	fmt.Println(s.Pop())
	fmt.Println(s)
	// Output:
	// 5
	// [4 3 2 1]
}

func ExampleStack_IsEmpty() {
	s := stk.NewStack[int]()
	fmt.Println(s.IsEmpty())

	s.Push(1)
	fmt.Println(s.IsEmpty())
	// Output:
	// true
	// false
}

func ExampleStack_Peek() {
	s := stk.NewStack(1, 2, 3, 4, 5)
	fmt.Println(s.Peek())
	fmt.Println(s)
	// Output:
	// 5
	// [5 4 3 2 1]
}

func ExampleStack_Len() {
	s := stk.NewStack(1, 2, 3, 4, 5)
	fmt.Println(s.Len())
	// Output: 5
}

func ExampleStack_All() {
	s := stk.NewStack(1, 2, 3, 4, 5)
	s.All()(func(v int) bool {
		fmt.Println(v)
		return true
	})
	fmt.Println(s.IsEmpty())
	// Output:
	// 5
	// 4
	// 3
	// 2
	// 1
	// true
}
