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

func TestIsFakePart2(t *testing.T) {
	var testCases = []struct {
		number   int
		expected bool
	}{
		{1212, true},
		{5, false},
		{123123, true},
		{11, true},
		{111, true},
		{1111, true},
		{1234, false},
		{12341234, true},
		{121212, true},
		{121213, false},
	}

	for _, tc := range testCases {
		result := isFakePart2(tc.number)
		if result != tc.expected {
			t.Errorf("isFake(%d) = %v; want %v", tc.number, result, tc.expected)
		}
	}
}
