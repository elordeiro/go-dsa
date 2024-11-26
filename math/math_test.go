package math_test

import (
	"testing"

	"github.com/elordeiro/goext/math"
	"github.com/elordeiro/goext/seqs"
)

func TestSum(t *testing.T) {
	want := 15
	seq := seqs.Range(1, 6)
	got := math.Sum(seq)

	if got != want {
		t.Errorf("Sum(%s) = %d want %d", seqs.String(seq), want, got)
	}
}

func TestProduct(t *testing.T) {
	want := 120
	seq := seqs.Range(1, 6)
	got := math.Product(seq)

	if got != want {
		t.Errorf("Product(%s) = %d want %d", seqs.String(seq), want, got)
	}
}

func TestMin(t *testing.T) {
	tests := []struct {
		i, j, want int
	}{
		{1, 5, 1},
		{-5, -1, -5},
	}

	for _, tc := range tests {
		seq := seqs.Range(tc.i, tc.j)
		got := math.Min(seq)

		if got != tc.want {
			t.Errorf("Min(%s) = %d want %d", seqs.String(seq), got, tc.want)
		}
	}
}

func TestMax(t *testing.T) {
	tests := []struct {
		i, j, want int
	}{
		{1, 6, 5},
		{-5, -1, -2},
	}

	for _, tc := range tests {
		seq := seqs.Range(tc.i, tc.j)
		got := math.Max(seq)

		if got != tc.want {
			t.Errorf("Min(%s) = %d want %d", seqs.String(seq), got, tc.want)
		}
	}
}
