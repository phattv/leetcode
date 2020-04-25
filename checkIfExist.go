package main

import (
	"fmt"
)

/*
https://leetcode.com/explore/learn/card/fun-with-arrays/527/searching-for-items-in-an-array/3250/
Check If N and Its Double Exist
Given an array arr of integers, check if there exists two integers N and M such that N is the double of M ( i.e. N = 2 * M).
More formally check if there exists two indices i and j such that :

i != j
0 <= i, j < arr.length
arr[i] == 2 * arr[j]

Example 1:
Input: arr = [10,2,5,3]
Output: true
Explanation: N = 10 is the double of M = 5,that is, 10 = 2 * 5.

Example 2:
Input: arr = [7,1,14,11]
Output: true
Explanation: N = 14 is the double of M = 7,that is, 14 = 2 * 7.

Example 3:
Input: arr = [3,1,7,11]
Output: false
Explanation: In this case does not exist N and M, such that N = 2 * M.

Constraints:
2 <= arr.length <= 500
-10^3 <= arr[i] <= 10^3
*/
func main() {
	// input := []int{10, 2, 5, 3}
	// fmt.Println(checkIfExist(input))

	// input2 := []int{1, 7, 14, 11}
	// fmt.Println(checkIfExist(input2))

	// input3 := []int{3, 1, 7, 11}
	// fmt.Println(checkIfExist(input3))

	input4 := []int{-778, -481, 842, 495, 44, 1000, -572, 977, 240, -116, 673, 997, -958, -539, -964, -187, -701, -928, 472, 965, -672, -88, 443, 36, 388, -127, 115, 704, -549, 1000, 998, 291, 633, 423, 57, -77, -543, 72, 328, -938, -192, 382, 179}
	fmt.Println(checkIfExist(input4))
}

func checkIfExist(arr []int) bool {
	if arr == nil || len(arr) < 2 || len(arr) > 500 {
		return false
	}

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == 2*arr[j] || 2*arr[i] == arr[j] {
				return true
			}
		}
	}

	return false
}
