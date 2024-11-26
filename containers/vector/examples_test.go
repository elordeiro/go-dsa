package vector_test

import (
	"fmt"

	"github.com/elordeiro/goext/containers/vector"
)

func ExampleNew() {
	v := vector.New(1, 2, 3)
	fmt.Println(v)
	// Output: [1 2 3]
}

func ExampleVector_At() {
	v := vector.New(1, 2, 3)
	fmt.Println(v.At(1))
	// Output: 2
}

func ExampleVector_Back() {
	v := vector.New(1, 2, 3)
	fmt.Println(v.Back())
	// Output: 3
}

func ExampleVector_Front() {
	v := vector.New(1, 2, 3)
	fmt.Println(v.Front())
	// Output: 1
}

func ExampleVector_IsEmpty() {
	v := vector.New[int]()
	fmt.Println(v.IsEmpty())
	v.Push(1)
	fmt.Println(v.IsEmpty())
	// Output:
	// true
	// false
}

func ExampleVector_Len() {
	v := vector.New(1, 2, 3)
	fmt.Println(v.Len())
	// Output: 3
}

func ExampleVector_Reverse() {
	v := vector.New(1, 2, 3)
	v.Reverse()
	fmt.Println(v)
	// Output: [3 2 1]
}

func ExampleVector_Set() {
	v := vector.New(1, 2, 3)
	v.Set(1, 4)
	fmt.Println(v)
	// Output: [1 4 3]
}

func ExampleVector_Swap() {
	v := vector.New(1, 2, 3, 4, 5, 6)
	v.Swap(1, 4)
	fmt.Println(v)
	// Output: [1 5 3 4 2 6]
}

func ExampleVector_Clear() {
	v := vector.New(1, 2, 3)
	v.Clear()
	fmt.Println(v)
	// Output: [0 0 0]
}

func ExampleVector_Concat() {
	v := vector.New(1, 2, 3)
	v.Concat(vector.New(4, 5, 6))
	fmt.Println(v)
	// Output: [1 2 3 4 5 6]
}

func ExampleVector_Copy() {
	v := vector.New(1, 2, 3)
	v2 := vector.New(4, 5, 6)
	v.Copy(v2)
	fmt.Println(v)
	// Output: [4 5 6]
}

func ExampleVector_Cut() {
	v := vector.New(1, 2, 3, 4, 5, 6)
	v.Cut(1, 4)
	fmt.Println(v)
	// Output: [1 5 6]
}

func ExampleVector_Insert() {
	v := vector.New(1, 2, 3)
	v.Insert(1, 4)
	fmt.Println(v)
	// Output: [1 4 2 3]
}

func ExampleVector_Pop() {
	v := vector.New(1, 2, 3)
	fmt.Println(v.Pop())
	fmt.Println(v)
	// Output:
	// 3
	// [1 2]
}

func ExampleVector_PopAt() {
	v := vector.New(1, 2, 3)
	fmt.Println(v.PopAt(1))
	fmt.Println(v)
	// Output:
	// 2
	// [1 3]
}

func ExampleVector_Push() {
	v := vector.New[int]()
	v.Push(1)
	v.Push(2)
	v.Push(3)
	fmt.Println(v)
	// Output: [1 2 3]
}

func ExampleVector_String() {
	v := vector.New(1, 2, 3)
	fmt.Println(v)
	// Output: [1 2 3]
}

func ExampleVector_Values() {
	v := vector.New(1, 2, 3)
	for val := range v.Values() {
		fmt.Println(val)
	}
	// Output:
	// 1
	// 2
	// 3
}

func ExampleVector_All() {
	v := vector.New(1, 2, 3)
	for idx, val := range v.All() {
		fmt.Println(idx, val)
	}
	// Output:
	// 0 1
	// 1 2
	// 2 3
}

func ExampleVector_Backwards() {
	v := vector.New(1, 2, 3)
	for val := range v.Backwards() {
		fmt.Println(val)
	}
	// Output:
	// 3
	// 2
	// 1
}
