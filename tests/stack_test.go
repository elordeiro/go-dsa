package tests

import (
	"testing"

	"github.com/elordeiro/go-container/stack"
)

func TestStack(t *testing.T) {
	s := stack.NewStack()

	if !s.IsEmpty() {
		t.Error("New stack should be empty")
	}

	for i := 0; i < 100; i++ {
		s.Push(i)
	}

	if s.Len() != 100 {
		t.Error("Stack should have 100 elements")
	}

	for i := 99; i >= 0; i-- {
		if s.Peek() != i {
			t.Error("Peek should return the top element")
		}
		if s.Pop() != i {
			t.Error("Pop should return the top element")
		}
	}

	if !s.IsEmpty() {
		t.Error("Stack should be empty")
	}
}
