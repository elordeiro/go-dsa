package set

import "testing"

func TestSet(t *testing.T) {
	tests := []struct {
		nums      []int
		expecteds []bool
	}{
		{[]int{1, 1, 2, 2, 3, 3, 7}, []bool{true, false, true, false, true, false, false}},
	}

	s := NewSet()
	for i := range 4 {
		s.add(i + 1)
	}
	var testname string
	for _, test := range tests {
		for i := range len(test.nums) {
			testname = "Contains"
			t.Run(testname, func(t *testing.T) {
				actual := s.contains(test.nums[i])
				expected := test.expecteds[i]
				if actual != expected {
					t.Errorf("Actual   : %v", actual)
					t.Errorf("Expected : %v", expected)
				}
			})
			s.remove(test.nums[i])
		}
	}
}
