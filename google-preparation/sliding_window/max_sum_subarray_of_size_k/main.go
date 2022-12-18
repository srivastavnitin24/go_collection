package main

import (
	"fmt"
	"math"
)

// Given an array of positive numbers and a positive number ‘k,’
// find the maximum sum of any contiguous subarray of size ‘k’.
// Input: [2, 1, 5, 1, 3, 2], k=3
// Output: 9
// Explanation: Subarray with maximum sum is [5, 1, 3].

// Input: [2, 3, 4, 1, 5], k=2
// Output: 7
// Explanation: Subarray with maximum sum is [3, 4].
func main() {
	arr := []int{2, 3, 4, 1, 5} // len:5-k+1
	k := 3

	max := Max_sum_subarray_optimal(arr, k)
	fmt.Println("MAX Sub Array Optimal : ", max)
}

func Max_sum_subarray_optimal(arr []int, k int) int {

	slidingStart, sum, max := 0, 0, 0
	for i := 0; i < len(arr); i++ {
		sum = sum + arr[i]
		if i >= k-1 {
			// if max < sum {
			// 	max = sum
			// }
			max = int(math.Max(float64(max), float64(sum)))
			sum = sum - arr[slidingStart]
			slidingStart = slidingStart + 1
		}
	}
	return max
}

//non optimal code
// func Max_sum_subarray(arr []int, k int) int {

// 	sum, max := 0, 0
// 	for i := 0; i < len(arr)-k+1; i++ {
// 		for j := i; j < i+k; j++ {
// 			sum = sum + arr[j]
// 		}
// 		if max < sum {
// 			max = sum
// 		}
// 		sum = 0
// 	}
// 	return max
// }
