package main

import (
	"fmt"
)

/*
https://leetcode.com/explore/learn/card/fun-with-arrays/523/conclusion/3270/

Find All Numbers Disappeared in an Array
Given an array of integers where 1 ≤ a[i] ≤ n (n = size of array), some elements appear twice and others appear once.
Find all the elements of [1, n] inclusive that do not appear in this array.
Could you do it without extra space and in O(n) runtime? You may assume the returned list does not count as extra space.

Example:
Input:
[4,3,2,7,8,2,3,1]

Output:
[5,6]

[4 3 2 -7 8 2 3 1]
[4 3 -2 -7 8 2 3 1]
[4 -3 -2 -7 8 2 3 1]
[4 -3 -2 -7 8 2 -3 1]
[4 -3 -2 -7 8 2 -3 -1]
[4 -3 -2 -7 8 2 -3 -1]
[4 -3 -2 -7 8 2 -3 -1]
[-4 -3 -2 -7 8 2 -3 -1]
             * *
*/
func main() {
	input := []int{4, 3, 2, 7, 8, 2, 3, 1}
	fmt.Println(findDisappearedNumbers(input))
}

func findDisappearedNumbers(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		j := abs(nums[i]) - 1
		nums[j] = abs(nums[j]) * -1
	}

	res := []int{}
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			res = append(res, i+1)
		}
	}
	return res
}

func abs(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}
