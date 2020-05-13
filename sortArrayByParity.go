package main

import "fmt"

/*
https://leetcode.com/explore/learn/card/fun-with-arrays/511/in-place-operations/3260/

Sort Array By Parity
Given an array A of non-negative integers, return an array consisting of all the even elements of A, followed by all the odd elements of A.

You may return any answer array that satisfies this condition.

Example 1:

Input: [3,1,2,4]
Output: [2,4,3,1]
The outputs [4,2,3,1], [2,4,1,3], and [4,2,1,3] would also be accepted.

Note:
1 <= A.length <= 5000
0 <= A[i] <= 5000
*/
func main() {
	input := []int{3, 1, 2, 4}
	fmt.Println(sortArrayByParity(input))
}

func sortArrayByParity(A []int) []int {
	if A == nil {
		return nil
	}

	N := len(A)
	if N < 1 || N > 5000 {
		return nil
	}

	// leftPointer := 0
	// rightPointer := N - 1
	// for leftPointer != rightPointer {
	// 	if A[leftPointer]%2 == 0 {
	// 		leftPointer++
	// 	} else {
	// 		if A[rightPointer]%2 == 0 {
	// 			// manual swap
	// 			// temp := A[leftPointer]
	// 			// A[leftPointer] = A[rightPointer]
	// 			// A[rightPointer] = temp
	// 			// swap
	// 			A[leftPointer], A[rightPointer] = A[rightPointer], A[leftPointer]
	// 			leftPointer++
	// 		} else {
	// 			rightPointer--
	// 		}
	// 	}
	// }

	// simplified
	evenIndex := 0
	for i := 0; i < N; i++ {
		if A[i]%2 == 0 {
			A[evenIndex], A[i] = A[i], A[evenIndex]
			evenIndex++
		}
	}

	return A
}
