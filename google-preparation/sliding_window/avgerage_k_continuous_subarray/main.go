package main

import "fmt"

// main func for sliding window problems
func main() {
	continous_k_subarray_average()

}

func continous_k_subarray_average() {

	//Given an array, find the average of all subarrays of ‘K’ contiguous elements in it.
	//Array: [1, 3, 2, 6, -1, 4, 1, 8, 2], K=5
	arr := []int{1, 3, 2, 6, -1, 4, 1, 8, 2} //9-k+1
	k := 5

	res := subArrAvgOptimal(arr, k)
	fmt.Println("Result Optimal ", res)
}

// Array: [1, 3, 2, 6, -1, 4, 1, 8, 2], K=5
// Here is the final output containing the averages of all subarrays of size 5:
// Output: [2.2, 2.8, 2.4, 3.6, 2.8]
func subArrAvgOptimal(arr []int, k int) []float64 {
	avgArr := make([]float64, 0)
	if k == 0 {
		fmt.Println("Cannot find the avg k > 0")
		return avgArr
	}

	sum := 0.0
	slidingStart := 0
	for i := 0; i < len(arr); i++ {
		sum = sum + float64(arr[i])
		if i >= k-1 {
			avgArr = append(avgArr, sum/float64(k))
			sum = sum - float64(arr[slidingStart])
			slidingStart = slidingStart + 1
		}
	}

	return avgArr

}

//non optimal code
// func subArrAvg(arr []int, k int) []float64 {
// 	avgArr := make([]float64, 0)
// 	if k == 0 {
// 		fmt.Println("Cannot find the avg k > 0")
// 		return avgArr
// 	}

// 	for i := 0; i < len(arr)-k+1; i++ {
// 		sum := 0.0
// 		for j := i; j < i+k; j++ {
// 			sum = sum + float64(arr[j])
// 		}
// 		avgArr = append(avgArr, sum/float64(k))
// 	}

// 	return avgArr

// }
