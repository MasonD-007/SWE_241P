package main

import (
	"math/rand"
	"reflect"
	"sort"
	"testing"
)

func TestSearchRange(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{"sample from main", []int{4, 9, 10, 13, 17, 17, 19, 21}, 17, []int{4, 5}},
		{"empty array", []int{}, 5, []int{-1, -1}},
		{"nil array", nil, 5, []int{-1, -1}},
		{"single element found", []int{5}, 5, []int{0, 0}},
		{"single element not found", []int{5}, 3, []int{-1, -1}},
		{"single element target larger", []int{5}, 7, []int{-1, -1}},
		{"all elements equal target", []int{2, 2, 2, 2, 2}, 2, []int{0, 4}},
		{"all elements equal, not target", []int{2, 2, 2, 2, 2}, 3, []int{-1, -1}},
		{"target at start", []int{1, 1, 2, 3, 4}, 1, []int{0, 1}},
		{"target at end", []int{1, 2, 3, 4, 4}, 4, []int{3, 4}},
		{"unique target in middle", []int{1, 3, 5, 7, 9}, 5, []int{2, 2}},
		{"first element unique", []int{1, 3, 5, 7, 9}, 1, []int{0, 0}},
		{"last element unique", []int{1, 3, 5, 7, 9}, 9, []int{4, 4}},
		{"target smaller than all", []int{5, 6, 7}, 1, []int{-1, -1}},
		{"target larger than all", []int{5, 6, 7}, 10, []int{-1, -1}},
		{"target between elements", []int{1, 3, 5, 7}, 4, []int{-1, -1}},
		{"two elements, both match", []int{3, 3}, 3, []int{0, 1}},
		{"two elements, second matches", []int{1, 3}, 3, []int{1, 1}},
		{"negative numbers", []int{-5, -3, -3, -3, 0, 2}, -3, []int{1, 3}},
		{"negative target missing", []int{-5, -3, 0, 2}, -4, []int{-1, -1}},
		{"zero target", []int{-1, 0, 0, 1}, 0, []int{1, 2}},
		{"duplicates around target", []int{1, 2, 2, 3, 3, 3, 4, 4}, 3, []int{3, 5}},
		{"large values", []int{1 << 40, 1 << 40, 1 << 50}, 1 << 40, []int{0, 1}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := searchRange(tc.nums, tc.target)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("searchRange(%v, %d) = %v, want %v", tc.nums, tc.target, got, tc.want)
			}
		})
	}
}

func TestSearchRangeDoesNotMutateInput(t *testing.T) {
	nums := []int{1, 2, 2, 2, 3}
	orig := append([]int(nil), nums...)
	searchRange(nums, 2)
	if !reflect.DeepEqual(nums, orig) {
		t.Errorf("input mutated: got %v, want %v", nums, orig)
	}
}

// naiveSearchRange is a linear-scan reference implementation.
func naiveSearchRange(nums []int, target int) []int {
	res := []int{-1, -1}
	for i, v := range nums {
		if v == target {
			if res[0] == -1 {
				res[0] = i
			}
			res[1] = i
		}
	}
	return res
}

func TestSearchRangeAgainstReference(t *testing.T) {
	rng := rand.New(rand.NewSource(1))
	for iter := 0; iter < 2000; iter++ {
		n := rng.Intn(30)
		nums := make([]int, n)
		for i := range nums {
			nums[i] = rng.Intn(15) - 5
		}
		sort.Ints(nums)
		target := rng.Intn(19) - 7

		got := searchRange(nums, target)
		want := naiveSearchRange(nums, target)
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("searchRange(%v, %d) = %v, want %v", nums, target, got, want)
		}
	}
}

// Runs in O(log n): a huge array should be handled instantly.
func TestSearchRangeLargeInput(t *testing.T) {
	const n = 1_000_000
	nums := make([]int, n)
	for i := range nums {
		nums[i] = i / 10 // each value appears 10 times
	}
	got := searchRange(nums, 50000)
	want := []int{500000, 500009}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestSearchMatrix(t *testing.T) {
	standard := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}

	tests := []struct {
		name   string
		matrix [][]int
		target int
		want   bool
	}{
		{"sample from main", standard, 3, true},
		{"sample not found", standard, 13, false},
		{"nil matrix", nil, 1, false},
		{"empty matrix", [][]int{}, 1, false},
		{"empty first row", [][]int{{}}, 1, false},
		{"single cell found", [][]int{{5}}, 5, true},
		{"single cell not found", [][]int{{5}}, 4, false},
		{"single row found", [][]int{{1, 3, 5, 7}}, 5, true},
		{"single row not found", [][]int{{1, 3, 5, 7}}, 4, false},
		{"single column found", [][]int{{1}, {3}, {5}, {7}}, 7, true},
		{"single column not found", [][]int{{1}, {3}, {5}, {7}}, 2, false},
		{"first element", standard, 1, true},
		{"last element", standard, 60, true},
		{"first of a middle row", standard, 10, true},
		{"last of a middle row", standard, 20, true},
		{"smaller than all", standard, 0, false},
		{"larger than all", standard, 61, false},
		{"between rows", standard, 21, false},
		{"gap at row boundary low side", standard, 8, false},
		{"negatives found", [][]int{{-10, -5}, {-3, 0}, {2, 4}}, -3, true},
		{"negatives not found", [][]int{{-10, -5}, {-3, 0}, {2, 4}}, -4, false},
		{"duplicates within row", [][]int{{1, 1, 1}, {2, 2, 3}}, 2, true},
		{"large values", [][]int{{1 << 40, 1 << 41}, {1 << 50, 1 << 55}}, 1 << 50, true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := searchMatrix(tc.matrix, tc.target); got != tc.want {
				t.Errorf("searchMatrix(%v, %d) = %v, want %v", tc.matrix, tc.target, got, tc.want)
			}
		})
	}
}

func TestSearchMatrixDoesNotMutateInput(t *testing.T) {
	matrix := [][]int{{1, 3}, {5, 7}}
	searchMatrix(matrix, 5)
	want := [][]int{{1, 3}, {5, 7}}
	for i := range want {
		for j := range want[i] {
			if matrix[i][j] != want[i][j] {
				t.Fatalf("matrix mutated at [%d][%d]: got %d, want %d", i, j, matrix[i][j], want[i][j])
			}
		}
	}
}

// Every value present is found; every value absent (including gaps) is not.
func TestSearchMatrixExhaustive(t *testing.T) {
	rng := rand.New(rand.NewSource(1))
	for iter := 0; iter < 500; iter++ {
		m, n := rng.Intn(6)+1, rng.Intn(6)+1
		matrix := make([][]int, m)
		present := map[int]bool{}
		v := rng.Intn(10) - 5
		for i := range matrix {
			matrix[i] = make([]int, n)
			for j := range matrix[i] {
				matrix[i][j] = v
				present[v] = true
				v += rng.Intn(3) + 1 // strictly increasing, with gaps
			}
		}
		for target := -10; target <= v+2; target++ {
			if got := searchMatrix(matrix, target); got != present[target] {
				t.Fatalf("searchMatrix(%v, %d) = %v, want %v", matrix, target, got, present[target])
			}
		}
	}
}

func TestSearchMatrixLargeInput(t *testing.T) {
	const m, n = 1000, 1000
	matrix := make([][]int, m)
	for i := range matrix {
		matrix[i] = make([]int, n)
		for j := range matrix[i] {
			matrix[i][j] = (i*n + j) * 2 // even numbers only
		}
	}
	if !searchMatrix(matrix, 1_999_998) {
		t.Error("expected last element to be found")
	}
	if !searchMatrix(matrix, 1_000_000) {
		t.Error("expected middle element to be found")
	}
	if searchMatrix(matrix, 1_000_001) {
		t.Error("expected odd value to be absent")
	}
}
