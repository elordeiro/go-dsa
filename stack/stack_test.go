package stack

import (
	"math/rand/v2"
	"testing"
)

func TestStack(t *testing.T) {
	goStack := []int{}
	customStack := Stack()
	for range 10 {
		num := rand.N(100)
		goStack = append(goStack, num)
		customStack.Push(num)
	}
	var testname string
	for !customStack.IsEmpty() {
		num1 := goStack[len(goStack)-1]
		num2 := customStack.Pop()
		testname = "Test Stack Pop"
		t.Run(testname, func(t *testing.T) {
			if num1 != num2 {
				t.Errorf("Actual   : %v", num1)
				t.Errorf("Expected : %v", num2)
			}
		})
		goStack = goStack[:len(goStack)-1]
	}
}
