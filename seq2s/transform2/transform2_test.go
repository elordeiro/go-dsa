package transform2_test

import (
	"slices"
	"testing"

	"github.com/elordeiro/goext/seq2s"
	"github.com/elordeiro/goext/seq2s/transform2"
	"github.com/elordeiro/goext/seqs"
)

func TestBackwards(t *testing.T) {
	want := seq2s.FromSlice([][2]int{{4, 50}, {3, 40}, {2, 30}, {1, 20}, {0, 10}})
	seq := seqs.Enumerate(0, seqs.Range(10, 60, 10))
	got := transform2.Backwards(seq)

	if !seq2s.Equal(got, want) {
		t.Errorf("Backwards(%s) = %s, want %s",
			seq2s.String(seq), seq2s.String(got), seq2s.String(want),
		)
	}
}

func TestDropWhile(t *testing.T) {
	want := seq2s.FromSlice([][2]int{{2, 30}, {3, 40}, {4, 50}})
	seq := seqs.Enumerate(0, seqs.Range(10, 60, 10))

	lessThan2 := func(k, v int) bool {
		return k < 2
	}
	got := transform2.DropWhile(seq, lessThan2)

	if !seq2s.Equal(got, want) {
		t.Errorf("DropWhile(%s, %s) = %s, want %s",
			seq2s.String(seq), "lessThan2()", seq2s.String(got), seq2s.String(want),
		)
	}
}

func TestFilter(t *testing.T) {
	want := seq2s.FromSlice([][2]int{{0, 10}, {2, 30}, {4, 50}})
	seq := seqs.Enumerate(0, seqs.Range(10, 60, 10))

	isEven := func(k, v int) bool {
		return k%2 == 0
	}
	got := transform2.Filter(seq, isEven)

	if !seq2s.Equal(got, want) {
		t.Errorf("Filter(%s, %s) = %s, want %s",
			seq2s.String(seq), "isEven()", seq2s.String(got), seq2s.String(want),
		)
	}
}

func TestForEach(t *testing.T) {
	want := seq2s.FromSlice([][2]int{{0, 10}, {1, 20}, {2, 30}, {3, 40}, {4, 50}})
	seq := seqs.Enumerate(0, seqs.Range(10, 60, 10))

	var slice [][2]int
	transform2.ForEach(seqs.Enumerate(0, seqs.Range(10, 60, 10)), func(k, v int) {
		slice = append(slice, [2]int{k, v})
	})

	got := seq2s.FromSlice(slice)
	if !seq2s.Equal(got, want) {
		t.Errorf("ForEach(%s, %s) = %s, want %s",
			seq2s.String(seq), "append()", seq2s.String(got), seq2s.String(want),
		)
	}
}

func TestMap(t *testing.T) {
	want := seq2s.FromSlice([][2]int{{0, 100}, {1, 200}, {2, 300}, {3, 400}, {4, 500}})
	seq := seqs.Enumerate(0, seqs.Range(1, 6))

	valTimes100 := func(k, v int) (int, int) {
		return k, v * 100
	}
	got := transform2.Map(seq, valTimes100)

	if !seq2s.Equal(got, want) {
		t.Errorf("Map(%s, %s) = %s, want %s",
			seq2s.String(seq), "valTimes100()", seq2s.String(got), seq2s.String(want),
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

	seq := seqs.Enumerate(0, seqs.Range(1, 6))
	for _, tc := range tests {
		got = false
		for _, v := range transform2.OnEmpty(seq, callback) {
			if v == tc.i {
				break
			}
		}
		if got != tc.want {
			t.Errorf("OnEmpty(%s, %s) = %v, want %v",
				seq2s.String(seq), "callback()", got, tc.want,
			)
		}
	}
}

func TestReduce(t *testing.T) {
	tests := []struct {
		vals         [][2]int
		want1, want2 int
	}{
		{
			[][2]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}, {6, 7}, {7, 8}, {8, 9}, {9, 10}},
			55, 45,
		},
		{
			[][2]int{{0, 10}, {1, 20}, {2, 30}, {3, 40}, {4, 50}},
			150, 10,
		},
	}

	sum := func(a, b int) int {
		return a + b
	}

	for _, tc := range tests {
		seq := seq2s.FromSlice(tc.vals)
		got := transform2.Reduce(seq, sum)
		if got != tc.want1 {
			t.Errorf("Reduce(%s, %s) = %d, want %d",
				seq2s.String(seq), "sum()", got, tc.want1,
			)
		}

		seq = transform2.SwapKV(seq2s.FromSlice(tc.vals))
		got = transform2.Reduce(seq, sum)
		if got != tc.want2 {
			t.Errorf("Reduce(%s, %s) = %d, want %d",
				seq2s.String(seq), "sum()", got, tc.want2,
			)
		}
	}
}

func TestRotate(t *testing.T) {
	tests := []struct {
		n    int
		want [][2]int
	}{
		{
			2, [][2]int{{2, 30}, {3, 40}, {4, 50}, {0, 10}, {1, 20}},
		},
		{
			7, [][2]int{{0, 10}, {1, 20}, {2, 30}, {3, 40}, {4, 50}},
		},
	}

	for _, tc := range tests {
		want := seq2s.FromSlice(tc.want)
		seq := seqs.Enumerate(0, seqs.Range(10, 60, 10))
		got := transform2.Rotate(tc.n, seq)
		if !seq2s.Equal(got, want) {
			t.Errorf("Rotate(%d, %s) = %s, want %s",
				tc.n, seq2s.String(seq), seq2s.String(got), seq2s.String(want),
			)
		}
	}
}

func TestSwapKV(t *testing.T) {
	seq1 := slices.Values([]int{1, 2, 3, 4, 5})
	seq2 := slices.Values([]string{"one", "two", "three", "four", "five"})
	want := seqs.Zip(seq1, seq2)
	seq := seqs.Zip(seq2, seq1)
	got := transform2.SwapKV(seq)

	if !seq2s.Equal(got, want) {
		t.Errorf("SwapKV(%s) = %s, want %s",
			seq2s.String(seq), seq2s.String(got), seq2s.String(want))
	}
}

func TestTakeWhile(t *testing.T) {
	tests := []struct {
		predicate func(int, int) bool
		predName  string
		want      [][2]int
	}{
		{
			func(k, v int) bool {
				return k <= 2
			},
			"v <= 2",
			[][2]int{{0, 10}, {1, 20}, {2, 30}},
		},
		{
			func(k, v int) bool {
				return k > 2
			},
			"k > 2",
			[][2]int{},
		},
	}

	for _, tc := range tests {
		want := seq2s.FromSlice(tc.want)
		seq := seqs.Enumerate(0, seqs.Range(10, 60, 10))
		got := transform2.TakeWhile(seq, tc.predicate)
		if !seq2s.Equal(got, want) {
			t.Errorf("TakeWhile(%s, %s) = %s, want %s",
				seq2s.String(seq), tc.predName, seq2s.String(got), seq2s.String(want),
			)
		}
	}
}

func TestWith(t *testing.T) {
	tests := []struct {
		vals [][2]int
		want []int
	}{
		{
			[][2]int{{0, 10}, {1, 20}, {2, 30}, {3, 40}, {4, 50}},
			[]int{10, 30, 50},
		},
		{
			[][2]int{},
			[]int{},
		},
	}

	var slice []int
	takeEvens := func(k, v int) {
		if k%2 == 0 {
			slice = append(slice, v)
		}
	}
	for _, tc := range tests {
		slice = []int{}
		seq := seq2s.FromSlice(tc.vals)
		want := slices.Values(tc.want)
		transform2.With(seq, takeEvens)(func(k, v int) bool {
			return true
		})
		got := slices.Values(slice)
		if !seqs.Equal(got, want) {
			t.Errorf("With(%s, %s) = %v, want %v",
				seq2s.String(seq2s.FromSlice(tc.vals)), "takeEvens()", seqs.String(got), seqs.String(want),
			)
		}
	}
}
