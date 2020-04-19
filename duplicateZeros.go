package main

import "log"

/*
https://leetcode.com/explore/learn/card/fun-with-arrays/525/inserting-items-into-an-array/3245/
Duplicate Zeros
Given a fixed length array arr of integers, duplicate each occurrence of zero, shifting the remaining elements to the right.
Note that elements beyond the length of the original array are not written.
Do the above modifications to the input array in place, do not return anything from your function.


Example 1:
Input: [1,0,2,3,0,4,5,0]
Output: null
Explanation: After calling your function, the input array is modified to: [1,0,0,2,3,0,0,4]

Example 2:
Input: [1,2,3]
Output: null
Explanation: After calling your function, the input array is modified to: [1,2,3]


Note:

1 <= arr.length <= 10000
0 <= arr[i] <= 9
*/
func main() {
	arr := []int{0, 4, 1, 0, 0, 8, 0, 0, 3}
	// 0, 4, 1, 0, 0, 8, 0, 0, 3
	// ^
	// 0, 0, 4, 1, 0, 0, 8, 0, 0
	//             ^
	// 0, 0, 4, 1, 0, 0, 0, 8, 0
	//                   ^
	// 0, 0, 4, 1, 0, 0, 0, 0, 8 (expected)
	duplicateZeros(arr)
	log.Println(arr)
}

// runtime: 12.5 ms
// func duplicateZeros(arr []int) {
// 	arrLength := len(arr)
// 	index := 0
// 	for index < arrLength {
// 		if arr[index] == 0 {
// 			// Shift element 1 position to the right
// 			for index2 := arrLength - 2; index2 > index; index2-- {
// 				arr[index2+1] = arr[index2]
// 			}
// 			// Duplicate zero
// 			if index < arrLength-1 {
// 				arr[index+1] = 0
// 			}
// 			// Skip next index
// 			index += 2
// 		} else {
// 			index++
// 		}
// 	}
// }

// runtime: 4ms
func duplicateZeros(arr []int) {
	length := len(arr)
	originalArray := make([]int, length)
	copy(originalArray, arr)

	index := 0
	for _, value := range originalArray {
		arr[index] = value
		index++

		if index == length {
			break
		}

		if value == 0 {
			arr[index] = 0
			index++
		}

		if index == length {
			break
		}
	}
}
