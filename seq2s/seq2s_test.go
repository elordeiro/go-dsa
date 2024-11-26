package seq2s_test

import (
	"fmt"
	"slices"
	"testing"

	"github.com/elordeiro/goext/containers/stack"
	"github.com/elordeiro/goext/containers/tuples"
	"github.com/elordeiro/goext/seq2s"
	"github.com/elordeiro/goext/seqs"
	"github.com/elordeiro/goext/seqs/transform"
)

func TestNilSeq2(t *testing.T) {
	seq2s.NilSeq2[int, int](nil)
}

func TestEqual(t *testing.T) {
	tests := []struct {
		i, j, k, l int
		want       bool
	}{
		{0, 5, 0, 5, true},
		{1, 5, 0, 5, false},
		{0, 5, 1, 5, false},
		{0, 5, 0, 4, false},
		{0, 0, 0, 0, true},
		{1, 1, 1, 1, true},
	}
	for _, tc := range tests {
		got := seq2s.Equal(seqs.Enumerate(0, seqs.Range(tc.i, tc.j)), seqs.Enumerate(0, seqs.Range(tc.k, tc.l)))
		if got != tc.want {
			t.Errorf("Equal(%d, %d, %d, %d) = %t, want %t", tc.i, tc.j, tc.k, tc.l, got, tc.want)
		}
	}
}

func TestEqualFunc(t *testing.T) {
	tests := []struct {
		i, j, k, l int
		want       bool
	}{
		{0, 5, 0, 5, true},
		{1, 5, 0, 5, false},
		{0, 5, 1, 5, false},
		{0, 5, 0, 4, false},
		{0, 0, 0, 0, true},
		{1, 1, 1, 1, true},
	}
	for _, tc := range tests {
		got := seq2s.EqualFunc(seqs.Enumerate(0, seqs.Range(tc.i, tc.j)), seqs.Enumerate(0, seqs.Range(tc.k, tc.l)), func(p1, p2 tuples.Pair[int, int]) bool {
			return p1 == p2
		})
		if got != tc.want {
			t.Errorf("Equal(%d, %d, %d, %d) = %t, want %t", tc.i, tc.j, tc.k, tc.l, got, tc.want)
		}
	}
}

func TestEqualUnordered(t *testing.T) {
	tests := []struct {
		seq1, seq2 []tuples.Pair[int, string]
		want       bool
	}{
		{
			tuples.Pairs(tuples.NewPair(0, "a"), tuples.NewPair(1, "b"), tuples.NewPair(2, "c")),
			tuples.Pairs(tuples.NewPair(0, "a"), tuples.NewPair(1, "b"), tuples.NewPair(2, "c")),
			true,
		},
		{
			tuples.Pairs(tuples.NewPair(0, "a"), tuples.NewPair(1, "b"), tuples.NewPair(2, "c")),
			tuples.Pairs(tuples.NewPair(1, "b"), tuples.NewPair(2, "c"), tuples.NewPair(0, "a")),
			true,
		},
		{
			tuples.Pairs(tuples.NewPair(0, "a"), tuples.NewPair(1, "b"), tuples.NewPair(2, "c")),
			tuples.Pairs(tuples.NewPair(0, "a"), tuples.NewPair(1, "b")),
			false,
		},
		{
			tuples.Pairs(tuples.NewPair(0, "a"), tuples.NewPair(1, "b"), tuples.NewPair(2, "c")),
			tuples.Pairs(tuples.NewPair(0, "a"), tuples.NewPair(1, "b"), tuples.NewPair(2, "d")),
			false,
		},
		{
			tuples.Pairs(tuples.NewPair(0, "a"), tuples.NewPair(1, "b"), tuples.NewPair(2, "c")),
			tuples.Pairs(
				tuples.NewPair(0, "a"), tuples.NewPair(1, "b"), tuples.NewPair(2, "c"), tuples.NewPair(3, "d"),
			),
			false,
		},
		{
			tuples.Pairs[int, string](),
			tuples.Pairs[int, string](),
			true,
		},
	}
	for _, tc := range tests {
		seq1 := transform.Unpair(slices.Values(tc.seq1))
		seq2 := transform.Unpair(slices.Values(tc.seq2))
		got := seq2s.EqualUnordered(
			seq1, seq2,
		)
		if got != tc.want {
			t.Errorf("EqualUnordered(%s, %s) = %t, want %t", seq2s.String(seq1), seq2s.String(seq2), got, tc.want)
		}
	}
}

func TestFromSlice(t *testing.T) {
	want := seqs.Enumerate(0, seqs.Range(10, 60, 10))
	slice := [][2]int{{0, 10}, {1, 20}, {2, 30}, {3, 40}, {4, 50}}
	got := seq2s.FromSlice(slice)
	if !seq2s.Equal(got, want) {
		t.Errorf("FromSlice(%v) = %s, want %s", slice, seq2s.String(got), seq2s.String(want))
	}
}

func TestKeys(t *testing.T) {
	want := slices.Values([]int{0, 1, 2, 3, 4})
	seq2 := seqs.Enumerate(0, seqs.Range(10, 60, 10))
	got := seq2s.Keys(seq2)

	if !seqs.Equal(got, want) {
		t.Errorf("Keys(%s) = %s, want %s", seq2s.String(seq2), seqs.String(got), seqs.String(want))
	}
}

func TestLen(t *testing.T) {
	tests := []struct {
		i, j, want int
	}{
		{0, 5, 5},
		{1, 5, 4},
		{5, 0, 5},
		{5, 1, 4},
		{0, 0, 0},
		{1, 1, 0},
	}
	for _, tc := range tests {
		got := seq2s.Len(seqs.Enumerate(0, seqs.Range(tc.i, tc.j)))
		if got != tc.want {
			t.Errorf("Len(%d, %d) = %d, want %d", tc.i, tc.j, got, tc.want)
		}
	}
}

func TestString(t *testing.T) {
	want := "=>[(0 1) (1 2) (2 3) (3 4) (4 5)]"
	seq := seqs.Enumerate(0, seqs.Range(1, 6))
	got := fmt.Sprint(seq2s.String(seq))

	if got != want {
		t.Errorf("String(%s) = %s, want %s", seq2s.String(seq), got, want)
	}
}

func TestValues(t *testing.T) {
	want := slices.Values([]int{10, 20, 30, 40, 50})
	seq2 := seqs.Enumerate(0, seqs.Range(10, 60, 10))
	got := seq2s.Values(seq2)

	if !seqs.Equal(got, want) {
		t.Errorf("Values(%s) = %s, want %s", seq2s.String(seq2), seqs.String(got), seqs.String(want))
	}
}

func TestChain(t *testing.T) {
	want := seq2s.FromSlice(
		[][2]int{{0, 10}, {1, 20}, {2, 30}, {3, 40}, {4, 50}, {5, 60}, {6, 70}, {7, 80}, {8, 90}},
	)
	seq1 := seqs.Enumerate(0, seqs.Range(10, 60, 10))
	seq2 := seqs.Enumerate(5, seqs.Range(60, 100, 10))
	got := seq2s.Chain(seq1, seq2)

	if !seq2s.Equal(got, want) {
		t.Errorf("Chain2(%s, %s) = %s, want %s",
			seq2s.String(seq1), seq2s.String(seq2), seq2s.String(got), seq2s.String(want))
	}
}

func TestCollect(t *testing.T) {
	want := [][2]int{{0, 10}, {1, 20}, {2, 30}, {3, 40}, {4, 50}}
	seq2 := seqs.Enumerate(0, seqs.Range(10, 60, 10))
	got := seq2s.Collect(seq2)

	if !slices.EqualFunc(got, want, func(s1, s2 [2]int) bool {
		return s1 == s2
	}) {
		t.Errorf("Collect2(%s) = %v, want %v", seq2s.String(seq2), got, want)
	}
}

func CollectPairs(t *testing.T) {
	want := tuples.Pairs(
		tuples.NewPair(0, 10),
		tuples.NewPair(1, 20),
		tuples.NewPair(2, 30),
		tuples.NewPair(3, 40),
		tuples.NewPair(4, 50),
	)
	seq2 := seqs.Enumerate(0, seqs.Range(10, 60, 10))
	got := seq2s.CollectPairs(seq2)

	if !slices.Equal(got, want) {
		t.Errorf("CollectPairs(%s) = %s, want %s", seq2s.String(seq2), got, want)
	}
}

func TestCycle(t *testing.T) {
	want := [][]int{{0, 10}, {1, 20}, {2, 30}, {0, 10}, {1, 20}, {2, 30}, {0, 10}, {1, 20}, {2, 30}}
	got := [][]int{}
	seq := seqs.Enumerate(0, seqs.Range(10, 40, 10))
	for i, v := range seq2s.Cycle(seq) {
		if len(got) == 9 {
			break
		}
		got = append(got, []int{i, v})
	}

	if !slices.EqualFunc(got, want, func(s1, s2 []int) bool {
		return slices.Equal(s1, s2)
	}) {
		t.Errorf("Cycle2(%s) = %v, want %v", seq2s.String(seq), got, want)
	}
}

func TestEnumerate(t *testing.T) {
	want := slices.All(
		[]tuples.Pair[string, int]{
			tuples.NewPair("a", 10),
			tuples.NewPair("b", 20),
			tuples.NewPair("c", 30),
			tuples.NewPair("d", 40),
			tuples.NewPair("e", 50),
		},
	)
	seq := slices.Values(
		tuples.Pairs(
			tuples.NewPair("a", 10),
			tuples.NewPair("b", 20),
			tuples.NewPair("c", 30),
			tuples.NewPair("d", 40),
			tuples.NewPair("e", 50),
		),
	)
	got := seq2s.Enumerate(0, transform.Unpair(seq))

	if !seq2s.Equal(got, want) {
		t.Errorf("Enumerate2(0, %s) = %s, want %s",
			seqs.String(seq), seq2s.String(got), seq2s.String(want),
		)
	}
}

func TestMultiUse(t *testing.T) {
	want := seq2s.FromSlice([][2]int{{0, 50}, {1, 40}, {2, 30}, {3, 20}, {4, 10}})
	stk := stack.New(10, 20, 30, 40, 50)
	got := seq2s.MultiUse(seqs.Enumerate(0, stk.Drain()))

	if stk.Len() != 0 {
		t.Errorf("stk.Len() = %d, want 0", stk.Len())
	}

	if !seq2s.Equal(got, want) {
		t.Errorf("MultiUse(%s) = %v, want %v", seq2s.String(want), seq2s.String(got), seq2s.String(want))
	}

	if !seq2s.Equal(got, want) {
		t.Errorf("MultiUse(%s) = %v, want %v", seq2s.String(want), seq2s.String(got), seq2s.String(want))
	}
}

func TestRepeat(t *testing.T) {
	tests := []struct {
		key, val, n int
	}{
		{4, 5, 5},
		{4, 5, -1},
	}

	for _, tc := range tests {
		count := 0
		for k, v := range seq2s.Repeat(tc.key, tc.val, tc.n) {
			if k != tc.key || v != tc.val {
				t.Errorf("Repeat(%d, %d) = %d, want %d", tc.val, tc.n, v, tc.val)
			}
			count++
			if count == 100 {
				break
			}
		}

		if tc.n > 0 && count != tc.n {
			t.Errorf("Repeat(%d, %d) = %d, want %d", tc.val, tc.n, count, tc.n)
		} else if tc.n < 0 && count != 100 {
			t.Errorf("Repeat(%d, %d) = %d, want %d", tc.val, tc.n, count, 100)
		}
	}
}

func TestSeqRange(t *testing.T) {
	tests := []struct {
		start, end int
		want       [][2]int
	}{
		{
			0, 4,
			[][2]int{{0, 10}, {1, 20}, {2, 30}, {3, 40}},
		},
		{
			1, 4,
			[][2]int{{1, 20}, {2, 30}, {3, 40}},
		},
		{
			0, 10,
			[][2]int{{0, 10}, {1, 20}, {2, 30}, {3, 40}, {4, 50}},
		},
		{
			0, 0,
			[][2]int{},
		},
		{
			1, 0,
			[][2]int{{1, 20}, {2, 30}, {3, 40}, {4, 50}},
		},
	}

	seq2 := seqs.Enumerate(0, seqs.Range(10, 60, 10))
	for _, tc := range tests {
		got := seq2s.SeqRange(tc.start, tc.end, seq2)
		want := seq2s.FromSlice(tc.want)
		if !seq2s.Equal(got, want) {
			t.Errorf("SeqRange2(%d, %d, %s) = %s, want %s",
				tc.start, tc.end, seq2s.String(seq2), seq2s.String(got), seq2s.String(want),
			)
		}
	}
}

func TestZip(t *testing.T) {
	want := transform.Unpair(
		slices.Values(
			tuples.Pairs(
				tuples.NewPair(tuples.NewPair("a", 10), tuples.NewPair("b", 20)),
				tuples.NewPair(tuples.NewPair("c", 30), tuples.NewPair("d", 40)),
				tuples.NewPair(tuples.NewPair("e", 50), tuples.NewPair("f", 60)),
			),
		),
	)

	s1 := transform.Unpair(
		slices.Values(
			tuples.Pairs(
				tuples.NewPair("a", 10),
				tuples.NewPair("c", 30),
				tuples.NewPair("e", 50),
			),
		),
	)
	s2 := transform.Unpair(
		slices.Values(
			tuples.Pairs(
				tuples.NewPair("b", 20),
				tuples.NewPair("d", 40),
				tuples.NewPair("f", 60),
			),
		),
	)
	got := seq2s.Zip(s1, s2)

	if !seq2s.Equal(got, want) {
		t.Errorf("Zip2(%s, %s) = %s, want %s",
			seq2s.String(s1), seq2s.String(s2), seq2s.String(got), seq2s.String(want),
		)
	}
}
