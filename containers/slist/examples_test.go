package slist_test

import (
	"fmt"

	"github.com/elordeiro/goext/containers/slist"
)

func ExampleNew() {
	list := slist.New(1, 2, 3)
	for val := range list.All() {
		fmt.Println(val)
	}
	// Output:
	// 1
	// 2
	// 3
}

func ExampleList_Append() {
	list := slist.New(1, 2, 3)
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
	list := slist.New(1, 2, 3)
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
	l := slist.New(1, 2, 3)
	l = slist.Remove(l, 2)
	for val := range l.All() {
		fmt.Println(val)
	}
	// Output:
	// 1
	// 3
}

func ExampleRemoveAll() {
	l := slist.New(1, 2, 2, 3)
	l = slist.RemoveAll(l, 2)
	for val := range l.All() {
		fmt.Println(val)
	}
	// Output:
	// 1
	// 3
}

func ExampleCompare() {
	l1 := slist.New(1, 2, 3)
	l2 := slist.New(1, 2, 3)
	fmt.Println(slist.Compare(l1, l2))
	// Output:
	// true
}

func ExampleList_Len() {
	l := slist.New(1, 2, 3)
	fmt.Println(l.Len())
	// Output:
	// 3
}

func ExampleList_All() {
	l := slist.New(1, 2, 3)
	for val := range l.All() {
		fmt.Println(val)
	}
	// Output:
	// 1
	// 2
	// 3
}

func ExampleList_String() {
	l := slist.New(1, 2, 3)
	fmt.Println(l)
	// Output:
	// ->[1 2 3]
}
