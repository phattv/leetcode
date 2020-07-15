package main

import (
	"fmt"
)

/*
https://leetcode.com/explore/learn/card/fun-with-arrays/527/searching-for-items-in-an-array/3251/
Given an array A of integers, return true if and only if it is a valid mountain array.

Recall that A is a mountain array if and only if:

A.length >= 3
There exists some i with 0 < i < A.length - 1 such that:
A[0] < A[1] < ... A[i-1] < A[i]
A[i] > A[i+1] > ... > A[A.length - 1]

https://assets.leetcode.com/uploads/2019/10/20/hint_valid_mountain_array.png

Example 1:
Input: [2,1]
Output: false

Example 2:
Input: [3,5,5]
Output: false

Example 3:
Input: [0,3,2,1]
Output: true

Note:
0 <= A.length <= 10000
0 <= A[i] <= 10000
*/
func main() {
	input := []int{0, 1, 2, 8, 9}
	fmt.Println(validMountainArray(input))
}

func validMountainArray(A []int) bool {
	if A == nil || len(A) < 3 || len(A) > 10000 {
		return false
	}

	N := len(A)
	i := 0

	// walk up
	for i+1 < N && A[i] < A[i+1] {
		i++
	}

	// peak can't be first or last
	if i == 0 || i == N-1 {
		return false
	}

	// walk down
	for i+1 < N && A[i] > A[i+1] {
		i++
	}

	return i == N-1
}
