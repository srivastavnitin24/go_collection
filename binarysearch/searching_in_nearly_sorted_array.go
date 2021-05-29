package binarysearch

func searchingInNearlySortedArray(arr []int, find int) int {
	if len(arr) == 0 || len(arr) == 1 {
		return -1
	}
	start, end := 0, len(arr)-1
	for start <= end {
		mid := start + (end-start)/2
		if arr[mid] == find {
			return mid
		} else if mid-1 >= start && arr[mid-1] == find {
			return mid - 1
		} else if mid+1 <= end && arr[mid+1] == find {
			return mid + 1
		}

		if arr[mid] < find {
			start = mid + 2
		}
		if arr[mid] > find {
			end = mid - 2
		}
	}

	return -1
}
