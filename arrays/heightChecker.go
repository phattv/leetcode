package main

import (
	"fmt"
	"sort"
)

/*
https://leetcode.com/explore/learn/card/fun-with-arrays/523/conclusion/3228/
Height Checker
Students are asked to stand in non-decreasing order of heights for an annual photo.

Return the minimum number of students that must move in order for all students to be standing in non-decreasing order of height.

Notice that when a group of students is selected they can reorder in any possible way between themselves and the non selected students remain on their seats.


Example 1:
Input: heights = [1,1,4,2,1,3]
Output: 3
Explanation:
Current array : [1,1,4,2,1,3]
Target array  : [1,1,1,2,3,4]
On index 2 (0-based) we have 4 vs 1 so we have to move this student.
On index 4 (0-based) we have 1 vs 3 so we have to move this student.
On index 5 (0-based) we have 3 vs 4 so we have to move this student.

Example 2:
Input: heights = [5,1,2,3,4]
Output: 5

Example 3:
Input: heights = [1,2,3,4,5]
Output: 0

Constraints:
1 <= heights.length <= 100
1 <= heights[i] <= 100
*/
func main() {
	input := []int{5, 1, 2, 3, 4}
	fmt.Println(heightChecker(input))
}

func heightChecker(heights []int) int {
	// check constraints
	count := 0
	if heights == nil {
		return 0
	}
	N := len(heights)
	if N < 1 || N > 100 {
		return 0
	}

	copied := make([]int, N)
	copy(copied, heights)
	sort.Ints(copied)
	for i := 0; i < N; i++ {
		if heights[i] != copied[i] {
			count++
		}
	}

	return count
}
