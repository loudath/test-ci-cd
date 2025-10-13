package main

import "testing"

func TestAdd(t *testing.T) {
	type test struct {
		a, b   int
		result int
	}

	tests := []test{
		{a: 1, b: 2, result: 3},
		{a: -1, b: 1, result: 0},
		{a: 0, b: 0, result: 0},
	}

	for _, tc := range tests {
		got := Add(tc.a, tc.b)
		if got != tc.result {
			t.Errorf("Add(%d, %d) = %d; want %d", tc.a, tc.b, got, tc.result)
		}
	}
}
