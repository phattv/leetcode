package main

import (
	"fmt"
)

/*
https://leetcode.com/explore/learn/card/fun-with-arrays/511/in-place-operations/3157/
Given an array nums, write a function to move all 0's to the end of it while maintaining the relative order of the non-zero elements.

Example:
Input: [0,1,0,3,12]
Output: [1,3,12,0,0]

Note:
You must do this in-place without making a copy of the array.
Minimize the total number of operations.
*/
func main() {
	input := []int{0, 1, 0, 3, 12}
	moveZeroes(input)
	fmt.Println(input)
}

func moveZeroes(nums []int) {
	if nums == nil {
		return
	}

	N := len(nums)
	anchor := 0

	for i := 0; i < N; i++ {
		if nums[i] != 0 {
			// swap
			temp := nums[anchor]
			nums[anchor] = nums[i]
			nums[i] = temp
			anchor++
		}
	}
}
