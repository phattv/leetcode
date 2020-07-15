package main

import "fmt"

/*
https://leetcode.com/explore/learn/card/fun-with-arrays/523/conclusion/3231/

Given a non-empty array of integers, return the third maximum number in this array. If it does not exist, return the maximum number. The time complexity must be in O(n).

Example 1:
Input: [3, 2, 1]
Output: 1
Explanation: The third maximum is 1.

Example 2:
Input: [1, 2]
Output: 2
Explanation: The third maximum does not exist, so the maximum (2) is returned instead.

Example 3:
Input: [2, 2, 3, 1]
Output: 1
Explanation: Note that the third maximum here means the third maximum distinct number.
Both numbers with value 2 are both considered as second maximum.
*/
func main() {
	input := []int{1, -2147483648, 2}
	fmt.Println(thirdMax(input))
}

func thirdMax(nums []int) int {
	N := len(nums)
	if N < 2 {
		return nums[0]
	}
	if N < 3 {
		if nums[0] > nums[1] {
			return nums[0]
		}
		return nums[1]
	}

	defaultMin := -9999999999
	max := nums[0]
	secondToMax := defaultMin
	thirdToMax := defaultMin
	for i := 1; i < N; i++ {
		num := nums[i]
		if num > max {
			thirdToMax = secondToMax
			secondToMax = max
			max = num
			continue
		}
		if num > secondToMax && num < max {
			thirdToMax = secondToMax
			secondToMax = num
			continue
		}
		if num > thirdToMax && num < secondToMax {
			thirdToMax = num
		}
	}
	fmt.Println(max)
	fmt.Println(secondToMax)
	fmt.Println(thirdToMax)

	if thirdToMax != defaultMin {
		return thirdToMax
	}
	return max
}
