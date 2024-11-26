package set_test

import (
	"slices"
	"testing"

	"github.com/elordeiro/goext/containers/set"
	"github.com/elordeiro/goext/seqs"
)

func TestNewSet(t *testing.T) {
	tests := []struct {
		vals []int
		want []int
	}{
		{[]int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{1}, []int{1}},
		{[]int{}, []int{}},
	}

	for _, tc := range tests {
		s := set.New(tc.vals...)
		got := s.All()
		for v := range got {
			if !slices.Contains(tc.want, v) {
				t.Errorf("New(%v) = %v, want %v", tc.vals, seqs.String(got), tc.want)
				break
			}
		}
	}
}

func TestAdd(t *testing.T) {
	tests := []struct {
		init []int
		add  []int
		want []int
	}{
		{[]int{}, []int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{1, 2, 3}, []int{4, 5, 6}, []int{1, 2, 3, 4, 5, 6}},
		{[]int{1, 2, 3}, []int{4}, []int{1, 2, 3, 4}},
		{[]int{}, []int{}, []int{}},
	}

	for _, tc := range tests {
		s := set.New(tc.init...)
		s.Add(tc.add...)
		got := s.All()
		for v := range got {
			if !slices.Contains(tc.want, v) {
				t.Errorf("{%v}.Add(%v) = %v, want %v", tc.init, tc.add, seqs.String(got), tc.want)
				break
			}
		}
	}
}

func TestRemove(t *testing.T) {
	tests := []struct {
		init []int
		del  int
		want []int
	}{
		{[]int{1, 2, 3}, 1, []int{2, 3}},
		{[]int{1, 2, 3}, 2, []int{1, 3}},
		{[]int{1, 2, 3}, 3, []int{1, 2}},
		{[]int{1, 2, 3}, 4, []int{1, 2, 3}},
	}

	for _, tc := range tests {
		s := set.New(tc.init...)
		s.Remove(tc.del)
		got := s.All()
		for v := range got {
			if !slices.Contains(tc.want, v) {
				t.Errorf("{%v}.Remove(%v) = %v, want %v", tc.init, tc.del, seqs.String(got), tc.want)
				break
			}
		}
	}
}

func TestContains(t *testing.T) {
	tests := []struct {
		init []int
		val  int
		want bool
	}{
		{[]int{1, 2, 3}, 1, true},
		{[]int{1, 2, 3}, 2, true},
		{[]int{1, 2, 3}, 3, true},
		{[]int{1, 2, 3}, 4, false},
		{[]int{}, 1, false},
	}

	for _, tc := range tests {
		s := set.New(tc.init...)
		got := s.Contains(tc.val)
		if got != tc.want {
			t.Errorf("{%v}.Contains(%v) = %v, want %v", tc.init, tc.val, got, tc.want)
		}
	}
}

func TestIsEmpty(t *testing.T) {
	tests := []struct {
		init []int
		want bool
	}{
		{[]int{1, 2, 3}, false},
		{[]int{1}, false},
		{[]int{}, true},
	}

	for _, tc := range tests {
		s := set.New(tc.init...)
		got := s.IsEmpty()
		if got != tc.want {
			t.Errorf("{%v}.IsEmpty() = %v, want %v", tc.init, got, tc.want)
		}
	}
}

func TestEqual(t *testing.T) {
	tests := []struct {
		s1   []int
		s2   []int
		want bool
	}{
		{[]int{1, 2, 3}, []int{1, 2, 3}, true},
		{[]int{1, 2, 3}, []int{3, 2, 1}, true},
		{[]int{1, 2, 3}, []int{1, 2, 3, 4}, false},
		{[]int{1, 2, 3}, []int{1, 2}, false},
		{[]int{}, []int{}, true},
	}

	for _, tc := range tests {
		s1 := set.New(tc.s1...)
		s2 := set.New(tc.s2...)
		got := s1.Equal(s2)
		if got != tc.want {
			t.Errorf("{%v}.Equal(%v) = %v, want %v", tc.s1, tc.s2, got, tc.want)
		}
	}
}

func TestUnion(t *testing.T) {
	tests := []struct {
		s1   []int
		s2   []int
		want []int
	}{
		{[]int{1, 2, 3}, []int{3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{[]int{1, 2, 3}, []int{4, 5, 6}, []int{1, 2, 3, 4, 5, 6}},
		{[]int{1, 2, 3}, []int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{}, []int{}, []int{}},
	}

	for _, tc := range tests {
		s1 := set.New(tc.s1...)
		s2 := set.New(tc.s2...)
		union := s1.Union(s2)
		got := union.All()
		for v := range got {
			if !slices.Contains(tc.want, v) {
				t.Errorf("{%v}.Union(%v) = %v, want %v", tc.s1, tc.s2, seqs.String(got), tc.want)
				break
			}
		}
	}
}

func TestIntersection(t *testing.T) {
	tests := []struct {
		s1   []int
		s2   []int
		want []int
	}{
		{[]int{1, 2, 3}, []int{3, 4, 5}, []int{3}},
		{[]int{1, 2, 3}, []int{4, 5, 6}, []int{}},
		{[]int{1, 2, 3}, []int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{}, []int{}, []int{}},
	}

	for _, tc := range tests {
		s1 := set.New(tc.s1...)
		s2 := set.New(tc.s2...)
		intersection := s1.Intersection(s2)
		got := intersection.All()
		for v := range got {
			if !slices.Contains(tc.want, v) {
				t.Errorf("{%v}.Intersection(%v) = %v, want %v", tc.s1, tc.s2, seqs.String(got), tc.want)
				break
			}
		}
	}
}

func TestDifference(t *testing.T) {
	tests := []struct {
		s1   []int
		s2   []int
		want []int
	}{
		{[]int{1, 2, 3}, []int{3, 4, 5}, []int{1, 2}},
		{[]int{1, 2, 3}, []int{4, 5, 6}, []int{1, 2, 3}},
		{[]int{1, 2, 3}, []int{1, 2, 3}, []int{}},
		{[]int{}, []int{}, []int{}},
	}

	for _, tc := range tests {
		s1 := set.New(tc.s1...)
		s2 := set.New(tc.s2...)
		difference := s1.Difference(s2)
		got := difference.All()
		for v := range got {
			if !slices.Contains(tc.want, v) {
				t.Errorf("{%v}.Difference(%v) = %v, want %v", tc.s1, tc.s2, seqs.String(got), tc.want)
				break
			}
		}
	}
}
