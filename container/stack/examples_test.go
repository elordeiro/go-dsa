package containers_test

import (
	"fmt"

	stk "github.com/elordeiro/go/container/stack"
)

func ExampleStack() {
	s := stk.NewStack(1, 2, 3, 4, 5)
	fmt.Println(s)
	// Output: [5 4 3 2 1]
}
