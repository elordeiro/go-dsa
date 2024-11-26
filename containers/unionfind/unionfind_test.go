package unionfind_test

import (
	"testing"

	"github.com/elordeiro/goext/containers/set"
	"github.com/elordeiro/goext/containers/unionfind"
)

func TestUnionFind(t *testing.T) {
	uf := unionfind.New[int]()
	uf.MakeSet(1)
	uf.MakeSet(2)
	uf.MakeSet(3)
	uf.MakeSet(4)
	uf.MakeSet(5)

	uf.Union(1, 2)
	uf.Union(3, 4)
	uf.Union(2, 3)
	uf.Union(1, 3)

	if !uf.Connected(1, 4) {
		t.Errorf("Connected(1, 4) = false, want true")
	}

	if uf.Connected(1, 5) {
		t.Errorf("Connected(1, 5) = true, want false")
	}

	if !uf.Connected(2, 3) {
		t.Errorf("Connected(2, 3) = false, want true")
	}

	if uf.Connected(4, 5) {
		t.Errorf("Connected(4, 5) = true, want false")
	}

	groups := uf.All()
	group1 := set.New(1, 2, 3, 4)
	group2 := set.New(5)

	for group := range groups {
		if !group.Equal(group1) && !group.Equal(group2) {
			t.Errorf("Groups() = %v, want %v", group, group1.Union(group2))
		}
	}
}
