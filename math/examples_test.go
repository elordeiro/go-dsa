package math_test

import (
	"fmt"
	"slices"

	"github.com/elordeiro/goext/math"
)

func ExampleSum() {
	seq := slices.Values([]int{1, 2, 3, 4, 5})
	fmt.Println(math.Sum(seq))
	// Output: 15
}

func ExampleProduct() {
	seq := slices.Values([]int{1, 2, 3, 4, 5})
	fmt.Println(math.Product(seq))
	// Output: 120
}

func ExampleMax() {
	seq := slices.Values([]int{1, 2, 3, 4, 5})
	fmt.Println(math.Max(seq))
	// Output: 5
}

func ExampleMin() {
	seq := slices.Values([]int{1, 2, 3, 4, 5})
	fmt.Println(math.Min(seq))
	// Output: 1
}
