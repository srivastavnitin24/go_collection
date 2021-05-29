package binarysearch

func floorofelementinsortedarray(arr []int, find int) int {
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
		if arr[mid] < find {
			res = arr[mid]
			start = mid + 1
		}
		if arr[mid] > find {
			end = mid - 1
		}
	}

	return res
}
