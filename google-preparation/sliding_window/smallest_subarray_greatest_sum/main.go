package main

import (
	"fmt"
	"math"
)

// main func for sliding window problems
func main() {
	smallest_subarray_greatest_sum()

}

// find the length of the smallest contiguous subarray whose sum is greater than or equal to ‘S’.
// Return 0 if no such subarray exists.

// Input: [2, 1, 5, 2, 3, 2], S=7
// Output: 2
// Explanation: The smallest subarray with a sum greater than or equal to ‘7’ is [5, 2].

// Input: [2, 1, 5, 2, 8], S=7
// Output: 1
// Explanation: The smallest subarray with a sum greater than or equal to ‘7’ is [8].
func smallest_subarray_greatest_sum() {
	// arr := []int{2, 1, 5, 2, 3, 2}
	// k := 7

	arr := []int{2, 1, 5, 2, 8}
	k := 7

	// arr := []int{3, 4, 1, 1, 6}
	// k := 8

	res := optimal(arr, k)
	fmt.Println("Result : ", res)
}

func optimal(arr []int, k int) int {
	slidingStart, sum, minArr := 0, 0, math.MaxInt
	for i := 0; i < len(arr); i++ {
		sum = sum + arr[i]
		for sum >= k {
			minArr = int(math.Min(float64(minArr), float64(i-slidingStart+1)))
			sum = sum - arr[slidingStart]
			slidingStart = slidingStart + 1
		}
	}
	if minArr == math.MaxInt {
		return 0
	}
	return minArr
}
