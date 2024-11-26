package search_test

import (
	"slices"
	"testing"

	"github.com/elordeiro/goext/algorithms/search"
	"github.com/elordeiro/goext/containers/tuples"
	"github.com/elordeiro/goext/seq2s"
	"github.com/elordeiro/goext/seq2s/transform2"
	"github.com/elordeiro/goext/seqs"
	"github.com/elordeiro/goext/seqs/transform"
)

func TestAll(t *testing.T) {
	tests := []struct {
		s         []int
		predicate func(i int) bool
		want      bool
	}{
		{[]int{1, 2, 3, 4, 5}, nil, true},
		{[]int{0, 0, 0, 0, 0}, nil, false},
		{[]int{0, 0, 1, 0, 0}, nil, false},
		{[]int{0, 1, 2, 3, 4, 5}, nil, false},
		{[]int{1, 2, 3, 4, 5}, func(i int) bool { return i < 5 }, false},
		{[]int{0, 2, 4, 6, 8}, func(i int) bool { return i%2 == 0 }, true},
	}

	for _, tc := range tests {
		seq := slices.Values(tc.s)

		var got bool
		if tc.predicate == nil {
			got = search.All(seq)
		} else {
			got = search.All(seq, tc.predicate)
		}

		if got != tc.want {
			t.Errorf("All(%s) = %t want %t", seqs.String(seq), got, tc.want)
		}
	}
}

func TestAny(t *testing.T) {
	tests := []struct {
		s         []int
		predicate func(i int) bool
		want      bool
	}{
		{[]int{1, 2, 3, 4, 5}, nil, true},
		{[]int{0, 0, 0, 0, 0}, nil, false},
		{[]int{0, 0, 1, 0, 0}, nil, true},
		{[]int{0, 1, 2, 3, 4, 5}, nil, true},
		{[]int{1, 2, 3, 4, 5}, func(i int) bool { return i < 5 }, true},
		{[]int{0, 2, 4, 6, 8}, func(i int) bool { return i%2 == 0 }, true},
	}

	for _, tc := range tests {
		seq := slices.Values(tc.s)

		var got bool
		if tc.predicate == nil {
			got = search.Any(seq)
		} else {
			got = search.Any(seq, tc.predicate)
		}

		if got != tc.want {
			t.Errorf("All(%s) = %t want %t", seqs.String(seq), got, tc.want)
		}
	}
}

func TestNone(t *testing.T) {
	tests := []struct {
		s         []int
		predicate func(i int) bool
		want      bool
	}{
		{[]int{1, 2, 3, 4, 5}, nil, false},
		{[]int{0, 0, 0, 0, 0}, nil, true},
		{[]int{0, 0, 1, 0, 0}, nil, false},
		{[]int{0, 1, 2, 3, 4, 5}, nil, false},
		{[]int{1, 2, 3, 4, 5}, func(i int) bool { return i < 5 }, false},
		{[]int{0, 2, 4, 6, 8}, func(i int) bool { return i%2 == 0 }, false},
	}

	for _, tc := range tests {
		seq := slices.Values(tc.s)

		var got bool
		if tc.predicate == nil {
			got = search.None(seq)
		} else {
			got = search.None(seq, tc.predicate)
		}

		if got != tc.want {
			t.Errorf("All(%s) = %t want %t", seqs.String(seq), got, tc.want)
		}
	}
}

func TestUnpair(t *testing.T) {
	seq1 := slices.Values([]int{1, 2, 3, 4, 5})
	seq2 := slices.Values([]string{"a", "b", "c", "d", "e"})
	want := seqs.Zip(seq1, seq2)
	seq := slices.Values(
		tuples.Pairs(
			tuples.NewPair(1, "a"),
			tuples.NewPair(2, "b"),
			tuples.NewPair(3, "c"),
			tuples.NewPair(4, "d"),
			tuples.NewPair(5, "e"),
		),
	)
	got := transform.Unpair(seq)

	if !seq2s.Equal(got, want) {
		t.Errorf("Unpair(%s) = %s, want %s", seqs.String(seq), seq2s.String(got), seq2s.String(want))
	}
}

func TestSwapKV(t *testing.T) {
	seq1 := slices.Values([]int{1, 2, 3, 4, 5})
	seq2 := slices.Values([]string{"a", "b", "c", "d", "e"})
	want := seqs.Zip(seq1, seq2)
	seq := seqs.Zip(seq2, seq1)
	got := transform2.SwapKV(seq)

	if !seq2s.Equal(got, want) {
		t.Errorf("SwapKV(%s) = %s, want %s", seq2s.String(seq), seq2s.String(got), seq2s.String(want))
	}
}
