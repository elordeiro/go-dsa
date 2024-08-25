package stack_test

import (
	"fmt"

	stk "github.com/elordeiro/go/container/stack"
)

func ExampleStack() {
	s := stk.NewStack(1, 2, 3, 4, 5)
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
	s := stk.NewStack(1, 2, 3)
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
	s := stk.NewStack(1, 2, 3, 4, 5)
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	// Output:
	// 5
	// 4
	// 3
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
	s.Pop()
	fmt.Println(s.Peek())
	// Output:
	// 5
	// 4
}

func ExampleStack_Len() {
	s := stk.NewStack(1, 2, 3, 4, 5)
	fmt.Println(s.Len())
	// Output: 5
}

func ExampleStack_Enumerate() {
	s := stk.NewStack(1, 2, 3, 4, 5)
	s.Enumerate(0)(func(i int, v int) bool {
		fmt.Println(i, v)
		return true
	})
	fmt.Println(s.IsEmpty())
	// Output:
	// 0 5
	// 1 4
	// 2 3
	// 3 2
	// 4 1
	// true
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
