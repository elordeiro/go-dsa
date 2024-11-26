package avl_test

import (
	"fmt"

	"github.com/elordeiro/goext/containers/avl"
)

func ExampleNew() {
	b := avl.New(3, 2, 1, 4, 5)
	fmt.Println(b)
	// Output: ^[1 2 3 4 5]
}

func ExampleTree_Insert() {
	b := avl.New(3, 2, 1, 4, 5)
	b.Insert(0)
	fmt.Println(b)
	// Output: ^[0 1 2 3 4 5]
}

func ExampleTree_Delete() {
	b := avl.New(3, 2, 1, 4, 5)
	b.Delete(3)
	fmt.Println(b)
	// Output: ^[1 2 4 5]
}

func ExampleTree_Search() {
	b := avl.New(3, 2, 1, 4, 5)
	fmt.Println(b.Search(4))
	// Output: ^[3 4 5]
}

func ExampleMin() {
	b := avl.New(3, 2, 1, 4, 5)
	fmt.Println(b.Min().Value())
	// Output: 1
}

func ExampleMax() {
	b := avl.New(3, 2, 1, 4, 5)
	fmt.Println(b.Max().Value())
	// Output: 5
}

func ExampleTree_Inorder() {
	b := avl.New(3, 2, 1, 4, 5)
	fmt.Print(b)
	// Output: ^[1 2 3 4 5]
}

func ExampleTree_Preorder() {
	b := avl.New(3, 2, 4, 1, 5)
	fmt.Print(b.StringOrder(b.Preorder))
	// Output: ^[3 2 1 4 5]
}

func ExampleTree_Postorder() {
	b := avl.New(3, 2, 4, 1, 5)
	fmt.Print(b.StringOrder(b.Postorder))
	// Output: ^[1 2 5 4 3]
}

func ExampleTree_Levelorder() {
	b := avl.New(3, 2, 4, 1, 5)
	fmt.Print(b.StringOrder(b.Levelorder))
	// Output: ^[3 2 4 1 5]
}

func ExampleTree_String() {
	b := avl.New(3, 2, 1, 4, 5)
	fmt.Println(b)
	// Output: ^[1 2 3 4 5]
}
