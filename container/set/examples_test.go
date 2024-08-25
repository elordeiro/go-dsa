package set_test

import (
	"fmt"

	"github.com/elordeiro/go/container/set"
)

func ExampleNewSet() {
	s := set.NewSet(1, 2, 3)
	for v := range s.All() {
		fmt.Println(v)
	}
	// Unordered output:
	// 3
	// 1
	// 2
}

func ExampleSet_Len() {
	s := set.NewSet(1, 2, 3)
	fmt.Println(s.Len())
	// Output: 3
}

func ExampleSet_Add() {
	s := set.NewSet(1, 2, 3)
	s.Add(4)
	for v := range s.All() {
		fmt.Println(v)
	}
	// Unordered output:
	// 3
	// 1
	// 2
	// 4
}

func ExampleSet_Remove() {
	s := set.NewSet(1, 2, 3)
	s.Remove(2)
	for v := range s.All() {
		fmt.Println(v)
	}
	// Unordered output:
	// 3
	// 1
}

func ExampleSet_Contains() {
	s := set.NewSet(1, 2, 3)
	fmt.Println(s.Contains(1))
	// Output: true
}

func ExampleSet_IsEmpty() {
	s := set.NewSet[int]()
	fmt.Println(s.IsEmpty())
	s.Add(1)
	fmt.Println(s.IsEmpty())
	// Output:
	// true
	// false
}

func ExampleSet_Union() {
	s1 := set.NewSet(1, 2, 3)
	s2 := set.NewSet(3, 4, 5)
	union := s1.Union(s2)
	for v := range union.All() {
		fmt.Println(v)
	}
	// Unordered output:
	// 3
	// 1
	// 2
	// 4
	// 5
}

func ExampleSet_Intersection() {
	s1 := set.NewSet(1, 2, 3)
	s2 := set.NewSet(3, 4, 5)
	intersection := s1.Intersection(s2)
	for v := range intersection.All() {
		fmt.Println(v)
	}
	// Output: 3
}

func ExampleSet_Difference() {
	s1 := set.NewSet(1, 2, 3)
	s2 := set.NewSet(3, 4, 5)
	difference := s1.Difference(s2)
	for v := range difference.All() {
		fmt.Println(v)
	}
	// Unordered output:
	// 1
	// 2
}

func ExampleSet_All() {
	s := set.NewSet(1, 2, 3)
	for v := range s.All() {
		fmt.Println(v)
	}
	// Unordered output:
	// 3
	// 1
	// 2
}
