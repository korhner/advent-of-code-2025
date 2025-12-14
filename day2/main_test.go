package main

import "testing"

func TestIsFake(t *testing.T) {
	var testCases = []struct {
		number   int
		expected bool
	}{
		{1212, true},
		{5, false},
		{123123, true},
	}

	for _, tc := range testCases {
		result := isFake(tc.number)
		if result != tc.expected {
			t.Errorf("isFake(%d) = %v; want %v", tc.number, result, tc.expected)
		}
	}
}
