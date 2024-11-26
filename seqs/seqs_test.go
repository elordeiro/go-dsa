package seqs_test

import (
	"bytes"
	"fmt"
	"iter"
	"log"
	"os"
	"slices"
	"strings"
	"testing"

	"github.com/elordeiro/goext/containers/stack"
	"github.com/elordeiro/goext/seq2s"
	"github.com/elordeiro/goext/seqs"
)

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
		got := seqs.Equal(seqs.Range(tc.i, tc.j), seqs.Range(tc.k, tc.l))
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
		got := seqs.EqualFunc(seqs.Range(tc.i, tc.j), seqs.Range(tc.k, tc.l), func(i, j int) bool {
			return i == j
		})
		if got != tc.want {
			t.Errorf("Equal(%d, %d, %d, %d) = %t, want %t", tc.i, tc.j, tc.k, tc.l, got, tc.want)
		}
	}
}

func TestEqualUnordered(t *testing.T) {
	tests := []struct {
		i, j, k, l int
		want       bool
	}{
		{0, 5, 0, 5, true},
		{0, 5, 4, -1, true},
		{4, -1, 0, 5, true},
		{1, 5, 0, 5, false},
		{0, 5, 1, 5, false},
		{0, 5, 0, 4, false},
		{0, 0, 0, 0, true},
		{1, 1, 1, 1, true},
	}
	for _, tc := range tests {
		got := seqs.EqualUnordered(seqs.Range(tc.i, tc.j), seqs.Range(tc.k, tc.l))
		if got != tc.want {
			t.Errorf("Equal(%d, %d, %d, %d) = %t, want %t", tc.i, tc.j, tc.k, tc.l, got, tc.want)
		}
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
		got := seqs.Len(seqs.Range(tc.i, tc.j))
		if got != tc.want {
			t.Errorf("Len(%d, %d) = %d, want %d", tc.i, tc.j, got, tc.want)
		}
	}
}

func TestString(t *testing.T) {
	want := "=>[1 2 3 4 5]"
	got := fmt.Sprint(seqs.String(seqs.Range(1, 6)))

	if got != want {
		t.Errorf("want %s but got %s", got, want)
	}
}

func TestChain(t *testing.T) {
	want := slices.Values([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	seq1 := seqs.Range(1, 6)
	seq2 := seqs.Range(6, 10)
	got := seqs.Chain(seq1, seq2)

	if !seqs.Equal(got, want) {
		t.Errorf("Chain(%s, %s) = %v, want %v",
			seqs.String(seq1), seqs.String(seq2), seqs.String(got), seqs.String(want))
	}
}

func TestCollect(t *testing.T) {
	want := []int{10, 20, 30, 40, 50}
	seq := seqs.Range(10, 60, 10)
	got := seqs.Collect(seq)

	if !slices.Equal(got, want) {
		t.Errorf("Collect(%s) = %v, want %v", seqs.String(seq), got, want)
	}
}

func TestCount(t *testing.T) {
	tests := []struct {
		start, step int
		want        []int
	}{
		{0, 0, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{0, 2, []int{0, 2, 4, 6, 8, 10, 12, 14, 16, 18}},
		{10, -1, []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}},
	}

	var slice []int
	for _, tc := range tests {
		slice = []int{}
		var seq iter.Seq[int]
		if tc.step == 0 {
			seq = seqs.Count(tc.start)
		} else {
			seq = seqs.Count(tc.start, tc.step)
		}

		for v := range seq {
			slice = append(slice, v)
			if len(slice) == 10 {
				break
			}
		}

		want := slices.Values(tc.want)
		got := slices.Values(slice)
		if !seqs.Equal(got, want) {
			t.Errorf("Count(%d, %d) = %v, want %v", tc.start, tc.step, seqs.String(got), seqs.String(want))
		}
	}
}

func TestCycle(t *testing.T) {
	want := []int{1, 2, 3, 1, 2, 3, 1, 2, 3}
	got := []int{}

	for v := range seqs.Cycle(seqs.Range(1, 4)) {
		if len(got) == 9 {
			break
		}
		got = append(got, v)
	}

	if !slices.Equal(got, want) {
		t.Errorf("Cycle(%s) = %v, want %v", seqs.String(seqs.Range(1, 4)), got, want)
	}
}

func TestEnumerate(t *testing.T) {
	want := seq2s.FromSlice([][2]int{{2, 10}, {3, 20}, {4, 30}, {5, 40}, {6, 50}})
	seq := seqs.Range(10, 60, 10)
	got := seqs.Enumerate(2, seq)

	if !seq2s.Equal(got, want) {
		t.Errorf("Enumerate(2, %s) = %s, want %s",
			seqs.String(seq), seq2s.String(got), seq2s.String(want),
		)
	}
}

func TestMultiUse(t *testing.T) {
	want := slices.Values([]int{50, 40, 30, 20, 10})
	stk := stack.New(10, 20, 30, 40, 50)
	got := seqs.MultiUse(stk.Drain())

	if stk.Len() != 0 {
		t.Errorf("stk.Len() = %d, want 0", stk.Len())
	}

	if !seqs.Equal(got, want) {
		t.Errorf("MultiUse(%s) = %v, want %v", seqs.String(want), seqs.String(got), seqs.String(want))
	}

	if !seqs.Equal(got, want) {
		t.Errorf("MultiUse(%s) = %v, want %v", seqs.String(want), seqs.String(got), seqs.String(want))
	}
}

func TestRange(t *testing.T) {
	tests := []struct {
		vals, want []int
		err        string
	}{
		{
			[]int{10},
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			"",
		},
		{
			[]int{10, 20},
			[]int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19},
			"",
		},
		{
			[]int{0, 10, 2},
			[]int{0, 2, 4, 6, 8},
			"",
		},
		{
			[]int{10, 20, 2},
			[]int{10, 12, 14, 16, 18},
			"",
		},
		{
			[]int{10, 0},
			[]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			"",
		},
		{
			[]int{20, 10},
			[]int{20, 19, 18, 17, 16, 15, 14, 13, 12, 11},
			"",
		},
		{
			[]int{10, 0, -2},
			[]int{10, 8, 6, 4, 2},
			"",
		},
		{
			[]int{20, 10, -2},
			[]int{20, 18, 16, 14, 12},
			"",
		},
		{
			[]int{0, 10, -1},
			[]int{},
			"empty iterator in Range(); start < end && step < 0",
		},
		{
			[]int{10, 20, -1},
			[]int{},
			"empty iterator in Range(); start < end && step < 0",
		},
		{
			[]int{10, 0, 1},
			[]int{},
			"empty iterator in Range(); start > end && step >= 0",
		},
		{
			[]int{20, 10, 1},
			[]int{},
			"empty iterator in Range(); start > end && step >= 0",
		},
		{
			[]int{0, 10, 0},
			[]int{},
			"infinite loop in Range(); step == 0",
		},
		{
			[]int{10, 0, 0},
			[]int{},
			"infinite loop in Range(); step == 0",
		},
	}

	for _, tc := range tests {
		stdout := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w
		log.SetOutput(os.Stdout)

		got := seqs.Range(tc.vals[0], tc.vals[1:]...)

		w.Close()
		var buf bytes.Buffer
		buf.ReadFrom(r)

		err := buf.String()
		os.Stdout = stdout

		if tc.err != "" {
			if err == "" {
				t.Errorf("Range(%v) = %v, want %v", tc.vals, err, tc.err)
				continue
			}

			var prefix string
			if strings.HasPrefix(tc.err, "infinite") {
				prefix = "infinite"
			} else {
				prefix = "empty"
			}

			err = err[strings.Index(err, prefix) : len(err)-1]
			if tc.err != err {
			}
			continue
		}

		if !seqs.Equal(got, slices.Values(tc.want)) {
			t.Errorf("Range(%v) = %s, want %s", tc.vals, seqs.String(got), seqs.String(slices.Values(tc.want)))
		}
	}

}

func TestRepeat(t *testing.T) {
	tests := []struct {
		val, n int
	}{
		{5, 5},
		{5, -1},
	}

	for _, tc := range tests {
		count := 0
		for v := range seqs.Repeat(tc.val, tc.n) {
			if v != tc.val {
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
		want       []int
	}{
		{
			0, 4,
			[]int{10, 20, 30, 40},
		},
		{
			1, 4,
			[]int{20, 30, 40},
		},
		{
			0, 10,
			[]int{10, 20, 30, 40, 50},
		},
		{
			0, 0,
			[]int{},
		},
		{
			1, 0,
			[]int{20, 30, 40, 50},
		},
		{
			2, 0,
			[]int{30, 40, 50},
		},
		{
			3, 1,
			[]int{40, 50},
		},
	}

	seq := seqs.Range(10, 60, 10)
	for _, tc := range tests {
		got := seqs.SeqRange(tc.start, tc.end, seq)
		if !seqs.Equal(got, slices.Values(tc.want)) {
			t.Errorf("SeqRange(%d, %d, %s) = %s, want %s",
				tc.start, tc.end, seqs.String(seq), seqs.String(got), seqs.String(slices.Values(tc.want)),
			)
		}
	}
}

func TestZip(t *testing.T) {
	want := seq2s.FromSlice([][2]int{{10, 20}, {20, 30}, {30, 40}, {40, 50}})
	got := seqs.Zip(seqs.Range(10, 50, 10), seqs.Range(20, 60, 10))

	if !seq2s.Equal(got, want) {
		t.Errorf("Zip(%s, %s) = %s, want %s",
			seqs.String(seqs.Range(10, 50, 10)),
			seqs.String(seqs.Range(20, 60, 10)),
			seq2s.String(got), seq2s.String(want),
		)
	}
}
