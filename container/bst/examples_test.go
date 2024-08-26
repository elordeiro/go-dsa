package bst_test

import (
	"fmt"

	bt "github.com/elordeiro/go/container/bst"
)

func ExampleNewBst() {
	b := bt.NewBst(3, 2, 1, 4, 5)
	fmt.Println(b)
	// Output: /\[1 2 3 4 5]
}

func ExampleTree_Insert() {
	b := bt.NewBst(3, 2, 1, 4, 5)
	b.Insert(0)
	fmt.Println(b)
	// Output: /\[0 1 2 3 4 5]
}

func ExampleTree_Delete() {
	b := bt.NewBst(3, 2, 1, 4, 5)
	b.Delete(3)
	fmt.Println(b)
	// Output: /\[1 2 4 5]
}

func ExampleTree_Search() {
	b := bt.NewBst(3, 2, 1, 4, 5)
	fmt.Println(b.Search(3))
	// Output: true
}

func ExampleTree_Enumerate() {
	b := bt.NewBst(3, 2, 1, 4, 5)
	for i, v := range b.Enumerate(0, b.Inorder()) {
		fmt.Println(i, v)
	}
	// Output:
	// 0 1
	// 1 2
	// 2 3
	// 3 4
	// 4 5
}

func ExampleTree_Inorder() {
	b := bt.NewBst(3, 2, 1, 4, 5)
	fmt.Print(b.StringSeq(b.Inorder()))
	// Output: /\[1 2 3 4 5]
}

func ExampleTree_Preorder() {
	b := bt.NewBst(3, 2, 4, 1, 5)
	fmt.Print(b.StringSeq(b.Preorder()))
	// Output: /\[3 2 1 4 5]
}

func ExampleTree_Postorder() {
	b := bt.NewBst(3, 2, 4, 1, 5)
	fmt.Print(b.StringSeq(b.Postorder()))
	// Output: /\[1 2 5 4 3]
}

func ExampleTree_Levelorder() {
	b := bt.NewBst(3, 2, 4, 1, 5)
	fmt.Print(b.StringSeq(b.Levelorder()))
	// Output: /\[3 2 4 1 5]
}

func ExampleTree_StringSeq() {
	b := bt.NewBst(3, 2, 1, 4, 5)
	fmt.Println(b.StringSeq(b.Inorder()))
	// Output: /\[1 2 3 4 5]
}

func ExampleTree_String() {
	b := bt.NewBst(3, 2, 1, 4, 5)
	fmt.Println(b)
	// Output: /\[1 2 3 4 5]
}
