package main

import "testing"

func TestLine_Intersects(t *testing.T) {

	tests := []struct {
		Line1    *Line
		Line2    *Line
		expected bool
	}{
		// Line1 vertical, Line2 horizontal
		// intersecting
		{&Line{Start: &Point{X: 2, Y: 1}, End: &Point{X: 2, Y: 4}}, &Line{Start: &Point{X: 1, Y: 3}, End: &Point{X: 4, Y: 3}}, true},
		// intersecting at endpoint
		{&Line{Start: &Point{X: 2, Y: 1}, End: &Point{X: 2, Y: 4}}, &Line{Start: &Point{X: 1, Y: 4}, End: &Point{X: 4, Y: 4}}, true},
		// intersecting at endpoint with endpoint
		{&Line{Start: &Point{X: 2, Y: 1}, End: &Point{X: 2, Y: 4}}, &Line{Start: &Point{X: 2, Y: 4}, End: &Point{X: 4, Y: 4}}, true},
		// non-intersecting, above
		{&Line{Start: &Point{X: 2, Y: 1}, End: &Point{X: 2, Y: 4}}, &Line{Start: &Point{X: 1, Y: 5}, End: &Point{X: 5, Y: 5}}, false},
		// non-intersecting, left
		{&Line{Start: &Point{X: 2, Y: 1}, End: &Point{X: 2, Y: 4}}, &Line{Start: &Point{X: 0, Y: 2}, End: &Point{X: 1, Y: 2}}, false},

		// Line1 horizontal, Line2 vertical
		// intersecting
		{&Line{Start: &Point{X: 1, Y: 3}, End: &Point{X: 4, Y: 3}}, &Line{Start: &Point{X: 2, Y: 1}, End: &Point{X: 2, Y: 4}}, true},
		// intersecting at endpoint
		{&Line{Start: &Point{X: 1, Y: 4}, End: &Point{X: 4, Y: 4}}, &Line{Start: &Point{X: 2, Y: 1}, End: &Point{X: 2, Y: 4}}, true},
		// intersecting at endpoint with endpoint
		{&Line{Start: &Point{X: 2, Y: 4}, End: &Point{X: 4, Y: 4}}, &Line{Start: &Point{X: 2, Y: 1}, End: &Point{X: 2, Y: 4}}, true},
		// non-intersecting, above
		{&Line{Start: &Point{X: 1, Y: 5}, End: &Point{X: 5, Y: 5}}, &Line{Start: &Point{X: 2, Y: 1}, End: &Point{X: 2, Y: 4}}, false},
		// non-intersecting, left
		{&Line{Start: &Point{X: 0, Y: 2}, End: &Point{X: 1, Y: 2}}, &Line{Start: &Point{X: 2, Y: 1}, End: &Point{X: 2, Y: 4}}, false},

		// Both lines horizontal
		// intersecting
		{&Line{Start: &Point{X: 1, Y: 3}, End: &Point{X: 4, Y: 3}}, &Line{Start: &Point{X: 3, Y: 3}, End: &Point{X: 5, Y: 3}}, true},
		// non-intersecting
		{&Line{Start: &Point{X: 1, Y: 3}, End: &Point{X: 2, Y: 3}}, &Line{Start: &Point{X: 3, Y: 3}, End: &Point{X: 5, Y: 3}}, false},

		// Both lines vertical
		// intersecting
		{&Line{Start: &Point{X: 2, Y: 1}, End: &Point{X: 2, Y: 4}}, &Line{Start: &Point{X: 2, Y: 3}, End: &Point{X: 2, Y: 5}}, true},
		// non-intersecting
		{&Line{Start: &Point{X: 2, Y: 1}, End: &Point{X: 2, Y: 2}}, &Line{Start: &Point{X: 2, Y: 3}, End: &Point{X: 2, Y: 5}}, false},
	}
	for _, tt := range tests {
		if got := tt.Line1.Intersects(tt.Line2); got != tt.expected {
			t.Errorf("Intersects() = %v, want %v", got, tt.expected)
		}
	}
}
