package main

import "testing"

func TestMinMoves(t *testing.T) {
	tests := []struct {
		name   string
		queens []int
		want   int
	}{
		{"example 1: diagonal", []int{1, 2, 3, 4, 5, 6, 7, 8}, 7},
		{"example 2: all in row 1", []int{1, 1, 1, 1, 1, 1, 1, 1}, 7},
		{"already a solution", []int{1, 5, 8, 6, 3, 7, 2, 4}, 0},
		{"another solution", []int{4, 2, 7, 3, 6, 8, 5, 1}, 0},
		{"one queen moved from a solution", []int{1, 5, 8, 6, 3, 7, 2, 1}, 1},
		{"all in row 8", []int{8, 8, 8, 8, 8, 8, 8, 8}, 7},
		{"reverse diagonal", []int{8, 7, 6, 5, 4, 3, 2, 1}, 7},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := minMoves(tc.queens); got != tc.want {
				t.Errorf("minMoves(%v) = %d, want %d", tc.queens, got, tc.want)
			}
		})
	}
}

func TestMinMovesDoesNotModifyInput(t *testing.T) {
	q := []int{1, 1, 1, 1, 1, 1, 1, 1}
	minMoves(q)
	for i, v := range q {
		if v != 1 {
			t.Errorf("input modified at %d: %v", i, q)
		}
	}
}

func TestMinMovesBounds(t *testing.T) {
	// A valid placement is always reachable with at most 8 moves (one per column).
	got := minMoves([]int{3, 3, 3, 3, 3, 3, 3, 3})
	if got < 0 || got > 8 {
		t.Errorf("moves out of range: %d", got)
	}
}
