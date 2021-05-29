package binarysearch

func countElementInSortedArray(arr []int, find int) int {
	firstOccurrence, lastOccurrence := 0, 0
	if len(arr) == 0 {
		return 0
	}
	if len(arr) == 1 {
		if arr[0] == find {
			return 1
		}
		return 0
	} else if len(arr) > 1 {
		l, u := 0, len(arr)-1
		for l <= u {
			mid := l + (u-l)/2
			if arr[mid] == find {
				firstOccurrence = mid
				lastOccurrence = mid
				u = mid - 1 //for 1st occurrence
				l = mid + 1 //for last occurrence
			}
			if arr[mid] > find {
				u = mid - 1
			}
			if arr[mid] < find {
				l = mid + 1
			}
		}
	}
	return lastOccurrence - firstOccurrence + 1
}
