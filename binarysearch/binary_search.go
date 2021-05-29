package binarysearch

func binarySearch(arr []int, find int) bool {
	sorted := ""
	if len(arr) == 0 {
		return false
	}
	if len(arr) == 1 {
		if arr[0] == find {
			return true
		}
		return false
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
				return true
			}
			if arr[mid] > find {
				if sorted == "increasing" {
					u = mid - 1 //for increasing sorted array
				} else {
					l = mid + 1 //for Decreasing sorted array
				}
			}
			if arr[mid] < find {
				if sorted == "decreasing" {
					l = mid + 1 //for increasing sorted array
				} else {
					u = mid - 1 //for Decreasing sorted array
				}
			}
		}
	}
	return false
}
