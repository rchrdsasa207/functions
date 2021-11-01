package main

import "testing"

// TestPow tests a Pow(x, y) function, which returns integer x**y.
func TestPow(t *testing.T) {
	for _, tc := range []struct {
		x, y uint
		want uint
	}{
		{x: 1, y: 2, want: 1},
		{x: 2, y: 2, want: 4},
		{x: 3, y: 2, want: 9},
		{x: 4, y: 2, want: 16},
		{x: 1, y: 3, want: 1},
		{x: 2, y: 4, want: 16},
		{x: 3, y: 5, want: 243},
		{x: 0, y: 2, want: 0},
		{x: 0, y: 5, want: 0},
	} {
		if got := pow(tc.x, tc.y); got != tc.want {
			t.Errorf("reverse(%v, %v) = %v, want = %v", tc.x, tc.y, got, tc.want)
		}
	}
}
