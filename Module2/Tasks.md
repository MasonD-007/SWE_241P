Task Description:  Complete the following tasks:

Task-1: Given an array of integer nums sorted in non-decreasing order, find a given target value's first and last position. Suppose the target is not found in the array; return [-1, -1]. You must write an algorithm with O(log n) runtime complexity.  Also, write additional sample test cases to validate your implementation.

Example 1:
Input: nums = [4,9,10,13,17,17,19,21], target = 17
Output: [4,5]

Example 2:
Input: nums = [2,4,6,8,10,14,16], target = 12
Output: [-1,-1]

Example 3:
Input: nums = [], target = 0
Output: [-1,-1]

 

Task-2:

You are given an m x n integer matrix with the following two properties:

Each row is sorted in non-decreasing order.
The first integer of each row is greater than the last integer of the previous row.
Given an integer target, return true if the target is in the matrix or false otherwise. Write a solution in O(log(m * n)) time complexity. Also, write additional sample test cases to validate your implementation.

Example 1:
Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
Output: true

Example 2:
Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13
Output: false
