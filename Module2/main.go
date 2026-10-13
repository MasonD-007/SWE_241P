package main

func main() {
	sample1 := []int{4, 9, 10, 13, 17, 17, 19, 21}
	result1 := searchRange(sample1, 17)
	println("Sample Test Case 1: ", result1[0], result1[1]) // Expected output: [4,5]

	sample2 := [][]int{{1, 3, 5}, {7, 9, 11}, {13, 15, 17}}
	result2 := searchMatrix(sample2, 9)
	println("Sample Test Case 2: ", result2) // Expected output: true
}

// Task-1: Given an array of integer nums sorted in non-decreasing order, find a given target value's first and last position. Suppose the target is not found in the array; return [-1, -1]. You must write an algorithm with O(log n) runtime complexity.  Also, write additional sample test cases to validate your implementation.

func searchRange(nums []int, target int) []int {
	res := []int{-1, -1}
	left, right := 0, len(nums)-1
	pos := -1

	// Find the first occurrence of the target
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			res[0] = mid
			pos = mid
			right = mid - 1
		}
	}

	if pos == -1 {
		return res // Target not found
	}

	// If the target is found, find the last occurrence
	left, right = pos, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			res[1] = mid
			left = mid + 1
		}
	}

	return res
}

// You are given an m x n integer matrix with the following two properties:
//
// Each row is sorted in non-decreasing order.
// The first integer of each row is greater than the last integer of the previous row.
// Given an integer target, return true if the target is in the matrix or false otherwise. Write a solution in O(log(m * n)) time complexity. Also, write additional sample test cases to validate your implementation.

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	rows, cols := len(matrix), len(matrix[0])
	left, right := 0, rows*cols-1

	for left <= right {
		mid := left + (right-left)/2
		midValue := matrix[mid/cols][mid%cols]
		if midValue == target {
			return true
		} else if midValue < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return false
}
