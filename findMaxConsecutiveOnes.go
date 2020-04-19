package main

import "fmt"

/*
https://leetcode.com/explore/learn/card/fun-with-arrays/521/introduction/3238/
Max Consecutive Ones

Given a binary array, find the maximum number of consecutive 1s in this array.

Example 1:
Input: [1,1,0,1,1,1]
Output: 3
Explanation: The first two digits or the last three digits are consecutive 1s.
		The maximum number of consecutive 1s is 3.

Note:
The input array will only contain 0 and 1.
The length of input array is a positive integer and will not exceed 10,000
*/
func main() {
	input := []int{1, 1, 0, 1, 1, 1}
	fmt.Println(findMaxConsecutiveOnes(input))
}

func findMaxConsecutiveOnes(nums []int) int {
	count := 0
	max := 0
	for _, num := range nums {
		if num == 1 {
			count++
			if count >= max {
				max = count
			}
		} else {
			count = 0
		}
	}
	return max
}
