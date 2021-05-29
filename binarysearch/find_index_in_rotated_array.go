package binarysearch

//To find the no of rotation in an sorted array
// 1: Find the smallest value index and return
// 2: If not 1st step then find the compare neighbour element and again search in unsorted array

func find_index_in_rotated_array(arr []int) int {
	if len(arr) == 0 || len(arr) == 1 {
		return -1
	}
	start, end := 0, len(arr)-1
	N := len(arr)
	for start <= end {
		mid := start + (end-start)/2
		prev := (mid + N - 1) % N
		next := (mid + 1) % N

		if arr[mid] < arr[prev] && arr[mid] < arr[next] {
			return mid
		}
		if arr[start] <= arr[mid] { //if array is sorted move to next half
			start = mid + 1
		} else if arr[mid] <= arr[end] { //if array is sorted move to previous half
			end = mid - 1
		}
	}
	return -1
}
