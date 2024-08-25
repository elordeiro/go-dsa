package set_test

import (
	"testing"

	"github.com/elordeiro/go/container/set"
)

func TestNewSet(t *testing.T) {
	s := set.NewSet(1, 2, 3)
	if s.Len() != 3 {
		t.Errorf("Expected 3, got %d", s.Len())
	}

	if !s.Contains(1) || !s.Contains(2) || !s.Contains(3) {
		t.Errorf("Expected true, got false")
	}
}

func TestAdd(t *testing.T) {
	s := set.NewSet(1, 2, 3)
	s.Add(4)
	if s.Len() != 4 {
		t.Errorf("Expected 4, got %d", s.Len())
	}

	if !s.Contains(4) {
		t.Errorf("Expected true, got false")
	}
}

func TestRemove(t *testing.T) {
	s := set.NewSet(1, 2, 3)
	s.Remove(2)
	if s.Len() != 2 {
		t.Errorf("Expected 2, got %d", s.Len())
	}

	if s.Contains(2) {
		t.Errorf("Expected false, got true")
	}
}

func TestContains(t *testing.T) {
	s := set.NewSet(1, 2, 3)
	if !s.Contains(1) || !s.Contains(2) || !s.Contains(3) {
		t.Errorf("Expected true, got false")
	}
}

func TestIsEmpty(t *testing.T) {
	s := set.NewSet[int]()
	if !s.IsEmpty() {
		t.Errorf("Expected true, got false")
	}

	s.Add(1)
	if s.IsEmpty() {
		t.Errorf("Expected false, got true")
	}
}

func TestUnion(t *testing.T) {
	s1 := set.NewSet(1, 2, 3)
	s2 := set.NewSet(3, 4, 5)
	union := s1.Union(s2)
	if union.Len() != 5 {
		t.Errorf("Expected 5, got %d", union.Len())
	}

	if !union.Contains(1) || !union.Contains(2) || !union.Contains(3) || !union.Contains(4) || !union.Contains(5) {
		t.Errorf("Expected true, got false")
	}
}

func TestIntersection(t *testing.T) {
	s1 := set.NewSet(1, 2, 3)
	s2 := set.NewSet(3, 4, 5)
	intersection := s1.Intersection(s2)
	if intersection.Len() != 1 {
		t.Errorf("Expected 1, got %d", intersection.Len())
	}

	if !intersection.Contains(3) {
		t.Errorf("Expected true, got false")
	}
}

func TestDifference(t *testing.T) {
	s1 := set.NewSet(1, 2, 3)
	s2 := set.NewSet(3, 4, 5)
	difference := s1.Difference(s2)
	if difference.Len() != 2 {
		t.Errorf("Expected 2, got %d", difference.Len())
	}

	if !difference.Contains(1) || !difference.Contains(2) {
		t.Errorf("Expected true, got false")
	}
}

func TestEnumerate(t *testing.T) {
	result := []struct {
		i int
		v string
	}{}
	s := set.NewSet("a", "b", "c")
	for i, item := range s.Enumerate(0) {
		result = append(result, struct {
			i int
			v string
		}{i, item})
	}

	if len(result) != 3 {
		t.Errorf("Expected 3, got %d", len(result))
	}
}
