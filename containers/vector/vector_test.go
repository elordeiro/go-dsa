package vector_test

import (
	"slices"
	"testing"

	"github.com/elordeiro/goext/containers/vector"
	"github.com/elordeiro/goext/seq2s"
	"github.com/elordeiro/goext/seqs"
)

func TestNew(t *testing.T) {
	v := vector.New(1, 2, 3)

	want := []int{1, 2, 3}
	if !slices.Equal(v, want) {
		t.Errorf("got %v, want %v", v, want)
	}
}

func TestAt(t *testing.T) {
	v := vector.New(1, 2, 3)

	for i, want := range seqs.Enumerate(0, seqs.Range(1, 4)) {
		got := v.At(i)
		if got != want {
			t.Errorf("At(%d) = %d, want %d", i, v.At(i), want)
		}
	}
}

func TestBack(t *testing.T) {
	v := vector.New(1, 2, 3)

	want := 3
	if v.Back() != want {
		t.Errorf("Back() = %d, want %d", v.Back(), want)
	}
}

func TestFront(t *testing.T) {
	v := vector.New(1, 2, 3)

	want := 1
	if v.Front() != want {
		t.Errorf("Front() = %d, want %d", v.Front(), want)
	}
}

func TestIsEmpty(t *testing.T) {
	v := vector.New[int]()

	if !v.IsEmpty() {
		t.Errorf("IsEmpty() = false, want true")
	}

	v.Push(1)

	if v.IsEmpty() {
		t.Errorf("IsEmpty() = true, want false")
	}
}

func TestLen(t *testing.T) {
	v := vector.New(1, 2, 3)

	want := 3
	if v.Len() != want {
		t.Errorf("Len() = %d, want %d", v.Len(), want)
	}
}

func TestReverse(t *testing.T) {
	v := vector.New(1, 2, 3, 4, 5, 6)
	v.Reverse()

	want := vector.New(6, 5, 4, 3, 2, 1)
	if !slices.Equal(v, want) {
		t.Errorf("Reverse() = %v, want %v", v, want)
	}
}

func TestSet(t *testing.T) {
	v := vector.New(1, 2, 3)
	v.Set(1, 4)

	want := vector.New(1, 4, 3)
	if !slices.Equal(v, want) {
		t.Errorf("Set(1, 4) = %v, want %v", v, want)
	}
}

func TestSwap(t *testing.T) {
	v := vector.New(1, 2, 3, 4, 5, 6)
	v.Swap(1, 4)

	want := vector.New(1, 5, 3, 4, 2, 6)
	if !slices.Equal(v, want) {
		t.Errorf("Swap(1, 4) = %v, want %v", v, want)
	}
}

func TestClear(t *testing.T) {
	v := vector.New(1, 2, 3, 4, 5, 6)
	v.Clear()

	want := vector.New(0, 0, 0, 0, 0, 0)
	if !slices.Equal(v, want) {
		t.Errorf("Clear() = %v, want %v", v, want)
	}
}

func TestConcat(t *testing.T) {
	v1 := vector.New(1, 2, 3)
	v2 := vector.New(4, 5, 6)

	v1.Concat(v2)

	want := vector.New(1, 2, 3, 4, 5, 6)
	if !slices.Equal(v1, want) {
		t.Errorf("Concat() = %v, want %v", v1, want)
	}
}

func TestCopy(t *testing.T) {
	v1 := vector.New(1, 2, 3)
	v2 := vector.New(4, 5, 6)

	v1.Copy(v2)

	want := vector.New(4, 5, 6)
	if !slices.Equal(v1, want) {
		t.Errorf("Copy() = %v, want %v", v1, want)
	}
}

func TestCut(t *testing.T) {
	v := vector.New(1, 2, 3, 4, 5, 6)
	v.Cut(2, 4)

	want := vector.New(1, 2, 5, 6)
	if !slices.Equal(v, want) {
		t.Errorf("Cut(2, 4) = %v, want %v", v, want)
	}
}

func TestInsert(t *testing.T) {
	v := vector.New(1, 2, 3)
	v.Insert(1, 0, 0)

	want := vector.New(1, 0, 0, 2, 3)
	if !slices.Equal(v, want) {
		t.Errorf("Insert(1, 0, 0) = %v, want %v", v, want)
	}

	v.Insert(0, 0, 0)

	want = vector.New(0, 0, 1, 0, 0, 2, 3)
	if !slices.Equal(v, want) {
		t.Errorf("Insert(0, 0, 0) = %v, want %v", v, want)
	}

	v.Insert(7, 0, 0)

	want = vector.New(0, 0, 1, 0, 0, 2, 3, 0, 0)
	if !slices.Equal(v, want) {
		t.Errorf("Insert(8, 0, 0) = %v, want %v", v, want)
	}
}

func TestPop(t *testing.T) {
	v := vector.New(1, 2, 3, 4, 5, 6)
	got := v.Pop()

	want := 6
	if got != want {
		t.Errorf("Pop() = %d, want %d", got, want)
	}
}

func TestPopAt(t *testing.T) {
	v := vector.New(1, 2, 3, 4, 5, 6)
	got := v.PopAt(2)

	want := 3
	if got != want {
		t.Errorf("PopAt(2) = %d, want %d", got, want)
	}

	wantVec := vector.New(1, 2, 4, 5, 6)
	if !slices.Equal(v, wantVec) {
		t.Errorf("got %v, want %v", v, wantVec)
	}
}

func TestPush(t *testing.T) {
	v := vector.New(1, 2, 3)
	v.Push(4)

	want := vector.New(1, 2, 3, 4)
	if !slices.Equal(v, want) {
		t.Errorf("Push(4) = %v, want %v", v, want)
	}
}

func TestValues(t *testing.T) {
	v := vector.New(1, 2, 3)

	want := slices.Values([]int{1, 2, 3})
	got := v.Values()
	if !seqs.Equal(got, want) {
		t.Errorf("Values() = %v, want %v", got, want)
	}
}

func TestAll(t *testing.T) {
	v := vector.New(1, 2, 3)

	want := slices.All([]int{1, 2, 3})
	got := v.All()
	if !seq2s.Equal(got, want) {
		t.Errorf("All() = %v, want %v", got, want)
	}
}

func TestBackwards(t *testing.T) {
	v := vector.New(1, 2, 3)

	want := slices.Values([]int{3, 2, 1})
	got := v.Backwards()
	if !seqs.Equal(got, want) {
		t.Errorf("Backwards() = %v, want %v", got, want)
	}
}
