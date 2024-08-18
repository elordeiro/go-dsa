package tests

import (
	"testing"

	dq "github.com/elordeiro/go-container/deque"
)

func TestDeque(t *testing.T) {
	dq := dq.NewDeque()
	for i := range 10 {
		dq.PushRight(i + 1)
	}
	tests := []int{1, 10, 2, 9, 3, 8, 4, 7, 5, 6}

	testname := "Deque Length"
	t.Run(testname, func(t *testing.T) {
		actual := dq.Length()
		if actual != 10 {
			t.Errorf("actual: %d, expected: %d", actual, 10)
		}
	})

	for i, test := range tests {
		if i%2 == 0 {
			testname := "PeekLeft"
			t.Run(testname, func(t *testing.T) {
				actual := dq.PeekLeft()
				if actual != test {
					t.Errorf("actual: %d, expected: %d", actual, test)
				}
			})
			testname = "PopLeft"
			t.Run(testname, func(t *testing.T) {
				actual := dq.PopLeft()
				if actual != test {
					t.Errorf("actual: %d, expected: %d", actual, test)
				}
			})
		} else {
			testname := "PeekRight"
			t.Run(testname, func(t *testing.T) {
				actual := dq.PeekRight()
				if actual != test {
					t.Errorf("actual: %d, expected: %d", actual, test)
				}
			})
			testname = "PopLeft"
			t.Run(testname, func(t *testing.T) {
				actual := dq.PopRight()
				if actual != test {
					t.Errorf("actual: %d, expected: %d", actual, test)
				}
			})
		}
	}

	testname = "Deque Length"
	t.Run(testname, func(t *testing.T) {
		actual := dq.Length()
		if actual != 0 {
			t.Errorf("actual: %d, expected: %d", actual, 0)
		}
	})
}
