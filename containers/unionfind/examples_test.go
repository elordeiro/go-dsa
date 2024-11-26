package unionfind_test

import (
	"fmt"

	"github.com/elordeiro/goext/containers/unionfind"
)

func ExampleNew() {
	uf := unionfind.New[int]()
	uf.MakeSet(1)
	uf.MakeSet(2)
	uf.MakeSet(3)
	for group := range uf.All() {
		fmt.Println(group)
	}
	// Unordered output:
	// {1}
	// {2}
	// {3}
}

func ExampleUnionFind_MakeSet() {
	uf := unionfind.New[int]()
	uf.MakeSet(1)
	uf.MakeSet(2)
	uf.MakeSet(3)
	for group := range uf.All() {
		fmt.Println(group)
	}
	// Unordered output:
	// {1}
	// {2}
	// {3}
}

func ExampleUnionFind_Union() {
	uf := unionfind.New[int]()
	uf.MakeSet(1)
	uf.MakeSet(2)
	uf.MakeSet(3)
	uf.Union(1, 2)
	uf.Union(2, 3)
	for group := range uf.All() {
		for el := range group {
			fmt.Println(el)
		}
	}
	// Unordered output:
	// 1
	// 2
	// 3
}

func ExampleUnionFind_Connected() {
	uf := unionfind.New[int]()
	uf.MakeSet(1)
	uf.MakeSet(2)
	uf.MakeSet(3)
	uf.Union(1, 2)
	fmt.Println(uf.Connected(1, 2))
	fmt.Println(uf.Connected(1, 3))
	// Output:
	// true
	// false
}

func ExampleUnionFind_All() {
	uf := unionfind.New[int]()
	uf.MakeSet(1)
	uf.MakeSet(2)
	uf.MakeSet(3)
	uf.Union(1, 2)
	uf.Union(2, 3)
	for group := range uf.All() {
		for el := range group {
			fmt.Println(el)
		}
	}
	// Unordered output:
	// 1
	// 2
	// 3
}

func ExampleUnionFind_String() {
	uf := unionfind.New[int]()
	uf.MakeSet(1)
	fmt.Println(uf)
	// Output: [~{1}]
}
