package main

import (
	"testing"
)

func equalSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func Test_buildMaxArray(t *testing.T) {
	var testCases = []struct {
		input    []int
		expected []int
	}{
		{[]int{2, 1, 5, 3, 4}, []int{5, 5, 5, 4, 4}},
		{[]int{1, 2, 3, 4, 5}, []int{5, 5, 5, 5, 5}},
		{[]int{5, 4, 3, 2, 1}, []int{5, 4, 3, 2, 1}},
		{[]int{3, 3, 3}, []int{3, 3, 3}},
		{[]int{}, []int{}},
	}

	for _, tc := range testCases {
		result := buildMaxArray(tc.input)
		if !equalSlices(result, tc.expected) {
			t.Errorf("buildMaxArray(%v) = %v; want %v", tc.input, result, tc.expected)
		}
	}
}

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
		result := findMaxJoltage(tc.input)
		if result != tc.expected {
			t.Errorf("findMaxJoltage(%v) = %d; want %d", tc.input, result, tc.expected)
		}
	}
}
