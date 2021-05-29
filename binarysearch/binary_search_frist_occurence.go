package binarysearch

func binarySearchFristOccurrence(arr []int, find int) int {
	res := -1
	sorted := ""
	if len(arr) == 0 {
		return -1
	}
	if len(arr) == 1 {
		if arr[0] == find {
			return -1
		}
		return -1
	} else if len(arr) > 1 {
		if arr[0] < arr[len(arr)-1] {
			sorted = "increasing"
		} else {
			sorted = "decreasing"
		}

		l, u := 0, len(arr)-1
		for l <= u {
			mid := l + (u-l)/2
			if arr[mid] == find {
				res = mid
				u = mid - 1 //for 1st occurrence
				//l = mid + 1 //for last occurrence
			}
			if arr[mid] > find {
				if sorted == "increasing" {
					u = mid - 1 //for increasing sorted array
				} else {
					l = mid + 1 //for Decreasing sorted array
				}
			}
			if arr[mid] < find {
				if sorted == "increasing" {
					l = mid + 1 //for increasing sorted array
				} else {
					u = mid - 1 //for Decreasing sorted array
				}
			}
		}
		return res
	}
	return -1
}
