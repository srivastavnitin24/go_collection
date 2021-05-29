package binarysearch

func findElementInRotatedSortedArray(arr []int, find int) int {
	index := find_index_in_rotated_array(arr)
	if index == -1 {
		return -1
	}
	start, end := 0, len(arr)-1

	x := binarySearchFindElement(arr[start:index], find)
	if x != -1 {
		return x
	}
	y := binarySearchFindElement(arr[index:end], find)
	if y != -1 {
		return y
	}
	return -1
}

func binarySearchFindElement(arr []int, find int) int {
	sorted := ""
	if len(arr) == 0 {
		return -1
	}
	if len(arr) == 1 {
		if arr[0] == find {
			return 0
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
				return mid
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
	return -1
}
