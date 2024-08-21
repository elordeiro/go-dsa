package iters_test

import (
	"slices"
	"testing"

	itr "github.com/elordeiro/go/iters"
)

func TestSeq(t *testing.T) {
	expected := []int{10, 20, 30, 40, 50}
	result := []int{}

	for v := range itr.ToSeq(slices.All([]int{10, 20, 30, 40, 50})) {
		result = append(result, v)
	}

	if len(expected) != len(result) {
		t.Error("Expected and result slices should have the same length")
	}

	for i := range expected {
		if expected[i] != result[i] {
			t.Errorf("Expected %d but got %d", expected[i], result[i])
		}
	}
}
func TestSeq2(t *testing.T) {
	expected := []struct {
		index int
		value int
	}{
		{0, 10},
		{1, 20},
		{2, 30},
		{3, 40},
		{4, 50},
	}

	result := []struct {
		index int
		value int
	}{}

	for i, v := range itr.ToSeq2(slices.Values([]int{10, 20, 30, 40, 50})) {
		result = append(result, struct {
			index int
			value int
		}{i, v})
	}

	if len(expected) != len(result) {
		t.Error("Expected and result slices should have the same length")
	}

	for i := range expected {
		if expected[i].index != result[i].index || expected[i].value != result[i].value {
			t.Errorf("Expected (%d, %d) but got (%d, %d)", expected[i].index, expected[i].value, result[i].index, result[i].value)
		}
	}
}

func TestRange(t *testing.T) {
	expected := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	result := []int{}

	for i := range itr.Range(10) {
		result = append(result, i)
	}

	if len(expected) != len(result) {
		t.Error("Expected and result slices should have the same length")
	}

	for i := range expected {
		if expected[i] != result[i] {
			t.Errorf("Expected %d but got %d", expected[i], result[i])
		}
	}

	expected = []int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19}
	result = []int{}

	for i := range itr.Range(10, 20) {
		result = append(result, i)
	}

	if len(expected) != len(result) {
		t.Error("Expected and result slices should have the same length")
	}

	for i := range expected {
		if expected[i] != result[i] {
			t.Errorf("Expected %d but got %d", expected[i], result[i])
		}
	}

	expected = []int{10, 12, 14, 16, 18}
	result = []int{}

	for i := range itr.Range(10, 20, 2) {
		result = append(result, i)
	}

	if len(expected) != len(result) {
		t.Error("Expected and result slices should have the same length")
	}

	for i := range expected {
		if expected[i] != result[i] {
			t.Errorf("Expected %d but got %d", expected[i], result[i])
		}
	}

	expected = []int{10, 8, 6, 4, 2}
	result = []int{}

	for i := range itr.Range(10, 0, -2) {
		result = append(result, i)
	}

	if len(expected) != len(result) {
		t.Error("Expected and result slices should have the same length")
	}

	for i := range expected {
		if expected[i] != result[i] {
			t.Errorf("Expected %d but got %d", expected[i], result[i])
		}
	}

	expected = []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	result = []int{}

	for i := range itr.Range(10, 0) {
		result = append(result, i)
	}

	if len(expected) != len(result) {
		t.Error("Expected and result slices should have the same length")
	}

	for i := range expected {
		if expected[i] != result[i] {
			t.Errorf("Expected %d but got %d", expected[i], result[i])
		}
	}
}

func TestCount(t *testing.T) {
	expected := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	result := []int{}

	for v := range itr.Count(0) {
		result = append(result, v)
		if len(result) == 10 {
			break
		}
	}

	if len(expected) != len(result) {
		t.Error("Expected and result slices should have the same length")
	}

	for i := range expected {
		if expected[i] != result[i] {
			t.Errorf("Expected %d but got %d", expected[i], result[i])
		}
	}
}
func TestEnumerate(t *testing.T) {
	expected := []struct {
		index int
		value int
	}{
		{2, 10},
		{3, 20},
		{4, 30},
		{5, 40},
		{6, 50},
	}

	result := []struct {
		index int
		value int
	}{}

	for i, v := range itr.SliceValues([]int{10, 20, 30, 40, 50}).Enumerate(2) {
		result = append(result, struct {
			index int
			value int
		}{i, v})
	}

	if len(expected) != len(result) {
		t.Error("Expected and result slices should have the same length")
	}

	for i := range expected {
		if expected[i].index != result[i].index || expected[i].value != result[i].value {
			t.Errorf("Expected (%d, %d) but got (%d, %d)", expected[i].index, expected[i].value, result[i].index, result[i].value)
		}
	}

	result = []struct {
		index int
		value int
	}{}

	for i, v := range itr.Enumerate(itr.SliceValues([]int{10, 20, 30, 40, 50}), 2) {
		result = append(result, struct {
			index int
			value int
		}{i, v})
	}

	if len(expected) != len(result) {
		t.Error("Expected and result slices should have the same length")
	}

	for i := range expected {
		if expected[i].index != result[i].index || expected[i].value != result[i].value {
			t.Errorf("Expected (%d, %d) but got (%d, %d)", expected[i].index, expected[i].value, result[i].index, result[i].value)
		}
	}
}

func TestZip(t *testing.T) {
	expected := []struct {
		value1 int
		value2 int
	}{
		{10, 20},
		{20, 30},
		{30, 40},
		{40, 50},
	}

	result := []struct {
		value1 int
		value2 int
	}{}

	for v1, v2 := range itr.Zip(slices.Values([]int{10, 20, 30, 40, 50}), slices.Values([]int{20, 30, 40, 50})) {
		result = append(result, struct {
			value1 int
			value2 int
		}{v1, v2})
	}

	if len(expected) != len(result) {
		t.Error("Expected and result slices should have the same length")
	}

	for i := range expected {
		if expected[i].value1 != result[i].value1 || expected[i].value2 != result[i].value2 {
			t.Errorf("Expected (%d, %d) but got (%d, %d)", expected[i].value1, expected[i].value2, result[i].value1, result[i].value2)
		}
	}
}

func TestFilter(t *testing.T) {
	expected := []int{2, 4, 6, 8, 10}
	result := []int{}

	filterFunc := func(v int) bool {
		return v%2 == 0
	}

	for v := range itr.Range(1, 11).Filter(filterFunc) {
		result = append(result, v)
	}

	if len(expected) != len(result) {
		t.Error("Expected and result slices should have the same length")
	}

	for i := range expected {
		if expected[i] != result[i] {
			t.Errorf("Expected %d but got %d", expected[i], result[i])
		}
	}
}

func TestMap(t *testing.T) {
	expected := []int{4, 9, 16, 25, 36}
	result := []int{}

	mapFunc := func(v int) int {
		return v * v
	}

	for v := range itr.Range(2, 7).Map(mapFunc) {
		result = append(result, v)
	}

	if len(expected) != len(result) {
		t.Error("Expected and result slices should have the same length")
	}

	for i := range expected {
		if expected[i] != result[i] {
			t.Errorf("Expected %d but got %d", expected[i], result[i])
		}
	}
}
func TestReduce(t *testing.T) {
	expected := 55
	result := itr.Range(1, 11).Reduce(func(acc, v int) int {
		return acc + v
	})

	if expected != result {
		t.Errorf("Expected %d but got %d", expected, result)
	}

	expected = 50
	result = itr.SliceValues([]int{20, 10, 50, 33, 40, 49}).Reduce(func(a, b int) int {
		if a > b {
			return a
		}
		return b
	})

	if expected != result {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}
func TestCycle(t *testing.T) {
	expected := []int{1, 2, 3, 1, 2, 3, 1, 2, 3}
	result := []int{}

	cycleFunc := func(v int) bool {
		result = append(result, v)
		return len(result) != 9
	}

	for v := range itr.Range(1, 4).Cycle() {
		if !cycleFunc(v) {
			break
		}
	}

	if len(expected) != len(result) {
		t.Error("Expected and result slices should have the same length")
	}

	for i := range expected {
		if expected[i] != result[i] {
			t.Errorf("Expected %d but got %d", expected[i], result[i])
		}
	}
}
func TestRepeat(t *testing.T) {
	expected := []int{5, 5, 5, 5, 5}
	result := []int{}

	for v := range itr.Repeat(5, 5) {
		result = append(result, v)
	}

	if len(expected) != len(result) {
		t.Error("Expected and result slices should have the same length")
	}

	for i := range expected {
		if expected[i] != result[i] {
			t.Errorf("Expected %d but got %d", expected[i], result[i])
		}
	}

	expected = []int{5, 5, 5, 5, 5}
	result = []int{}

	i := 0
	for v := range itr.Repeat(5) {
		if i == 5 {
			break
		}
		result = append(result, v)
		i++
	}

	if len(expected) != len(result) {
		t.Error("Expected and result slices should have the same length")
	}

	for i := range expected {
		if expected[i] != result[i] {
			t.Errorf("Expected %d but got %d", expected[i], result[i])
		}
	}
}
func TestChain(t *testing.T) {
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	result := []int{}

	iterator1 := itr.Range(1, 6)
	iterator2 := itr.Range(6, 11)
	iterator3 := itr.Range(11, 16)

	for v := range itr.Chain(iterator1, iterator2, iterator3) {
		result = append(result, v)
	}

	if len(expected) != len(result) {
		t.Error("Expected and result slices should have the same length")
	}

	for i := range expected {
		if expected[i] != result[i] {
			t.Errorf("Expected %d but got %d", expected[i], result[i])
		}
	}
}
func TestTake(t *testing.T) {
	expected := []int{10, 20, 30, 40, 50}
	result := []int{}

	for v := range itr.Range(10, 60, 10).Take(5) {
		result = append(result, v)
	}

	if len(expected) != len(result) {
		t.Error("Expected and result slices should have the same length")
	}

	for i := range expected {
		if expected[i] != result[i] {
			t.Errorf("Expected %d but got %d", expected[i], result[i])
		}
	}

	expected = []int{10, 20, 30}
	result = []int{}

	for v := range itr.Range(10, 60, 10).Take(3) {
		result = append(result, v)
	}

	if len(expected) != len(result) {
		t.Error("Expected and result slices should have the same length")
	}

	for i := range expected {
		if expected[i] != result[i] {
			t.Errorf("Expected %d but got %d", expected[i], result[i])
		}
	}

	expected = []int{}
	result = []int{}

	for v := range itr.Range(10, 60, 10).Take(0) {
		result = append(result, v)
	}

	if len(expected) != len(result) {
		t.Error("Expected and result slices should have the same length")
	}

	for i := range expected {
		if expected[i] != result[i] {
			t.Errorf("Expected %d but got %d", expected[i], result[i])
		}
	}
}
func TestDrop(t *testing.T) {
	expected := []int{30, 40, 50}
	result := []int{}

	for v := range itr.Range(10, 60, 10).Drop(2) {
		result = append(result, v)
	}

	if len(expected) != len(result) {
		t.Error("Expected and result slices should have the same length")
	}

	for i := range expected {
		if expected[i] != result[i] {
			t.Errorf("Expected %d but got %d", expected[i], result[i])
		}
	}

	expected = []int{10, 20, 30, 40, 50}
	result = []int{}

	for v := range itr.Range(10, 60, 10).Drop(0) {
		result = append(result, v)
	}

	if len(expected) != len(result) {
		t.Error("Expected and result slices should have the same length")
	}

	for i := range expected {
		if expected[i] != result[i] {
			t.Errorf("Expected %d but got %d", expected[i], result[i])
		}
	}
}
func TestTakeWhile(t *testing.T) {
	expected := []int{2, 4, 6, 8, 10}
	result := []int{}

	predicate := func(v int) bool {
		return v <= 10
	}

	for v := range itr.Range(2, 11, 2).TakeWhile(predicate) {
		result = append(result, v)
	}

	if len(expected) != len(result) {
		t.Error("Expected and result slices should have the same length")
	}

	for i := range expected {
		if expected[i] != result[i] {
			t.Errorf("Expected %d but got %d", expected[i], result[i])
		}
	}

	expected = []int{}
	result = []int{}

	predicate = func(v int) bool {
		return v > 10
	}

	for v := range itr.Range(1, 11).TakeWhile(predicate) {
		result = append(result, v)
	}

	if len(expected) != len(result) {
		t.Error("Expected and result slices should have the same length")
	}

	for i := range expected {
		if expected[i] != result[i] {
			t.Errorf("Expected %d but got %d", expected[i], result[i])
		}
	}
}
func TestDropWhile(t *testing.T) {
	expected := []int{30, 40, 50}
	result := []int{}

	predicate := func(v int) bool {
		return v < 30
	}

	for v := range itr.Range(10, 60, 10).DropWhile(predicate) {
		result = append(result, v)
	}

	if len(expected) != len(result) {
		t.Error("Expected and result slices should have the same length")
	}

	for i := range expected {
		if expected[i] != result[i] {
			t.Errorf("Expected %d but got %d", expected[i], result[i])
		}
	}
}
func TestWith(t *testing.T) {
	expected := []int{}
	result := []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5}

	withFunc := func() {
		result = result[:len(result)-1]
	}

	for range itr.Range(1, 6).With(withFunc) {
		result = result[:len(result)-1]
	}

	if len(expected) != len(result) {
		t.Error("Expected and result slices should have the same length")
	}

	for i := range expected {
		if expected[i] != result[i] {
			t.Errorf("Expected %d but got %d", expected[i], result[i])
		}
	}
}
func TestElse(t *testing.T) {
	expected := []int{1}
	result := []int{}

	iterator := itr.Range(1, 6)
	elseFunc := func() {
		t.Error("Else function should not be called")
	}

	for v := range iterator.Else(elseFunc) {
		result = append(result, v)
		break
	}

	if len(expected) != len(result) {
		t.Error("Expected and result slices should have the same length")
	}

	for i := range expected {
		if expected[i] != result[i] {
			t.Errorf("Expected %d but got %d", expected[i], result[i])
		}
	}

	expected = []int{1, 2, 3, 4, 5, 0}
	result = []int{}

	iterator = itr.Range(1, 6)
	elseFunc = func() {
		result = append(result, 0)
	}

	for v := range iterator.Else(elseFunc) {
		result = append(result, v)
	}

	if len(expected) != len(result) {
		t.Error("Expected and result slices should have the same length")
	}

	for i := range expected {
		if expected[i] != result[i] {
			t.Errorf("Expected %d but got %d", expected[i], result[i])
		}
	}
}
