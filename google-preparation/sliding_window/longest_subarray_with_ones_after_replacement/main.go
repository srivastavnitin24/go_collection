package main

import (
	"fmt"
	"math"
)

/*
Given an array containing 0s and 1s, if you are allowed to replace no more than ‘k’ 0s with 1s, find the length of the longest contiguous subarray having all 1s.

Input: Array=[0, 1, 1, 0, 0, 0, 1, 1, 0, 1, 1], k=2
Output: 6
Explanation: Replace the '0' at index 5 and 8 to have the longest contiguous subarray of 1s having length 6.

Input: Array=[0, 1, 0, 0, 1, 1, 0, 1, 1, 0, 0, 1, 1], k=3
Output: 9
Explanation: Replace the '0' at index 6, 9, and 10 to have the longest contiguous subarray of 1s having length 9.
*/
func main() {
	arr := []int{0, 1, 1, 0, 0, 0, 1, 1, 0, 1, 1}
	k := 2
	maxCountRepeatingElement := 0
	slidingStart := 0
	maxlen := 0
	m := make(map[int]int)
	for index, key := range arr {
		_, ok := m[key]
		if ok {
			m[key] += 1
		} else {
			m[key] = 1
		}
		maxCountRepeatingElement = int(math.Max(float64(maxCountRepeatingElement), float64(m[key])))
		if (index - slidingStart + 1 - maxCountRepeatingElement) > k {
			//char := m[key]
			_, ok := m[key]
			if ok {
				m[key] -= 1
			}
			slidingStart += 1
		}
		maxlen = int(math.Max(float64(maxlen), float64(index-slidingStart+1)))
	}
	fmt.Println("map : ", m)
	fmt.Println("maxlen : ", maxlen)
}
