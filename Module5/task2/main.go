package main

import (
	"fmt"
	"math"
)

func main() {
	// Example input: 1 5 8 6 3 7 2 4
	queens := []int{1, 5, 8, 6, 3, 7, 2, 4}
	moves := minMoves(queens)
	fmt.Println("Minimum moves required:", moves)
}

func minMoves(queens []int) int {
	res := math.MaxInt

	var backtrack func(col, moveCount int, rows, diag1, diag2 map[int]bool)
	backtrack = func(col, moveCount int, rows, diag1, diag2 map[int]bool) {
		if col == 8 {
			if moveCount < res {
				res = moveCount
			}
			return
		}

		for row := 1; row <= 8; row++ {
			if !rows[row] && !diag1[row+col] && !diag2[row-col] {
				// Calculate moves needed to place the queen in the new row
				newMoveCount := moveCount
				if queens[col] != row {
					newMoveCount++
				}

				if newMoveCount >= res {
					continue
				}
				// Place queen at (row, col)
				rows[row] = true
				diag1[row+col] = true
				diag2[row-col] = true

				backtrack(col+1, newMoveCount, rows, diag1, diag2)

				// Backtrack
				rows[row] = false
				diag1[row+col] = false
				diag2[row-col] = false
			}
		}
	}

	rows := make(map[int]bool)
	diag1 := make(map[int]bool)
	diag2 := make(map[int]bool)
	backtrack(0, 0, rows, diag1, diag2)

	return res
}
