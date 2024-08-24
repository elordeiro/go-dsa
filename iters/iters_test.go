package iters_test

import (
	"fmt"
	"slices"
	"strings"
	"testing"

	it "github.com/elordeiro/go/iters"
)

// Shorthand for Pair types. Last 2 letters are the types of the pair.
type (
	Pairii it.Pair[int, int]    // Pair of int, int
	Pairis it.Pair[int, string] // Pair of int, string
	Pairsi it.Pair[string, int] // Pair of string, int
)

func TestIterable(t *testing.T) {
	type IntSlice []int

	expected := []int{1, 2, 3, 4, 5}
	result := []int{}
	intSlice := IntSlice{1, 2, 3, 4, 5}
	slice := it.Iterable(intSlice)

	for v := range slice.Values() {
		result = append(result, v)
	}

	if len(expected) != len(result) {
		t.Error("Expected and result slices should have the same length")
	}

	for i, res := range result {
		if expected[i] != res {
			t.Errorf("Expected %d but got %d", expected[i], res)
		}
	}
}

func TestIterable2(t *testing.T) {
	type IntStringMap map[int]string

	expected := []it.Pair[int, string]{{0, "1"}, {1, "2"}, {2, "3"}, {3, "4"}, {4, "5"}}
	result := []it.Pair[int, string]{}
	intStringMap := IntStringMap{0: "1", 1: "2", 2: "3", 3: "4", 4: "5"}
	slice := it.Iterable2(intStringMap)

	for k, v := range slice.All() {
		result = append(result, it.Pair[int, string]{k, v})
	}

	if len(expected) != len(result) {
		t.Error("Expected and result slices should have the same length")
	}

	for _, res := range result {
		if !slices.Contains(expected, res) {
			t.Errorf("Expected %s but got %s", expected, result)
		}
	}
}

func TestSeq_Len(t *testing.T) {
	expected := 5
	result := it.Range(1, 6).Len()

	if expected != result {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}

func TestSeq_All(t *testing.T) {
	expected := [][]int{{0, 10}, {1, 20}, {2, 30}, {3, 40}, {4, 50}}
	result := [][]int{}
	seq := it.Iterable([]int{10, 20, 30, 40, 50}).Values()
	for i, v := range seq.All() {
		result = append(result, []int{i, v})
	}

	if len(expected) != len(result) {
		t.Error("Expected and result slices should have the same length")
	}

	for i, res := range result {
		if expected[i][0] != res[0] || expected[i][1] != res[1] {
			t.Errorf("Expected %d but got %d", expected[i], result[i])
		}
	}
}

func TestSeq2_Len(t *testing.T) {
	expected := 5
	result := it.Iterable2(map[int]string{0: "1", 1: "2", 2: "3", 3: "4", 4: "5"}).All().Len()

	if expected != result {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}

func TestSeq2_Keys(t *testing.T) {
	expected := []int{0, 1, 2, 3, 4}
	result := []int{}
	seq2 := it.Range(10, 60, 10).All()

	for k := range seq2.Keys() {
		result = append(result, k)
	}

	if len(expected) != len(result) {
		t.Error("Expected and result slices should have the same length")
	}

	for i, res := range result {
		if expected[i] != res {
			t.Errorf("Expected %d but got %d", expected[i], res)
		}
	}
}

func TestSeq2_Values(t *testing.T) {
	expected := []int{10, 20, 30, 40, 50}
	result := []int{}
	seq2 := it.Range(10, 60, 10).All()

	for v := range seq2.Values() {
		result = append(result, v)
	}

	if len(expected) != len(result) {
		t.Error("Expected and result slices should have the same length")
	}

	for i, res := range result {
		if expected[i] != res {
			t.Errorf("Expected %d but got %d", expected[i], res)
		}
	}
}

func TestSlice_Len(t *testing.T) {
	expected := 5
	result := it.Iterable([]int{1, 2, 3, 4, 5}).Len()

	if expected != result {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}

func TestMap_Len(t *testing.T) {
	expected := 5
	result := it.Iterable2(map[int]string{0: "1", 1: "2", 2: "3", 3: "4", 4: "5"}).Len()

	if expected != result {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}

func TestRange(t *testing.T) {
	expected := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	result := []int{}

	for i := range it.Range(10) {
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

	for i := range it.Range(10, 20) {
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

	for i := range it.Range(10, 20, 2) {
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

	for i := range it.Range(10, 0, -2) {
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

	for i := range it.Range(10, 0) {
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

	for v := range it.Count(0) {
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

func TestCountDown(t *testing.T) {
	expected := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	result := []int{}

	for v := range it.CountDown(10) {
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
	expected := []Pairii{
		{2, 10},
		{3, 20},
		{4, 30},
		{5, 40},
		{6, 50},
	}

	result := []Pairii{}
	s := it.Iterable([]int{10, 20, 30, 40, 50})
	for i, v := range it.Enumerate(2, s) {
		result = append(result, Pairii{i, v})
	}

	if len(expected) != len(result) {
		t.Error("Expected and result slices should have the same length")
	}

	for i := range result {
		if expected[i] != result[i] {
			t.Errorf("Expected %d but got %d", expected[i], result[i])
		}
	}
}

func TestEnumerateFunc2(t *testing.T) {
	expected := []it.Pair[int, it.Pair[string, int]]{
		{2, it.Pair[string, int]{"a", 10}},
		{3, it.Pair[string, int]{"b", 20}},
		{4, it.Pair[string, int]{"c", 30}},
		{5, it.Pair[string, int]{"d", 40}},
		{6, it.Pair[string, int]{"e", 50}},
	}

	result := []it.Pair[int, it.Pair[string, int]]{}

	slice := it.Iterable([]it.Pair[string, int]{{"a", 10}, {"b", 20}, {"c", 30}, {"d", 40}, {"e", 50}})
	kvSlice := it.Split(slice)
	for i, p := range it.Enumerate2(2, kvSlice) {
		result = append(result, it.Pair[int, it.Pair[string, int]]{i, p})
	}

	if len(expected) != len(result) {
		t.Error("Expected and result slices should have the same length")
	}

	for i := range result {
		if !slices.Contains(expected, result[i]) {
			t.Errorf("Expected %s but got %s", expected[i], result[i])
		}
	}
}

func TestZip(t *testing.T) {
	expected := []Pairii{
		{10, 20},
		{20, 30},
		{30, 40},
		{40, 50},
	}

	result := []Pairii{}

	for v1, v2 := range it.Zip(it.Range(10, 50, 10), it.Range(20, 60, 10)) {
		result = append(result, Pairii{v1, v2})
	}

	if len(expected) != len(result) {
		t.Error("Expected and result slices should have the same length")
	}

	for i := range result {
		if expected[i] != result[i] {
			t.Errorf("Expected %d but got %d", expected[i], result[i])
		}
	}
}

func TestZip2(t *testing.T) {
	expected := []it.Pair[it.Pair[string, int], it.Pair[string, int]]{
		{it.Pair[string, int]{"a", 10}, it.Pair[string, int]{"b", 20}},
		{it.Pair[string, int]{"c", 30}, it.Pair[string, int]{"d", 40}},
		{it.Pair[string, int]{"e", 50}, it.Pair[string, int]{"f", 60}},
	}

	result := []it.Pair[it.Pair[string, int], it.Pair[string, int]]{}

	s1 := it.Iterable([]it.Pair[string, int]{{"a", 10}, {"c", 30}, {"e", 50}})
	s2 := it.Iterable([]it.Pair[string, int]{{"b", 20}, {"d", 40}, {"f", 60}})

	for p1, p2 := range it.Zip2(it.Split(s1), it.Split(s2)) {
		result = append(result, it.Pair[it.Pair[string, int], it.Pair[string, int]]{p1, p2})
	}

	if len(expected) != len(result) {
		t.Error("Expected and result slices should have the same length")
	}

	for i := range result {
		if expected[i] != result[i] {
			t.Errorf("Expected %s but got %s", expected[i], result[i])
		}
	}
}

func TestRepeat(t *testing.T) {
	expected := []int{5, 5, 5, 5, 5}
	result := []int{}

	for v := range it.Repeat(5, 5) {
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

	result = []int{}

	i := 0
	for v := range it.Repeat(5, -1) {
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

func TestRepeat2(t *testing.T) {
	expected := []Pairii{
		{4, 5},
		{4, 5},
		{4, 5},
		{4, 5},
		{4, 5},
	}

	result := []Pairii{}

	for k, v := range it.Repeat2(4, 5, 5) {
		result = append(result, Pairii{k, v})
	}

	if len(expected) != len(result) {
		t.Error("Expected and result slices should have the same length")
	}

	for i := range result {
		if expected[i] != result[i] {
			t.Errorf("Expected %d but got %d", expected[i], result[i])
		}
	}

	result = []Pairii{}

	i := 0
	for k, v := range it.Repeat2(4, 5, -1) {
		if i == 5 {
			break
		}
		result = append(result, Pairii{k, v})
		i++
	}

	if len(expected) != len(result) {
		t.Error("Expected and result slices should have the same length")
	}

	for i := range result {
		if expected[i] != result[i] {
			t.Errorf("Expected %d but got %d", expected[i], result[i])
		}
	}
}

func TestCycle(t *testing.T) {
	expected := []int{1, 2, 3, 1, 2, 3, 1, 2, 3}
	result := []int{}

	for v := range it.Cycle(it.Range(1, 4)) {
		if len(result) == 9 {
			break
		}
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

func TestCycle2(t *testing.T) {
	expected := []Pairii{
		{0, 10},
		{1, 20},
		{2, 30},
		{0, 10},
		{1, 20},
		{2, 30},
		{0, 10},
		{1, 20},
		{2, 30},
	}

	result := []Pairii{}

	for i, v := range it.Cycle2(it.Range(10, 40, 10).All()) {
		if len(result) == 9 {
			break
		}
		result = append(result, Pairii{i, v})
	}

	if len(expected) != len(result) {
		t.Error("Expected and result slices should have the same length")
	}

	for i := range result {
		if expected[i] != result[i] {
			t.Errorf("Expected %d but got %d", expected[i], result[i])
		}
	}
}

func TestChain(t *testing.T) {
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	result := []int{}

	for v := range it.Chain(it.Range(1, 6), it.Range(6, 10)) {
		result = append(result, v)
	}

	if len(expected) != len(result) {
		t.Error("Expected and result slices should have the same length")
	}

	for i := range result {
		if expected[i] != result[i] {
			t.Errorf("Expected %d but got %d", expected[i], result[i])
		}
	}
}

func TestChain2(t *testing.T) {
	expected := []Pairii{
		{0, 10},
		{1, 20},
		{2, 30},
		{3, 40},
		{4, 50},
		{5, 60},
		{6, 70},
		{7, 80},
		{8, 90},
	}

	result := []Pairii{}

	for i, v := range it.Chain2(it.Range(10, 60, 10).All(), it.Enumerate(5, it.Range(60, 100, 10))) {
		result = append(result, Pairii{i, v})
	}

	if len(expected) != len(result) {
		t.Error("Expected and result slices should have the same length")
	}

	for i := range result {
		if expected[i] != result[i] {
			t.Errorf("Expected %d but got %d", expected[i], result[i])
		}
	}
}

func TestBackwards(t *testing.T) {
	expected := []int{5, 4, 3, 2, 1}
	result := []int{}

	for v := range it.Backwards(it.Range(1, 6)) {
		result = append(result, v)
	}

	if len(expected) != len(result) {
		t.Error("Expected and result slices should have the same length")
	}

	for i := range result {
		if expected[i] != result[i] {
			t.Errorf("Expected %d but got %d", expected[i], result[i])
		}
	}

	expected = []int{5, 4, 3, 2, 1}
	result = []int{}

	for v := range it.Backwards(it.Iterable([]int{1, 2, 3, 4, 5})) {
		result = append(result, v)
	}

	if len(expected) != len(result) {
		t.Error("Expected and result slices should have the same length")
	}

	for i := range result {
		if expected[i] != result[i] {
			t.Errorf("Expected %d but got %d", expected[i], result[i])
		}
	}
}

func TestBackwards2(t *testing.T) {
	expected := []Pairii{
		{4, 50},
		{3, 40},
		{2, 30},
		{1, 20},
		{0, 10},
	}

	result := []Pairii{}

	for k, v := range it.Backwards2(it.Range(10, 60, 10).All()) {
		result = append(result, Pairii{k, v})
	}

	if len(expected) != len(result) {
		t.Error("Expected and result slices should have the same length")
	}

	for i := range result {
		if expected[i] != result[i] {
			t.Errorf("Expected %d but got %d", expected[i], result[i])
		}
	}

	expected = []Pairii{
		{4, 50},
		{3, 40},
		{2, 30},
		{1, 20},
		{0, 10},
	}

	result = []Pairii{}
	m := it.Iterable2(map[int]int{0: 10, 1: 20, 2: 30, 3: 40, 4: 50})
	for k, v := range it.Backwards2(m) {
		result = append(result, Pairii{k, v})
	}

	if len(expected) != len(result) {
		t.Error("Expected and result slices should have the same length")
	}

	for i := range result {
		if !slices.Contains(expected, result[i]) {
			t.Errorf("Expected %d but got %d", expected[i], result[i])
		}
	}
}

func TestTake(t *testing.T) {
	expected := []int{10, 20, 30, 40, 50}
	result := []int{}

	for v := range it.Take(5, it.Range(10, 60, 10)) {
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

	for v := range it.Take(3, it.Range(10, 60, 10)) {
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

	for v := range it.Take(0, it.Range(10, 60, 10)) {
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

func TestTake2(t *testing.T) {
	expected := []Pairii{
		{0, 10},
		{1, 20},
		{2, 30},
	}
	result := []Pairii{}

	for k, v := range it.Take2(3, it.Range(10, 60, 10).All()) {
		result = append(result, Pairii{k, v})
	}

	if len(expected) != len(result) {
		t.Error("Expected and result slices should have the same length")
	}

	for i := range expected {
		if expected[i] != result[i] {
			t.Errorf("Expected %v but got %v", expected[i], result[i])
		}
	}

	expected = []Pairii{}
	result = []Pairii{}

	for k, v := range it.Take2(0, it.Range(10, 60, 10).All()) {
		result = append(result, Pairii{k, v})
	}

	if len(expected) != len(result) {
		t.Error("Expected and result slices should have the same length")
	}

	for i := range expected {
		if expected[i] != result[i] {
			t.Errorf("Expected %v but got %v", expected[i], result[i])
		}
	}
}

func TestDrop(t *testing.T) {
	expected := []int{30, 40, 50}
	result := []int{}

	for v := range it.Drop(2, it.Range(10, 60, 10)) {
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

	for v := range it.Drop(0, it.Range(10, 60, 10)) {
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

func TestDrop2(t *testing.T) {
	expected := []Pairii{
		{2, 30},
		{3, 40},
		{4, 50},
	}
	result := []Pairii{}

	for k, v := range it.Drop2(2, it.Range(10, 60, 10).All()) {
		result = append(result, Pairii{k, v})
	}

	if len(expected) != len(result) {
		t.Error("Expected and result slices should have the same length")
	}

	for i := range expected {
		if expected[i] != result[i] {
			t.Errorf("Expected %v but got %v", expected[i], result[i])
		}
	}

	expected = []Pairii{
		{0, 10},
		{1, 20},
		{2, 30},
		{3, 40},
		{4, 50},
	}
	result = []Pairii{}

	for k, v := range it.Drop2(0, it.Range(10, 60, 10).All()) {
		result = append(result, Pairii{k, v})
	}

	if len(expected) != len(result) {
		t.Error("Expected and result slices should have the same length")
	}

	for i := range expected {
		if expected[i] != result[i] {
			t.Errorf("Expected %v but got %v", expected[i], result[i])
		}
	}
}

func TestTakeBetween(t *testing.T) {
	expected := []int{20, 30, 40}
	result := []int{}

	for v := range it.TakeBetween(1, 4, it.Range(10, 60, 10)) {
		result = append(result, v)
	}

	if len(expected) != len(result) {
		t.Fatal("Expected and result slices should have the same length")
	}

	for i, res := range result {
		if expected[i] != res {
			t.Errorf("Expected %d but got %d", expected[i], res)
		}
	}
}

func TestTakeBetween2(t *testing.T) {
	expected := []Pairii{
		{1, 20},
		{2, 30},
		{3, 40},
	}
	result := []Pairii{}

	for k, v := range it.TakeBetween2(1, 4, it.Range(10, 60, 10).All()) {
		result = append(result, Pairii{k, v})
	}

	if len(expected) != len(result) {
		t.Error("Expected and result slices should have the same length")
	}

	for i, res := range result {
		if expected[i] != res {
			t.Errorf("Expected %v but got %v", expected[i], res)
		}
	}
}

func TestRotate(t *testing.T) {
	expected := []int{30, 40, 50, 10, 20}
	result := []int{}

	for v := range it.Rotate(2, it.Range(10, 60, 10)) {
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

	expected = []int{40, 50, 10, 20, 30}
	result = []int{}

	for v := range it.Rotate(-2, it.Range(10, 60, 10)) {
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

func TestRotate_overLen(t *testing.T) {
	expected := []int{20, 30, 40, 50, 10}
	result := []int{}

	for v := range it.Rotate(6, it.Range(10, 60, 10)) {
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

	expected = []int{50, 10, 20, 30, 40}
	result = []int{}

	for v := range it.Rotate(-6, it.Range(10, 60, 10)) {
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

func TestRotate2(t *testing.T) {
	expected := []Pairii{
		{2, 30},
		{3, 40},
		{4, 50},
		{0, 10},
		{1, 20},
	}
	result := []Pairii{}

	for k, v := range it.Rotate2(2, it.Range(10, 60, 10).All()) {
		result = append(result, Pairii{k, v})
	}

	if len(expected) != len(result) {
		t.Error("Expected and result slices should have the same length")
	}

	for i := range result {
		if expected[i] != result[i] {
			t.Errorf("Expected %v but got %v", expected[i], result[i])
		}
	}

	expected = []Pairii{
		{3, 40},
		{4, 50},
		{0, 10},
		{1, 20},
		{2, 30},
	}
	result = []Pairii{}

	for k, v := range it.Rotate2(-2, it.Range(10, 60, 10).All()) {
		result = append(result, Pairii{k, v})
	}

	if len(expected) != len(result) {
		t.Error("Expected and result slices should have the same length")
	}

	for i := range result {
		if expected[i] != result[i] {
			t.Errorf("Expected %v but got %v", expected[i], result[i])
		}
	}
}

func TestRotate2_overLen(t *testing.T) {
	expected := []Pairii{
		{1, 20},
		{2, 30},
		{3, 40},
		{4, 50},
		{0, 10},
	}
	result := []Pairii{}

	for k, v := range it.Rotate2(6, it.Range(10, 60, 10).All()) {
		result = append(result, Pairii{k, v})
	}

	if len(expected) != len(result) {
		t.Error("Expected and result slices should have the same length")
	}

	for i := range result {
		if expected[i] != result[i] {
			t.Errorf("Expected %v but got %v", expected[i], result[i])
		}
	}

	expected = []Pairii{
		{4, 50},
		{0, 10},
		{1, 20},
		{2, 30},
		{3, 40},
	}
	result = []Pairii{}

	for k, v := range it.Rotate2(-6, it.Range(10, 60, 10).All()) {
		result = append(result, Pairii{k, v})
	}

	if len(expected) != len(result) {
		t.Error("Expected and result slices should have the same length")
	}

	for i := range result {
		if expected[i] != result[i] {
			t.Errorf("Expected %v but got %v", expected[i], result[i])
		}
	}
}

func TestForEach(t *testing.T) {
	expected := []int{1, 4, 3, 16, 5, 36}
	result := []int{}

	it.ForEach(it.Range(1, 7), func(v int) {
		if v%2 == 0 {
			result = append(result, v*v)
		} else {
			result = append(result, v)
		}
	})

	if len(expected) != len(result) {
		t.Error("Expected and result slices should have the same length")
	}

	for i := range result {
		if expected[i] != result[i] {
			t.Errorf("Expected %d but got %d", expected[i], result[i])
		}
	}
}

func TestForEach2(t *testing.T) {
	expected := []Pairii{
		{0, 10},
		{1, 20},
		{2, 30},
		{3, 40},
		{4, 50},
	}
	result := []Pairii{}

	it.ForEach2(it.Range(10, 60, 10).All(), func(k, v int) {
		result = append(result, Pairii{k, v})
	})

	if len(expected) != len(result) {
		t.Error("Expected and result slices should have the same length")
	}

	for i := range result {
		if expected[i] != result[i] {
			t.Errorf("Expected %v but got %v", expected[i], result[i])
		}
	}
}

func TestFilter(t *testing.T) {
	expected := []int{2, 4, 6, 8, 10}
	result := []int{}

	filterFunc := func(v int) bool {
		return v%2 == 0
	}

	for v := range it.Filter(it.Range(1, 11), filterFunc) {
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

func TestFilter2(t *testing.T) {
	expected := []Pairii{
		{0, 10},
		{2, 30},
		{4, 50},
	}

	result := []Pairii{}

	filterFunc := func(k, v int) bool {
		return k%2 == 0
	}

	for k, v := range it.Filter2(it.Range(10, 60, 10).All(), filterFunc) {
		result = append(result, Pairii{k, v})
	}

	if len(expected) != len(result) {
		t.Error("Expected and result slices should have the same length")
	}

	for i := range result {
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

	for v := range it.Map(it.Range(2, 7), mapFunc) {
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

func TestMap2(t *testing.T) {
	expected := []Pairii{
		{0, 100},
		{1, 200},
		{2, 300},
		{3, 400},
		{4, 500},
	}

	result := []Pairii{}

	mapFunc := func(k, v int) (int, int) {
		return k, v * 100
	}

	for k, v := range it.Map2(it.Range(1, 6).All(), mapFunc) {
		result = append(result, Pairii{k, v})
	}

	if len(expected) != len(result) {
		t.Error("Expected and result slices should have the same length")
	}

	for i := range result {
		if expected[i] != result[i] {
			t.Errorf("Expected %d but got %d", expected[i], result[i])
		}
	}
}

func TestReduce(t *testing.T) {
	expected := 55
	result := it.Reduce(it.Range(1, 11), func(acc, v int) int {
		return acc + v
	})

	if expected != result {
		t.Errorf("Expected %d but got %d", expected, result)
	}

	expected = 50
	result = it.Reduce(it.Iterable([]int{20, 10, 50, 33, 40, 49}).Values(), func(a, b int) int {
		if a > b {
			return a
		}
		return b
	})

	if expected != result {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}

func TestReduce2(t *testing.T) {
	expected := 150
	m := it.Iterable2(map[int]int{0: 10, 1: 20, 2: 30, 3: 40, 4: 50})
	result := it.Reduce2(m.All(), func(acc, v int) int {
		return acc + v
	})

	if expected != result {
		t.Errorf("Expected %d but got %d", expected, result)
	}

	expected = 50
	result = it.Reduce2(m.All(), func(a, b int) int {
		if a > b {
			return a
		}
		return b
	})

	if expected != result {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}

func TestReduce2_second(t *testing.T) {
	expected := 10
	m := it.Iterable2(map[int]int{0: 10, 1: 20, 2: 30, 3: 40, 4: 50})
	result := it.Reduce(m.Keys(), func(acc, v int) int {
		return acc + v
	})

	if expected != result {
		t.Errorf("Expected %d but got %d", expected, result)
	}

	expected = 50
	result = it.Reduce(m.All(), func(a, b int) int {
		if a > b {
			return a
		}
		return b
	})

	if expected != result {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}

func TestTakeWhile(t *testing.T) {
	expected := []int{2, 4, 6, 8, 10}
	result := []int{}

	predicate := func(v int) bool {
		return v <= 10
	}

	for v := range it.TakeWhile(it.Range(2, 11, 2), predicate) {
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

	for v := range it.TakeWhile(it.Range(1, 11), predicate) {
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

func TestTakeWhile2(t *testing.T) {
	expected := []Pairii{
		{0, 10},
		{1, 20},
		{2, 30},
	}
	result := []Pairii{}

	predicate := func(k, v int) bool {
		return k <= 2
	}

	for k, v := range it.TakeWhile2(it.Range(10, 60, 10).All(), predicate) {
		result = append(result, Pairii{k, v})
	}

	if len(expected) != len(result) {
		t.Error("Expected and result slices should have the same length")
	}

	for i := range result {
		if expected[i] != result[i] {
			t.Errorf("Expected %v but got %v", expected[i], result[i])
		}
	}

	expected = []Pairii{}
	result = []Pairii{}

	predicate = func(k, v int) bool {
		return k > 2
	}

	for k, v := range it.TakeWhile2(it.Range(10, 60, 10).All(), predicate) {
		result = append(result, Pairii{k, v})
	}

	if len(expected) != len(result) {
		t.Error("Expected and result slices should have the same length")
	}

	for i := range result {
		if expected[i] != result[i] {
			t.Errorf("Expected %v but got %v", expected[i], result[i])
		}
	}
}

func TestDropWhile(t *testing.T) {
	expected := []int{30, 40, 50}
	result := []int{}

	predicate := func(v int) bool {
		return v < 30
	}

	for v := range it.DropWhile(it.Range(10, 60, 10), predicate) {
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

func TestDropWhile2(t *testing.T) {
	expected := []Pairii{
		{2, 30},
		{3, 40},
		{4, 50},
	}
	result := []Pairii{}

	predicate := func(k, v int) bool {
		return k < 2
	}

	for k, v := range it.DropWhile2(it.Range(10, 60, 10).All(), predicate) {
		result = append(result, Pairii{k, v})
	}

	if len(expected) != len(result) {
		t.Error("Expected and result slices should have the same length")
	}

	for i := range result {
		if expected[i] != result[i] {
			t.Errorf("Expected %v but got %v", expected[i], result[i])
		}
	}
}

func TestWith(t *testing.T) {
	expected := 30
	result := 0

	withFunc := func(v int) {
		if v%2 == 0 {
			result += v
		}
	}

	for range it.With(it.Range(1, 11), withFunc) {
	}

	if expected != result {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}

func TestWith2(t *testing.T) {
	expected := 140
	result := 0

	withFunc := func(k, v int) {
		if k%2 == 0 {
			result += k * v
		}
	}

	for range it.With2(it.Range(1, 11).All(), withFunc) {
	}

	if expected != result {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}

func TestOrElse(t *testing.T) {
	expected := []int{1}
	result := []int{}

	elseFunc := func() {
		t.Error("Else function should not be called")
	}

	for v := range it.OnEmpty(it.Range(1, 6), elseFunc) {
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

	elseFunc = func() {
		result = append(result, 0)
	}

	for v := range it.OnEmpty(it.Range(1, 6), elseFunc) {
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

func TestOrElse2(t *testing.T) {
	expected := []Pairii{{0, 1}}
	result := []Pairii{}

	elseFunc := func() {
		t.Error("Else function should not be called")
	}

	for k, v := range it.OnEmpty2(it.Range(1, 6).All(), elseFunc) {
		result = append(result, Pairii{k, v})
		break
	}

	if len(expected) != len(result) {
		t.Error("Expected and result slices should have the same length")
	}

	for i := range result {
		if expected[i] != result[i] {
			t.Errorf("Expected %v but got %v", expected[i], result[i])
		}
	}

	expected = []Pairii{{0, 1}, {1, 2}, {2, 3}, {3, 4}, {4, 5}, {0, 0}}
	result = []Pairii{}

	elseFunc = func() {
		result = append(result, Pairii{0, 0})
	}

	for k, v := range it.OnEmpty2(it.Range(1, 6).All(), elseFunc) {
		result = append(result, Pairii{k, v})
	}

	if len(expected) != len(result) {
		t.Error("Expected and result slices should have the same length")
	}

	for i := range result {
		if expected[i] != result[i] {
			t.Errorf("Expected %v but got %v", expected[i], result[i])
		}
	}
}

func TestSum(t *testing.T) {
	expected := 15
	result := it.Sum(it.Range(1, 6))

	if expected != result {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}

func TestProduct(t *testing.T) {
	expected := 120
	result := it.Product(it.Range(1, 6))

	if expected != result {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}

func TestMin(t *testing.T) {
	expected := 1
	result := it.Min(it.Range(1, 6))

	if expected != result {
		t.Errorf("Expected %d but got %d", expected, result)
	}

	expected = -15
	result = it.Min(it.Range(-15, -5))

	if expected != result {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}

func TestMax(t *testing.T) {
	expected := 5
	result := it.Max(it.Range(1, 6))

	if expected != result {
		t.Errorf("Expected %d but got %d", expected, result)
	}

	expected = -6
	result = it.Max(it.Range(-15, -5))

	if expected != result {
		t.Errorf("Expected %d but got %d", expected, result)
	}
}

func TestAll(t *testing.T) {
	expected := true
	result := it.All(it.Range(1, 6))

	if expected != result {
		t.Errorf("Expected %t but got %t", expected, result)
	}

	expected = false

	result = it.All(it.Range(1, 6), func(v int) bool {
		return v < 5
	})

	if expected != result {
		t.Errorf("Expected %t but got %t", expected, result)
	}

	expected = false
	slice := it.Iterable([]int{0, 1, 2, 3, 4, 5})
	result = it.All(slice.Values())

	if expected != result {
		t.Errorf("Expected %t but got %t", expected, result)
	}

	expected = true
	slice = it.Iterable([]int{1, 2, 3, 4, 5})
	result = it.All(slice.Values())

	if expected != result {
		t.Errorf("Expected %t but got %t", expected, result)
	}
}

func TestAny(t *testing.T) {
	expected := true
	result := it.Any(it.Range(1, 6))

	if expected != result {
		t.Errorf("Expected %t but got %t", expected, result)
	}

	expected = true

	result = it.Any(it.Range(1, 6), func(v int) bool {
		return v < 5
	})

	if expected != result {
		t.Errorf("Expected %t but got %t", expected, result)
	}

	expected = true
	slice := it.Iterable([]int{0, 1, 2, 3, 4, 5})
	result = it.Any(slice.Values())

	if expected != result {
		t.Errorf("Expected %t but got %t", expected, result)
	}

	expected = true
	slice = it.Iterable([]int{1, 2, 3, 4, 5})
	result = it.Any(slice.Values())

	if expected != result {
		t.Errorf("Expected %t but got %t", expected, result)
	}

	expected = false
	slice = it.Iterable([]int{0, 0, 0, 0, 0})
	result = it.Any(slice.Values())

	if expected != result {
		t.Errorf("Expected %t but got %t", expected, result)
	}
}

func TestNone(t *testing.T) {
	expected := false
	result := it.None(it.Range(1, 6))

	if expected != result {
		t.Errorf("Expected %t but got %t", expected, result)
	}

	expected = false

	result = it.None(it.Range(1, 6), func(v int) bool {
		return v < 5
	})

	if expected != result {
		t.Errorf("Expected %t but got %t", expected, result)
	}

	expected = false
	slice := it.Iterable([]int{0, 1, 2, 3, 4, 5})
	result = it.None(slice.Values())

	if expected != result {
		t.Errorf("Expected %t but got %t", expected, result)
	}

	expected = false
	slice = it.Iterable([]int{1, 2, 3, 4, 5})
	result = it.None(slice.Values())

	if expected != result {
		t.Errorf("Expected %t but got %t", expected, result)
	}

	expected = true
	slice = it.Iterable([]int{0, 0, 0, 0, 0})
	result = it.None(slice.Values())

	if expected != result {
		t.Errorf("Expected %t but got %t", expected, result)
	}
}

func TestSplit(t *testing.T) {
	expected := it.Iterable([]it.Pair[int, string]{
		{1, "a"},
		{2, "b"},
		{3, "c"},
		{4, "d"},
		{5, "e"},
	})
	result := []it.Pair[int, string]{}
	slice := it.Iterable([]it.Pair[int, string]{{1, "a"}, {2, "b"}, {3, "c"}, {4, "d"}, {5, "e"}})
	for k, v := range it.Split(slice) {
		result = append(result, it.Pair[int, string]{k, v})
	}

	if len(expected) != len(result) {
		t.Error("Expected and result slices should have the same length")
	}

	for i := range expected {
		if expected[i] != result[i] {
			t.Errorf("Expected %s but got %s", expected[i], result[i])
		}
	}
}

func TestSwapKV(t *testing.T) {
	expected := []Pairii{
		{10, 0},
		{20, 1},
		{30, 2},
		{40, 3},
		{50, 4},
	}
	result := []Pairii{}

	for k, v := range it.SwapKV(it.Range(10, 60, 10).All()) {
		result = append(result, Pairii{k, v})
	}

	if len(expected) != len(result) {
		t.Error("Expected and result slices should have the same length")
	}

	for i := range result {
		if expected[i] != result[i] {
			t.Errorf("Expected %v but got %v", expected[i], result[i])
		}
	}
}

func TestSeqString(t *testing.T) {
	expected := "Seq[int][1 2 3 4 5]"
	result := fmt.Sprint(it.Range(1, 6))

	if expected != result {
		t.Errorf("Expected %s but got %s", expected, result)
	}
}

func TestSeq2String(t *testing.T) {
	expected := "Seq2[int,int][0:1 1:2 2:3 3:4 4:5]"
	result := fmt.Sprint(it.Range(1, 6).All())

	if expected != result {
		t.Errorf("Expected %s but got %s", expected, result)
	}
}

func TestItSliceString(t *testing.T) {
	expected := "ItSlice[int][1 2 3 4 5]"
	result := fmt.Sprint(it.Iterable([]int{1, 2, 3, 4, 5}))

	if expected != result {
		t.Errorf("Expected %s but got %s", expected, result)
	}
}

func TestItMapString(t *testing.T) {
	expected := []string{"0:10", "1:20", "2:30", "3:40", "4:50"}
	m := it.Iterable2(map[int]int{0: 10, 1: 20, 2: 30, 3: 40, 4: 50})
	str := fmt.Sprint(m)
	begin := strings.Index(str, "]") + 2
	str = str[begin : len(str)-1]
	result := strings.Split(str, " ")

	if len(expected) != len(result) {
		t.Error("Expected and result slices should have the same length")
	}

	for i := range result {
		if !slices.Contains(expected, result[i]) {
			t.Errorf("Expected %s but got %s", expected, result)
		}
	}
}

func TestPairString(t *testing.T) {
	expected := "Seq2[int,iters.Pair[int,int]][0:0:1 1:1:2 2:2:3 3:3:4 4:4:5]"
	result := fmt.Sprint(it.Enumerate2(0, it.Range(1, 6).All()))

	if expected != result {
		t.Errorf("Expected %s but got %s", expected, result)
	}
}
