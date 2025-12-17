package main

import (
	"testing"
)

func Test_findMaxJoltage(t *testing.T) {
	var testCases = []struct {
		input    []int
		expected int
	}{
		{[]int{2, 1, 5, 3, 4}, 54},
		{[]int{1, 2, 3, 4, 5}, 45},
		{[]int{5, 4, 3, 2, 1}, 54},
		{[]int{8, 1, 1, 1, 9}, 89},
	}

	for _, tc := range testCases {
		result := findMaxJoltage(tc.input, 0, len(tc.input)-1, 2)
		if result != tc.expected {
			t.Errorf("findMaxJoltage(%v) = %d; want %d", tc.input, result, tc.expected)
		}
	}
}
