package transform_test

import (
	"slices"
	"testing"

	"github.com/elordeiro/goext/containers/tuples"
	"github.com/elordeiro/goext/seq2s"
	"github.com/elordeiro/goext/seq2s/transform2"
	"github.com/elordeiro/goext/seqs"
	"github.com/elordeiro/goext/seqs/transform"
)

func TestBackwards(t *testing.T) {
	want := slices.Values([]int{5, 4, 3, 2, 1})
	seq := seqs.Range(1, 6)
	got := transform.Backwards(seq)

	if !seqs.Equal(got, want) {
		t.Errorf("Backwards(%s) = %s, want %s", seqs.String(seq), seqs.String(got), seqs.String(want))
	}
}

func TestDropWhile(t *testing.T) {
	want := slices.Values([]int{30, 40, 50})
	lessThan30 := func(v int) bool {
		return v < 30
	}
	seq := seqs.Range(10, 60, 10)
	got := transform.DropWhile(seqs.Range(10, 60, 10), lessThan30)
	if !seqs.Equal(got, want) {
		t.Errorf("DropWhile(%s, %s) = %s, want %s",
			seqs.String(seq), "lessThan30()", seqs.String(got), seqs.String(want),
		)
	}
}
func TestFilter(t *testing.T) {
	want := slices.Values([]int{2, 4, 6, 8, 10})

	isEven := func(v int) bool {
		return v%2 == 0
	}
	seq := seqs.Range(1, 11)
	got := transform.Filter(seq, isEven)

	if !seqs.Equal(got, want) {
		t.Errorf("Filter(%s, %s) = %s, want %s",
			seqs.String(seq), "isEven()", seqs.String(got), seqs.String(want),
		)
	}
}

func TestForEach(t *testing.T) {
	want := slices.Values([]int{1, 4, 3, 16, 5, 36})
	slice := []int{}
	seq := seqs.Range(1, 7)

	doubleEvens := func(v int) {
		if v%2 == 0 {
			slice = append(slice, v*v)
		} else {
			slice = append(slice, v)
		}
	}

	transform.ForEach(seqs.Range(1, 7), doubleEvens)
	got := slices.Values(slice)
	if !seqs.Equal(got, want) {
		t.Errorf("ForEach(%s, %s) = %s, want %s",
			seqs.String(seq), "doubleEvens()", seqs.String(got), seqs.String(want),
		)
	}
}

func TestMap(t *testing.T) {
	want := slices.Values([]int{4, 9, 16, 25, 36})

	double := func(v int) int {
		return v * v
	}
	seq := seqs.Range(2, 7)
	got := transform.Map(seq, double)

	if !seqs.Equal(got, want) {
		t.Errorf("Map(%s, %s) = %s, want %s",
			seqs.String(seq), "double()", seqs.String(got), seqs.String(want),
		)
	}
}

func TestOnEmpty(t *testing.T) {
	tests := []struct {
		i    int
		want bool
	}{
		{1, false},
		{2, false},
		{10, true},
	}

	var got bool
	callback := func() {
		got = true
	}

	seq := seqs.Range(1, 6)
	for _, tc := range tests {
		got = false
		for v := range transform.OnEmpty(seq, callback) {
			if v == tc.i {
				break
			}
		}
		if got != tc.want {
			t.Errorf("OnEmpty(%s, %s) = %v, want %v", seqs.String(seq), "callback()", got, tc.want)
		}
	}

}

func TestReduce(t *testing.T) {
	tests := []struct {
		vals []int
		want int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 55},
		{[]int{10, 20, 30, 40, 50}, 150},
	}

	sum := func(a, b int) int {
		return a + b
	}

	for _, tc := range tests {
		seq := slices.Values(tc.vals)
		got := transform.Reduce(seq, sum)
		if got != tc.want {
			t.Errorf("Reduce(%s, %s) = %d, want %d",
				seqs.String(seq), "sum()", got, tc.want,
			)
		}
	}
}

func TestRotate(t *testing.T) {
	tests := []struct {
		n    int
		want []int
	}{
		{2, []int{30, 40, 50, 10, 20}},
		{6, []int{10, 20, 30, 40, 50}},
	}

	for _, tc := range tests {
		want := slices.Values(tc.want)
		seq := seqs.Range(10, 60, 10)
		got := transform.Rotate(tc.n, seq)
		if !seqs.Equal(got, want) {
			t.Errorf("Rotate(%d, %s) = %s, want %s",
				tc.n, seqs.String(seq), seqs.String(got), seqs.String(want),
			)
		}
	}
}

func TestTakeWhile(t *testing.T) {
	tests := []struct {
		i, j int
		want []int
	}{
		{0, 5, []int{0, 1, 2, 3, 4}},
		{5, 15, []int{5, 6, 7, 8, 9}},
		{11, 20, []int{}},
	}

	lessThan10 := func(v int) bool {
		return v < 10
	}

	for _, tc := range tests {
		want := slices.Values(tc.want)
		seq := seqs.Range(tc.i, tc.j)
		got := transform.TakeWhile(seq, lessThan10)
		if !seqs.Equal(got, want) {
			t.Errorf("TakeWhile(%s, %s) = %s, %s",
				seqs.String(seq), "lessThan10()", seqs.String(got), seqs.String(want),
			)
		}
	}
}

func TestUnpair(t *testing.T) {
	seq1 := slices.Values([]int{1, 2, 3, 4, 5})
	seq2 := slices.Values([]string{"one", "two", "three", "four", "five"})
	want := seqs.Zip(seq1, seq2)

	var pairSlice []tuples.Pair[int, string]
	transform2.ForEach(want, func(i int, a string) {
		pairSlice = append(pairSlice, tuples.NewPair(i, a))
	})
	seq := slices.Values(pairSlice)
	got := transform.Unpair(seq)

	if !seq2s.Equal(got, want) {
		t.Errorf("Unpair(%s) = %s, want %s",
			seqs.String(seq), seq2s.String(got), seq2s.String(want))
	}
}

func TestWith(t *testing.T) {
	tests := []struct {
		vals []int
		want []int
	}{
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []int{0, 2, 4, 6, 8}},
		{[]int{}, []int{}},
	}

	var slice []int
	takeEvens := func(v int) {
		if v%2 == 0 {
			slice = append(slice, v)
		}
	}

	for _, tc := range tests {
		slice = []int{}
		seq := slices.Values(tc.vals)
		want := slices.Values(tc.want)
		transform.With(seq, takeEvens)(func(v int) bool {
			return true
		})
		got := slices.Values(slice)
		if !seqs.Equal(got, want) {
			t.Errorf("With(%s, %s) = %s, want %s",
				seqs.String(seq), "takeEvens()", seqs.String(got), seqs.String(want),
			)
		}
	}
}
