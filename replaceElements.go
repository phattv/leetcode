package main

import "fmt"

/*
https://leetcode.com/explore/learn/card/fun-with-arrays/511/in-place-operations/3259/
Replace Elements with Greatest Element on Right Side
Given an array arr, replace every element in that array with the greatest element among the elements to its right, and replace the last element with -1.
After doing so, return the array.

Example 1:
Input: arr = [17,18,5,4,6,1]
Output: [18,6,6,6,1,-1]

Constraints:
1 <= arr.length <= 10^4
1 <= arr[i] <= 10^5
*/
func main() {
	input := []int{17, 18, 5, 4, 6, 1}
	fmt.Println(replaceElements(input))
}

func replaceElements(arr []int) []int {
	if arr == nil {
		return nil
	}

	N := len(arr)
	if N < 1 || N > 10000 {
		return nil
	}
	max := arr[N-1]
	arr[N-1] = -1

	for i := N - 2; i >= 0; i-- {
		prevMax := max
		if arr[i] > max {
			max = arr[i]
		}
		arr[i] = prevMax
	}

	return arr
}
