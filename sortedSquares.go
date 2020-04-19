package main

import (
	"fmt"
	"sort"
)

/*
https://leetcode.com/explore/learn/card/fun-with-arrays/521/introduction/3240/
Squares of a Sorted Array
Given an array of integers A sorted in non-decreasing order, return an array of the squares of each number, also in sorted non-decreasing order.


Example 1:
Input: [-4,-1,0,3,10]
Output: [0,1,9,16,100]

Example 2:
Input: [-7,-3,2,3,11]
Output: [4,9,9,49,121]


Note:
1 <= A.length <= 10000
-10000 <= A[i] <= 10000
A is sorted in non-decreasing order.
*/
func main() {
	input := []int{-4, -1, 0, 3, 10}
	fmt.Println(sortedSquares(input))
}

func sortedSquares(A []int) []int {
	for i, num := range A {
		A[i] = num * num
	}
	sort.Ints(A[:])
	return A
}
