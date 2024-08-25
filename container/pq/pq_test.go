package pq_test

import (
	"testing"

	pq "github.com/elordeiro/go/container/pq"
)

type pair struct {
	fruit    string
	priority int
}

func TestMaxPq(t *testing.T) {
	expected := []int{7, 5, 3, 1}
	test := []int{7, 3, 1, 5}

	pq := pq.NewMaxHeap(test...)

	t.Run("Testing max pq", func(t *testing.T) {
		for _, fruit := range expected {
			actual := pq.Pop()
			if fruit != actual {
				t.Errorf("Test Failed\n\tActual  : %v\n\tExpected: %v\n", fruit, actual)
			}
		}
	})
}

func TestMinPq(t *testing.T) {
	expected := []int{1, 3, 5, 7}
	test := []int{7, 3, 1, 5}

	pq := pq.NewMinHeap(test...)

	t.Run("Testing min pq", func(t *testing.T) {
		for _, fruit := range expected {
			actual := pq.Pop()
			if fruit != actual {
				t.Errorf("Test Failed\n\tActual  : %v\n\tExpected: %v\n", fruit, actual)
			}
		}
	})
}

func TestMaxPqFunc(t *testing.T) {
	expected := []string{"banana", "grape", "orange", "apple"}
	test := []pair{
		{"banana", 7},
		{"orange", 3},
		{"apple", 1},
		{"grape", 5},
	}

	pq := pq.NewPqFunc(func(item1, item2 pair) bool {
		return item1.priority > item2.priority
	}, test...)

	t.Run("Testing max pq func", func(t *testing.T) {
		for _, fruit := range expected {
			actual := pq.Pop().fruit
			if fruit != actual {
				t.Errorf("Test Failed\n\tActual  : %v\n\tExpected: %v\n", fruit, actual)
			}
		}
	})
}

func TestMinPqFunc(t *testing.T) {
	expected := []string{"apple", "orange", "grape", "banana"}
	test := []pair{
		{"banana", 7},
		{"orange", 3},
		{"apple", 1},
		{"grape", 5},
	}

	pq := pq.NewPqFunc(func(item1, item2 pair) bool {
		return item1.priority < item2.priority
	}, test...)

	t.Run("Testing min pq func", func(t *testing.T) {
		for _, fruit := range expected {
			actual := pq.Pop().fruit
			if fruit != actual {
				t.Errorf("Test Failed\n\tActual  : %v\n\tExpected: %v\n", fruit, actual)
			}
		}
	})
}

func TestPqAll(t *testing.T) {
	expected := []string{"apple", "banana", "grape", "orange"}
	test := []string{"banana", "orange", "apple", "grape"}

	pq := pq.NewMinHeap(test...)

	t.Run("Testing all pq", func(t *testing.T) {
		for i, actual := range pq.All() {
			if actual != expected[i] {
				t.Errorf("Test Failed\n\tActual  : %v\n\tExpected: %v\n", actual, expected[i])
			}
		}
	})
}

func TestPqUpdate(t *testing.T) {
	expected := []string{"apple", "banana", "grape", "orange"}
	test := []pair{
		{"banana", 7},
		{"orange", 3},
		{"apple", 1},
		{"grape", 5},
	}

	pq := pq.NewPqFunc(func(item1, item2 pair) bool {
		return item1.priority > item2.priority
	}, test...)

	pq.Update(2, pair{"apple", 10})

	t.Run("Testing update pq", func(t *testing.T) {
		for i, actual := range pq.All() {
			if actual.fruit != expected[i] {
				t.Errorf("Test Failed\n\tActual  : %v\n\tExpected: %v\n", actual, expected[i])
			}
		}
	})
}

func TestPqUpdateLess(t *testing.T) {
	expected := []string{"banana", "grape", "orange", "apple"}
	test := []pair{
		{"banana", 7},
		{"orange", 3},
		{"apple", 1},
		{"grape", 5},
	}

	pq := pq.NewPqFunc(func(item1, item2 pair) bool {
		return item1.priority < item2.priority
	}, test...)

	pq.UpdateLess(func(item1, item2 pair) bool {
		return item1.priority > item2.priority
	})

	t.Run("Testing update pq less", func(t *testing.T) {
		for i, actual := range pq.All() {
			if actual.fruit != expected[i] {
				t.Errorf("Test Failed\n\tActual  : %v\n\tExpected: %v\n", actual, expected[i])
			}
		}
	})
}

func TestPqUpdateAll(t *testing.T) {
	expected := []string{"grape", "apple", "orange", "banana"}
	test := []pair{
		{"banana", 7},
		{"orange", 3},
		{"apple", 1},
		{"grape", 5},
	}

	pq := pq.NewPqFunc(func(item1, item2 pair) bool {
		return item1.priority > item2.priority
	}, test...)

	pq.UpdateAll([]pair{
		{"banana", 1},
		{"orange", 3},
		{"apple", 5},
		{"grape", 7},
	}...)

	t.Run("Testing update all pq", func(t *testing.T) {
		for i, actual := range pq.All() {
			if actual.fruit != expected[i] {
				t.Errorf("Test Failed\n\tActual  : %v\n\tExpected: %v\n", actual, expected[i])
			}
		}
	})
}
