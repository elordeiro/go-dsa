package tests

import (
	"testing"

	"github.com/elordeiro/go/dsa/set"
)

func TestSet(t *testing.T) {
	s := set.NewSet[int]()

	if !s.IsEmpty() {
		t.Error("Expected empty set")
	}

	s.Add(1)
	s.Add(2)
	s.Add(3)

	if s.IsEmpty() {
		t.Error("Expected non-empty set")
	}

	if !s.Contains(1) {
		t.Error("Expected set to contain 1")
	}

	if !s.Contains(2) {
		t.Error("Expected set to contain 2")
	}

	if !s.Contains(3) {
		t.Error("Expected set to contain 3")
	}

	s.Remove(1)

	if s.Contains(1) {
		t.Error("Expected set to not contain 1")
	}

	if !s.Contains(2) {
		t.Error("Expected set to contain 2")
	}

	if !s.Contains(3) {
		t.Error("Expected set to contain 3")
	}

	s.Remove(2)

	if s.Contains(2) {
		t.Error("Expected set to not contain 2")
	}

	if !s.Contains(3) {
		t.Error("Expected set to contain 3")
	}

	if s.Len() != 1 {
		t.Error("Expected set to have length 1")
	}

	s.Remove(3)

	if s.Contains(3) {
		t.Error("Expected set to not contain 3")
	}

	if !s.IsEmpty() {
		t.Error("Expected empty set")
	}
}

func TestSetUnion(t *testing.T) {
	s1 := set.NewSet(1, 2, 3)
	s2 := set.NewSet(3, 4, 5)

	union := s1.Union(&s2)

	if !union.Contains(1) {
		t.Error("Expected union to contain 1")
	}

	if !union.Contains(2) {
		t.Error("Expected union to contain 2")
	}

	if !union.Contains(3) {
		t.Error("Expected union to contain 3")
	}

	if !union.Contains(4) {
		t.Error("Expected union to contain 4")
	}

	if !union.Contains(5) {
		t.Error("Expected union to contain 5")
	}
}

func TestSetIntersection(t *testing.T) {
	s1 := set.NewSet(1, 2, 3)
	s2 := set.NewSet(3, 4, 5)

	intersection := s1.Intersection(&s2)

	if !intersection.Contains(3) {
		t.Error("Expected intersection to contain 3")
	}

	if intersection.Contains(1) {
		t.Error("Expected intersection to not contain 1")
	}

	if intersection.Contains(2) {
		t.Error("Expected intersection to not contain 2")
	}

	if intersection.Contains(4) {
		t.Error("Expected intersection to not contain 4")
	}

	if intersection.Contains(5) {
		t.Error("Expected intersection to not contain 5")
	}
}

func TestSetDifference(t *testing.T) {
	s1 := set.NewSet(1, 2, 3)
	s2 := set.NewSet(3, 4, 5)

	difference := s1.Difference(&s2)

	if !difference.Contains(1) {
		t.Error("Expected difference to contain 1")
	}

	if !difference.Contains(2) {
		t.Error("Expected difference to contain 2")
	}

	if difference.Contains(3) {
		t.Error("Expected difference to not contain 3")
	}

	if difference.Contains(4) {
		t.Error("Expected difference to not contain 4")
	}

	if difference.Contains(5) {
		t.Error("Expected difference to not contain 5")
	}
}

func TestSetString(t *testing.T) {
	s := set.NewSet(1, 2, 3)

	expected := "[1 2 3]"
	if s.String() != expected {
		t.Errorf("Expected %s, got %s", expected, s.String())
	}
}

func TestSetItems(t *testing.T) {
	s := set.NewSet(1, 2, 3)

	items := s.Items()

	if len(items) != 3 {
		t.Error("Expected 3 items")
	}

	if items[0] != 1 {
		t.Error("Expected item at index 0 to be 1")
	}

	if items[1] != 2 {
		t.Error("Expected item at index 1 to be 2")
	}

	if items[2] != 3 {
		t.Error("Expected item at index 2 to be 3")
	}
}

func TestSetIter(t *testing.T) {
	s := set.NewSet(1, 2, 3)

	expected := 6
	sum := 0
	for item := range s.Keys() {
		sum += item
	}

	if sum != expected {
		t.Errorf("Expected sum of %d, got %d", expected, sum)
	}
}
