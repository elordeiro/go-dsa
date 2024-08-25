package list_test

import (
	"fmt"

	"github.com/elordeiro/go/container/list"
)

func ExampleNewList() {
	list := list.NewList(1, 2, 3)
	for val := range list.All() {
		fmt.Println(val)
	}
	// Output:
	// 1
	// 2
	// 3
}

func ExampleList_Append() {
	list := list.NewList(1, 2, 3)
	list = list.Append(4)
	for val := range list.All() {
		fmt.Println(val)
	}
	// Output:
	// 1
	// 2
	// 3
	// 4
}

func ExampleList_Prepend() {
	list := list.NewList(1, 2, 3)
	list = list.Prepend(0)
	for val := range list.All() {
		fmt.Println(val)
	}
	// Output:
	// 0
	// 1
	// 2
	// 3
}

func ExampleRemove() {
	l := list.NewList(1, 2, 3)
	l = list.Remove(l, 2)
	for val := range l.All() {
		fmt.Println(val)
	}
	// Output:
	// 1
	// 3
}

func ExampleRemoveAll() {
	l := list.NewList(1, 2, 2, 3)
	l = list.RemoveAll(l, 2)
	for val := range l.All() {
		fmt.Println(val)
	}
	// Output:
	// 1
	// 3
}

func ExampleCompare() {
	l1 := list.NewList(1, 2, 3)
	l2 := list.NewList(1, 2, 3)
	fmt.Println(list.Compare(l1, l2))
	// Output:
	// true
}

func ExampleList_Len() {
	l := list.NewList(1, 2, 3)
	fmt.Println(l.Len())
	// Output:
	// 3
}

func ExampleList_All() {
	l := list.NewList(1, 2, 3)
	for val := range l.All() {
		fmt.Println(val)
	}
	// Output:
	// 1
	// 2
	// 3
}

func ExampleList_Enumerate() {
	l := list.NewList(1, 2, 3)
	for i, val := range l.Enumerate(0) {
		fmt.Println(i, val)
	}
	// Output:
	// 0 1
	// 1 2
	// 2 3
}
