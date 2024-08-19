package tests

import (
	"testing"

	"github.com/elordeiro/go/dsa/stack"
)

func TestStack(t *testing.T) {
	s := stack.NewStack[int]()

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

	if s.Len() != 0 {
		t.Error("Stack should have 0 elements")
	}
}

func TestStackNonInt(t *testing.T) {
	arr := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}

	s := stack.NewStack("a", "b", "c", "d", "e", "f", "g", "h", "i", "j")

	for i := 10; i < 100; i++ {
		s.Push(arr[i%10])
	}

	if s.Len() != 100 {
		t.Error("Stack should have 100 elements")
	}

	for i := 99; i >= 0; i-- {
		if s.Peek() != arr[i%10] {
			t.Error("Peek should return the top element")
		}
		if s.Pop() != arr[i%10] {
			t.Error("Pop should return the top element")
		}
	}

	if !s.IsEmpty() {
		t.Error("Stack should be empty")
	}

	if s.Len() != 0 {
		t.Error("Stack should have 0 elements")
	}
}
