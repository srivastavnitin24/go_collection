package binarysearch

func ceilofelementinsortedarray(arr []int, find int) int {
	if len(arr) == 0 {
		return -1
	}
	res := -1
	start, end := 0, len(arr)-1
	for start <= end {
		mid := start + (end-start)/2
		if arr[mid] == find {
			return mid
		}
		//else {
		//	res = arr[mid]
		//}

		if arr[mid] < find {
			start = mid + 1
		}
		if arr[mid] > find {
			res = arr[mid]
			end = mid - 1
		}
	}

	return res
}
